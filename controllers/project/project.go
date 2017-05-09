package project

import (
	"fmt"
	ctl "lvbu/controllers"
	mper "lvbu/models/permission"
	mpro "lvbu/models/project"
	"lvbu/utils"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

type ProController struct {
	ctl.BaseController
}

func (c *ProController) Add() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid

	if !mper.Isperitem("proa", uid) { //项目添加(proa)和环境判断
		beego.Debug("动作:请求添加项目,权限验证失败")
		c.Abort("503")
	}
	if c.Ctx.Request.Method == "GET" {
		c.TplName = "project/project_add.tpl"
	} else if c.Ctx.Request.Method == "POST" {
		name := c.GetString("name")
		sign := c.GetString("sign")
		git := c.GetStrings("git")
		gituser := c.GetString("gituser") //选 填
		insfile := c.GetString("insfile")
		compile := c.GetString("compile")
		compilever := c.GetString("compilever")
		sharedpath := c.GetString("sharedpath")
		dns := c.GetString("dns")
		beego.Info("gitList:", git)
		contr := true
		if name == "" {
			c.Data["nameerr"] = "项目名称不能为空"
			contr = false
		} else if new(mpro.Project).Query().Filter("Name", name).Exist() {
			c.Data["nameerr"] = "项目名称重复"
			contr = false
		}
		git_contr := false
		for _, v_git := range git {
			if v_git != "" {
				git_contr = true
				break
			}
		}
		if git_contr == false {
			c.Data["giterr"] = "仓库地址不能为空"
			contr = git_contr
		}
		if compile == "" {
			c.Data["compileerr"] = "代码标识不能为空"
			contr = false
		}

		if sign == "" {
			c.Data["signerr"] = "唯一标识不能为空"
			contr = false
		} else if new(mpro.Project).Query().Filter("Sign", sign).Exist() {
			c.Data["signerr"] = "项目标识重复"
			contr = false
		}
		userinfo := strings.SplitN(gituser, ":", 2)
		if gituser != "" && len(userinfo) != 2 || gituser == userinfo[0] {
			c.Data["gitusererr"] = "格式输入有误"
			contr = false
		}
		if dns != "" && len(dns) > 16 {
			c.Data["dnserr"] = "格式输入有误"
			contr = false
		}
		tmp_git := utils.GitlisttoString(git)
		if contr {

			pro := mpro.Project{
				Name:       name,
				Sign:       sign,
				Git:        tmp_git,
				Compile:    compile,
				Compilever: compilever,
				Gituser:    userinfo[0],
				Gitpass:    userinfo[1],
				Sharedpath: sharedpath,
				Insfile:    insfile,
				Dns:        dns,
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
				Name:       name,
				Sign:       sign,
				Git:        tmp_git,
				Compile:    compile,
				Compilever: compilever,
				Gituser:    gituser,
				Insfile:    insfile,
				Sharedpath: sharedpath,
				Dns:        dns,
			}
			c.Data["pro"] = &pro
			c.TplName = "project/project_add.tpl"
		}

	}

}

