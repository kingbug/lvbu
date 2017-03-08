package utils

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	mpro "lvbu/models/project"
	"os"
	"os/exec"
	"sort"
	"strings"

	"github.com/astaxie/beego"
	"github.com/fsouza/go-dockerclient"
)

var (
	dockerbin      string
	dockerregsitry string
)

func InitDocker() {
	dockerbin = beego.AppConfig.String("Dockerbin")
	dockerregsitry = beego.AppConfig.String("Dockerregsitry")
	if dockerbin == "" {
		dockerbin = "docker"
	}
	if dockerregsitry == "" {
		beego.Error("app.conf -> key:Dockerregsitry is nil")
	}
	bash := exec.Command(dockerbin)
	if err := bash.Run(); err != nil {
		beego.Error("docker deam not found. 查看app.conf Dockerbin是否设置正确")
	}
}

//docker pull dockerrepositry:port/reimage:tag
func PullImage(image string, message chan string) (bool, error) {
	message <- "准备pull image:" + image
	beego.Debug("准备pull image:", image)
	bash := exec.Command(dockerbin, "pull", image)
	stdout, err := bash.StdoutPipe()
	if err != nil {
		beego.Error("StdoutPipeERROR:", err)
		return false, err
	}
	if starterr := bash.Start(); starterr != nil {
		beego.Error("下载镜像出错:", starterr)
		return false, err
	}
	reader := bufio.NewReader(stdout)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			beego.Debug("err2:", line, err2)
			message <- err2.Error()
			break
		}
		message <- line
	}
	bash.Wait()
	if bash.ProcessState.Success() {
		beego.Debug("pull,完成")
		return true, nil
	} else {
		return false, errors.New("PULL image error:" + bash.ProcessState.String())
	}
}

func BuildImage(node *mpro.Node, ver string, message chan string) error {

	pro := node.Pro
	mirr := node.Mir
	//git 切换
	//...end git
	port_list := GetPortList(node.Port)
	port := ""
	for _, v := range port_list {
		port = v + " " + port + " "
	}
	//删除项目忽略文件
	if node.Pro.Insfile != "" {
		if Gitchecver(node.Mac.Adminurl, ver, node.Pro.Insfile, message) != true {
			return errors.New("删除项目忽略文件失败")
		}
	}
	beego.Debug("port:", port)
	pro_path_name := Gittoname(pro.Git)
	pro_path := "code/"
	dockerfile_str := "FROM " + mirr.Hubaddress + "\n" +
		"ADD " + Gittoname(pro.Git) + " /cihi/run/  \n" +
		"EXPOSE " + port
	//此处要有锁
	dockerfile, fileerr := os.Create(pro_path + "/Dockerfile")
	if fileerr != nil && io.EOF != nil {
		beego.Error("Dockerfile 创建出错!", fileerr)
		return errors.New("Dockerfile 创建出错!" + fileerr.Error())
	}
	_, w_err := dockerfile.WriteString(dockerfile_str)
	if w_err != nil && io.EOF != w_err {
		beego.Error("写DOCKERFILE 出错:", w_err)
		return w_err
	}
	imagename := dockerregsitry + "/" + pro_path_name + ":" + node.CurVer
	beego.Debug(imagename)
	bash := exec.Command(dockerbin, "build", "-t="+imagename, ".")
	bash.Dir = pro_path
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
	bash.Run()
	//	bash.Wait()
	if bash.ProcessState.Success() {
		return nil
	} else {
		beego.Error("BUILD 镜像出错：", bash.ProcessState.String())
		return errors.New(bash.ProcessState.String())
	}
}

