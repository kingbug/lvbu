package node

import (
	"bytes"
	"encoding/json"
	"fmt"
	ctl "lvbu/controllers"
	mcn "lvbu/models/config"
	men "lvbu/models/env"
	mac "lvbu/models/machine"
	mir "lvbu/models/mirror"
	mper "lvbu/models/permission"
	mpro "lvbu/models/project"
	"lvbu/utils"
	"net/http"
	//	"os"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//	"github.com/fsouza/go-dockerclient"
	"github.com/gorilla/websocket"
)

type NodeController struct {
	ctl.BaseController
}

func (c *NodeController) Add() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isperitem("noda", uid) { //节点添加(noda)
		beego.Debug("动作:请求添加节点,权限验证失败")
		c.Abort("503")
	}
	proid, _ := c.GetInt(":proid")
	sign := c.GetString(":sign")
	envid := men.Getenvid(sign)
	pro := mpro.Project{Id: uint(proid)}
	env := men.Env{Id: uint(envid)}
	if proerr := pro.Read(); proerr != nil {
		beego.Error("读取项目信息出错:", proerr)
	}
	if enverr := env.Read(); enverr != nil {
		beego.Error("读取项目信息出错:", enverr)
	}
	var macs []mac.Machine
	if _, macserr := new(mac.Machine).Query().Filter("Env__Id", envid).All(&macs); macserr != nil {
		beego.Error("macserr", macserr)
	}
	if c.Ctx.Request.Method == "GET" {
		c.Data["pro"] = &pro
		env.Sign = strings.ToLower(env.Sign)
		c.Data["env"] = &env
		c.Data["macs"] = &macs
		c.TplName = "node/node_add.tpl"
	} else if c.Ctx.Request.Method == "POST" {
		//POST负责提交数据的合法性。，在websocket里面取出这些信息做逻辑
		unix := time.Now().UnixNano()
		rand_str := strconv.FormatInt(unix, 10)[0:12]
		n_sign := c.GetString("sign")
		n_machine, _ := c.GetInt("machine")
		n_port := c.GetString("port")
		n_mirr, _ := c.GetInt("mirr")
		if n_mirr == 0 || n_port == "" || n_machine == 0 || n_sign == "" {
			c.Data["json"] = map[string]interface{}{"message": "error", "data": "提交数据不完整"}
			c.ServeJSON()
			return
		}
		node := mpro.Node{
			Name: n_sign,
			Sign: n_sign,
			Pro: &mpro.Project{
				Id: uint(proid),
			},
			Mac: &mac.Machine{
				Id: uint(n_machine),
			},
			Mir: &mir.Mirror{
				Id: uint(n_mirr),
			},
			Port: n_port,
		}
		beego.Debug(n_port, n_mirr)
		utils.Conjoin(rand_str, &node)
		//提交后，页面js 发送websocket 开始接收信息
		c.Data["json"] = map[string]interface{}{"message": "success", "data": rand_str}
		c.ServeJSON()
		return
	}

}
func (c *NodeController) Edit() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isperitem("node", uid) { //节点修改(node)
		beego.Debug("动作:请求修改节点,权限验证失败")
		c.Abort("503")
	}
	proid := c.GetString(":proid")
	sign := c.GetString(":sign")
	id, _ := c.GetInt(":id")

	node := mpro.Node{
		Id: uint(id),
	}
	if err := node.Read(); err != nil {
		beego.Error("动作:数据库操作,查询节点信息失败:", err)
	}
	if err := node.Mir.Read(); err != nil {
		beego.Error("动作:数据库操作,查询节点镜像信息失败:", err)
	}
	if err := node.Mac.Read(); err != nil {
		beego.Error("动作:数据库操作,查询节点所属主机信息失败:", err)
	}
	if err := node.Mac.Env.Read(); err != nil {
		beego.Error("动作:数据库操作,查询节点所属主机环境信息失败:", err)
	}
	if err := node.Pro.Read(); err != nil {
		beego.Error("动作:数据库操作,查询节点所属项目信息失败:", err)
	}
	var macs []mac.Machine
	if _, macserr := new(mac.Machine).Query().Filter("Env__Id", node.Mac.Env.Id).All(&macs); macserr != nil {
		beego.Error("macserr", macserr)
	}
	node.Mac.Env.Sign = strings.ToLower(node.Mac.Env.Sign)
	c.Data["node"] = &node
	if c.Ctx.Request.Method == "GET" {
		c.Data["macs"] = &macs
		c.TplName = "node/node_edit.tpl"

	} else if c.Ctx.Request.Method == "POST" {
		node_sign := c.GetString("sign")
		node_port := c.GetString("port")
		node_Mir, _ := c.GetInt16("mirror")
		if node_sign == "" || node_port == "" || node_Mir == 0 {
			c.Data["message"] = "必填项不能留空"
			c.TplName = "node/node_edit.tpl"
			return
		}
		mirr := mir.Mirror{
			Id: uint(node_Mir),
		}
		node.Mir = &mirr
		node.Sign = node_sign
		node.Name = node_sign
		node.Port = node_port
		if err := node.Update(); err != nil {
			beego.Error("动作:数据库操作,修改节点信息失败:", err)
			c.Data["message"] = "数据库出错:" + err.Error()
			c.TplName = "node/node_edit.tpl"
			return
		}
		c.Redirect("/"+proid+"/"+sign+"/nodelist", 302)
	}

}

