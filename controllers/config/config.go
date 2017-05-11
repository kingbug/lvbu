package config

import (
	ctl "lvbu/controllers"
	mcn "lvbu/models/config"
	mper "lvbu/models/permission"
	mpro "lvbu/models/project"
	"lvbu/utils"
	"strings"

	"github.com/astaxie/beego"
)

type ConController struct {
	ctl.BaseController
}

func (c *ConController) List() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isperitem("cons", uid) { //配置[管理|查看](cons)
		beego.Debug("动作:请求配置列表,权限验证失败")
		c.Abort("503")
	}
	pro_id, _ := c.GetInt(":proid")
	pro := mpro.Project{Id: uint(pro_id)}
	pro_filename := c.GetString("filename")
	if proerr := pro.Read(); proerr != nil {
		beego.Error("动作：查询项目信息，配置请求项目信息出错:", proerr)
	}
	var conf []*mcn.Config
	if pro_filename == "" {
		if row, err := new(mcn.Config).Query().Filter("Pro__Id", pro_id).Filter("Filename", "").All(&conf); err != nil {
			if row == 0 {
				beego.Info("项目{Id:", pro_id, ", Name:", pro.Name, "},查询配置项为零")
			} else {
				beego.Error("动作：查询配置列表,数据库出错:", err)
			}
		}
	} else {
		if mpro.IsExistName(uint(pro_id), pro_filename) {
			// 取出固定文件名的配置项
			conf = mcn.GetConfforName(uint(pro_id), pro_filename)
		} else {
			c.Data["message"] = map[string]interface{}{"message": "error", "content": "没有该配置文件名!", "type": 2}
		}
	}
	c.Data["PRO_FILENAME"] = pro_filename
	c.Data["pro"] = &pro
	c.Data["conf"] = &conf
	c.TplName = "config/config_list.tpl"
}

func (c *ConController) Add() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isperitem("cona", uid) && !mper.Isuserper("DE", uid) { //配置添加(cona)并且要有开发环境操作权限
		beego.Debug("动作:请求添加配置项,权限验证失败")
		c.Abort("503")
	}
	pro_id, proerr := c.GetInt("pro_id")
	filename := c.GetString("filename")
	key := c.GetString("key")
	value := c.GetString("value")
	description := c.GetString("description")
	sign := c.GetString("sign")
	if proerr != nil || pro_id == 0 {
		beego.Error("动作：添加配置项,项目Id获取出错:", proerr)
	}
	//	sign := c.GetString("sign")
	//	env_id := men.Getenvid("sign")
	pro := mpro.Project{Id: uint(pro_id)}
	if key == "" {
		c.Data["json"] = map[string]interface{}{"message": "error", "content": "key 不能为空", "type": 2}
		c.ServeJSON()
		return
	}
	if new(mcn.Config).Query().Filter("Pro__Id", pro_id).Filter("Name", key).Filter("Filename", filename).Exist() {
		//如果同样的项目中已有该key := c.GetString("key"),返回警告
		c.Data["json"] = map[string]interface{}{"message": "error", "content": "已有相同的KEY" + key, "type": 2}
		c.ServeJSON()
		return
	}
	var conf mcn.Config
	if sign != "de" {
		c.Data["json"] = map[string]interface{}{"message": "error", "content": "非法操作", "type": 3, "confid": conf.Id}
		c.ServeJSON()
		return
	}
	//验证该项目是否有此文件名
	if !mpro.IsExistName(uint(pro_id), filename) {
		c.Data["json"] = map[string]interface{}{"message": "error", "content": "非法操作,没有此配置文件", "type": 3}
		c.ServeJSON()
		return
	}
	conf = mcn.Config{
		Name:        key,
		Dvalue:      value,
		Pro:         &pro,
		Dtstatus:    1,
		Description: description,
		Filename:    filename,
	}
	if conferr := conf.Insert(); conferr != nil {
		beego.Error("动作：添加配置项,数据库出错:", conferr)
		c.Data["json"] = map[string]interface{}{"message": "error", "content": conferr.Error(), "type": 3}
		c.ServeJSON()
		return
	} else {
		pro.Read()
		mcn.AutoChecknum(pro.Sign, strings.ToUpper(sign), filename)
		//操作记录
		c.Data["json"] = map[string]interface{}{"message": "success", "content": "添加成功", "confid": conf.Id, "type": 1}
		c.ServeJSON()
		return
	}

}

