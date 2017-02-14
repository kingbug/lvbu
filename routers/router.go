package routers

import (
	"lvbu/controllers"
	"lvbu/controllers/config"
	"lvbu/controllers/env"
	"lvbu/controllers/machine"
	"lvbu/controllers/mirror"
	"lvbu/controllers/node"
	"lvbu/controllers/project"
	"lvbu/controllers/record"
	"lvbu/controllers/sys"
	"lvbu/controllers/user"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//用户
	beego.Router("/login", &user.UserLoginController{}, "get:Get;post:Post")
	beego.Router("/index", &user.UserController{}, "*:Index")
	beego.Router("/profile", &user.UserController{}, "get:Profile;post:ProfilePost")
	beego.Router("/headimg", &user.UserController{}, "*:Headimg")
	beego.Router("/logout", &user.UserController{}, "*:Logout")
	beego.Router("/lockuser/:id([0-9]+)", &user.UserController{}, "*:Lock")
	beego.Router("/unlockuser/:id([0-9]+)", &user.UserController{}, "*:Unlock")
	beego.Router("/jqrmuser", &user.UserController{}, "*:Jqrmuser")

	//项目
	beego.Router("/prolist", &project.ProController{}, "*:List")
	beego.Router("/proedit/:id([0-9]+)", &project.ProController{}, "*:Edit")
	beego.Router("/proadd", &project.ProController{}, "*:Add")
	beego.Router("/prodel", &project.ProController{}, "post:Del")
	//节点
	beego.Router("/nodelist", &node.NodeController{}, "*:List")
	beego.Router("/nodedit", &node.NodeController{}, "*:Edit")
	beego.Router("/nodeadd", &node.NodeController{}, "*:Add")
	//主机
	beego.Router("/maclist", &machine.MacController{}, "*:List")
	beego.Router("/macedit/:id([0-9]+)", &machine.MacController{}, "*:Edit")
	beego.Router("/macadd", &machine.MacController{}, "*:Add")
	beego.Router("/macdel", &machine.MacController{}, "post:Del")
	//镜像
	beego.Router("/mirrlist", &mirror.MirController{}, "*:List")
	beego.Router("/mirredit", &mirror.MirController{}, "*:Edit")
	beego.Router("/mirradd", &mirror.MirController{}, "get:Add;post:AddPost")
	beego.Router("/mirrgroupadd", &mirror.MirController{}, "get:Gadd;post:GaddPost")
	beego.Router("/jqmirrlist", &mirror.MirController{}, "*:JQlist")
	beego.Router("/jqmirr", &mirror.MirController{}, "*:JQmirr")
	beego.Router("/jqrmmir", &mirror.MirController{}, "*:JQrmmirr")
	//环境
	beego.Router("/env", &env.EnvController{}, "*:List")
	//系统
	beego.Router("/usermanager", &sys.UserController{}, "*:List")
	beego.Router("/useradd", &sys.UserController{}, "get:AddGet;post:AddPost")
	beego.Router("/useredit/:id([0-9]+)", &sys.UserController{}, "get:EditGet;post:EditPost")
	beego.Router("/permanage/:id([0-9]+)", &sys.PerController{}, "Get:List;post:Post")
	beego.Router("/poslist", &sys.PosController{}, "*:List")
	//职位
	beego.Router("/jqaddpos", &sys.PosController{}, "*:Jqaddpos")
	beego.Router("/jqupdatepos", &sys.PosController{}, "*:Jqupdatepos")
	beego.Router("/jqrmpos/:id([0-9]+)", &sys.PosController{}, "*:Jqrmpos")
	beego.Router("/about", &sys.SysController{}, "*:About")
	beego.Router("/503", &sys.SysController{}, "*:Noentry")
	//记录
	beego.Router("/reclist", &record.RecController{}, "*:List")
	//配置
	beego.Router("/conlist", &config.ConController{}, "*:List")

}