func (c *NodeController) List() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isperitem("nods", uid) { //节点[管理|查看](nods)
		beego.Debug("动作:请求查看节点,权限验证失败")
		c.Abort("503")
	}
	proid, _ := c.GetInt(":proid")
	sign := c.GetString(":sign")
	envid := men.Getenvid(sign)
	var nodes []*mpro.Node
	if _, err := new(mpro.Node).Query().Filter("Pro__Id", uint(proid)).Filter("Mac__Env__Id", uint(envid)).All(&nodes); err != nil {
		beego.Error("动作：获取节点列表，数据库操作出错:", err)
	}
	pro := mpro.Project{Id: uint(proid)}
	if proerr := pro.Read(); proerr != nil {
		beego.Error("proerr:", proerr)
	}
	env := men.Env{Id: uint(envid)}
	if enverr := env.Read(); enverr != nil {
		beego.Error("enverr:", enverr)
	}
	env.Sign = strings.ToLower(env.Sign)
	c.Data["pro"] = &pro
	c.Data["env"] = &env
	c.Data["nodes"] = &nodes
	c.TplName = "node/node_list.tpl"
}

func (c *NodeController) Del() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	if !mper.Isperitem("nodd", uid) { //节点删除(nodd)
		beego.Debug("动作:请求查看节点,权限验证失败")
		c.Abort("503")
	}
	node_id, _ := c.GetUint16("node_id")
	if node_id == 0 {
		c.Abort("503")
		return
	}
	node := mpro.Node{
		Id: uint(node_id),
	}
	if err := node.Delete(); err != nil {
		c.Data["json"] = map[string]interface{}{"message": "error", "data": err.Error()}
		c.ServeJSON()
		return
	} else {
		c.Data["json"] = map[string]interface{}{"message": "success"}
		c.ServeJSON()
		return
	}
}

