package project

import (
	"errors"
	"strings"
	//	"lvbu/utils"
	"regexp"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Project struct {
	Id         uint      `orm:"pk;auto"` //
	Name       string    `orm:"size(50);unique"`
	Sign       string    `orm:"size(50)"`
	Compile    string    `orm:"size(50)"`
	Compilever string    `orm:"size(50)"`
	Git        string    `orm:"type(text)"`
	Gituser    string    `orm:"size(50)"`
	Gitpass    string    `orm:"size(50)"`
	Insfile    string    `orm:"size(500)"`
	Conffile   string    `orm:"type(text)"` //配置文件字符串,各个文件已","分割
	Conflist   []string  `orm:"-"`
	Filecount  int       `orm:"-"`          //前端页面判断为0时，显示默认文件 ，否则显示真实文件列表
	Sharedpath string    `orm:"type(text)"` //容器共享目录 -v
	Dns        string    `orm:"type(16)"`   //容器DNS，只支持一个
	Created    time.Time `orm:"auto_now_add;type(datetime)"`
	Updated    time.Time `orm:"auto_now;type(datetime)"`
}

func (m *Project) TableName() string {
	return beego.AppConfig.String("dbprefix") + "project"
}

func (m *Project) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Project) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Project) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Project) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Project) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func Getproject() []*Project {
	var pro []*Project
	if _, err := new(Project).Query().All(&pro, "Id", "Name", "Conffile"); err != nil {
		beego.Error("动作:数据库操作,查询项目列表出错:", err)
	}
	for _, p := range pro {
		p.Conflist = strings.Split(p.Conffile, ",")
		p.Filecount = len(strings.Split(p.Conffile, ","))
	}
	return pro
}

//返回项目所载节点数量
func Getprofornodecount(proid uint) int64 {
	var count int64
	var err error
	if count, err = new(Node).Query().Filter("Pro__Id", proid).Count(); err != nil {
		beego.Error("动作:数据库操作,查询项目节点数量出错:", err)
	}
	return count
}

//添加配置文件名
func AddConf(proid uint, filename string) error {
	if !RepComparison(filename) {
		return errors.New("文件名规则不匹配[0-9A-Za-z./]")
	}
	pro := &Project{Id: proid}
	if err := pro.Read(); err != nil {
		return err
	}
	pro.Conflist = strings.Split(pro.Conffile, ",")
	for _, v := range pro.Conflist {
		if v == filename {
			return errors.New("已有重复文件名")
		}
	}
	pro.Conffile = pro.Conffile + "," + filename
	beego.Info("项目:"+pro.Name+",当前配置文件:", pro.Conffile)
	return pro.Update()
}

//编辑配置文件名
func EditConf(proid uint, oldfile, filename string) error {
	if !RepComparison(filename) {
		return errors.New("文件名规则不匹配[0-9A-Za-z./]")
	}
	pro := &Project{Id: proid}
	if err := pro.Read(); err != nil {
		return err
	}
	beego.Info("项目:"+pro.Name+",修改前配置文件:", pro.Conffile)
	pro.Conflist = strings.Split(pro.Conffile, ",")
	pro.Conffile = ""
	for _, v := range pro.Conflist {
		if v == "" {
			continue
		}
		if v == oldfile {
			pro.Conffile = pro.Conffile + "," + filename
			continue
		}
		pro.Conffile = pro.Conffile + "," + v
	}
	beego.Info("项目:"+pro.Name+",修改后配置文件:", pro.Conffile)
	return pro.Update()
}

//删除配置文件名
func DelConf(proid uint, filename string) error {
	pro := &Project{Id: proid}
	if err := pro.Read(); err != nil {
		return err
	}
	pro.Conflist = strings.Split(pro.Conffile, ",")
	pro.Conffile = ""
	for _, conf := range pro.Conflist[1:len(pro.Conflist)] {
		if conf == filename {
			continue
		}
		pro.Conffile = pro.Conffile + "," + conf
	}
	pro.Conffile = pro.Conffile[:len(pro.Conffile)-1]
	return pro.Update()
}

func IsExistName(pid uint, filename string) bool {
	pro := &Project{Id: pid}
	if err := pro.Read(); err != nil {
		beego.Error(err)
	}
	pro.Conflist = strings.Split(pro.Conffile, ",")
	for _, v := range pro.Conflist {
		if v == filename {
			return true
		}
	}
	return false
}

//正则匹配
func RepComparison(filename string) bool {
	reg4 := regexp.MustCompile(`[\/\.]`)
	reg5 := regexp.MustCompile(`[[:alpha:][:digit:]\.\/]`)
	if len(reg5.FindAllString(filename, -1)) != len(filename) {
		return false
	}
	//开头不能为 ./
	if len(reg4.FindAllString(filename[:1], -1)) != 0 {
		return false
	}
	//结尾不能为 ./
	if len(reg4.FindAllString(filename[len(filename)-1:len(filename)], -1)) != 0 {
		return false
	}
	return true
}
