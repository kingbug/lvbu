package sys

import (
	ctl "lvbu/controllers"
)

type SysController struct {
	ctl.BaseController
}

func (c *SysController) About() {
	c.Data["uid"] = c.GetSession("uid")
	c.TplName = "sys/sys_about.tpl"
}
