package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/astaxie/beego"
)

type Conf struct {
	Key   string
	Value string
}

type Serconf struct {
	Data []Conf
}

func Serialijson(this *Serconf, breaks string) *bytes.Buffer {
	var data bytes.Buffer
	if breaks == "" {
		breaks = "\n"
	}
	data.WriteString("{")
	for _, val := range this.Data {
		data.WriteString("\"" + val.Key + "\":" + "\"" + val.Value + "\"" + breaks)
	}
	data.WriteString("}")
	return &data
}

func Serialiproper(this *Serconf, breaks string) *bytes.Buffer {
	var data bytes.Buffer
	if breaks == "" {
		breaks = "\n"
	}
	for _, val := range this.Data {
		data.WriteString(val.Key + "=" + val.Value + breaks)
	}
	return &data
}

func Serialixml(this *Serconf, breaks string) *bytes.Buffer {
	return nil
}

func GetConf(pro, env, ver, filetype, breaks string) *bytes.Buffer {
	confidir := "prohisconf"
	if mkerr := os.MkdirAll(confidir+"/"+pro, 0666); mkerr != nil {
		beego.Debug("动作:请求配置文件,创建目录出错:", mkerr)
	}
	filename := confidir + "/" + pro + "/" + env + "_" + ver + ".conf"
	var conffile *os.File
	defer conffile.Close()
	var fileerr error
	if conffile, fileerr = os.OpenFile(filename, os.O_RDONLY, 0); fileerr != nil && os.IsNotExist(fileerr) {
		//		if conffile, fileerr = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666); fileerr != nil {
		//			beego.Error("动作:请求配置文件,创建文件出错:", fileerr)
		//		}
		//		conf, err := new(config.Config).GetRead(pro)
		//		if err != nil {
		//			beego.Error("查找真实KEY，VALUE出错:", err)
		//		}
		//		jsons, _ := json.Marshal(conf)
		//		if _, werr := conffile.Write(jsons); werr != nil {
		//			beego.Debug(jsons)
		//			beego.Error("配置文件写入文件错误:", werr)
		//		}
	}
	var conf Serconf
	data, rerr := ioutil.ReadAll(conffile)
	if rerr != nil {
		beego.Error("ioutil.ReadAll,err:", rerr)
	}
	jserr := json.Unmarshal(data, &conf)
	if jserr != nil {
		beego.Error("Unmarshal error:", jserr)
	}

	beego.Debug(&conf)
	// 正式转换格式文件
	switch filetype {
	case "json":
		return Serialijson(&conf, breaks)
	case "properties":
		return Serialiproper(&conf, breaks)
	default:
		return nil
	}
}
