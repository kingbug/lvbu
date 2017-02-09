package sys

import (
	ctl "lvbu/controllers"
	mper "lvbu/models/permission"
	muser "lvbu/models/user"

	"github.com/astaxie/beego"
)

type PosController struct {
	ctl.BaseController
}

func (c *PosController) List() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isposition("OS", uid) { //只有运维经理(OS)，可以 添加，修改职位资料
		//		s_url := c.Ctx.Request.Referer()
		beego.Debug("动作:请求职位列表,权限验证失败")
		c.Abort("503") //因是JS请求并不会跳转，但也不会往下执行
	}
	c.TplName = "user/user_index.tpl"
}

func (c *PosController) Jqupdatepos() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isposition("OS", uid) { //只有运维经理(OS)，可以 添加，修改职位资料
		//		s_url := c.Ctx.Request.Referer()
		beego.Debug("动作:请求修改职位,权限验证失败")
		c.Abort("503") //因是JS请求并不会跳转，但也不会往下执行
	}
	posid, err := c.GetInt("posid")
	if err == nil {
		var pos muser.Position
		beego.Debug("jq请求修改职位Id:", posid)
		pos.Id = uint(posid)
		err = pos.Read()
		//err := mirs.Query().Filter("Id", mirid).One(&mirs)
		if err != nil {
			beego.Error("查询职位时出错:", err.Error())
			c.Data["json"] = "多人操作中，该条信息已删除，请刷新当前页面，已获取最新:" + err.Error()
			c.ServeJSON()
			return
		}

		beego.Debug("JS请求更新前职位信息,", pos)
		pos.Name = c.GetString("posname")
		pos.Sign = c.GetString("possign")
		beego.Debug("JS请求更新职位 ,name:", pos.Name, ", Sign:", pos.Sign)
		pos.Update()
		beego.Debug(pos)
		c.Data["json"] = "success"
		c.ServeJSON()
		return

	} else {
		beego.Error(err.Error())
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
}

func (c *PosController) Jqrmpos() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isposition("OS", uid) { //只有运维经理(OS)，可以 添加，修改职位资料
		//		s_url := c.Ctx.Request.Referer()
		beego.Debug("动作:请求删除职位,权限验证失败")
		c.Abort("503") //因是JS请求并不会跳转，但也不会往下执行
	}
	posid, err := c.GetInt(":id")
	if err == nil {
		var pos muser.Position
		pos.Id = uint(posid)
		beego.Debug("删除职位Id:", posid)
		if err := pos.Delete(); err != nil {
			c.Data["json"] = err.Error()
			c.ServeJSON()
			return
		}
		c.Data["json"] = "success"
		c.ServeJSON()
		return

	} else {
		c.Data["json"] = nil
		c.ServeJSON()
		return
	}
}

func (c *PosController) Jqaddpos() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isposition("OS", uid) { //只有运维经理(OS)，可以 添加，修改任意用户资料
		//		s_url := c.Ctx.Request.Referer()
		beego.Debug("动作:请求添加职位,权限验证失败")
		c.Abort("503") //因是JS请求并不会跳转，但也不会往下执行
	}
	name := c.GetString("name")
	sign := c.GetString("sign")

	if name != "" && sign != "" {
		var pos muser.Position
		pos.Name = name
		pos.Sign = sign
		beego.Debug("添加职位:", pos)
		if err := pos.Insert(); err != nil {
			c.Data["json"] = err.Error()
			c.ServeJSON()
			return
		}
		c.Data["json"] = map[string]interface{}{"status": "success", "id": pos.Id}
		c.ServeJSON()
		return

	} else {
		c.Data["json"] = "职位名或标识不能为空"
		c.ServeJSON()
		return
	}
}
