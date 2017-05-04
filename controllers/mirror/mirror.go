package mirror

import (
	//"encoding/json"
	"fmt"
	ctl "lvbu/controllers"
	"lvbu/models/mirror"
	mper "lvbu/models/permission"
	"strconv"

	"github.com/astaxie/beego"
)

type MirController struct {
	ctl.BaseController
}

func (c *MirController) Gadd() {
	c.Data["uid"] = c.GetSession("uid")
	c.TplName = "mirror/mirrorgroup_add.tpl"
}

func (c *MirController) GaddPost() {
	c.Data["uid"] = c.GetSession("uid")
	mirror_group_name := c.GetString("mirrorgroupname")
	if mirror_group_name == "" {
		fmt.Println("(mirrorgroupname)输入为空:", mirror_group_name)
		c.Data["emessage"] = "输入错误或不能为空"
		c.TplName = "mirror/mirrorgroup_add.tpl"
	} else {
		beego.Debug("开始写入数据库")
		var mirror_group mirror.Mirrorgroup
		mirror_group.Name = mirror_group_name
		beego.Debug("mirror_group_name:", mirror_group_name)
		if err := mirror_group.ReadOrCreate(); err != nil {
			beego.Debug("写入数据库出错:", err)
			c.Data["emessage"] = "数据库出错" + err.Error() + "，请重试"
			c.TplName = "mirror/mirrorgroup_add.tpl"
		} else {
			c.Redirect("mirrlist", 302)
		}

	}

}

func (c *MirController) Get() {
	c.Data["uid"] = c.GetSession("uid")
	c.TplName = "mirror/mirrorgroup_add.tpl"
}

//添加镜像页面GET
func (c *MirController) Add() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isperitem("mira", uid) { //镜像添加(mira)
		beego.Debug("动作:请求添加镜像,权限验证失败")
		c.Abort("503")
	}
	var mirror_groups []mirror.Mirrorgroup
	new(mirror.Mirrorgroup).Query().All(&mirror_groups)
	c.Data["mirgs"] = mirror_groups
	c.TplName = "mirror/mirror_add.tpl"
}

//添加镜像页面Post
func (c *MirController) AddPost() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isperitem("mira", uid) { //镜像添加(mira)
		beego.Debug("动作:请求添加镜像POST,权限验证失败")
		c.Abort("503")
	}
	mirname := c.GetString("mirname")
	mirrorgroup := c.GetString("mirrorgroup")
	hubaddress := c.GetString("hubaddress")
	//判断是否为空已放在JS里面
	mir := new(mirror.Mirror)
	mir.Name = mirname
	id, _ := strconv.Atoi(mirrorgroup)
	//mirg := mirror.Mirrorgroup{Id: uint(id)}.Read()
	mir.Mirrorgroup = &mirror.Mirrorgroup{Id: uint(id)}
	mir.Hubaddress = hubaddress
	if err := mir.Insert(); err != nil {
		beego.Debug("插入镜像时，数据库出错")
		c.Data["emessage"] = "插入镜像时，数据库出错"
		c.TplName = "mirror/mirror_add.tpl"
	} else {
		c.Redirect("mirrlist?mirgid="+mirrorgroup, 302)
	}

}

func (c *MirController) Edit() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isperitem("mire", uid) { //镜像修改(mire)
		beego.Debug("动作:请求镜像修改,权限验证失败")
		c.Abort("503")
	}
	c.TplName = "mirror/mirror_edit.tpl"
}