func PushImages(image, node_ver string, message chan string) error {
	bash := exec.Command(dockerbin, "push", dockerregsitry+"/"+Gittoname(image)+":"+node_ver)
	stdout, stdouterr := bash.StdoutPipe()
	if stdouterr != nil {
		beego.Info("Error:", stdouterr)
	}
	if starterr := bash.Start(); starterr != nil {
		beego.Error("上传镜像出错:", starterr)
		return starterr
	}
	reader := bufio.NewReader(stdout)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		message <- line
	}
	bash.Wait()
	Delimage(dockerregsitry+"/"+Gittoname(image), node_ver, message) //这个不做正常运行标识
	if bash.ProcessState.Success() {
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
	if gitpath == "" {
		return errors.New("git仓库地址无效," + giturl)
	}
	dir, _ := os.Getwd()
	pro_path := "code/" + gitpath
	if is, _ := PathExists(pro_path); is {
		if err := os.Chdir(pro_path); err != nil {

			beego.Error("打开路径:"+dir+"/"+pro_path, err)
			return err
		}
		bash := exec.Command("git", "pull", "origin", "master")
		if err := bash.Run(); err != nil {
			beego.Error(gitpath, "同步出错:", err)
			return err
		} else {
			message <- "同步完成"
			return nil
		}
	}

	bash := exec.Command("git", "clone", giturl)
	bash.Dir = "code/"
	stdout, stderr := bash.StdoutPipe()
	if stderr != nil {
		beego.Info("Error:", stderr)
	}
	if starterr := bash.Start(); starterr != nil {
		beego.Error("克隆仓库出错:", starterr)
		return starterr
	}
	reader := bufio.NewReader(stdout)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			beego.Error("err2:", err2)
			break
		}
		message <- line
	}
	bash.Wait()
	beego.Info(bash.ProcessState.String())
	message <- "克隆完成"
	return nil
}

//git checkout version
func Gitchecver(giturl, version, insfile string, message chan string) bool {
	gitpath := Gittoname(giturl)
	pro_path := "code/" + gitpath
	if is, _ := PathExists(pro_path); is {
		//切换 工作目录
		if err := os.Chdir(pro_path); err != nil {
			dir, _ := os.Getwd()
			beego.Error("打开路径:"+dir+"/"+pro_path, err)
			return false
		}
		bash := exec.Command("git", "checkout", "master")
		if err := bash.Run(); err != nil {
			beego.Error(gitpath, "切换版本(master)出错:", err)
			return false
		}
		bash = exec.Command("git", "checkout", "version")
		if err := bash.Run(); err != nil {
			beego.Error(gitpath, "切换版本出错:", err)
			return false
		}

		filelist := strings.Split(insfile, "\n")
		for _, v := range filelist {
			if err := os.Remove(v); err != nil {
				beego.Error("删除忽略文件或目录"+v+"出错：", err)
			}

		}
		message <- "success"
		return true

	} else {
		return false
	}
}

func GitTags(giturl string) []string {
	gitpath := Gittoname(giturl)
	pro_path := "code/" + gitpath
	var tags []string
	var buf bytes.Buffer
	if is, _ := PathExists(pro_path); is {
		bash := exec.Command("git", "tag")
		bash.Dir = pro_path
		bash.Stdout = &buf
		if err := bash.Run(); err != nil {
			beego.Error(gitpath, "列出版本(master)出错:", err)
		}
		tags_str := buf.String()
		tags = strings.Split(tags_str, "\n")

	}
	sort.Sort(sort.Reverse(sort.StringSlice(tags)))
	return tags
}

