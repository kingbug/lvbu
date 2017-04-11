package utils

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	mpro "lvbu/models/project"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

	"github.com/astaxie/beego"
	"github.com/fsouza/go-dockerclient"
)

var (
	dockerbin         string
	dockerregsitry    string
	antbin            string
	configurl         string //配置文件下载地址
	configprocotol    string //配置文件下载协议[http|https]
	dockervolumepath  string //容器在主机上共享的主目录（主目录/容器名/指定路径）
	containerdir      string //容器工作目录
	PD                string //PathDelimiter
	PathPD            string //环境变量分隔符
	EXECPATH          string //main.go 执行目录
	Java7home         string
	Java8home         string
	Composerbin       string
	ComposerPackagist string //composer国内镜像地址
	Tomcat6home       string
)

func InitDocker() {
	dockerbin = beego.AppConfig.String("Dockerbin")
	dockerregsitry = beego.AppConfig.String("Dockerregsitry")

	antbin = beego.AppConfig.String("Antbin")
	configurl = beego.AppConfig.String("ConfigServerIp")
	configprocotol = beego.AppConfig.String("ConfigServerProtocol")
	dockervolumepath = beego.AppConfig.String("DockerVolumePath")

	Composerbin = beego.AppConfig.String("Composerbin")
	ComposerPackagist = beego.AppConfig.String("ComposerPackagist")
	Java7home = beego.AppConfig.String("Java7Home")
	Java8home = beego.AppConfig.String("Java8Home")
	Tomcat6home = beego.AppConfig.String("Tomcat6Home")

	if dockerbin == "" {
		dockerbin = "docker"
	}
	if dockerregsitry == "" {
		beego.Error("app.conf -> key:Dockerregsitry is nil")
	}
	if antbin == "" {
		antbin = "ant"
	}
	if configurl == "" {
		beego.Error("ConfigServerUrl not found. 查看app.conf ConfigServerUrl是否设置正确")
	} else {
		configurl = configurl + ":" + beego.AppConfig.String("httpport")
	}
	if dockervolumepath == "" {
		dockervolumepath = "/lvbu"
	} else if dockervolumepath[:1] != "/" {
		panic("app.conf -> DockerVolumePath 必须使用绝对路径")
	}
	bash := exec.Command(dockerbin)
	if err := bash.Run(); err != nil {
		beego.Error("docker deam not found. 查看app.conf Dockerbin是否设置正确")
	}
	containerdir = "/cihi/run"
	if runtime.GOOS == "linux" {
		PD = "/"
		PathPD = ":"
	} else {
		PD = "\\"
		PathPD = ";"
	}

	if Composerbin == "" {
		Composerbin = "composer"
	}
	if Tomcat6home == "" {
		Tomcat6home = ".profile/tomcat6"
	}
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	f, _ := os.Stat(path)
	if !f.IsDir() {
		pathindex := strings.LastIndex(path, PD)
		EXECPATH = path[:pathindex+1]
	} else {
		EXECPATH = path
	}
}

