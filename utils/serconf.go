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

func Serialijson(this *Serconf, breaks string) (*bytes.Buffer, error) {
	var data bytes.Buffer
	if breaks == "" {
		breaks = "\n"
	}
	data.WriteString("{")
	for _, val := range this.Data {
		data.WriteString("\"" + val.Key + "\":" + "\"" + val.Value + "\"," + breaks)
	}
	data.WriteString("}")
	return &data, nil
}

func Serialiproper(this *Serconf, breaks string) (*bytes.Buffer, error) {
	var data bytes.Buffer
	if breaks == "" {
		breaks = "\n"
	}
	for _, val := range this.Data {
		data.WriteString(val.Key + "=" + val.Value + breaks)
	}
	return &data, nil
}

func Serialixml(this *Serconf, breaks string) (*bytes.Buffer, error) {
	var confxml *bytes.Buffer
	return confxml, nil
}

func GetConf(sign, env, ver, filetype, breaks string) (*bytes.Buffer, error) {
	var conffortype *bytes.Buffer
	confidir := "prohisconf"
	pro_name := sign
	if mkerr := os.MkdirAll(confidir+"/"+pro_name, 0666); mkerr != nil {
		beego.Debug("动作:请求配置文件,创建目录出错:", mkerr)
	}
	filename := confidir + "/" + pro_name + "/" + env + "_" + ver + ".conf"
	var conffile *os.File
	defer conffile.Close()
	var serconf Serconf
	var fileerr error
	if conffile, fileerr = os.OpenFile(filename, os.O_RDONLY, 0); fileerr != nil && os.IsNotExist(fileerr) {
		//当文件不存在时，（这里假设项目从头到生命结束，都在本系统生态下），即时，当前版本，需要从数据库读取
		var oldverconf []orm.ParamsList
		env = strings.ToUpper(env)
		if env == "DE" {
			if _, err := new(mcn.Config).Query().Filter("Pro__Sign", pro_name).ValuesList(&oldverconf, "Name", "Dvalue"); err != nil {
				return conffortype, errors.New("生成老版本配置时，数据查询失败:" + err.Error())
			}
		} else if env == "QE" {
			if _, err := new(mcn.Config).Query().Filter("Pro__Sign", pro_name).ValuesList(&oldverconf, "Name", "Tvalue"); err != nil {
				return conffortype, errors.New("生成老版本配置时，数据查询失败:" + err.Error())
			}
		} else { //"OE"
			if _, err := new(mcn.Config).Query().Filter("Pro__Sign", pro_name).ValuesList(&oldverconf, "Name", "Ovalue"); err != nil {
				return conffortype, errors.New("生成老版本配置时，数据查询失败:" + err.Error())
			}
		}
		var conf []Conf
		for _, v := range oldverconf {
			conf = append(conf, Conf{Key: fmt.Sprintf("%s", v[0]), Value: fmt.Sprintf("%s", v[1])})
		}
		serconf = Serconf{
			Data: conf,
		}

	} else {
		data, rerr := ioutil.ReadAll(conffile)
		if rerr != nil {
			beego.Error("ioutil.ReadAll,err:", rerr)
			return conffortype, rerr
		}
		jserr := json.Unmarshal(data, &serconf)
		if jserr != nil {
			beego.Error("Unmarshal error:", jserr)
			return conffortype, jserr
		}
	}
	beego.Debug(&serconf)
	// 正式转换格式文件
	switch filetype {
	case "json":
		return Serialijson(&serconf, breaks)
	case "properties":
		return Serialiproper(&serconf, breaks)
	default:
		return conffortype, nil
	}
}

func Makejsonconf(sign, env, ver string, conf []Conf) error {
	serconf := Serconf{
		Data: conf,
	}
	confidir := "prohisconf"
	b, berr := json.Marshal(serconf)
	if berr != nil {
		return errors.New("序列化json错误," + berr.Error())
	}
	path := confidir + "/" + sign
	if ex, _ := PathExists(path); !ex {
		beego.Debug("目录不存在")
		if err := os.MkdirAll(path, 0755); err != nil {
			return errors.New("项目配置目录不存在,尝试创建时出错:" + err.Error())
		}
	} else {
		beego.Debug("目录存在")
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
	beego.Debug(serconf)
	return nil
}
