package user

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Position struct {
	Id         uint      `orm:"pk;auto"`
	Name       string    `orm:"size(50)"`
	Sign       string    `orm:"size(10)"`
	Permission string    `orm:"size(200)"`
	Created    time.Time `orm:"auto_now_add;type(datetime)"`
	Updated    time.Time `orm:"auto_now;type(datetime)"`
}

func (m *Position) TableName() string {
	return beego.AppConfig.String("dbprefix") + "position"
}

func (m *Position) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Position) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Position) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Position) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Position) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
