package config

import (
	ctl "lvbu/controllers"
)

type ConController struct {
	ctl.BaseController
}

func (c *ConController) List() {
	c.Data["uid"] = c.GetSession("uid")
	c.TplName = "config/config_list.tpl"
}
