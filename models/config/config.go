package config

import (
	mpro "lvbu/models/project"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Config struct {
	Id          uint          `orm:"pk;auto"`
	Name        string        `orm:"size(50)"`
	Description string        `orm:"size(200)"`
	Dvalue      string        `orm:"type(text)"`
	Tvalue      string        `orm:"type(text)"`
	Ovalue      string        `orm:"type(text)"`
	Dtstatus    int           `orm:column(默认0无状态,1为已修改,2,3为已删除)`
	Tostatus    int           `orm:column(默认0无状态,1为已修改,3为已删除)`
	Filename    string        `orm:"size(200)"column(一个文件时前端显示'默认文件')`
	Pro         *mpro.Project `orm:"rel(fk)"`
	Content     string        `orm:"size(200)"`
	Created     time.Time     `orm:"auto_now_add;type(datetime)"`
}

func (m *Config) TableName() string {
	return beego.AppConfig.String("dbprefix") + "config"
}

func (m *Config) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Config) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Config) GetRead(proid string) ([]*Config, error) {
	var pro mpro.Project
	if err := new(mpro.Project).Query().Filter("Sign", proid).One(&pro, "Id"); err != nil {
		return nil, err
	}
	var conf []*Config
	if _, cerr := new(Config).Query().Filter("Pro__Id", pro.Id).Filter("Tostatus__lt", 3).All(conf); cerr != nil {
		return nil, cerr
	}
	return conf, nil
}

func (m *Config) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Config) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Config) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func GetConfforName(pid uint, filename string) []*Config {
	var confs []*Config
	if _, err := new(Config).Query().Filter("Pro__Id", pid).Filter("Filename", filename).All(&confs); err != nil {
		return nil
	}
	return confs
}
