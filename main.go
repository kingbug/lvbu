package main //package main

import (
	_ "lvbu/init"
	_ "lvbu/routers"
	"runtime"

	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	orm.RunSyncdb("default", false, true)
	beego.Run()

}
