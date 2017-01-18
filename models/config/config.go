package config

import (
	men "lvbu/models/env"
	"time"
)

type Config struct {
	Id       uint   `orm:"pk;auto"`
	Name     string `orm:"size(50)"`
	Dvalue   string `orm:"size(200)"`
	Tvalue   string `orm:"size(200)"`
	Ovalue   string `orm:"size(200)"`
	Dtstatus bool
	Tostatus bool
	Content  string    `orm:"size(200)"`
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
}
type Configver struct {
	Id    uint     `orm:"pk;auto"`
	Name  string   `orm:"size(50)"`
	Value string   `orm:"size(200)"`
	Env   *men.Env `orm:"rel(fk)"`
	Ver   string   `orm:"size(50)"`
}