func (c *ConController) Edit() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isperitem("cone", uid) { //配置添加(cona)
		beego.Debug("动作:请求修改配置项,权限验证失败")
		c.Abort("503")
	}
	conf_id, conferr := c.GetInt("conf_id")
	value := c.GetString("value")
	sign := c.GetString("sign")
	//环境权限验证
	if !mper.Isuserper(strings.ToUpper(sign), uid) {
		beego.Debug("动作:请求修改配置项,环境权限(" + sign + ")验证失败")
		c.Abort("503")
	}

	if conferr != nil || conf_id == 0 {
		beego.Error("动作：修改配置项,项目Id获取出错:", conferr)
	}
	//	sign := c.GetString("sign")
	//	env_id := men.Getenvid("sign")
	conf := mcn.Config{Id: uint(conf_id)}

	if confreaderr := conf.Read(); confreaderr != nil {
		beego.Error("动作：修改配置项,数据库出错:", confreaderr.Error())
	}
	if sign == "de" {
		conf.Dvalue = value
		conf.Dtstatus = 1
	} else if sign == "qe" {
		conf.Tvalue = value
		conf.Dtstatus = 0
		conf.Tostatus = 1

	} else if sign == "oe" {
		conf.Ovalue = value
		conf.Tostatus = 0
	} else {
		c.Data["json"] = map[string]interface{}{"message": "success", "content": "非法操作", "type": 3}
		beego.Info("越级修改:", conf.Name, "失败")
		c.ServeJSON()
		return
	}
	if conferr := conf.Update(); conferr != nil {
		beego.Error("动作：修改配置项,数据库出错:", conferr)
		c.Data["json"] = map[string]interface{}{"message": "success", "content": conferr.Error(), "type": 3}
		c.ServeJSON()
		return
	} else {
		conf.Pro.Read()
		mcn.AutoChecknum(conf.Pro.Sign, strings.ToUpper(sign), conf.Filename)
		//操作记录
		c.Data["json"] = map[string]interface{}{"message": "success", "content": "修改成功", "type": 1}
		c.ServeJSON()
		return
	}

}

func (c *ConController) Del() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isperitem("cond", uid) { //配置删除 (cond)
		beego.Debug("动作:请求删除配置项,权限验证失败")
		c.Abort("503")
	}
	conf_id, conferr := c.GetInt("conf_id")
	sign := c.GetString("sign")
	//环境权限验证
	if !mper.Isuserper(strings.ToUpper(sign), uid) {
		beego.Debug("动作:请求修改配置项,环境权限验证失败")
		c.Abort("503")
	}
	if conferr != nil || conf_id == 0 {
		beego.Error("动作：删除配置项,项目Id获取出错:", conferr)
	}
	//	sign := c.GetString("sign")
	//	env_id := men.Getenvid("sign")
	conf := mcn.Config{Id: uint(conf_id)}
	if confreaderr := conf.Read(); confreaderr != nil {
		beego.Error("动作：删除前提取配置项,数据库出错:", confreaderr.Error())
	}
	if sign == "de" {
		conf.Dtstatus = 2 //2 表示只是开发删除（针对，开发不可见) ,配置项删除动作只是更改可见状态，数据并不清空 和删除
	} else if sign == "qe" && conf.Dtstatus == 2 {
		conf.Dtstatus = 3 //3 表示测试删除(开发和测试匀看不到)
	} else if sign == "oe" {
		conf.Tostatus = 3 // 生产环境删除
	}
	if conferr := conf.Update(); conferr != nil {
		beego.Error("动作：开发删除配置项,数据库出错:", conferr)
		c.Data["json"] = map[string]interface{}{"message": "error", "content": conferr.Error(), "type": 2}
		c.ServeJSON()
		return
	} else {
		conf.Pro.Read()
		mcn.AutoChecknum(conf.Pro.Sign, strings.ToUpper(sign), conf.Filename)
		//操作记录
		c.Data["json"] = map[string]interface{}{"message": "success", "content": "删除成功", "type": 1}
		c.ServeJSON()
		return
	}
}