func Compilecode(compile, compilever, md5path string, message chan string) error {
	message <- "开始编译代码"
	var cmd = []string{"ls -l", "ant"}
	var env []string
	if strings.Contains(compile, "JAVA") && compilever != "" {
		if compile == "" {
			env = []string{}
		} else {
			file, _ := exec.LookPath(os.Args[0])
			lvbupath, _ := filepath.Abs(file)
			pathindex := strings.LastIndex(lvbupath, PD)
			lvbupath = lvbupath[:pathindex+1]
			antbin := lvbupath + "/.profile/apache-ant-1.9.7/bin:"
			antlib := lvbupath + "/.profile/apache-ant-1.9.7/lib:"
			javahome := lvbupath + compilever
			beego.Debug("JAVAHOME=", javahome)
			path := os.Getenv("PATH")
			path = javahome + "/bin" + PathPD + antbin + path
			classpath := javahome + "/lib/dt.jar" + PathPD + javahome + "/lib/tools.jar" + PathPD + antlib
			env = []string{"JAVA_HOME=" + javahome, "PATH=" + path, "CLASSPATH=" + classpath}
			beego.Debug("JAVA_HOME=", javahome)
			beego.Debug("PATH=", path)
			beego.Debug("CLASSPATH=", classpath)
		}
	} else if compile == "PHP" && compilever != "" && compilever != "web" {
		message <- "PHP--> 开始安装依赖(composer install)"
		var composerpackagistcmd string
		if ComposerPackagist == "" {
			composerpackagistcmd = Composerbin + " config github-oauth.github.com \"54f16c9921a3ee75aef9d0b3c83a92c77bb867b4\""
		} else {
			composerpackagistcmd = Composerbin + " config repo.packagist " + Composerbin + " https://packagist.phpcomposer.com"
		}
		cmd = []string{composerpackagistcmd, Composerbin + " install", "ls -lh"}
		path := os.Getenv("PATH")
		path = compilever + PathPD + path
		composerhome := md5path
		env = []string{"PATH=" + path, "COMPOSER_HOME=" + composerhome}
	} else {
		message <- "不需要编译,退出编译状态"
		return nil
	}
	var buf bytes.Buffer
	for _, v := range cmd {
		bash := exec.Command("/bin/bash", "-c", v)
		bash.Dir = md5path
		bash.Env = env
		bash.Stdout = os.Stdout
		bash.Stderr = &buf
		if err := bash.Run(); err != nil {
			return errors.New(err.Error() + buf.String())
		}
		if !bash.ProcessState.Success() {
			return errors.New(bash.ProcessState.String() + buf.String())
		}
	}
	return nil
}

//docker pull dockerrepositry:port/reimage:tag
func PullImage(image string, message chan string) (bool, error) {
	message <- "准备pull image:" + image + " ..."
	var buferr bytes.Buffer
	beego.Debug("准备pull image:", image)
	bash := exec.Command(dockerbin, "pull", image)
	bash.Stderr = &buferr
	//	stdout, err := bash.StdoutPipe()
	//	if err != nil {
	//		beego.Error("StdoutPipeERROR:", err)
	//		return false, err
	//	}
	if starterr := bash.Run(); starterr != nil {
		beego.Error("下载镜像出错:", starterr, buferr.String())
		return false, errors.New(starterr.Error() + ";" + buferr.String())
	}
	//	reader := bufio.NewReader(stdout)
	//	for {
	//		line, err2 := reader.ReadString('\n')
	//		if err2 != nil || io.EOF == err2 {
	//			beego.Debug("err2:", line, err2)
	//			message <- err2.Error()
	//			break
	//		}
	//		message <- line
	//	}
	//	bash.Wait()
	if bash.ProcessState.Success() {
		beego.Debug("pull,完成")
		message <- "下载镜像完成"
		return true, nil
	} else {
		return false, errors.New("PULL image error:" + bash.ProcessState.String())
	}
}