func (c *NodeController) Wsadd() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	ws, err := websocket.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil, 1024, 1024)
	var message = make(chan string, 10)
	defer ws.Close()
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(c.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup Websocket connection:", err)
		return
	}

	mt, ms, err := ws.ReadMessage()
	beego.Debug("messageType:", mt)
	beego.Debug("message:", ms)
	var bu bytes.Buffer
	bu.Write(ms)
	m := bu.String()
	beego.Debug("[]bytes to string:", m)
	if mt == -1 {
		utils.Conleave(m)
		beego.Error("messageError:", err)
		return
	}

	node := utils.Getcon(m).(*mpro.Node)
	mach := node.Mac //主机
	mirr := node.Mir //镜像
	pro := node.Pro  //项目
	mach.Read()
	mirr.Read()
	pro.Read()
	beego.Debug("主机管理地址:", mach.Adminurl,
		"镜像地址:", mirr.Hubaddress,
		"项目git:", pro.Git,
		"gituser:", pro.Gituser,
		"gitpass:", pro.Gitpass)

	utils.Conleave(m)
	go func() {
		if contr, err := utils.PullImage(mirr.Hubaddress, message); !contr {
			message <- err.Error()
		}
		if err := utils.Gitpull(pro.Git, message); err != nil {
			message <- err.Error()
		}
		message <- "success"

	}()
	beego.Info("websocket 接收到(node):", node)
	defer ws.Close()
	for {
		select {
		case mes := <-message:
			var buf bytes.Buffer
			buf.WriteString(mes)
			if ws.WriteMessage(websocket.TextMessage, buf.Bytes()) != nil {
				ws.Close()
			}
			if mes == "success" {
				ws.Close()
				if err := node.Insert(); err != nil {
					beego.Error("动作:数据库操作，添加节点失败：", err)
				}
				beego.Debug("处理结束")
				break
			}
		}

	}
	beego.Info("message 循环已退出")
}

func (c *NodeController) Wsdeploy() {
	ws, err := websocket.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil, 1024, 1024)
	proid, interr := c.GetInt("proid")
	if interr != nil || proid == 0 {
		beego.Error("从URL地址获取项目ID出错,proid:", proid, ",error:", interr)
	}
	var message = make(chan string, 10)
	defer ws.Close()
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(c.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup Websocket connection:", err)
		return
	}
	var exitstats = make(chan bool)
	go func() { //循环检测节点容器状态
		var nodes []*mpro.Node
		if _, err := new(mpro.Node).Query().Filter("Pro__Id", proid).All(&nodes); err != nil {
			beego.Error("查询数据库出错", err)
		} else if len(nodes) > 0 {
			for _, node := range nodes {
				if err := node.Mac.Read(); err != nil {
					beego.Error("查询节点所属主机出错", err)
				}
			}
		DONE:
			for {

				select {
				case code := <-exitstats:
					if code {
						beego.Debug("退出状态检测")
						break DONE
					}
				default:
					containers, err := utils.Cliinspectcon(nodes)
					if err != nil {
						mes := "检测节点状态出错"
						var buf bytes.Buffer
						buf.WriteString(mes)
						if err := ws.WriteMessage(websocket.TextMessage, buf.Bytes()); err != nil {
							beego.Error("发送websocket出错:", err)
							break DONE
						}
						break DONE
					}
					beego.Debug("containers.len:", len(containers))
					for _, value := range containers {
						if value == nil {
							beego.Info("空对象")
							continue
						}
						event := utils.Event{
							Type:           utils.EVENT_STAT,
							Containerid:    value.ID,
							Containerstats: value.State.Running,
						}
						data, err := json.Marshal(event)
						if err != nil {
							beego.Error("Fail to marshal event:", err)
							continue
						}
						if ws.WriteMessage(websocket.TextMessage, data) != nil {
						}
					}
					time.Sleep(5 * time.Second)
				}

			}
		}

	}()

