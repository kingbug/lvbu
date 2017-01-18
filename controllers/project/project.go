package project

import (
	ctl "lvbu/controllers"
	mpro "lvbu/models/project"
)

type ProController struct {
	ctl.BaseController
}

func (c *ProController) Add() {
	c.Data["uid"] = c.GetSession("uid")
	c.TplName = "project/project_add.tpl"
}
func (c *ProController) Edit() {
	c.Data["uid"] = c.GetSession("uid")
	c.TplName = "project/project_edit.tpl"
}

func (c *ProController) List() {
	c.Data["uid"] = c.GetSession("uid")
	var pro []mpro.Project
	new(mpro.Project).Query().All(&pro)
	c.Data["pros"] = pro
	c.TplName = "project/project_list.tpl"
}
