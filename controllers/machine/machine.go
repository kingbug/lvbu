package machine

import (
	"fmt"
	ctl "lvbu/controllers"
	mac "lvbu/models/machine"
	mper "lvbu/models/permission"

	"github.com/astaxie/beego"
)

type MacController struct {
	ctl.BaseController
}

func (c *MacController) Add() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isperitem("maca", uid) { //主机添加(maca)
		beego.Debug("动作:请求添加主机,权限验证失败")
		c.Abort("503")
	}
	c.TplName = "machine/machine_add.tpl"
}
func (c *MacController) Edit() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isperitem("mace", uid) { //主机编辑(mace)
		beego.Debug("动作:请求修改主机,权限验证失败")
		c.Abort("503")
	}
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
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isperitem("macs", uid) { //主机[管理|查看](macs)
		beego.Debug("动作:请求主机列表,权限验证失败")
		c.Abort("503")
	}
	var macd, macq, maco []mac.Machine
	new(mac.Machine).Query().Filter("env_id", 1).All(&macd)
	new(mac.Machine).Query().Filter("env_id", 2).All(&macq)
	new(mac.Machine).Query().Filter("env_id", 3).All(&maco)
	c.Data["macd"] = macd
	c.Data["macq"] = macq
	c.Data["maco"] = maco
	c.TplName = "machine/machine_list.tpl"
}
