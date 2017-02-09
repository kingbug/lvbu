package node

import (
	ctl "lvbu/controllers"
	mper "lvbu/models/permission"

	"github.com/astaxie/beego"
)

type NodeController struct {
	ctl.BaseController
}

func (c *NodeController) Add() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isperitem("noda", uid) { //节点添加(noda)
		beego.Debug("动作:请求添加节点,权限验证失败")
		c.Abort("503")
	}
	c.TplName = "node/node_add.tpl"
}
func (c *NodeController) Edit() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isperitem("node", uid) { //节点修改(node)
		beego.Debug("动作:请求修改节点,权限验证失败")
		c.Abort("503")
	}
	c.TplName = "node/node_edit.tpl"
}

func (c *NodeController) List() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isperitem("nods", uid) { //节点[管理|查看](nods)
		beego.Debug("动作:请求查看节点,权限验证失败")
		c.Abort("503")
	}
	c.TplName = "node/node_list.tpl"
}
