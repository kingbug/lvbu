package mirror

import (
	"errors"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Mirror struct {
	Id          uint
	Name        string       `orm:"size(50)"`
	Hubaddress  string       `orm:"size(100)"`
	Mirrorgroup *Mirrorgroup `orm:"null;rel(fk);on_delete(set_null)"`
	Updated     time.Time    `orm:"auto_now;type(datetime)"`
	Created     time.Time    `orm:"auto_now_add;type(datetime)"`
}

func (m *Mirror) TableName() string {
	return beego.AppConfig.String("dbprefix") + "mirror"
}

func (m *Mirror) ReadOrCreate() error {
	if created, _, err := orm.NewOrm().ReadOrCreate(m, "Name"); err == nil {
		if created {
			return nil
		} else {
			return errors.New("已有相同名称")

		}
	} else {
		return err
	}
}

func (m *Mirror) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Mirror) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Mirror) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Mirror) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Mirror) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func Getmir(groupid uint) ([]*Mirror, error) {
	var mirs []*Mirror
	if _, err := new(Mirror).Query().Filter("Mirrorgroup__Id", groupid).All(&mirs); err != nil {
		return mirs, err
	}
	return mirs, nil
}
