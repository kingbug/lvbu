package sys

import (
	ctl "lvbu/controllers"
	muser "lvbu/models/user"

	"github.com/astaxie/beego"
)

type PosController struct {
	ctl.BaseController
}

func (c *PosController) List() {
	c.Data["uid"] = c.GetSession("uid")
	c.TplName = "sys/sys_about.tpl"
}

func (c *PosController) Jqupdatepos() {
	posid, err := c.GetInt("posid")
	if err == nil {
		var pos muser.Position
		beego.Debug("jq请求修改镜像Id:", posid)
		pos.Id = uint(posid)
		err = pos.Read()
		//err := mirs.Query().Filter("Id", mirid).One(&mirs)
		if err != nil {
			beego.Error("查询镜像时出错:", err.Error())
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
