package controllers

import (
	//"fmt"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}
type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "sys/index.tpl"
}
func (c *BaseController) Prepare() {
	_, ok := (c.GetSession("uid")).(uint)
	if !ok && c.Ctx.Request.RequestURI != "/login" {
		c.Data["redirect"] = c.Ctx.Request.RequestURI
		c.Redirect("/login", 302)
	}
}
