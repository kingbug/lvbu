package mirror

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MirrorGroup struct {
	Id      uint      `orm:"pk;auto"` //
	Name    string    `orm:"size(50);unique"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
}

func (m *MirrorGroup) TableName() string {
	return beego.AppConfig.String("dbprefix") + "mirgroup"
}

func (m *MirrorGroup) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *MirrorGroup) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *MirrorGroup) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *MirrorGroup) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *MirrorGroup) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