func (c *MirController) List() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isperitem("mirs", uid) { //镜像查看(mirs)
		beego.Debug("动作:请求镜像列表,权限验证失败")
		c.Abort("503")
	}

	//取出所有镜像类别
	var mirror_groups []mirror.Mirrorgroup
	row, _ := new(mirror.Mirrorgroup).Query().All(&mirror_groups)
	fmt.Println(mirror_groups)
	c.Data["mirgs"] = mirror_groups
	// 如果get有mirgid 类别ID说明是从添加镜像页面过来，取出该类别的所有镜像，优先显示当前页列出该 "mirgid"类别镜像
	var mirgid uint
	if re_mirgid := c.GetString("mirgid"); re_mirgid != "" {
		a, _ := (strconv.ParseInt(re_mirgid, 10, 0))
		mirgid = uint(a)
		beego.Warning("从添加页面过来...")
		c.Data["isadd"] = true //添加镜像页面标识
	} else {
		if row > 0 {
			mirgid = mirror_groups[0].Id
			beego.Debug("默认类别Id:", mirgid)
		}

	}
	c.Data["mirgid"] = mirgid
	var mirs []*mirror.Mirror
	if row, _ := new(mirror.Mirror).Query().Filter("Mirrorgroup__Id", mirgid).All(&mirs); row > 0 {
		c.Data["mirs"] = mirs
	} else {
		beego.Warning("未获得该类别下镜像!!!, Mirrorgroup__Id:", mirgid)
		c.TplName = "mirror/mirror_list.tpl"
	}

	c.TplName = "mirror/mirror_list.tpl"
}

//JQuery 请求镜像类别列表
func (c *MirController) JQlist() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isperitem("mirs", uid) { //镜像查看(mirs)
		beego.Debug("动作:请求镜像列表,权限验证失败")
		c.Abort("503")
	}
	mirgid := c.GetString("mirgid")
	if mirgid != "" {
		var mirs []*mirror.Mirror
		new(mirror.Mirror).Query().Filter("Mirrorgroup__Id", mirgid).All(&mirs)

		beego.Debug(mirs)
		c.Data["json"] = &mirs
		c.ServeJSON()
		return

	} else {
		c.Data["json"] = nil
		c.ServeJSON()
		return
	}
}

//JQuery 请求修改镜像信息
func (c *MirController) JQmirr() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isperitem("mire", uid) { //镜像修改(mire)
		beego.Debug("动作:请求镜像修改,权限验证失败")
		c.Abort("503")
	}
	mirid, err := c.GetInt("mirid")
	if err == nil {
		var mirs mirror.Mirror
		//mirid, _ := strconv.ParseInt(mirid, 10, 64)
		beego.Debug("jq请求修改镜像Id:", mirid)
		mirs.Id = uint(mirid)
		err := mirs.Read()
		//err := mirs.Query().Filter("Id", mirid).One(&mirs)
		if err != nil {
			beego.Error("查询镜像时出错:", err.Error())
			c.Data["json"] = map[string]interface{}{"message": "error", "content": "数据库出错:" + err.Error(), "type": 2}
			c.ServeJSON()
			return
		}

		beego.Debug("JS请求更新前镜像,", mirs)
		mirs.Name = c.GetString("mirname")
		mirs.Hubaddress = c.GetString("mirhubaddress")
		beego.Debug("JS请求更新镜像,name:", mirs.Name, ", hubaddress:", mirs.Hubaddress)
		mirs.Update()
		beego.Debug(mirs)
		c.Data["json"] = map[string]interface{}{"message": "success", "content": "成功更新信息镜像", "type": 1}
		c.ServeJSON()
		return

	} else {
		beego.Error(err.Error())
		c.Data["json"] = map[string]interface{}{"message": "error", "content": "传参错误", "type": 3}
		c.ServeJSON()
		return
	}
}

//JQuery 请求删除镜像信息
func (c *MirController) JQrmmirr() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isperitem("mird", uid) { //镜像删除(mird)
		beego.Debug("动作:请求删除镜像,权限验证失败")
		c.Abort("503")
	}
	mirid, err := c.GetInt("mirid")
	if err == nil {
		var mirs mirror.Mirror
		mirs.Id = uint(mirid)
		beego.Debug("删除镜像Id:", mirid)
		if err := mirs.Delete(); err != nil {
			c.Data["json"] = map[string]interface{}{"message": "error", "content": "数据库出错:" + err.Error(), "type": 2}
			c.ServeJSON()
			return
		}
		c.Data["json"] = map[string]interface{}{"message": "success", "content": "成功删除镜像", "type": 1}
		c.ServeJSON()
		return

	} else {
		c.Data["json"] = map[string]interface{}{"message": "error", "content": "传参错误", "type": 3}
		c.ServeJSON()
		return
	}
}
