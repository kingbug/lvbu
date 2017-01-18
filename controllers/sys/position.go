package sys

import (
	ctl "lvbu/controllers"
)

type PosController struct {
	ctl.BaseController
}

func (c *PosController) List() {
	c.Data["uid"] = c.GetSession("uid")
	c.TplName = "sys/sys_about.tpl"
}