//客户端PULL IMAGE
func Clipullimage(adminurl, repository, tag string, message chan string) error {
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
	repository = dockerregsitry + "/" + Gittoname(repository)
	beego.Debug("repository:", repository, "tag:", tag)
	var buf bytes.Buffer
	var pullopts docker.PullImageOptions
	pullopts.Repository = repository
	pullopts.Tag = tag
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

//主机容器列表信息
func Clilistcons(adminurl string) ([]docker.APIContainers, error) {
	endpoint := "tcp://" + adminurl
	client, err := docker.NewClient(endpoint)
	if err != nil {
		beego.Error("连接客户端", adminurl, ",", err)
		var c []docker.APIContainers
		return c, err
	}
	containers, err := client.ListContainers(docker.ListContainersOptions{})
	if err != nil {
		beego.Error("获取容器列表出错", adminurl, ",", err)
		var c []docker.APIContainers
		return c, err
	}
	return containers, nil
}

//容器信息
func Cliinspectcon(nodes []*mpro.Node) ([]*docker.Container, error) {

	c := make([]*docker.Container, len(nodes))
	errindex := 0
	for k, node := range nodes {
		endpoint := "tcp://" + node.Mac.Adminurl
		client, err := docker.NewClient(endpoint)
		if err != nil {
			beego.Error("连接客户端", node.Mac.Adminurl, ",", err)
			return c, err
		}
		if node.DocId == "" {
			beego.Debug("该节点未初始化,node:", node.Name)
			errindex = errindex + 1 //如果这个错误索引的话，在调用遍历 c 时，会抛空指针异常的
			continue
		}
		container, err := client.InspectContainer(node.DocId)
		if err != nil {
			beego.Error("获取容器列表出错跳过", node.Mac.Adminurl, ",", err)
			errindex = errindex + 1 //如果这个错误索引的话，在调用遍历 c 时，会抛空指针异常的
			continue
		}
		c[k-errindex] = container
	}
	if len(c) < 1 {
		return c, errors.New("容器列表长度为零，请通知管理员,这并不是一个BUG")
	}
	return c, nil
}

//客户端创建容器
func Clicreatecon(adminurl, port, ver, giturl string) (string, error) {
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
	//	abc := docker.PortBinding{
	//		HostPort: "33",
	//	}
	var exposedports = make(map[docker.Port]struct{})
	for _, v := range GetPortList(port) {
		docport := docker.Port(v + "/tcp")
		exposedports[docport] = struct{}{}
	}
	image := dockerregsitry + "/" + Gittoname(giturl) + ":" + ver
	beego.Debug("giturl:", giturl)
	beego.Debug("image:", image)
	//	client.SkipServerVersionCheck = true
	// Reading logs from container a84849 and sending them to buf.
	conf := docker.Config{
		AttachStderr: true,
		AttachStdin:  false,
		AttachStdout: true,
		Tty:          true,
		OpenStdin:    true,
		Image:        image,
		ExposedPorts: exposedports,
	}

	//	conf.ExposedPorts = map[docker.Port]struct{}{
	//		"89/tcp": {},
	//		"88/tcp": {},
	//	}
	var portbinding = make(map[docker.Port][]docker.PortBinding)
	portmap, porterr := Getportmap(port)
	if porterr != nil {
		return "", porterr
	}
	for hostport, conport := range portmap {
		portbinding[docker.Port(conport+"/tcp")] = []docker.PortBinding{
			docker.PortBinding{
				HostPort: hostport,
			},
		}
	}
	//	port := map[docker.Port][]docker.PortBinding{
	//		"89/tcp": {
	//			docker.PortBinding{
	//				HostPort: "99",
	//			},
	//		},
	//		"88/tcp": {
	//			docker.PortBinding{
	//				HostPort: "98",
	//			},
	//		},
	//	}
	hostconfig := docker.HostConfig{
		PortBindings: portbinding,
	}

	createconopts := docker.CreateContainerOptions{
		Name:       Gittoname(giturl),
		Config:     &conf,
		HostConfig: &hostconfig,
	}

	container, err := client.CreateContainer(createconopts)
	if err != nil {
		beego.Error("创建容器失败:", err)
		return "", errors.New("创建容器失败:" + err.Error())
	}
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

//客户端删除容器
func Clidelcon(adminurl, conid string) error {
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
	removeopts := docker.RemoveContainerOptions{
		ID:            conid,
		RemoveVolumes: false, //是否删除卷
		Force:         true,  //是否把正在运行的容器，删除
	}
	err = client.RemoveContainer(removeopts)
	if err != nil {
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
	return false, err
}
