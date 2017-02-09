package config

import (
	ctl "lvbu/controllers"
	mper "lvbu/models/permission"

	"github.com/astaxie/beego"
)

type ConController struct {
	ctl.BaseController
}

func (c *ConController) List() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isperitem("cons", uid) { //配置[管理|查看](cons)
		beego.Debug("动作:请求配置列表,权限验证失败")
		c.Abort("503")
	}
	c.TplName = "config/config_list.tpl"
}
