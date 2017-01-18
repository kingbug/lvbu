package permission

import (
	"fmt"
	"lvbu/models/user"
	"lvbu/utils"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Permenu struct {
	Id   uint   `orm:"pk;auto"`
	Name string `orm:"size(50)"`
}
type Peritem struct {
	Id   uint     `orm:"pk;auto"`
	Name string   `orm:"size(50)"`
	Sign string   `orm:"size(10)"`
	Menu *Permenu `orm:"rel(fk)"`
}
type Pershow struct {
	Id    uint
	Name  string
	Sign  string
	Check string
}

func (m *Peritem) TableName() string {
	return beego.AppConfig.String("dbprefix") + "peritem"
}
func (m *Permenu) TableName() string {
	return beego.AppConfig.String("dbprefix") + "permenu"
}
func (m *Peritem) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}
func (m *Peritem) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}
func (m *Peritem) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Peritem) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Peritem) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func (m *Permenu) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Permenu) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Permenu) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Permenu) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}
func (m *Permenu) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
func Getmenu() []Permenu {
	var menu []Permenu
	err := utils.GetCache("permenu", &menu)
	if err != nil {
		new(Permenu).Query().OrderBy("Id").All(&menu)
		utils.SetCache("permenu", menu)
	}
	return menu
}
func Getpospers(peid, poid uint) []Pershow {
	var peritems []Peritem
	var pershow []Pershow
	new(Peritem).Query().Filter("menu_id", peid).All(&peritems)
	posper := Getposper(poid)
	for _, v := range peritems {
		var pers Pershow
		pers.Id = v.Id
		pers.Name = v.Name
		pers.Sign = v.Sign
		if strings.Contains(posper, v.Sign) {
			pers.Check = "checked"
		}
		pershow = append(pershow, pers)
	}
	return pershow
}
func Getposper(poid uint) string {
	var posper string

	var position user.Position
	position.Id = poid
	position.Read()
	posper = position.Permission

	return posper
}
func Isperitem(name string, uid uint) bool {
	ss := false
	err := utils.GetCache("show."+name+fmt.Sprintf("%d", uid), &ss)
	if err != nil {
		var user user.User
		user.Id = uid
		user.Read("Id")
		user.Position.Read("Id")
		if strings.Contains(user.Position.Permission, name) {
			ss = true
		} else {
			ss = false
		}
		utils.SetCache("show."+name+fmt.Sprintf("%d", uid), ss)
	}
	return ss

}
func Isuserper(name string, uid uint) bool {
	ss := false
	err := utils.GetCache("userper."+name+fmt.Sprintf("%d", uid), &ss)
	if err != nil {
		var user user.User
		user.Id = uid
		user.Read("Id")
		if strings.Contains(user.Permission, name) {
			ss = true
		} else {
			ss = false
		}
		utils.SetCache("userper."+name+fmt.Sprintf("%d", uid), ss)
	}
	return ss

}