func BuildImage(node *mpro.Node, version []string, md5path []string, insfile string, message chan string) error {
	message <- "BUILDING 镜像..."
	beego.Debug("路径:", EXECPATH)
	pro := node.Pro
	mirr := node.Mir
	port_list := GetPortList(node.Port)
	port := ""
	var buferr bytes.Buffer
	for _, v := range port_list {
		port = v + " " + port + " "
	}

	for _, v := range md5path {
		var paths []string
		if pro.Compile != "" && pro.Compile == "JAVASE" { //删除build/, build.xml, src
			paths = []string{"build", "build.xml", "src/", ".gitignore"}
		} else if pro.Compile != "" && pro.Compile == "JAVAEE" {
			paths = []string{"build", "build.xml", "src/", "WebContent", "config/", ".gitignore"}
		}
		filelist := strings.Split(insfile, "\n")
		paths = append(paths, filelist...)
		for _, path := range paths {
			if path == "" {
				continue
			}
			if path[:1] == "/" {
				path = path[1:]
			}
			if err := os.RemoveAll(v + PD + path); err != nil {
				message <- "删除忽略文件或目录" + v + PD + path + "出错：" + err.Error()
				beego.Error("删除忽略文件或目录"+v+PD+path+"出错：", err)
			}
		}

	}

	beego.Debug("port:", port)

	gitname := Gittoname(pro.Git)
	addprocode := ""
	for _, v := range gitname {
		if pro.Compilever == "web" {
			v = v + PD + "web"
		}
		addprocode = "ADD " + v + " /cihi/run/ \n" +
			addprocode + "\n"
	}
	beego.Debug("gitname:", gitname)
	dockerfile_str := "FROM " + mirr.Hubaddress + "\n" +
		addprocode +
		"ENV LANG=en_US.UTF-8 \n" +
		"ENV TZ=Asia/Shanghai \n" +
		"ENV CONFIGPROTOCOL  " + configprocotol + "\n" + //项目名
		"ENV CONFIGURL  " + configurl + "\n" + //项目名
		"ENV PRONAME  " + pro.Sign + "\n" + //项目名
		"ENV PROVERSION " + VerlisttoString(version) + "\n" + //版本号
		"EXPOSE " + port
	//
	md5pathlen := len(md5path[0])
	buildpath := md5path[0][:md5pathlen-len(gitname[0])]
	dockerfilepath := buildpath + "Dockerfile"
	beego.Debug("Dockerfile", dockerfilepath)
	if ok, err := PathExists(buildpath); ok {
		beego.Debug("目录存在")
	} else {
		beego.Error("目录不存在", err)
	}
	dockerfile, fileerr := os.Create(dockerfilepath)
	beego.Debug("路径:", EXECPATH)
	if fileerr != nil && io.EOF != nil {
		beego.Error("Dockerfile 创建出错!", fileerr)
		return errors.New("Dockerfile 创建出错!" + fileerr.Error())
	}
	_, w_err := dockerfile.WriteString(dockerfile_str)
	if w_err != nil && io.EOF != w_err {
		beego.Error("写DOCKERFILE 出错:", w_err)
		return w_err
	}
	containername := ""
	for _, v := range gitname {
		containername = containername + v
	}
	imagename := dockerregsitry + "/" + containername + ":" + VerlisttoString(version)
	imagename = strings.ToLower(imagename)
	beego.Debug(imagename)
	bash := exec.Command(dockerbin, "build", "-t="+imagename, ".")
	bash.Dir = buildpath
	//	stdout, stdouterr := bash.StdoutPipe()
	//	stderr, stderrerr := bash.StderrPipe()
	//	if stdouterr != nil {
	//		beego.Info("Error:", stdouterr)
	//	}
	//	if stderrerr != nil {
	//		beego.Info("Error:", stderrerr)
	//	}
	//	if starterr := bash.Start(); starterr != nil {
	//		beego.Error("下载镜像出错:", starterr)
	//		return errors.New("下载镜像出错:" + starterr.Error())
	//	}
	//	reader := bufio.NewReader(stdout)
	//	reader_err := bufio.NewReader(stderr)
	//	for {
	//		lineerr, err3 := reader_err.ReadString('\n')
	//		line, err2 := reader.ReadString('\n')
	//		if err2 != nil || err3 != nil || io.EOF == err2 || io.EOF == err3 {
	//			break
	//		}
	//		message <- line + lineerr
	//	}
	bash.Stdout = os.Stdout
	bash.Stderr = &buferr
	if err := bash.Run(); err != nil {
		return errors.New(err.Error() + buferr.String())
	}
	//	bash.Wait()
	if bash.ProcessState.Success() {
		message <- "BUILD 镜像完成"
		return nil
	} else {
		beego.Error("BUILD 镜像出错：", bash.ProcessState.String())
		return errors.New(bash.ProcessState.String())
	}
}

func DelNoneImage(adminurl string, message chan string) error {

	if adminurl == "" { //本机操作
		message <- "本机操作，删除失败镜像"
		var buferr bytes.Buffer
		cmd := "docker rmi `docker images | grep \"<none>\" | awk '{print $3}'`"
		bash := exec.Command("/bin/bash", "-c", cmd)
		bash.Stdout = os.Stdout
		bash.Stderr = &buferr
		if err := bash.Run(); err != nil {
			beego.Info("[警告]删除失败镜像出错:" + err.Error() + buferr.String())
		}
		message <- "删除镜像完成，[信息]+ err.Error() + buferr.String()"
		return nil
	} else { // end 本机操作
		return nil
	}
}

