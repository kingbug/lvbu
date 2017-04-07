package utils

import (
	"bytes"
	"encoding/gob"
	"errors"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/memcache"
	_ "github.com/astaxie/beego/cache/redis"
)

var cc cache.Cache
var cache_expire int

func InitCache() {
	cacheConfig := beego.AppConfig.String("cache")
	cache_time, errs := beego.AppConfig.Int("cache_expire")
	if errs != nil {
		cache_expire = 60
	} else {
		cache_expire = cache_time
	}

	switch cacheConfig {
	case "redis":
		initRedis()
		break
	case "memcache":
		initMemcache()
		break
	case "memory":
		initMemory()
		break
	default:
		cc = nil
		break
	}
}

func initMemcache() {
	var err error
	cc, err = cache.NewCache("memcache", `{"conn":"`+beego.AppConfig.String("memcache_host")+`"}`)

	if err != nil {
		beego.Info(err)
	}
}
func initMemory() {
	var err error
	cc, err = cache.NewCache("memory", `{"interval":`+beego.AppConfig.String("cache_expire")+`}`)

	if err != nil {
		beego.Info(err)
	}
}

func initRedis() {
	// cc = &cache.Cache{}
	var err error

	defer func() {
		if r := recover(); r != nil {
			//fmt.Println("initial redis error caught: %v\n", r)
			cc = nil
		}
	}()
	cc, err = cache.NewCache("redis", `{"conn":"`+beego.AppConfig.String("redis_host")+`","password":"`+beego.AppConfig.String("redis_passwd")+`"}`)
	if err != nil {
		//fmt.Println(err)
	}
}

func SetCache(key string, value interface{}) error {
	if cc == nil {
		return errors.New("cc is nil")
	}
	data, err := Encode(value)
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			//fmt.Println("set cache error caught: %v\n", r)
			cc = nil
		}
	}()
	timeouts := time.Duration(cache_expire) * time.Second
	err = cc.Put(key, data, timeouts)
	if err != nil {
		//fmt.Println("Cache失败，key:", key)
		return err
	} else {
		//fmt.Println("Cache成功，key:", key)
		return nil
	}
}
func GetCache(key string, to interface{}) error {
	if cc == nil {
		return errors.New("cc is nil")
	}

	defer func() {
		if r := recover(); r != nil {
			//fmt.Println("get cache error caught: %v\n", r)
			cc = nil
		}
	}()
	data := cc.Get(key)
	if data == nil {
		return errors.New("Cache不存在")
	}
	// log.Pinkln(data)
	err := Decode(data.([]byte), to)
	if err != nil {
		//fmt.Println("获取Cache失败", key, err)
	} else {
		//fmt.Println("获取Cache成功", key)
	}
	return err
}

func DelCache(key string) error {
	if cc == nil {
		return errors.New("cc is nil")
	}

	defer func() {
		if r := recover(); r != nil {
			//fmt.Println("get cache error caught: %v\n", r)
			cc = nil
		}
	}()
	err := cc.Delete(key)
	if err != nil {
		return errors.New("Cache删除失败")
	} else {
		//fmt.Println("删除Cache成功 " + key)
		return nil
	}
}

func ClearAll() error {
	return cc.ClearAll()
}

// --------------------
// Encode
// 用gob进行数据编码
//
func Encode(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// -------------------
// Decode
// 用gob进行数据解码
//
func Decode(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}
