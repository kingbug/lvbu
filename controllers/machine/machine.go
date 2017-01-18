package machine

import (
	"fmt"
	ctl "lvbu/controllers"
	mac "lvbu/models/machine"
)

type MacController struct {
	ctl.BaseController
}

func (c *MacController) Add() {
	c.Data["uid"] = c.GetSession("uid")
	c.TplName = "machine/machine_add.tpl"
}
func (c *MacController) Edit() {
	if c.Ctx.Request.Method == "GET" {
		c.Data["uid"] = c.GetSession("uid")
		mid := c.GetString(":id")
		fmt.Println(mid)
		var machine mac.Machine
		new(mac.Machine).Query().Filter("Id", mid).One(&machine)
		fmt.Println(machine)
		c.Data["mac"] = machine
		c.TplName = "machine/machine_edit.tpl"
	} else {
		c.TplName = "machine/machine_edit.tpl"
	}
}

func (c *MacController) List() {
	c.Data["uid"] = c.GetSession("uid")
	var macd, macq, maco []mac.Machine
	new(mac.Machine).Query().Filter("env_id", 1).All(&macd)
	new(mac.Machine).Query().Filter("env_id", 2).All(&macq)
	new(mac.Machine).Query().Filter("env_id", 3).All(&maco)
	c.Data["macd"] = macd
	c.Data["macq"] = macq
	c.Data["maco"] = maco
	c.TplName = "machine/machine_list.tpl"
}
