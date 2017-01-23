package init

import (
	"fmt"
	mcn "lvbu/models/config"
	men "lvbu/models/env"
	mac "lvbu/models/machine"
	mir "lvbu/models/mirror"
	mper "lvbu/models/permission"
	mpro "lvbu/models/project"
	mur "lvbu/models/user"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func InitSql() {
	user := beego.AppConfig.String("mysqluser")
	passwd := beego.AppConfig.String("mysqlpass")
	host := beego.AppConfig.String("mysqlurls")
	port, err := beego.AppConfig.Int("mysqlport")
	dbname := beego.AppConfig.String("mysqldb")
	prefix := beego.AppConfig.String("mysqlpre")
	if nil != err {
		port = 3306
	}
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//orm.RegisterDataBase("default", "mysql", "root:@/blog?charset=utf8", 30)
	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&loc=Local", user, passwd, host, port, dbname))
	//orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", user, passwd, host, port, dbname))
	orm.RegisterModelWithPrefix(prefix, new(mir.Mirrorgroup), new(mir.Mirror), new(mur.User), new(mur.Position), new(men.Env), new(mcn.Config), new(mcn.Configver), new(mac.Machine), new(mper.Permenu), new(mper.Peritem), new(mpro.Project))

}
