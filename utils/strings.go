package utils

import (
	"crypto/rand"
	"errors"
	"math/big"

	"net/smtp"
	"regexp"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/zheng-ji/goSnowFlake"
)

/*
func AjaxMsg(code int, message string) {
	this.Data["json"] = map[string]interface{}{"code": code, "message": message}
	this.ServeJSON()
	return
}
*/
//字串截取
func SubString(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func GetFileSuffix(s string) string {
	re, _ := regexp.Compile(".(jpg|jpeg|png|gif|exe|doc|docx|ppt|pptx|xls|xlsx)")
	suffix := re.ReplaceAllString(s, "")
	return suffix
}

func RandInt64(min, max int64) int64 {
	maxBigInt := big.NewInt(max)
	i, _ := rand.Int(rand.Reader, maxBigInt)
	if i.Int64() < min {
		RandInt64(min, max)
	}
	return i.Int64()
}

func Strim(str string) string {
	str = strings.Replace(str, "\t", "", -1)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "\r", "", -1)
	return str
}

func Unicode(rs string) string {
	json := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			json += string(r)
		} else {
			json += "\\u" + strconv.FormatInt(int64(rint), 16)
		}
	}
	return json
}

func HTMLEncode(rs string) string {
	html := ""
	for _, r := range rs {
		html += "&#" + strconv.Itoa(int(r)) + ";"
	}
	return html
}

/**
 *  to: example@example.com;example1@163.com;example2@sina.com.cn;...
 *  subject:The subject of mail
 *  body: The content of mail
 */
func SendMail(to string, subject string, body string) error {
	user := beego.AppConfig.String("mailfrom")
	password := beego.AppConfig.String("mailpassword")
	host := beego.AppConfig.String("mailhost")

	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	content_type = "Content-type:text/html;charset=utf-8"

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func SnowFlakeId() int64 {
	iw, _ := goSnowFlake.NewIdWorker(1)
	if id, err := iw.NextId(); err != nil {
		return 0
	} else {
		return id
	}
}

func Seek(a, b int) int {
	return a % b
}

func Gittoname(giturl string) string {
	end := strings.LastIndex(giturl, ".")
	start := strings.LastIndex(giturl, "/")
	if start == -1 || end == -1 {
		return ""
	}
	return giturl[start+1 : end]
}

//example: 80:90, 88:89 返回90, 89
func GetPortList(port string) []string {
	buf_port := strings.Split(port, ",")
	var tmp_port []string
	for _, v := range buf_port {
		v = strings.Replace(v, " ", "", -1)
		v = strings.Replace(v, "\n", "", -1)
		tmp_port = append(tmp_port, strings.Split(v, ":")[1])
	}
	return tmp_port
}

//example: 80:90, 88:89 返回map["80"]{"90"},["88"]{"89"}
func Getportmap(port string) (map[string]string, error) {
	buf_port := strings.Split(port, ",")
	beego.Debug("buf_port(len):", len(buf_port))
	var tmp_port = make(map[string]string)
	for k, v := range buf_port {
		beego.Debug("k", k)
		v = strings.Replace(v, " ", "", -1)
		v = strings.Replace(v, "\n", "", -1)
		tmp_pairs := strings.Split(v, ":")
		beego.Debug("tmp_paris(len):", len(tmp_pairs), "value:", tmp_pairs)
		if len(tmp_pairs) != 2 {
			return make(map[string]string), errors.New(port + "格式错误")
		}
		beego.Debug(tmp_pairs[0], ",", tmp_pairs[1])
		key := strings.Replace(strings.Replace(tmp_pairs[0], " ", "", -1), "\n", "", -1)
		value := strings.Replace(strings.Replace(tmp_pairs[1], " ", "", -1), "\n", "", -1)

		tmp_port[key] = value
	}
	return tmp_port, nil
}
