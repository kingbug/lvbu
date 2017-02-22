package project

import (
	//	"lvbu/utils"
	"time"

	mac "lvbu/models/machine"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Node struct {
	Id      uint         `orm:"pk;auto"` //
	Name    string       `orm:"size(50);unique"`
	Sign    string       `orm:"size(50)"`
	Pro     *Project     `orm:"rel(fk)"`
	DocId   string       `orm:"size(100)"`
	Mac     *mac.Machine `orm:"rel(fk)"`
	Created time.Time    `orm:"auto_now_add;type(datetime)"`
}

func (m *Node) TableName() string {
	return beego.AppConfig.String("dbprefix") + "node"
}

func (m *Node) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Node) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Node) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Node) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Node) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func Getnode(proid uint) []*Node {
	var node []*Node
	if _, err := new(Node).Query().Filter("Pro__Id", proid).All(&node); err != nil {
		beego.Error("动作:数据库操作,查询项目所载节点列表出错:", err)
	}
	return node
}
