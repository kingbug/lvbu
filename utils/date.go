package utils

import (
	//"fmt"
	"time"
)

type EventType int

const (
	EVENT_NODE_STAT   = iota //0状态检测
	EVENT_PRO_STAT           //1项目状态检测
	EVENT_ERROR              //2错误
	EVENT_MESSAGE            //3信息
	EVENT_UPDATE_NODE        //4部署完成时，容器ID更改，这时，需要主动推给WEB客户端做相应改变
)

type STATSTYPE int

const (
	STATS_RUNNING    = iota //0正在运行
	STATS_RESTARTING        //1正在重启	html 隐式提示
	STATS_EXIT              //2停止运行	html 黄灯闪烁
	STATS_EXIST             //3不存在		html 红灯闪烁
)

type Event struct {
	Type           EventType
	Message        string
	Envsign        string
	Proid          string
	Proallnodes    int
	Prorunnodes    int
	Nodeid         string
	Containerid    string
	Containerstats STATSTYPE
	Error          string
}

/*
func GetDate(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 15:04")
}
func GetDateMH(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("01-02 03:04")
}*/

func GetDateFormat(timestamp int64, format string) string {
	if timestamp <= 0 {
		return ""
	}
	tm := time.Unix(timestamp, 0)
	return tm.Format(format)
}

func GetDate(timestamp int64) string {
	if timestamp <= 0 {
		return ""
	}
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02")
}

func GetDateMH(timestamp int64) string {
	if timestamp <= 0 {
		return ""
	}
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 15:04")
}

func GetTimeParse(times string) int64 {
	if "" == times {
		return 0
	}
	loc, _ := time.LoadLocation("Local")
	parse, _ := time.ParseInLocation("2006-01-02 15:04", times, loc)
	return parse.Unix()
}

func GetDateParse(dates string) int64 {
	if "" == dates {
		return 0
	}
	loc, _ := time.LoadLocation("Local")
	parse, _ := time.ParseInLocation("2006-01-02", dates, loc)
	return parse.Unix()
}
