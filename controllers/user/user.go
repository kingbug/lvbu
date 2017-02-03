package user

import (
	//"fmt"
	ctl "lvbu/controllers"
	muser "lvbu/models/user"
	"lvbu/utils"

	"github.com/astaxie/beego"
)

type UserController struct {
	ctl.BaseController
}
type UserLoginController struct {
	beego.Controller
}

func (c *UserController) Logout() {
	c.SetSession("uid", nil)
	c.Redirect("/login", 302)

}
func (c *UserLoginController) Get() {
	beego.Debug("已有你的SESSION状态(uid):", c.GetSession("uid"))
	if c.GetSession("uid") != nil {
		c.Redirect("/index", 302)
		return
	}
	c.TplName = "user/user_login.tpl"
}
func (c *UserLoginController) Post() {
	var user muser.User
	username := c.GetString("username")
	passwd := c.GetString("passwd")
	user.UserName = username
	user.Read("UserName")
	if user.Passwd == utils.Md5(passwd) {
		c.SetSession("uid", user.Id)
		c.Redirect("/index", 302)
	} else {
		c.Data["message"] = "帐号或密码错误"
		c.TplName = "user/user_login.tpl"
	}
}
func (c *UserController) Index() {
	c.Data["uid"] = c.GetSession("uid")
	c.TplName = "user/user_index.tpl"
}
func (c *UserController) Profile() {
	c.Data["uid"] = c.GetSession("uid")
	var uid uint
	var user muser.User
	switch s := c.GetSession("uid").(type) {
	case int:
		uid = uint(s)
	case string:
		beego.Debug("session.Id 类型string")
		return
	case uint:
		uid = s
		beego.Debug("OK!")
	}
	if uid != 0 {

		user.Id = uint(uid)
		user.Read()
		c.Data["user"] = user
	} else {
		beego.Debug("user.Id不正常")
	}
	beego.Debug(user)
	c.TplName = "user/user_profile.tpl"
}

func (c *UserController) ProfilePost() {
	c.Data["uid"] = c.GetSession("uid")
	var uid uint
	switch s := c.GetSession("uid").(type) {
	case int:
		uid = uint(s)
	case string:
		beego.Debug("session.Id 类型string")
		return
	case uint:
		uid = s
		beego.Debug("OK!")
	}
	if uid != 0 {
		var user muser.User
		user.Id = uint(uid)
		user.Read()
		if phone := c.GetString("phone"); phone != "" {
			user.Phone = phone
		}
		if email := c.GetString("email"); email != "" {
			user.Email = email
		}

		if pw := c.GetString("password"); pw != "" {
			user.Passwd = utils.Md5(pw)
		}
		user.Update()
		c.Data["user"] = user
		c.Data["message"] = "修改成功!"
	} else {
		beego.Debug("user.Id不正常")
	}

	c.TplName = "user/user_profile.tpl"
}

func (c *UserController) Headimg() {
	c.Data["uid"] = c.GetSession("uid")
	c.TplName = "user/user_headimg.tpl"
}

func (c *UserController) Lock() {
	var user muser.User
	tmp, _ := c.GetUint16(":id")
	user.Id = uint(tmp)
	user.Status = 1

	if err := user.Update("Status"); err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	} else {
		c.Data["json"] = "lock"
		c.ServeJSON()
		return
	}

}
func (c *UserController) Unlock() {
	var user muser.User
	tmp, _ := c.GetUint16(":id")
	user.Id = uint(tmp)
	user.Status = 0

	if err := user.Update("Status"); err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	} else {
		c.Data["json"] = "unlock"
		c.ServeJSON()
		return
	}
}

func (c *UserController) Jqrmuser() {
	userid, _ := c.GetInt("userid")
	beego.Debug("删除用户Id:", userid)
	var user muser.User
	user.Id = uint(userid)
	if err := user.Delete(); err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	} else {
		c.Data["json"] = "success"
		c.ServeJSON()
		return
	}

}
