package init

import (
	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
)

func Initlog() {
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)
	logs.Async(1e3)
	beego.SetLogger(logs.AdapterMultiFile, `{"filename":"test.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)
}
