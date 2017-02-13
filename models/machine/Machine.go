package machine

import (
	"errors"
	"lvbu/models/env"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Machine struct {
	Id       uint      `orm:"pk;auto"`
	Name     string    `orm:"size(50)"`
	Ipaddr1  string    `orm:"size(50)"`
	Ipaddr2  string    `orm:"size(50)"`
	Env      *env.Env  `orm:"rel(fk)"`
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"auto_now;type(datetime)"`
	Adminurl string    `orm:"size(100)"`
	Content  string    `orm:"size(100)"`
	Status   uint8     `orm:"default(0)"`
}

func (m *Machine) TableName() string {
	return beego.AppConfig.String("dbprefix") + "machine"
}

func (m *Machine) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Machine) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Machine) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Machine) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Machine) Getenvsign() (string, error) {
	if m.Id == 0 { //uint 默认为零，但数据库Id是从1 开始的
		return "", errors.New("Machine.Id为空")
	}
	m.Read()
	return m.Env.Sign, nil
}

func (m *Machine) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func GetMacforenv(macid uint) string {
	if macid == 0 {
		beego.Debug("动作：查询主机环境出错")
	}
	mach := Machine{Id: macid}
	mach.Read()
	mach.Env.Read()
	return mach.Env.Name
}