func PushImages(giturl string, version []string, message chan string) error {
	message <- "正在上传镜像..."
	var buferr bytes.Buffer
	gitname := Gittoname(giturl)
	containername := ""
	for _, v := range gitname {
		containername = containername + v
	}
	bash := exec.Command(dockerbin, "push", dockerregsitry+"/"+strings.ToLower(containername)+":"+strings.ToLower(VerlisttoString(version)))
	bash.Stdout = os.Stdout
	bash.Stderr = &buferr
	//	if stdouterr != nil {
	//		beego.Info("Error:", stdouterr)
	//	}
	if starterr := bash.Run(); starterr != nil {
		beego.Error("上传镜像出错:", starterr, buferr.String())
		return errors.New(starterr.Error() + buferr.String())
	}
	//	reader := bufio.NewReader(stdout)
	//	for {
	//		line, err2 := reader.ReadString('\n')
	//		if err2 != nil || io.EOF == err2 {
	//			break
	//		}
	//		message <- line
	//	}
	//	bash.Wait()
	Delimage(dockerregsitry+"/"+strings.ToLower(containername), strings.ToLower(VerlisttoString(version)), message) //这个不做正常运行标识
	if bash.ProcessState.Success() {
		message <- "上传镜像完成"
		return nil
	} else {
		beego.Error("上传 镜像出错：", bash.ProcessState.String())
		return errors.New(bash.ProcessState.String())
	}

}

func Delimage(image, node_ver string, message chan string) error {
	message <- "准备删除缓冲镜像" + image
	bash := exec.Command("bash", "-c", dockerbin+" rmi `"+dockerbin+" images |grep "+image+" | grep "+node_ver+" | awk '{ print $3 }'`")
	beego.Debug("`"+dockerbin, "images", "|grep", image, "|awk", "'{", "print", "$3", "}'`")
	stdout, outerr := bash.StdoutPipe()
	bash.Stderr = os.Stdout
	if outerr != nil {
		beego.Info("Error:", outerr)
	}
	if starterr := bash.Start(); starterr != nil {
		beego.Error("删除镜像出错:", starterr)
		return starterr
	}
	readout := bufio.NewReader(stdout)
	for {
		line, err2 := readout.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}

		message <- line
	}
	bash.Wait()
	if bash.ProcessState.Success() {
		message <- "删除缓冲镜像" + image + "成功"
		return nil
	} else {
		beego.Error("删除缓冲镜像出错", bash.ProcessState.String())
		return errors.New(bash.ProcessState.String())
	}
}

func Gitpull(giturl string, message chan string) error {
	message <- "开始git 操作"
	gitpath := Gittoname(giturl)
	beego.Debug("解析git仓库地址:", gitpath)
	if gitpath == nil {
		return errors.New("git仓库地址无效," + giturl)
	}
	var buferr bytes.Buffer
	var pro_path = make([]string, len(gitpath))
	for k, v := range gitpath {
		pro_path[k] = EXECPATH + "/code/" + v + "/"
	}
	beego.Debug("项目路径", pro_path)
	for k, v_pro_path := range pro_path {
		if is, _ := PathExists(v_pro_path); is {
			message <- "项目存在,准备同步"
			cmd := []string{"git reset --hard master", "git pull origin master", "git fetch --tags"}
			for _, v := range cmd {
				bash := exec.Command("/bin/bash", "-c", v)
				bash.Dir = v_pro_path
				bash.Stderr = &buferr
				if err := bash.Run(); err != nil {
					beego.Error(gitpath, "仓库"+gitpath[k]+"同步出错:", err, buferr.String())
					message <- "仓库" + gitpath[k] + "同步出错,准备重新克隆"
					bash = exec.Command("/bin/bash", "-c", "git clone -b master "+strings.Split(giturl, ",")[k])
					bash.Stderr = &buferr
					bash.Dir = EXECPATH + "/code"
					if err := bash.Run(); err != nil {
						message <- "仓库" + gitpath[k] + "克隆出错:"
						beego.Error("仓库" + gitpath[k] + "克隆出错:" + err.Error() + buferr.String())
						return err
					}
				}
			} // end for
			message <- "仓库" + gitpath[k] + "同步完成"
			continue
		} //end 路径是否存在
		message <- "仓库" + gitpath[k] + "项目不存在,开始克隆"
		beego.Debug("仓库" + gitpath[k] + "开始克隆")
		bash := exec.Command("/bin/bash", "-c", "git clone -b master "+strings.Split(giturl, ",")[k])
		bash.Dir = EXECPATH + "/code"
		bash.Stdout = os.Stdout
		bash.Stderr = &buferr
		if err := bash.Run(); err != nil {
			beego.Debug("仓库" + gitpath[k] + "克隆出错:" + buferr.String())
			return errors.New("仓库" + gitpath[k] + "克隆出错:" + err.Error() + buferr.String())
		}
		beego.Info(bash.ProcessState.String())
		message <- "仓库" + gitpath[k] + "克隆完成"
	}

	return nil
}

