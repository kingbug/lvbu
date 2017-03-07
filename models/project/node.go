package project

import (
	"time"

	mac "lvbu/models/machine"
	mir "lvbu/models/mirror"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Node struct {
	Id      uint         `orm:"pk;auto"` //
	Name    string       `orm:"size(50)"`
	Sign    string       `orm:"size(50)"`
	Pro     *Project     `orm:"rel(fk)"`
	DocId   string       `orm:"size(100)"`
	Port    string       `orm:"size(100)"`
	Mir     *mir.Mirror  `orm:"rel(fk)"`
	Mac     *mac.Machine `orm:"rel(fk)"`
	CurVer  string       `orm:"size(50)"`
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

//返回proid 项目，env_sign环境下的所有节点
func Getnode(proid uint, env_sign string) []*Node {
	var node []*Node
	if _, err := new(Node).Query().Filter("Pro__Id", proid).Filter("Mac__Env__Sign", env_sign).All(&node); err != nil {
		beego.Error("动作:数据库操作,查询项目所载节点列表出错:", err)
	}
	for _, v := range node {
		if err := v.Mac.Read(); err != nil {
			beego.Error("动作:数据库操作,查询节点所属主机出错:", err)
		}
	}
	return node
}
