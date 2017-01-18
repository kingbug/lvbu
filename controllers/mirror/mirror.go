package mirror

import (
	"fmt"
	ctl "lvbu/controllers"
	"lvbu/models/mirror"
)

type MirController struct {
	ctl.BaseController
}

func (c *MirController) Gadd() {
	c.Data["uid"] = c.GetSession("uid")
	c.TplName = "mirror/mirrorgroup_add.tpl"
}
func (c *MirController) Add() {
	c.Data["uid"] = c.GetSession("uid")
	c.TplName = "mirror/mirror_add.tpl"
}
func (c *MirController) Edit() {
	c.Data["uid"] = c.GetSession("uid")
	c.TplName = "mirror/mirror_edit.tpl"
}

func (c *MirController) List() {
	c.Data["uid"] = c.GetSession("uid")
	c.TplName = "mirror/mirror_list.tpl"
}
func (c *MirController) Post() {
	c.Data["uid"] = c.GetSession("uid")
	mirgname := c.GetString("mirgname")
	fmt.Println(mirgname)
	var mgoup mirror.MirrorGroup
	mgoup.Name = mirgname
	mgoup.Insert()
	c.TplName = "mirror/mirror_list.tpl"
}
