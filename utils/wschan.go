package utils

import (
	"container/list"
	"strconv"

	"github.com/astaxie/beego"
)

type Subscribe struct {
	Id  int64
	Obj interface{}
}

var (
	Subscribes = list.New()
)

func Conjoin(id string, obj interface{}) {
	key, _ := strconv.ParseInt(id, 10, 64)
	Subscribes.PushBack(Subscribe{Id: key, Obj: obj})
	beego.Debug("okkkk...")
}

func Getcon(id string) interface{} {
	key, _ := strconv.ParseInt(id, 10, 64)
	for sub := Subscribes.Front(); sub != nil; sub = sub.Next() {
		if sub.Value.(Subscribe).Id == key {
			return sub.Value.(Subscribe).Obj
		}
	}
	return nil
}

func Conleave(id string) {
	key, _ := strconv.ParseInt(id, 10, 64)
	beego.Debug(key)
	for sub := Subscribes.Front(); sub != nil; sub = sub.Next() {
		beego.Debug("id:", sub.Value.(Subscribe).Id)
		if sub.Value.(Subscribe).Id == key {
			Subscribes.Remove(sub)

			beego.Info("Remove:", key)
			break
		}
	}
	beego.Debug("Len:", Subscribes.Len())
}