//git checkout version
func Gitchecver(giturl string, version []string, message chan string) bool {
	gitpath := Gittoname(giturl)
	if gitpath == nil {
		beego.Error("git仓库地址无效," + giturl)
		return false
	}
	var buferr bytes.Buffer
	var pro_path = make([]string, len(gitpath))
	for k, v := range gitpath {
		pro_path[k] = EXECPATH + "code/" + v + "/"
	}
	for k, v_pro_path := range pro_path {
		var check_ver string
		beego.Debug("k:", k)
		for _, git_check_ver := range version {

			if git_check_ver == "" {
				continue
			}

			gitname_ver := strings.Split(git_check_ver, "_")
			if gitpath[k] == gitname_ver[0] {
				check_ver = gitname_ver[1]
			}
		}
		if is, _ := PathExists(v_pro_path); is {
			bash := exec.Command("/bin/bash", "-c", "git checkout master")
			bash.Dir = v_pro_path
			bash.Stderr = &buferr
			if err := bash.Run(); err != nil {
				beego.Error(pro_path, "切换版本出错:", err)
				message <- "切换版本出错:" + err.Error() + buferr.String() + "命令：git checkout master"
				return false
			}
			bash = exec.Command("/bin/bash", "-c", "git checkout "+check_ver)
			bash.Dir = v_pro_path
			bash.Stderr = &buferr
			if err := bash.Run(); err != nil {
				beego.Error(pro_path, "切换版本出错:", err)
				message <- "切换版本出错:" + err.Error() + buferr.String() + "命令：git checkout " + check_ver
				return false
			}
			beego.Debug("bash.Dir=", bash.Dir)

			continue

		} else {
			message <- "动作：切换项目TAG项目路径不存在"
			beego.Debug("动作：切换项目TAG项目路径不存在")
			continue
		}
	}
	return true

}

func GitTags(giturl string) (map[string][]string, error) {
	beego.Debug("准备执行git tag")
	gitpath := Gittoname(giturl)
	var pro_path = make([]string, len(gitpath))
	for k, v := range gitpath {
		pro_path[k] = EXECPATH + "/code/" + v + "/"
	}
	var tags []string
	var buf bytes.Buffer
	var buferr bytes.Buffer
	message := make(chan string, 1)
	go func() {
		for {
			beego.Info(<-message)
		}

	}()
	if err := Gitpull(giturl, message); err != nil {
		return nil, err
	}
	var pro_gittag = make(map[string][]string)
	for k, v := range pro_path {
		//这里多线程
		if is, _ := PathExists(v); is {
			bash := exec.Command("git", "tag")
			bash.Dir = v
			bash.Stdout = &buf
			bash.Stderr = &buferr
			if err := bash.Run(); err != nil {
				beego.Error(v, "列出版本(master)出错:", err)
				beego.Debug("命令标准错误:", buferr.String())
				return nil, errors.New(err.Error() + buferr.String())
			}

			tags_str := buf.String()
			//			beego.Debug(tags_str)
			tags = strings.Split(tags_str, "\n")
			beego.Debug("命令退出状态:", bash.ProcessState.String())
			sort.Sort(sort.Reverse(sort.StringSlice(tags)))
			beego.Debug("tags:", tags)
			m, ok := pro_gittag[gitpath[k]]
			if !ok {
				m = tags
			} else {
				beego.Warning("同一个项目获取节点版本号时，发现有相同仓库名的节点:", v)
			}
			pro_gittag[gitpath[k]] = m
			buf.Reset()
		}
	}
	return pro_gittag, nil
}

