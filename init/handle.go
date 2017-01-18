package init

import (
	"html/template"
	"net/http"

	"github.com/astaxie/beego"
)

func Inithandle() {
	beego.ErrorHandler("404", page_not_found)
	beego.ErrorHandler("401", page_note_permission)
}
func page_not_found(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("404.tpl").ParseFiles("views/common/404.tpl")
	data := make(map[string]interface{})
	t.Execute(rw, data)
}

func page_note_permission(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("401.tpl").ParseFiles("views/401.tpl")
	data := make(map[string]interface{})
	t.Execute(rw, data)
}