RECEIVE:
	//tmp_ms: example 90:v8.8.8
	for {
		mt, tmp_ms, _ := ws.ReadMessage()
		beego.Debug("messageType:", mt)
		beego.Debug("message:", tmp_ms)
		if mt == -1 {
			beego.Debug("正在尝试关闭当前websocket连接")
			exitstats <- true //控制检测节点状态for退出
			if err := ws.Close(); err != nil {
				beego.Info("关闭websocket连接出错：", err)
			}
			break RECEIVE
		}
		buf_ms := strings.Split(string(tmp_ms), ":")
		node_id := buf_ms[0]
		node_ver := buf_ms[1]
		var ms int
		if len(tmp_ms) != 0 {
			ms, _ = strconv.Atoi(node_id)
		}
		node := mpro.Node{
			Id: uint(ms),
		}
		//node下个版本会放在list里面,如在一个页面频繁操作,就多次操作同一样的sql,并且在这同时,修改node节点的可能性不大
		if err := node.Read(); err != nil {
			beego.Error("动作:数据库操作,查询节点信息失败:", err)
		}
		if err := node.Mir.Read(); err != nil {
			beego.Error("动作:数据库操作,查询节点镜像信息失败:", err)
		}
		if err := node.Mac.Read(); err != nil {
			beego.Error("动作:数据库操作,查询节点所属主机信息失败:", err)
		}
		if err := node.Mac.Env.Read(); err != nil {
			beego.Error("动作:数据库操作,查询节点所属主机环境信息失败:", err)
		}
		if err := node.Pro.Read(); err != nil {
			beego.Error("动作:数据库操作,查询节点项目信息失败:", err)
		}
		go func() { //生成当前版本JSON文件，以便版本回退时使用
			if node.CurVer != node_ver {

				var oldverconf []orm.ParamsList
				if node.Mac.Env.Sign == "DE" {
					if _, err := new(mcn.Config).Query().Filter("Pro__Id", node.Pro.Id).ValuesList(&oldverconf, "Name", "Dvalue"); err != nil {
						message <- "生成老版本配置时，数据查询失败:" + err.Error()
					}
				} else if node.Mac.Env.Sign == "QE" {
					if _, err := new(mcn.Config).Query().Filter("Pro__Id", node.Pro.Id).ValuesList(&oldverconf, "Name", "Tvalue"); err != nil {
						message <- "生成老版本配置时，数据查询失败:" + err.Error()
					}
				} else { //"OE"
					if _, err := new(mcn.Config).Query().Filter("Pro__Id", node.Pro.Id).ValuesList(&oldverconf, "Name", "Ovalue"); err != nil {
						message <- "生成老版本配置时，数据查询失败:" + err.Error()
					}
				}
				var conf []utils.Conf
				for _, v := range oldverconf {
					conf = append(conf, utils.Conf{Key: fmt.Sprintf("%s", v[0]), Value: fmt.Sprintf("%s", v[1])})
				}
				if err := utils.Makejsonconf(node.Pro.Sign, node.Mac.Env.Sign, node.CurVer, conf); err != nil {
					message <- "生成老版本配置失败, error:" + err.Error()
				}
			}

		}()
		go func() { //部署容器
			if node.DocId != "" {
				if err := utils.Clidelcon(node.Mac.Adminurl, node.DocId); err != nil {
					message <- "Info:" + err.Error()
				}
				//客户端创建镜像
			}

			//BuildImage
			if err := utils.BuildImage(&node, node_ver, message); err != nil {
				message <- "镜像BUILD失败,error:" + err.Error()
			} else if err := utils.PushImages(node.Pro.Git, node_ver, message); err != nil {
				//上传镜像
				message <- "上传镜像失败,error:" + err.Error()
			} else if err := utils.Clipullimage(node.Mac.Adminurl, node.Pro.Git, node_ver, message); err != nil {
				//客户端下载镜像
				message <- "客户端下载镜像失败,error:" + err.Error()
			} else if node_docid, createerr := utils.Clicreatecon(node.Mac.Adminurl, node.Port, node_ver, node.Pro.Git); createerr != nil {
				message <- "客户端创建镜像失败,error:" + createerr.Error()
			} else if err := utils.Clistartcon(node.Mac.Adminurl, node_docid); err != nil {
				node.DocId = node_docid
				node.CurVer = node_ver
				if err := node.Update(); err != nil {
					beego.Error("动作:数据库操作，添加节点版本失败：", err)
					//如果数据操作错误，就需要把刚刚创建 的容器删除了， 不然下次再次部署时，提示已有相同容器
					if err := utils.Clidelcon(node.Mac.Adminurl, node.DocId); err != nil {
						message <- "回滚事件删除容器error:" + err.Error()
					}
				}
				message <- "error:" + err.Error()
			} else {
				node.DocId = node_docid
				node.CurVer = node_ver
				if err := node.Update(); err != nil {
					beego.Error("动作:数据库操作，添加节点版本失败：", err)
					//如果数据操作错误，就需要把刚刚创建 的容器删除了， 不然下次再次部署时，提示已有相同容器
					if err := utils.Clidelcon(node.Mac.Adminurl, node.DocId); err != nil {
						message <- "回滚事件删除容器error:" + err.Error()
					}
				}
				message <- "success"
			}

		}()
		node.CurVer = node_ver
		for {
			contron := false
			select {
			case mes := <-message:
				event := utils.Event{
					Type:    utils.EVENT_MESSAGE,
					Message: mes,
				}
				data, err := json.Marshal(event)
				if err != nil {
					beego.Error("序列化EVENT TO json 出错:", err)
				}
				if ws.WriteMessage(websocket.TextMessage, data) != nil {
				}
				if strings.Contains(mes, "error") {
					beego.Debug("处理结束err")
					contron = true
				}
				if mes == "success" {

					beego.Debug("处理结束")
					contron = true
				}
			}
			if contron {
				beego.Error("message 循环已退出0")
				break
			}
		}
		beego.Info("message 循环已退出00")
	}
	beego.Info("websocket 函数体退出")
}

