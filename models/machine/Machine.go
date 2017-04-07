package machine

import (
	"errors"
	"lvbu/models/env"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Machine struct {
	Id        uint      `orm:"pk;auto"`
	Name      string    `orm:"size(50)"`
	Ipaddr1   string    `orm:"size(50)"` //外网
	Ipaddr2   string    `orm:"size(50)"`
	Env       *env.Env  `orm:"rel(fk)"`
	Adminurl  string    `orm:"size(100)"`
	Content   string    `orm:"size(100)"`
	Status    uint8     `orm:"default(0)"`
	Openports string    `orm:"type(text)"`
	Created   time.Time `orm:"auto_now_add;type(datetime)"`
	Updated   time.Time `orm:"auto_now;type(datetime)"`
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

func GetMacforenvname(macid uint) string {
	if macid == 0 {
		beego.Error("动作：查询主机环境出错")
	}
	mach := Machine{Id: macid}
	mach.Read()
	mach.Env.Read()
	return mach.Env.Name
}

//返回主机开放端口列表
func (m *Machine) getOpenports() []string {
	var ports []string
	if m.Openports == "" {
		return ports
	}
	ports = strings.Split(m.Openports, ",")
	return ports
}

//添加主机开放端口记录
func (m *Machine) Addport(ports []string) (bool, []string) {
	var tmp_ports []string
	if len(ports) == 0 {
		return true, tmp_ports
	}
	tmp_open_port := m.Openports
	ok := true
	for _, port := range ports {
		for _, v := range m.getOpenports() {
			if port == v {
				ok = false
				tmp_ports = append(tmp_ports, port)
			} else {
				m.Openports = port + "," + m.Openports
			}
		}
	}
	if !ok {
		m.Openports = tmp_open_port
		return ok, tmp_ports
	}
	return ok, tmp_ports
}

//删除主机开放端口记录
func (m *Machine) Delport(ports []string) {
	var tmp_ports string
	for _, port := range m.getOpenports() {
		if port == "" {
			continue
		}
		ishas := true
		for _, del_port := range ports {
			if del_port == port {
				ishas = false
				break
			}
		}
		if ishas {
			tmp_ports = port + "," + tmp_ports
		}
	}
	if tmp_ports != "" {
		m.Openports = tmp_ports[:len(tmp_ports)-1]
	}

}

//给定SIGN(OE,QE,DE) 返回所有主机列表
func (m *Machine) GetMacforenv(sign string) []*Machine {
	var macs []*Machine
	if _, err := orm.NewOrm().QueryTable(m).Filter("Env__Sign", strings.ToUpper(sign)).All(&macs, "Id", "Adminurl", "Status"); err != nil {
		beego.Error("动作数据库操作,查询环境", sign, "主机列表出错了")
		return macs
	}
	return macs
}