func (c *ConController) Sync() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isperitem("cone", uid) { //配置编辑(cone)
		beego.Debug("动作:请求同步(编辑)配置项,权限验证失败")
		c.Abort("503")
	}
	conf_id, conferr := c.GetInt("conf_id")
	sign := c.GetString("sign")
	//环境权限验证
	if !mper.Isuserper(strings.ToUpper(sign), uid) {
		beego.Debug("动作:请求同步配置项,环境权限(" + sign + ")验证失败")
		c.Abort("503")
	}
	if conferr != nil || conf_id == 0 {
		beego.Error("动作：同步配置项,项目Id获取出错:", conferr)
	}
	//	sign := c.GetString("sign")
	//	env_id := men.Getenvid("sign")
	conf := mcn.Config{Id: uint(conf_id)}
	if confreaderr := conf.Read(); confreaderr != nil {
		beego.Error("动作：同步配置项,数据库出错:", confreaderr.Error())
	}
	var per string
	var data string
	if sign == "qe" {
		per = "测试"
		if conf.Dtstatus == 2 { //2为开发已删除
			conf.Dtstatus = 3 //3为测试删除状态
			conf.Tostatus = 1 //告诉生产环境，有改动
		} else if conf.Dtstatus == 1 {
			conf.Tvalue = conf.Dvalue
			conf.Dtstatus = 0 //改为无状态
			conf.Tostatus = 1
			data = conf.Dvalue
		}

	} else if sign == "oe" {
		per = "运维"
		if conf.Dtstatus == 3 {
			conf.Tostatus = 3 //线上环境删除
		} else if conf.Tostatus == 1 {
			conf.Ovalue = conf.Tvalue
			conf.Tostatus = 0
			data = conf.Tvalue
		}

	} else {
		c.Data["json"] = map[string]interface{}{"message": "success", "content": "不能越级同步", "type": 2}
		c.ServeJSON()
		return
	}
	if conferr := conf.Update(); conferr != nil {
		beego.Error("动作："+per+"同步配置项,数据库出错:", conferr)
		c.Data["json"] = map[string]interface{}{"message": "success", "content": conferr.Error(), "type": 2}
		c.ServeJSON()
		return
	} else {
		conf.Pro.Read()
		mcn.AutoChecknum(conf.Pro.Sign, strings.ToUpper(sign), conf.Filename)
		//操作记录
		c.Data["json"] = map[string]interface{}{"message": "success", "content": "同步成功", "type": 1, "data": data}
		c.ServeJSON()
		return
	}
}

func (c *ConController) Ignore() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isperitem("cone", uid) { //配置编辑 (cone)
		beego.Debug("动作:请求忽略配置项,权限验证失败")
		c.Abort("503")
	}
	conf_id, conferr := c.GetInt("conf_id")
	sign := c.GetString("sign")
	//环境权限验证
	if !mper.Isuserper(strings.ToUpper(sign), uid) {
		beego.Debug("动作:请求忽略配置项,环境权限验证失败")
		c.Abort("503")
	}
	if conferr != nil || conf_id == 0 {
		beego.Error("动作：忽略配置项,项目Id获取出错:", conferr)
	}
	//	sign := c.GetString("sign")
	//	env_id := men.Getenvid("sign")
	conf := mcn.Config{Id: uint(conf_id)}
	if confreaderr := conf.Read(); confreaderr != nil {
		beego.Error("动作：忽略前提取配置项,数据库出错:", confreaderr.Error())
	}
	if sign == "qe" && conf.Dtstatus == 1 {
		conf.Dtstatus = 0
		conf.Tostatus = 1
	} else if sign == "oe" {
		conf.Tostatus = 0 // 生产环境删除
	}
	if conferr := conf.Update(); conferr != nil {
		beego.Error("动作：开发删除配置项,数据库出错:", conferr)
		c.Data["json"] = map[string]interface{}{"message": "error", "content": conferr.Error(), "type": 2}
		c.ServeJSON()
		return
	} else {
		//操作记录
		c.Data["json"] = map[string]interface{}{"message": "success", "content": "已忽略", "type": 1}
		c.ServeJSON()
		return
	}
}

func (c *ConController) Download() {
	pro := c.GetString("pro") //PRONAME
	filetype := c.GetString("filetype")
	env := c.GetString("env")               //
	version := c.GetString("version")       //PROVERSION
	conffilename := c.GetString("filename") //可为空，空时为默认文件
	round, _ := c.GetInt("round")           //"1"为轮循，空为正常获取
	line := c.GetString("line")             //空
	if pro != "" && filetype != "" && env != "" && version != "" {
		filename := "prohisconf/" + pro + "_" + env + "_" + version + "_" + filetype + "_" + conffilename + ".conf"
		file, err := utils.GetConf(pro, env, version, filetype, conffilename, round, line)
		if err != nil {
			c.Data["data"] = err.Error()
			beego.Info("配置文件下载出错:", err)
		} else {
			if file != nil {
				c.Data["data"] = string(file.Bytes()[:])
			} else {
				c.Data["data"] = ""
				beego.Warning("配置文件为空")
			}
			beego.Info("配置文件下载:" + filename)
		}

	} else {
		beego.Info("配置文件下载,缺少参数")
	}
	c.TplName = "common/blank.tpl"
}