func (c *NodeController) Jnodeopera() {
	uid := c.GetSession("uid").(uint)
	c.Data["uid"] = uid
	env_sign := c.GetString("env_sign")
	if !mper.Isuserper(strings.ToUpper(env_sign), uid) {
		beego.Debug("动作:请求操作节点,环境验证失败")
		c.Abort("503")
	}
	if !mper.Isperitem("node", uid) { //节点删除(nodd)
		beego.Debug("动作:请求查看节点,权限验证失败")
		c.Abort("503")
	}
	node_id, _ := c.GetInt("id")
	signal := c.GetString("signal")
	node := mpro.Node{
		Id: uint(node_id),
	}
	if err := node.Read(); err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]string{"message": "error", "data": err.Error()}
		c.ServeJSON()
		return
	}
	if err := node.Mac.Read(); err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]string{"message": "error", "data": err.Error()}
		c.ServeJSON()
		return
	}
	if node.DocId == "" {
		beego.Error("容器ID为空")
		c.Data["json"] = map[string]string{"message": "error", "data": "容器ID为空"}
		c.ServeJSON()
		return
	}

	if signal == "START" {
		if err := utils.Clistartcon(node.Mac.Adminurl, node.DocId); err != nil {
			c.Data["json"] = map[string]string{"message": "error", "data": err.Error()}
		} else {
			c.Data["json"] = map[string]string{"message": "success", "data": "成功"}
		}
	} else if signal == "STOP" {
		if err := utils.Clistopcon(node.Mac.Adminurl, node.DocId); err != nil {
			c.Data["json"] = map[string]string{"message": "error", "data": err.Error()}
		} else {
			c.Data["json"] = map[string]string{"message": "success", "data": "成功"}
		}
	} else if signal == "RESTART" {
		if err := utils.Clistartcon(node.Mac.Adminurl, node.DocId); err != nil {
			c.Data["json"] = map[string]string{"message": "error", "data": err.Error()}
		} else {
			c.Data["json"] = map[string]string{"message": "success", "data": "成功"}
		}
	}
	c.ServeJSON()
	return
}

func (c *NodeController) Jstopnode() {

}

func (c *NodeController) Jrestartnode() {

}
