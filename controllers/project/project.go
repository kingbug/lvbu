package project

import (
	ctl "lvbu/controllers"
	mper "lvbu/models/permission"
	mpro "lvbu/models/project"
	"lvbu/utils"
	"strings"

	"github.com/astaxie/beego"
)

type ProController struct {
	ctl.BaseController
}

func (c *ProController) Add() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isperitem("proa", uid) { //项目添加(proa)
		beego.Debug("动作:请求添加项目,权限验证失败")
		c.Abort("503")
	}
	if c.Ctx.Request.Method == "GET" {
		c.TplName = "project/project_add.tpl"
	} else if c.Ctx.Request.Method == "POST" {
		name := c.GetString("name")
		sign := c.GetString("sign")
		git := c.GetString("git")
		gituser := c.GetString("gituser") //选 填
		insfile := c.GetString("insfile")
		contr := true
		if name == "" {
			c.Data["nameerr"] = "项目名称不能为空"
			contr = false
		} else if new(mpro.Project).Query().Filter("Name", name).Exist() {
			c.Data["nameerr"] = "项目名称重复"
			contr = false
		}
		if git == "" {
			c.Data["giterr"] = "仓库地址不能为空"
			contr = false
		}
		if sign == "" {
			c.Data["signerr"] = "唯一标识不能为空"
			contr = false
		}
		userinfo := strings.SplitN(gituser, ":", 2)
		if gituser != "" && len(userinfo) != 2 || gituser == userinfo[0] {
			c.Data["gitusererr"] = "格式输入有误"
			contr = false
		}

		if contr {
			pro := mpro.Project{
				Name:    name,
				Sign:    sign,
				Git:     git,
				Gituser: userinfo[0],
				Gitpass: userinfo[1],
				Insfile: insfile,
			}

			if proerr := pro.Insert(); proerr != nil {
				beego.Error("动作：数据库操作, 添加项目出错:", proerr)
			} else {
				//操作记录
			}
			beego.Debug(pro)
			c.SetSession("newpro", pro.Id) //跳转项目列表后，突出显示新添加项目
			c.Redirect("/prolist", 302)
		} else {
			pro := mpro.Project{
				Name:    name,
				Sign:    sign,
				Git:     git,
				Gituser: gituser,
				Insfile: insfile,
			}
			c.Data["pro"] = &pro
			c.TplName = "project/project_add.tpl"
		}

	}

}
func (c *ProController) Edit() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isperitem("proe", uid) { //项目编辑(proe)
		beego.Debug("动作:请求编辑项目,权限验证失败")
		c.Abort("503")
	}
	if c.Ctx.Request.Method == "GET" {
		id, _ := c.GetInt16(":id")
		var pro mpro.Project
		pro.Id = uint(id)
		if proerr := pro.Read(); proerr != nil {
			beego.Error("动作：请求编辑项目, 数据库出错:", proerr)
			if strings.Contains(proerr.Error(), "no row found") {
				c.Data["message"] = "不要玩老道，城墙很厚的,随便打洞"
			}
			c.TplName = "project/project_edit.tpl"
		} else {
			pro.Gituser = pro.Gituser + ":" + pro.Gitpass
			c.Data["pro"] = &pro
			beego.Debug(pro)
			c.TplName = "project/project_edit.tpl"
		}

	} else if c.Ctx.Request.Method == "POST" {
		id, iderr := c.GetInt16("id")
		if iderr != nil {
			c.Abort("503") //除非在html源码里改，否则不会出错
		}
		name := c.GetString("name")
		sign := c.GetString("sign")
		git := c.GetString("git")
		gituser := c.GetString("gituser")
		insfile := c.GetString("insfile")
		contr := true
		if name == "" {
			c.Data["nameerr"] = "项目名称不能为空"
			contr = false
		} else {
			var tmp_pro mpro.Project
			new(mpro.Project).Query().Filter("Name", name).One(&tmp_pro)
			if tmp_pro.Id != uint(id) {
				c.Data["nameerr"] = "项目名称重复"
				contr = false
			}
		}
		if git == "" {
			c.Data["giterr"] = "仓库地址不能为空"
			contr = false
		}
		if sign == "" {
			c.Data["signerr"] = "唯一标识不能为空"
			contr = false
		}
		userinfo := strings.SplitN(gituser, ":", 2)
		if gituser != "" && len(userinfo) != 2 || gituser == userinfo[0] {
			c.Data["gitusererr"] = "格式输入有误"
			contr = false
		}
		if contr {
			pro := mpro.Project{Id: uint(id)}
			if readerr := pro.Read(); readerr != nil {
				beego.Error("动作：数据库操作, 查询项目出错:", readerr)
			}
			pro.Name = name
			pro.Sign = sign
			pro.Git = git
			pro.Gituser = userinfo[0]
			pro.Gitpass = userinfo[1]
			pro.Insfile = insfile
			if proerr := pro.Update(); proerr != nil {
				beego.Error("动作：数据库操作, 修改项目出错:", proerr)
			} else {
				//操作记录
			}
			beego.Debug(pro)
			c.Redirect("/prolist", 302)
		} else {
			pro := mpro.Project{
				Name:    name,
				Sign:    sign,
				Git:     git,
				Gituser: gituser,
				Insfile: insfile,
			}
			c.Data["pro"] = &pro
			c.TplName = "project/project_edit.tpl"
		}

	}

}

func (c *ProController) List() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isperitem("pros", uid) { //项目[管理|查看](pros)
		beego.Debug("动作:请求项目列表,权限验证失败")
		c.Abort("503")
	}
	c.Data["newpro"] = c.GetSession("newpro") //新添加项目提示
	c.TplName = "project/project_list.tpl"
}

func (c *ProController) Del() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isperitem("prod", uid) { //项目[管理|查看](pros)
		beego.Debug("动作:请求项目列表,权限验证失败")
		c.Abort("503")
	}
	var pro mpro.Project
	id, _ := c.GetUint16("id")
	pro.Id = uint(id)

	if err := pro.Delete(); err != nil {
		beego.Debug("project:", pro, "err:", err)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	} else {
		//操作记录
		beego.Info("删除项目:{Id :", id, "}")
		c.Data["json"] = "success" //OK
		c.ServeJSON()
		return
	}
}

func (c *ProController) Verlist() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isperitem("pros", uid) { //项目[管理|查看](pros)
		beego.Debug("动作:请求项目列表,权限验证失败")
		c.Abort("503")
	}
	var pro mpro.Project
	id, _ := c.GetUint16("id")
	pro.Id = uint(id)
	if err := pro.Read(); err != nil {
		beego.Error("动作:数据库操作,查询项目出错:", err)
		c.Data["json"] = map[string]interface{}{"message": "error", "data": err.Error()}
		c.ServeJSON()
		return
	}
	tags := utils.GitTags(pro.Git)
	c.Data["json"] = map[string]interface{}{"message": "success", "data": tags}
	c.ServeJSON()
	return
}