func (c *ProController) FileAdd() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid

	if !mper.Isperitem("proe", uid) { //项目编辑(proe)
		beego.Debug("动作:请求编辑项目,权限验证失败")
		c.Abort("503")
	}
	if c.Ctx.Request.Method == "POST" {
		id, _ := c.GetInt16("pid")
		filename := c.GetString("filename")
		if filename == "" {
			c.Data["json"] = map[string]interface{}{"message": "error", "content": "添加失败：文件名不能为空!", "type": 2}
			c.ServeJSON()
			return
		}
		if err := mpro.AddConf(uint(id), filename); err != nil {
			c.Data["json"] = map[string]interface{}{"message": "error", "content": "添加失败:" + err.Error(), "type": 3}
			c.ServeJSON()
			return
		} else {
			c.Data["json"] = map[string]interface{}{"message": "success", "content": "添加成功!", "type": 1}
			c.ServeJSON()
			return
		}
	} else {
		c.Data["json"] = map[string]interface{}{"message": "error", "content": "非法提交!", "type": 2}
		c.ServeJSON()
		return
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
		git := c.GetStrings("git")
		gituser := c.GetString("gituser")
		insfile := c.GetString("insfile")
		compile := c.GetString("compile")
		compilever := c.GetString("compilever")
		sharedpath := c.GetString("sharedpath")
		dns := c.GetString("dns")
		contr := true
		if name == "" {
			c.Data["nameerr"] = "项目名称不能为空"
			contr = false
		} else {
			var tmp_project mpro.Project
			new(mpro.Project).Query().Filter("Name", name).One(&tmp_project)
			if tmp_project.Id != uint(id) {
				c.Data["nameerr"] = "项目名称重复"
				contr = false
			}

		}
		git_contr := false
		for _, v_git := range git {
			if v_git != "" {
				git_contr = true
				break
			}
		}
		if git_contr == false {
			c.Data["giterr"] = "仓库地址不能为空"
			contr = git_contr
		}
		if compile == "" {
			c.Data["compileerr"] = "代码标识不能为空"
			contr = false
		}
		//		if compile != "PHP" && compilever == "" {
		//			c.Data["compileerr"] = "代码标识版本不能为空"
		//			contr = false
		//		}
		userinfo := strings.SplitN(gituser, ":", 2)
		if gituser != "" && len(userinfo) != 2 || gituser == userinfo[0] {
			c.Data["gitusererr"] = "格式输入有误"
			contr = false
		}
		if dns != "" && len(dns) > 16 {
			c.Data["dnserr"] = "格式输入有误"
			contr = false
		}
		tmp_git := utils.GitlisttoString(git)
		pro := mpro.Project{Id: uint(id)}
		if readerr := pro.Read(); readerr != nil {
			beego.Error("动作：数据库操作, 查询项目出错:", readerr)
		}
		if contr {
			pro.Name = name
			pro.Git = tmp_git
			pro.Gituser = userinfo[0]
			pro.Gitpass = userinfo[1]
			pro.Insfile = insfile
			pro.Compile = compile
			pro.Compilever = compilever
			pro.Sharedpath = sharedpath
			pro.Dns = dns
			if proerr := pro.Update(); proerr != nil {
				beego.Error("动作：数据库操作, 修改项目出错:", proerr)
			} else {
				//操作记录
			}
			beego.Debug(pro)
			c.Redirect("/prolist", 302)
		} else {
			pro := mpro.Project{
				Name:       name,
				Git:        tmp_git,
				Gituser:    gituser,
				Compile:    compile,
				Compilever: compilever,
				Insfile:    insfile,
				Sharedpath: sharedpath,
				Dns:        dns,
			}
			c.Data["pro"] = &pro
			c.TplName = "project/project_edit.tpl"
		}

	}

}

func (c *ProController) FileEdit() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid

	if !mper.Isperitem("proe", uid) { //项目编辑(proe)
		beego.Debug("动作:请求编辑项目,权限验证失败")
		c.Abort("503")
	}
	if c.Ctx.Request.Method == "POST" {
		id, _ := c.GetInt16("pid")
		oldfile := c.GetString("oldfile")
		filename := c.GetString("newfilename")
		if oldfile == "" {
			c.Data["json"] = map[string]interface{}{"message": "error", "content": "修改失败：默认文件不能修改文件名!", "type": 2}
			c.ServeJSON()
			return
		}
		if filename == "" {
			c.Data["json"] = map[string]interface{}{"message": "error", "content": "修改失败：参数不足!", "type": 2}
			c.ServeJSON()
			return
		}

		if err := mpro.EditConf(uint(id), oldfile, filename); err != nil {
			c.Data["json"] = map[string]interface{}{"message": "error", "content": "修改失败:" + err.Error(), "type": 3}
			c.ServeJSON()
			return
		} else {
			c.Data["json"] = map[string]interface{}{"message": "success", "content": "修改成功!", "type": 1, "filename": filename}
			c.ServeJSON()
			return
		}
	} else {
		c.Data["json"] = map[string]interface{}{"message": "error", "content": "非法提交!", "type": 2}
		c.ServeJSON()
		return
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
		c.Data["json"] = map[string]interface{}{"message": "error", "content": "出错:<" + err.Error() + ">", "type": 2}
		c.ServeJSON()
		return
	} else {
		//操作记录
		beego.Info("删除项目:{Id :", id, "}")
		c.Data["json"] = map[string]interface{}{"message": "success", "content": "删除成功!", "type": 1} //OK
		c.ServeJSON()
		return
	}
}

