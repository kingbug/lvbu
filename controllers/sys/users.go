package sys

import (
	ctl "lvbu/controllers"
	muser "lvbu/models/user"
)

type UserController struct {
	ctl.BaseController
}

func (c *UserController) Add() {
	c.Data["uid"] = c.GetSession("uid")
	c.TplName = "sys/user_add.tpl"
}
func (c *UserController) Edit() {
	c.Data["uid"] = c.GetSession("uid")
	c.TplName = "sys/user_edit.tpl"
}
func (c *UserController) List() {
	c.Data["uid"] = c.GetSession("uid")
	var user []muser.User
	var postions []muser.Position
	new(muser.User).Query().All(&user)
	new(muser.Position).Query().All(&postions)
	c.Data["users"] = user
	c.Data["poss"] = postions
	c.TplName = "sys/user_list.tpl"
}
