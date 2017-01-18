package record

import (
	ctl "lvbu/controllers"
)

type RecController struct {
	ctl.BaseController
}

func (c *RecController) List() {
	c.Data["uid"] = c.GetSession("uid")
	c.TplName = "record/record_list.tpl"
}