//客户端PULL IMAGE
func Clipullimage(adminurl, giturl, tag string, message chan string) error {
	//	certPath := "conf/cert.pem"
	//	keyPath := "conf/key.pem"
	//	caPath := "conf/ca.pem"
	endpoint := "tcp://" + adminurl
	//	endpoint = "tcp://dockerreg.cihi.cn:2375"
	//	client, err := docker.NewTLSClient(endpoint, certPath, keyPath, caPath)
	client, err := docker.NewClient(endpoint)
	if err != nil {
		beego.Error("连接客户端", adminurl, ",", err)
		return err
	}
	client.SkipServerVersionCheck = true
	gitname := Gittoname(giturl)
	containername := ""
	for _, v := range gitname {
		containername = containername + v
	}
	repository := dockerregsitry + "/" + containername
	beego.Debug("repository:", repository, "tag:", strings.ToLower(tag))
	repository = strings.ToLower(repository)
	var buf bytes.Buffer
	var pullopts docker.PullImageOptions
	pullopts.Repository = repository
	pullopts.Tag = strings.ToLower(tag)
	pullopts.Registry = ""
	pullopts.RawJSONStream = true
	pullopts.OutputStream = &buf
	var auth docker.AuthConfiguration
	auth.Email = "zhaoyc1990@163.com"
	//	contr := make(chan error)
	if err = client.PullImage(pullopts, auth); err != nil {
		beego.Error(" 客户端PULL IMAGE 出错：", err)
		return errors.New(" 客户端PULL IMAGE 出错：" + err.Error())
	} else {
		message <- "客户端PULL IMAGE" + repository + ":" + tag + "完成"
		return nil
	}
	//	go func() {
	//		if err = client.PullImage(pullopts, auth); err != nil {
	//			beego.Error(" 客户端PULL IMAGE 出错：", err)
	//			contr <- err
	//		} else {
	//			contr <- nil
	//		}
	//	}()
	//	var delim byte = '\n'
	//	for {
	//		contron := false
	//		select {
	//		case pullerr := <-contr:
	//			beego.Error(pullerr)
	//			return pullerr

	//		default:
	//			if buf.Len() == 0 {
	//				continue
	//			}
	//			line, err := buf.ReadString(delim)

	//			if err != nil {
	//				if err == io.EOF {
	//					continue
	//				}
	//				beego.Info("for break")
	//				contron = true
	//			}
	//			message <- line
	//			if strings.Contains(line, "Status") {
	//				contron = true
	//			}
	//		}
	//		if contron {
	//			beego.Info("真正break")
	//			break
	//		}

	//	}

	//	beego.Info("结束 循环")
	//	return <-contr
}

////主机正在运行的容器列表
func Clilistcons(adminurl string) ([]docker.APIContainers, error) {
	endpoint := "tcp://" + adminurl
	client, err := docker.NewClient(endpoint)
	client.SkipServerVersionCheck = true
	if err != nil {
		beego.Error("连接客户端", adminurl, ",", err)
		var c []docker.APIContainers
		return c, err
	}
	listconopt := docker.ListContainersOptions{
		All: true,
		//		Filters: map[string][]string{
		//			"status": []string{"running"},
		//		},
	}
	containers, err := client.ListContainers(listconopt)
	if err != nil {
		beego.Error("获取容器列表出错", adminurl, ",", err)
		var c []docker.APIContainers
		return c, err
	}
	return containers, nil
}