func (c *ProController) FileDel() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid

	if !mper.Isperitem("proe", uid) { //项目编辑(proe)
		beego.Debug("动作:请求删除配置文件,权限验证失败")
		c.Abort("503")
	}
	if c.Ctx.Request.Method == "POST" {
		id, _ := c.GetInt16("pid")
		filename := c.GetString("filename")
		undo, _ := c.GetBool("undo")
		if filename == "" {
			c.Data["json"] = map[string]interface{}{"message": "error", "content": "删除失败：参数不足!", "type": 2}
			c.ServeJSON()
			return
		}
		if filename == "default" {
			c.Data["json"] = map[string]interface{}{"message": "error", "content": "删除失败：默认文件不能删除!", "type": 2}
			c.ServeJSON()
			return
		}
		if undo == true {
			//一个假消息
			beego.Info("删除配置文件，假消息已发送")
			c.Data["json"] = map[string]interface{}{"message": "success", "content": "删除成功!", "type": 1}
			c.ServeJSON()
			return
		}
		if err := mpro.DelConf(uint(id), filename); err != nil {
			c.Data["json"] = map[string]interface{}{"message": "error", "content": "删除失败:" + err.Error(), "type": 3}
			c.ServeJSON()
			return
		} else {
			beego.Info("删除配置文件，消息已接受")
			c.Data["json"] = map[string]interface{}{"message": "success", "content": "删除成功!", "type": 1}
			c.ServeJSON()
			return
		}
	} else {
		c.Data["json"] = map[string]interface{}{"message": "error", "content": "非法提交!", "type": 2}
		c.ServeJSON()
		return
	}

}

//版本列表
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
		c.Data["json"] = map[string]interface{}{"message": "error", "content": "查询所属项目出错:" + err.Error(), "type": 3}
		c.ServeJSON()
		return
	} else {
		row, err := new(mpro.Node).Query().Filter("Pro__Id", pro.Id).Count()
		if err != nil {
			beego.Error("动作:数据库操作,查询项目节点异常error:", err, "row:", row)
			c.Data["json"] = map[string]interface{}{"message": "error", "content": "查询兄弟节点数量出错:" + err.Error(), "type": 3}
			c.ServeJSON()
			return
		} else if row < 1 {
			c.Data["json"] = map[string]interface{}{"message": "success", "content": []string{"无可用节点"}, "type": 1}
			c.ServeJSON()
			return
		}

	}
	tags, err := utils.GitTags(pro.Git)
	if err != nil {
		beego.Error("error:", err)
		c.Data["json"] = map[string]interface{}{"message": "error", "content": err.Error(), "type": 2}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"message": "success", "content": tags, "type": 1}
	c.ServeJSON()
	return
}

func (c *ProController) Wsproject() {
	ws, err := websocket.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil, 1024, 1024)
	if err != nil {
		beego.Error("websocket 连接创建失败")
	}
	t := time.Now().UnixNano()
	md5id := utils.Md5(fmt.Sprintf("%d", t)) //为保持同一个项目部署的交叉执行,每一个生成不同的目录操作
	md5id = md5id[:10]
	utils.Join(md5id, ws)
	defer utils.Leave(md5id)
	for {
		mt, tmp_ms, _ := ws.ReadMessage()
		beego.Debug("messageType:", mt)
		beego.Debug("message:", string(tmp_ms))
		if mt == -1 {
			beego.Debug("正在尝试关闭当前websocket连接")
			if err := ws.Close(); err != nil {
				beego.Info("关闭websocket连接出错：", err)
			}
			break
		}

	} //ws.ReadMessage()
	beego.Debug("退出项目状态检测")
}
