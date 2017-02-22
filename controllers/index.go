package controllers

import (
	"strings"
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
	if strings.Split(c.Ctx.Request.RequestURI, "?")[0] == "/confdown" {
		return
	}
	_, ok := (c.GetSession("uid")).(uint)
	if !ok && c.Ctx.Request.RequestURI != "/login" {
		c.SetSession("redirect", c.Ctx.Request.RequestURI)
		c.Redirect("/login", 302)
	}
}