//容器信息
func Cliinspectcon(nodes []*mpro.Node) ([]*Event, error) {

	var events []*Event
	errindex := 0
	for _, node := range nodes {
		endpoint := "tcp://" + node.Mac.Adminurl
		client, err := docker.NewClient(endpoint)
		if err != nil {
			beego.Error("连接客户端", node.Mac.Adminurl, ",", err)
			return events, err
		}
		if node.DocId == "" {
			beego.Debug("该节点未初始化,node:", node.Name)
			errindex = errindex + 1 //如果这个错误索引的话，在调用遍历 c 时，会抛空指针异常的
			continue
		}
		container, err := client.InspectContainer(node.DocId)
		var event Event
		if err != nil {
			nosuch := &docker.NoSuchContainer{ID: node.DocId}
			if err.Error() == nosuch.Error() {
				event = Event{
					Type:           EVENT_NODE_STAT,
					Containerid:    node.DocId,
					Containerstats: STATS_EXIST,
				}
				beego.Info("容器不存在 ：", node.DocId)
			} else {
				beego.Error("获取容器列表出错跳过", node.Mac.Adminurl, ",", err)
				errindex = errindex + 1 //如果这个错误索引的话，在调用遍历 c 时，会抛空指针异常的
				continue
			}
		} else {
			var stats STATSTYPE
			if container.State.Restarting {
				stats = STATS_RESTARTING
			} else if container.State.Running {
				stats = STATS_RUNNING
			} else {
				stats = STATS_EXIT
			}
			event = Event{
				Type:           EVENT_NODE_STAT,
				Containerid:    container.ID,
				Containerstats: stats,
			}
		}

		events = append(events, &event)
	}
	if len(events) < 1 {
		return events, errors.New("容器列表长度为零，请通知管理员,这并不是一个BUG")
	}
	return events, nil
}

//客户端创建容器
func Clicreatecon(adminurl, port, ver, giturl, env, sharedpath, dns string) (string, error) {
	//	certPath := "conf/cert.pem"
	//	keyPath := "conf/key.pem"
	//	caPath := "conf/ca.pem"
	//	endpoint := "tcp://" + adminurl
	endpoint := "tcp://" + adminurl
	//	client, err := docker.NewTLSClient(endpoint, certPath, keyPath, caPath)
	client, err := docker.NewClient(endpoint)
	if err != nil {
		beego.Error("连接客户端", adminurl, ",", err)
		return "", err
	}

	gitname := Gittoname(giturl)
	containername := ""
	for _, v := range gitname {
		containername = containername + v
	}
	var exposedports = make(map[docker.Port]struct{})
	for _, v := range GetPortList(port) {
		docport := docker.Port(v + "/tcp")
		exposedports[docport] = struct{}{}
	}
	listshared := strings.Split(sharedpath, "\n")
	var binds []string
	beego.Debug("volum:", listshared)
	for _, v := range listshared {
		if v == "" {
			continue
		}
		if v[:1] != "/" { //如果不是绝对路径,就加上"/cihi/run".本项目团队每个镜像的工作目录都在这里
			v = containerdir + "/" + v
		}
		binds = append(binds, dockervolumepath+"/"+containername+v+":"+v)
	}
	beego.Debug("Binds:", binds, "len:", len(binds))
	image := dockerregsitry + "/" + containername + ":" + ver
	image = strings.ToLower(image)
	beego.Debug("giturl:", giturl)
	beego.Debug("image:", image)
	//	client.SkipServerVersionCheck = true
	// Reading logs from container a84849 and sending them to buf.
	dockerenv := []string{"PROENV=" + env}
	conf := docker.Config{
		AttachStderr: true,
		AttachStdin:  false,
		AttachStdout: true,
		Tty:          true,
		OpenStdin:    true,
		Image:        image,
		Env:          dockerenv,
		ExposedPorts: exposedports,
	}
	var portbinding = make(map[docker.Port][]docker.PortBinding)
	portmap, porterr := Getportmap(port)
	if porterr != nil {
		return "", porterr
	}
	for conport, hostport := range portmap {
		var hostbindports []docker.PortBinding
		for _, port := range hostport {
			bindport := docker.PortBinding{HostPort: port}
			hostbindports = append(hostbindports, bindport)
		}
		portbinding[docker.Port(conport+"/tcp")] = hostbindports
		//		portbinding[docker.Port(conport+"/tcp")] = []docker.PortBinding{
		//			docker.PortBinding{
		//				HostPort: hostport,
		//			},
		//		}
	}
	var hostconfig docker.HostConfig
	var dnslist []string
	dns = strings.Replace(dns, " ", "", -1)
	dns = strings.Replace(dns, "\n", "", -1)
	if dns != "" {
		dnslist = []string{dns}
	}
	hostconfig = docker.HostConfig{
		PortBindings:  portbinding,
		RestartPolicy: docker.AlwaysRestart(),
		Binds:         binds,
		DNS:           dnslist,
	}

	createconopts := docker.CreateContainerOptions{
		Name:       containername,
		Config:     &conf,
		HostConfig: &hostconfig,
	}

	container, err := client.CreateContainer(createconopts)
	if err != nil {
		beego.Error("创建容器失败:", err)
		return "", errors.New("创建容器失败:" + err.Error())
	}
	//	message <- "容器:" + container.ID[:10] + "创建完成"
	return container.ID, nil
}

