package machine

import (
	"fmt"
	ctl "lvbu/controllers"
	men "lvbu/models/env"
	mac "lvbu/models/machine"
	mper "lvbu/models/permission"
	"strings"

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
	if c.Ctx.Request.Method == "GET" {
		c.TplName = "machine/machine_add.tpl"
	} else if c.Ctx.Request.Method == "POST" {
		name := c.GetString("name")
		ipaddr1 := c.GetString("ipaddr1") //
		ipaddr2 := c.GetString("ipaddr2") //内， 外网两选 一
		port, err := c.GetInt("port")
		interf := c.GetString("interface")
		env_id, _ := c.GetInt("env")
		content := c.GetString("context") //可空
		var adminurl string
		contrl := true //最后验证必填项是否留空
		if err != nil {
			beego.Debug("动作：添加主机,端口填写错误")
			c.Data["porterr"] = "必须为数字"
			contrl = false
		} else {
			if interf == "1" && ipaddr1 != "" {
				adminurl = ipaddr1 + ":" + fmt.Sprintf("%d", port)
			} else if ipaddr2 != "" && interf == "2" {
				adminurl = ipaddr2 + ":" + fmt.Sprintf("%d", port)
			} else {
				c.Data["ipaddr2err"] = "所选管理接口和所填(内，外网)接口,不一致"
			}
		}
		beego.Debug("admin:", adminurl)
		if name == "" {
			c.Data["nameerr"] = "不能为空"
			contrl = false
		}
		if env_id == 0 {
			c.Data["enverr"] = "环境必选"
			contrl = false
		}

		var host mac.Machine
		env := men.Env{
			Id: uint(env_id),
		}
		host.Name = name
		host.Ipaddr1 = ipaddr1
		host.Ipaddr2 = ipaddr2
		host.Env = &env
		host.Adminurl = adminurl
		host.Content = content
		if contrl { //验证通过
			if err = host.Insert(); err != nil {
				beego.Debug("动作:添加主机,数据库出错:", err)
			}
			c.Redirect("/maclist", 302)
		} else {
			//记录
			c.Data["mac"] = &host
			c.TplName = "machine/machine_add.tpl"
		}

	}

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
		mid, _ := c.GetInt(":id")
		fmt.Println(mid)
		var machine mac.Machine
		machine.Id = uint(mid)
		machine.Read()
		fmt.Println(machine)
		c.Data["mac"] = machine
		c.TplName = "machine/machine_edit.tpl"
		//c.Redirect("/maclist", 302)
	} else {
		id, err := c.GetInt("id")
		name := c.GetString("name")
		ipaddr1 := c.GetString("ipaddr1")
		ipaddr2 := c.GetString("ipaddr2")
		adminurl := c.GetString("adminurl")
		content := c.GetString("content")
		contr := true
		var message string
		if name == "" {
			message = "1.主机名不能为空\n"
			contr = false
		}
		if ipaddr1 == "" && ipaddr2 == "" || adminurl == "" {
			c.Data["message"] = message + "2.管理地址为空或内外网地址匀为空"
			if err != nil || id == 0 {
				beego.Info("动作:修改主机信息，id解析出错:", err)
				c.Redirect("/maclist", 302)
			} else {
				mach := mac.Machine{Id: uint(id)}
				mach.Read()
				c.Data["mac"] = &mach
				c.TplName = "machine/machine_edit.tpl"
				return
			}
		}
		if ipaddr1 == "" {
			if contr := strings.Contains(adminurl, ipaddr2); !contr {
				c.Data["message"] = "管理地址必须使用内网IP或外网IP其一"
			}
		}
		if ipaddr2 == "" {
			if !strings.Contains(adminurl, ipaddr1) {
				contr = false
				c.Data["message"] = "管理地址必须使用内网IP或外网IP其一"
			}
		}
		if !strings.Contains(adminurl, ipaddr1) || !strings.Contains(adminurl, ipaddr2) {
			contr = false
			c.Data["message"] = "管理地址必须使用内网IP或外网IP其一"
		}
		if contr {
			mach := mac.Machine{Id: uint(id)}
			mach.Read()
			mach.Name = name
			mach.Adminurl = adminurl
			mach.Ipaddr1 = ipaddr1
			mach.Ipaddr2 = ipaddr2
			mach.Content = content
			if err = mach.Update(); err != nil {
				beego.Error("动作:数据库操作,修改主机信息出错:", err)
			}
			//添加操作记录
			c.Redirect("/maclist", 302)
		} else {
			mach := mac.Machine{Id: uint(id)}
			mach.Read()
			c.Data["mac"] = &mach
			c.TplName = "machine/machine_edit.tpl"
		}
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

//ajax 请求
func (c *MacController) Del() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isperitem("macd", uid) { //主机[管理|查看](macs)
		beego.Debug("动作:请求主机列表,权限验证失败")
		c.Abort("503")
	}
	mac_id_tmp, err := c.GetInt("id")
	if err != nil {
		beego.Debug("非法提交")
		c.Abort("503")
	}
	mac_id := uint(mac_id_tmp)
	var host mac.Machine
	host.Id = mac_id
	var sign string
	if sign, err = host.Getenvsign(); err != nil {
		c.Abort("500")
	}
	if !mper.Isuserper(sign, uid) {
		c.Abort("503")
	}
	if err = host.Delete(); err != nil {
		beego.Debug("主机信息删除失败:", err)
		c.Abort("503")
	}
	//成功返回
	c.Data["json"] = "success"
	c.ServeJSON()
	return

}
