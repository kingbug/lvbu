package env

import (
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Env struct {
	Id      uint      `orm:"pk;auto"`
	Name    string    `orm:"size(50)"`
	Sign    string    `orm:"size(10)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func (m *Env) TableName() string {
	return beego.AppConfig.String("dbprefix") + "env"
}

func (m *Env) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Env) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Env) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Env) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Env) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func Getenv() []*Env {
	var envs []*Env
	var env Env
	_, err := env.Query().All(&envs, "Id", "Name")
	if err != nil {
		beego.Error("查询数据库出错:" + err.Error())
	}
	return envs
}

//通过 user表的Permission字段查找env表的Id 列表
func GetEnvIdList(perm string) []uint {
	var envidlist []uint
	a := strings.Split(perm, ",")
	for k, value := range a {
		var posi Env
		if value != "" {
			if err := new(Env).Query().Filter("Sign", value).One(&posi); err != nil {
				beego.Error("查询数据表"+new(Env).TableName()+"出错:"+err.Error()+"遍历", k)
				continue
			} else {
				envidlist = append(envidlist, posi.Id)
			}
		}

	}
	beego.Debug(envidlist)
	return envidlist
}
