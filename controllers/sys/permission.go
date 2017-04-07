package sys

import (
	"fmt"
	ctl "lvbu/controllers"
	mper "lvbu/models/permission"
	muser "lvbu/models/user"
	"lvbu/utils"

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
	if err := pos.Update("Permission"); err != nil {
		beego.Debug("更新数据出错", err)
		c.Redirect("/permanage/"+fmt.Sprint(posid), 302)
	}
	if err := utils.ClearAll(); err != nil {
		beego.Debug("清空缓存出错", err)
	}
	c.Redirect("/permanage/"+fmt.Sprint(posid), 302)
}
