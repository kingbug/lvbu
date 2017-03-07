package mirror

import (
	"errors"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Mirrorgroup struct {
	Id   uint
	Name string `orm:"size(50)"`
	//Updated time.Time `orm:"auto_now;type(datetime)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
}

func (m *Mirrorgroup) TableName() string {
	return beego.AppConfig.String("dbprefix") + "mirrorgroup"
}

func (m *Mirrorgroup) ReadOrCreate() error {
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

func (m *Mirrorgroup) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Mirrorgroup) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Mirrorgroup) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Mirrorgroup) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Mirrorgroup) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func Getmirgroup() ([]*Mirrorgroup, error) {
	var mirs []*Mirrorgroup
	if _, err := new(Mirrorgroup).Query().All(&mirs); err != nil {
		return mirs, err
	}
	return mirs, nil
}
