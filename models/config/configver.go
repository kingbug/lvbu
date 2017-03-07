package config

import (
	mpro "lvbu/models/project"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Configver struct {
	Id   uint          `orm:"pk;auto"`
	File string        `orm:"size(200)"`
	Pro  *mpro.Project `orm:"rel(fk)"`
	Ver  string        `orm:"size(50)"`
}

func (m *Configver) TableName() string {
	return beego.AppConfig.String("dbprefix") + "configver"
}

func (m *Configver) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Configver) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Configver) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Configver) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Configver) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
