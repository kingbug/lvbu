package sys

import (
	"fmt"
	ctl "lvbu/controllers"
	men "lvbu/models/env"
	mper "lvbu/models/permission"
	muser "lvbu/models/user"
	"lvbu/utils"
	"strconv"

	"github.com/astaxie/beego"
)

type UserController struct {
	ctl.BaseController
}

func (c *UserController) AddGet() {
	c.Data["uid"] = c.GetSession("uid")
	c.TplName = "sys/user_add.tpl"

}
func (c *UserController) AddPost() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isposition("OS", uid) { //只有运维经理，可以 添加，修改任意用户资料
		//		s_url := c.Ctx.Request.Referer()
		beego.Debug("动作:请求添加用户,权限验证失败")
		c.Abort("503")
	}
	var message string
	con := true
	var user muser.User
	if user.UserName = c.GetString("username"); user.UserName == "" {
		c.Data["usernameerr"] = "用户名（登录名）必填"
		con = false
	}
	if user.Passwd = c.GetString("passwd"); user.Passwd == "" {
		c.Data["passwderr"] = "密码必填"
		con = false
	}
	if user.Nick = c.GetString("nick"); user.Nick == "" {
		c.Data["nickerr"] = "姓名必填"
		con = false
	}
	if sex, err := c.GetInt("sex"); err != nil {
		c.Data["sexerr"] = "不要恶搞网页源代码"

		con = false
	} else {
		user.Sex = uint(sex)
	}
	c.Data["usersex"], _ = c.GetInt("sex")
	if user.Phone = c.GetString("phone"); user.Phone == "" {
		c.Data["phoneerr"] = "电话必填"
		con = false
	}
	if user.Email = c.GetString("email"); user.Email == "" {
		c.Data["emailerr"] = "邮箱必填"
		con = false
	}
	var position muser.Position
	id, err := c.GetInt("position")
	if err != nil {
		c.Data["positionerr"] = "职务选择错误,请通知管理员" + err.Error()
		con = false
	} else {
		if id < 1 {
			c.Data["positionerr"] = "职务为空"
			con = false
		}
		position.Id = uint(id)
		beego.Info("添加用户["+user.UserName+"]职务Id:", id)
		user.Position = &position
	}

	permissions := c.GetStrings("permission") //接收页面传过来的数组
	for _, value := range permissions {       //遍历，从数据库查到 env 表中 sign 字段
		var env men.Env
		if id, err := strconv.Atoi(value); err != nil {
			c.Data["enverr"] = "不要恶捣网页源代码哦 + 必选多选"
			beego.Warning("添加用户页面,转换permission to int字段出错：" + err.Error())
			con = false
		} else if id > 0 {
			env.Id = uint(id)
			err := env.Read()
			if err != nil {
				message = "服务出错:" + err.Error() + message
				beego.Error("添加页面, 查询表" + env.TableName() + "出错:" + err.Error())
				con = false
			} else {
				beego.Debug(env)
				if user.Permission == "" {
					user.Permission = env.Sign
				} else {
					user.Permission = user.Permission + "," + env.Sign
				}
			}
		} else {
			c.Data["enverr"] = "环境必选"
			con = false
		}
	}
	beego.Debug(user.Permission)
	c.Data["user"] = &user
	if con != true {
		c.Data["message"] = "添加用户失败,有必填项留空" + message
		beego.Error("添加用户失败,有必填项留空", user)
		c.TplName = "sys/user_add.tpl"
	} else {
		user.Passwd = utils.Md5(user.Passwd)
		err = user.Insert()
		if err != nil {
			c.Data["message"] = "添加用户失败:" + err.Error() + message
			beego.Error("添加用户失败,数据库出错:"+err.Error(), user)
			c.TplName = "sys/user_add.tpl"
		} else {
			beego.Info("添加新用户:", user)
			c.Redirect("usermanager", 302)
		}
	}

}
func (c *UserController) EditGet() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isposition("OS", uid) { //只有运维经理，可以 添加，修改任意用户资料
		//		s_url := c.Ctx.Request.Referer()
		beego.Debug("动作:请求修改用户,权限验证失败")
		c.Abort("503")
	}
	id, _ := c.GetInt(":id")
	var user muser.User
	user.Id = uint(id)
	if err := user.Read(); err != nil {
		c.Data["message"] = "查询数据库出错"
		beego.Error("查询用户Id: ", id, "出错:"+err.Error())
	}
	c.Data["user"] = &user
	c.Data["usersex"] = fmt.Sprintf("%d", user.Sex)
	beego.Debug("修改用户资料Id:", id)
	c.TplName = "sys/user_edit.tpl"
}
func (c *UserController) EditPost() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isposition("OS", uid) { //只有运维经理，可以 添加，修改任意用户资料
		//		s_url := c.Ctx.Request.Referer()
		beego.Debug("动作:修改用户POST,权限验证失败")
		c.Abort("503")
	}
	uuid, _ := c.GetInt(":id")
	uid = uint(uuid)
	var message string
	con := true
	var user muser.User
	if user.UserName = c.GetString("username"); user.UserName == "" {
		c.Data["usernameerr"] = "用户名（登录名）必填"
		con = false
	}
	if user.Passwd = c.GetString("passwd"); user.Passwd == "" {
		c.Data["passwderr"] = "密码必填"
		con = false
	}
	if user.Nick = c.GetString("nick"); user.Nick == "" {
		c.Data["nickerr"] = "姓名必填"
		con = false
	}
	if sex, err := c.GetInt("sex"); err != nil {
		c.Data["sexerr"] = "不要恶搞网页源代码"

		con = false
	} else {
		user.Sex = uint(sex)
		fmt.Println("性别:", user.Sex)
	}
	c.Data["usersex"], _ = c.GetInt("sex")
	if user.Phone = c.GetString("phone"); user.Phone == "" {
		c.Data["phoneerr"] = "电话必填"
		con = false
	}
	if user.Email = c.GetString("email"); user.Email == "" {
		c.Data["emailerr"] = "邮箱必填"
		con = false
	}
	if status, err := c.GetInt("status"); err != nil {
		c.Data["emailerr"] = "不要恶捣网页源代码哦"
		con = false
	} else {
		user.Status = status
	}
	var position muser.Position
	id, err := c.GetInt("position")
	if err != nil {
		c.Data["positionerr"] = "职务选择错误,请通知管理员" + err.Error()
		con = false
	} else {
		if id < 1 {
			c.Data["positionerr"] = "职务为空"
			con = false
		}
		position.Id = uint(id)
		beego.Info("添加用户["+user.UserName+"]职务Id:", id)
		user.Position = &position
	}

	permissions := c.GetStrings("permission") //接收页面传过来的数组
	for _, value := range permissions {       //遍历，从数据库查到 env 表中 sign 字段
		var env men.Env
		if id, err := strconv.Atoi(value); err != nil {
			c.Data["enverr"] = "不要恶捣网页源代码哦 + 必选多选"
			beego.Warning("添加用户页面,转换permission to int字段出错：" + err.Error())
			con = false
		} else if id > 0 {
			env.Id = uint(id)
			err := env.Read()
			if err != nil {
				message = "服务出错:" + err.Error() + message
				beego.Error("添加页面, 查询表" + env.TableName() + "出错:" + err.Error())
				con = false
			} else {
				beego.Debug(env)
				if user.Permission == "" {
					user.Permission = env.Sign
				} else {
					user.Permission = user.Permission + "," + env.Sign
				}
			}
		} else {
			c.Data["enverr"] = "环境必选"
			con = false
		}
	}
	beego.Debug(user.Permission)
	var copyuser muser.User
	copyuser.Id = uid
	copyuser.Read()
	c.Data["user"] = &copyuser
	if con != true {
		c.Data["message"] = "修改用户资料失败,有必填项留空" + message
		beego.Error("修改用户失败,有必填项留空", user)
		c.TplName = "sys/user_edit.tpl"
	} else {
		user.Passwd = utils.Md5(user.Passwd)
		user.Id = uid
		user.Created = copyuser.Created
		err = user.Update()
		if err != nil {
			c.Data["message"] = "修改用户资料户失败:" + err.Error() + message
			beego.Error("修改用户资料失败,数据库出错:"+err.Error(), user)
			c.TplName = "sys/user_edit.tpl"
		} else {
			beego.Info("修改用户资料:", user)
			c.Redirect("/usermanager", 302)
		}
	}
}

func (c *UserController) List() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isposition("OS", uid) { //只有运维经理，可以 添加，修改任意用户资料
		//		s_url := c.Ctx.Request.Referer()
		beego.Debug("动作:请求用户列表,权限验证失败")
		c.Abort("503")
	}
	var user []muser.User
	var postions []muser.Position
	new(muser.User).Query().All(&user)
	new(muser.Position).Query().All(&postions)
	c.Data["users"] = user
	c.Data["poss"] = postions
	c.TplName = "sys/user_list.tpl"
}
