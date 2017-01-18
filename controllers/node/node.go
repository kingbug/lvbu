package node

import (
	ctl "lvbu/controllers"
)

type NodeController struct {
	ctl.BaseController
}

func (c *NodeController) Add() {
	c.Data["uid"] = c.GetSession("uid")
	c.TplName = "node/node_add.tpl"
}
func (c *NodeController) Edit() {
	c.Data["uid"] = c.GetSession("uid")
	c.TplName = "node/node_edit.tpl"
}

func (c *NodeController) List() {
	c.Data["uid"] = c.GetSession("uid")
	c.TplName = "node/node_list.tpl"
}
