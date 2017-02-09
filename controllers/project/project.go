package project

import (
	ctl "lvbu/controllers"
	mper "lvbu/models/permission"
	mpro "lvbu/models/project"

	"github.com/astaxie/beego"
)

type ProController struct {
	ctl.BaseController
}

func (c *ProController) Add() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isperitem("proa", uid) { //项目添加(proa)
		beego.Debug("动作:请求添加项目,权限验证失败")
		c.Abort("503")
	}
	c.TplName = "project/project_add.tpl"
}
func (c *ProController) Edit() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isperitem("proe", uid) { //项目编辑(proe)
		beego.Debug("动作:请求编辑项目,权限验证失败")
		c.Abort("503")
	}
	c.TplName = "project/project_edit.tpl"
}

func (c *ProController) List() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isperitem("pros", uid) { //项目[管理|查看](pros)
		beego.Debug("动作:请求项目列表,权限验证失败")
		c.Abort("503")
	}
	var pro []mpro.Project
	new(mpro.Project).Query().All(&pro)
	c.Data["pros"] = pro
	c.TplName = "project/project_list.tpl"
}
