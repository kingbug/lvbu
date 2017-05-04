package user

import (
	"fmt"
	"lvbu/utils"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id         uint   `orm:"pk;auto"` //用户id
	UserName   string `orm:"size(50);unique"`
	Passwd     string `orm:"size(50)"`
	Nick       string `orm:"size(50)"`
	Sex        uint
	Email      string    `orm:"size(50)"`
	Phone      string    `orm:"size(50)"`
	Avatar     string    `orm:"size(100)"`
	Position   *Position `orm:"rel(fk)"`
	Permission string    `orm:"size(200)"`
	Created    time.Time `orm:"auto_now_add;type(date)"`
	Updated    time.Time `orm:"auto_now;type(date)"`
	Status     int       `orm:"default(0)"`
}

func (m *User) TableName() string {
	return beego.AppConfig.String("dbprefix") + "user"
}

func (m *User) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *User) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *User) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
func (m *User) Clscache() error {
	utils.DelCache("user.id." + fmt.Sprintf("%d", m.Id))
	return nil
}

func Getuser(id uint) User {
	var user User
	err := utils.GetCache("user.id."+fmt.Sprintf("%d", id), &user)
	if err != nil {
		user = User{Id: id}
		user.Read()
		fmt.Println(user.Position)
		utils.SetCache("user.id."+fmt.Sprintf("%d", id), user)
	}
	return user
}
func Getposname(id uint) string {
	var pos string
	err := utils.GetCache("pos.name."+fmt.Sprintf("%d", id), &pos)
	if err != nil {
		positions := Position{Id: id}
		positions.Read()
		pos = positions.Name
		utils.SetCache("pos.name."+fmt.Sprintf("%d", id), pos)
	}
	return pos
}

func Count() int64 {
	count, _ := new(User).Query().Count()
	return count
}
