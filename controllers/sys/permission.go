package sys

import (
	"fmt"
	ctl "lvbu/controllers"
	mper "lvbu/models/permission"
	muser "lvbu/models/user"

	"github.com/astaxie/beego"
)

type PerController struct {
	ctl.BaseController
}

func (c *PerController) List() {
	c.Data["uid"] = c.GetSession("uid")
	pp, _ := c.GetUint16(":id")
	posid := uint(pp)
	c.Data["permenu"] = mper.Getmenu()
	c.Data["posid"] = posid
	c.TplName = "sys/sys_per.tpl"
}
func (c *PerController) Post() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isposition("OS", uid) { //运维经理(OS)
		beego.Debug("动作:请求修改职位权限,权限验证失败")
		c.Abort("503")
	}
	pp, _ := c.GetUint16(":id")
	var per string
	var pos muser.Position
	posid := uint(pp)
	pos.Id = posid
	ss := c.GetStrings("pers")
	fmt.Println(ss)
	for _, k := range ss {
		per += (k + ",")
	}
	pos.Permission = per
	pos.Update("Permission")
	c.Redirect("/permanage/"+fmt.Sprint(posid), 302)
}