//客户端开启容器
func Clistartcon(adminurl, conid string) error {
	//	certPath := "conf/cert.pem"
	//	keyPath := "conf/key.pem"
	//	caPath := "conf/ca.pem"
	endpoint := "tcp://" + adminurl
	//	client, err := docker.NewTLSClient(endpoint, certPath, keyPath, caPath)
	client, err := docker.NewClient(endpoint)
	if err != nil {
		beego.Error("连接客户端", adminurl, ",", err)
		return err
	}
	err = client.StartContainer(conid, nil)
	if err != nil {
		beego.Error("容器", conid, "启动失败:", err)
		return err
	} else {
		return nil
	}

}

//客户端停止容器
func Clistopcon(adminurl, conid string) error {
	//	certPath := "conf/cert.pem"
	//	keyPath := "conf/key.pem"
	//	caPath := "conf/ca.pem"
	endpoint := "tcp://" + adminurl
	//	client, err := docker.NewTLSClient(endpoint, certPath, keyPath, caPath)
	client, err := docker.NewClient(endpoint)
	if err != nil {
		beego.Error("连接客户端", adminurl, ",", err)
		return err
	}
	err = client.StopContainer(conid, 3)
	if err != nil {
		beego.Error("容器", conid, "启动失败:", err)
		return err
	} else {
		return nil
	}
}

func Clirestartcon(adminurl, conid string) error {
	//	certPath := "conf/cert.pem"
	//	keyPath := "conf/key.pem"
	//	caPath := "conf/ca.pem"
	endpoint := "tcp://" + adminurl
	//	client, err := docker.NewTLSClient(endpoint, certPath, keyPath, caPath)
	client, err := docker.NewClient(endpoint)
	if err != nil {
		beego.Error("连接客户端", adminurl, ",", err)
		return err
	}
	err = client.RestartContainer(conid, 5)
	if err != nil {
		beego.Error("容器", conid, "启动失败:", err)
		return err
	} else {
		return nil
	}

}

//客户端删除容器
func Clidelcon(adminurl, conid string) error {
	//	certPath := "conf/cert.pem"
	//	keyPath := "conf/key.pem"
	//	caPath := "conf/ca.pem"
	if conid == "" {
		beego.Info("初始化容器,ID为空,return nil")
		return nil
	}
	endpoint := "tcp://" + adminurl
	//	client, err := docker.NewTLSClient(endpoint, certPath, keyPath, caPath)
	client, err := docker.NewClient(endpoint)
	if err != nil {
		beego.Error("连接客户端", adminurl, ",", err)
		return err
	}
	removeopts := docker.RemoveContainerOptions{
		ID:            conid,
		RemoveVolumes: false, //是否删除卷
		Force:         true,  //是否把正在运行的容器，删除
	}
	err = client.RemoveContainer(removeopts)
	if err != nil {
		nosuch := &docker.NoSuchContainer{ID: conid}
		if err.Error() == nosuch.Error() {
			beego.Info("容器Id:", conid, "已手动删除")
			return nil
		}
		beego.Error("容器", conid, "启动失败:", err)
		return err
	} else {
		return nil
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	beego.Debug("error:", err)
	return false, err
}
