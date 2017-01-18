package project

import (
	//	"lvbu/utils"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Project struct {
	Id      uint      `orm:"pk;auto"` //
	Name    string    `orm:"size(50);unique"`
	Sign    string    `orm:"size(50)"`
	Git     string    `orm:"size(100)"`
	Gituser string    `orm:"size(50)"`
	Gitpass string    `orm:"size(50)"`
	Insfile string    `orm:"size(500)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
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
