package env

import (
	"time"
)

type Env struct {
	Id      uint      `orm:"pk;auto"`
	Name    string    `orm:"size(50)"`
	Sign    string    `orm:"size(10)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}
