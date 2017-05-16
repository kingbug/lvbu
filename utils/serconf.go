package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	mcn "lvbu/models/config"
	"os"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Conf struct {
	Key   string
	Value string
}

type Serconf struct {
	Data []Conf
}

func Serialijson(this map[string]string) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	if data, err := json.Marshal(this); err != nil {
		return &buf, err
	} else if _, err := buf.Write(data); err != nil {
		return &buf, err
	} else {
		return &buf, nil
	}
}

func Serialiproper(this map[string]string) (*bytes.Buffer, error) {
	var data bytes.Buffer
	breaks := "\n"
	for k, v := range this {
		data.WriteString(k + "=" + v + breaks)
	}
	return &data, nil
}

func Serialixml(this *Serconf, breaks string) (*bytes.Buffer, error) {
	var confxml *bytes.Buffer
	return confxml, nil
}

func GetConf(sign, env, ver, filetype, conffilename string, round int, breaks string) (*bytes.Buffer, error) {
	if round != 0 {
		if num, ok := mcn.Checknum[sign+"_"+env+"_"+conffilename]; ok { //如果OK有值
			num = num + mcn.Modifynum[sign+"_"+env+"_"+conffilename]
			mnum := mcn.Modifynum[sign+"_"+env+"_"+conffilename]
			beego.Debug("num:", num)
			mkey := sign + "_" + env + "_" + conffilename
			beego.Debug("mcn.Modifynum["+mkey+"]", mnum)
			if round == num { //配置没改动
				beego.Debug("没改动")
				return nil, nil
			} else { //配置有改动
				beego.Debug("有改动")
			}
			mcn.Checknum[sign+"_"+env+"_"+conffilename] = num
		} else { //第一次
			beego.Debug("第一次")
			mcn.Checknum[sign+"_"+env+"_"+conffilename] = round + 1
			beego.Debug("项目:(标识)"+sign, "第一次获取环境", env, "配置")
		}

		mcn.Modifynum[sign+"_"+env+"_"+conffilename] = 0
	}
	var conffortype *bytes.Buffer
	confidir := "prohisconf"
	pro_name := sign
	if mkerr := os.MkdirAll(confidir+"/"+pro_name, 0666); mkerr != nil {
		beego.Debug("动作:请求配置文件,创建目录出错:", mkerr)
	}
	var filename string
	if conffilename == "" {
		filename = confidir + "/" + pro_name + "/" + env + "_" + ver + ".conf"
	} else {
		filename = confidir + "/" + pro_name + "/" + env + "_" + ver + "_" + filename + ".conf"
	}

	var conffile *os.File
	defer conffile.Close()
	var jsonfile = make(map[string]string)
	var fileerr error
	if conffile, fileerr = os.OpenFile(filename, os.O_RDONLY, 0); fileerr != nil && os.IsNotExist(fileerr) {
		//当文件不存在时，（这里假设项目从头到生命结束，都在本系统生态下），即时，当前版本，需要从数据库读取
		var oldverconf []orm.ParamsList
		env = strings.ToUpper(env)
		if env == "DE" {
			if _, err := new(mcn.Config).Query().Filter("Pro__Sign", pro_name).Filter("Filename", conffilename).Filter("Dtstatus__lt", 2).ValuesList(&oldverconf, "Name", "Dvalue"); err != nil {
				return conffortype, errors.New("获取最新配置时，数据查询失败:" + err.Error())
			}
		} else if env == "QE" {
			if _, err := new(mcn.Config).Query().Filter("Pro__Sign", pro_name).Filter("Filename", conffilename).Filter("Dtstatus__lt", 3).ValuesList(&oldverconf, "Name", "Tvalue"); err != nil {
				return conffortype, errors.New("获取最新配置时，数据查询失败:" + err.Error())
			}
		} else { //"OE"
			if _, err := new(mcn.Config).Query().Filter("Pro__Sign", pro_name).Filter("Filename", conffilename).Filter("Tostatus__lt", 3).ValuesList(&oldverconf, "Name", "Ovalue"); err != nil {
				return conffortype, errors.New("获取最新配置时，数据查询失败:" + err.Error())
			}
		}
		for _, v := range oldverconf {
			jsonfile[fmt.Sprintf("%s", v[0])] = fmt.Sprintf("%s", v[1])
		}
		if round != 0 {
			jsonfile["checknum"] = fmt.Sprintf("%d", mcn.Checknum[sign+"_"+env+"_"+conffilename])
		}

	} else {
		data, rerr := ioutil.ReadAll(conffile)
		if rerr != nil {
			beego.Error("ioutil.ReadAll,err:", rerr)
			return conffortype, rerr
		}
		jserr := json.Unmarshal(data, &jsonfile)
		if jserr != nil {
			beego.Error("Unmarshal error:", jserr)
			return conffortype, jserr
		}
	}
	beego.Debug(&jsonfile)
	// 正式转换格式文件
	switch filetype {
	case "json":
		return Serialijson(jsonfile)
	case "properties":
		return Serialiproper(jsonfile)
	default:
		return conffortype, nil
	}
}

func Makejsonconf(sign, env, ver string, conf map[string]string) error {

	confidir := "prohisconf"
	b, berr := json.Marshal(conf)
	if berr != nil {
		return errors.New("序列化json错误," + berr.Error())
	}
	path := confidir + "/" + sign
	if ex, _ := PathExists(path); !ex {
		beego.Debug("历史配置目录不存在")
		if err := os.MkdirAll(path, 0755); err != nil {
			return errors.New("项目配置目录不存在,尝试创建时出错:" + err.Error())
		}
	} else {
		beego.Debug("历史配置目录存在")
	}
	f, ferr := os.Create(confidir + "/" + sign + "/" + env + "_" + ver + ".conf")
	if ferr != nil {
		return errors.New("创建文件错误," + ferr.Error())
	}
	defer f.Close()
	if _, err := f.Write(b); err != nil {
		return errors.New("写文件错误," + err.Error())
	}
	f.Sync()
	f.Close()
	beego.Info("历史配置文件生成成功,文件名：" + confidir + "/" + sign + "/" + env + "_" + ver + ".conf")
	return nil
}
