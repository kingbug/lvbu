package env

import (
	ctl "lvbu/controllers"
)

type EnvController struct {
	ctl.BaseController
}

func (c *EnvController) Edit() {
	c.Data["uid"] = c.GetSession("uid")
	c.TplName = "env/env_edit.tpl"
}

func (c *EnvController) List() {
	c.Data["uid"] = c.GetSession("uid")
	c.TplName = "env/env_list.tpl"
}
