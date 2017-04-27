package node

import (
	"bytes"
	"encoding/json"
	"errors"
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

	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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
	//获取某一环境下所有主机
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
			c.Data["json"] = map[string]interface{}{"message": "error", "error": "提交数据不完整"}
			c.ServeJSON()
			return
		}
		mach := &mac.Machine{Id: uint(n_machine)}
		if err := mach.Read(); err != nil {
			c.Data["json"] = map[string]interface{}{"message": "error", "error": "数据库出错:" + err.Error()}
			c.ServeJSON()
			return
		} else { //检查该主机是否已有该端口存在
			if ok, ports := mach.Addport(utils.Getmacports(n_port)); !ok {
				c.Data["json"] = map[string]interface{}{"message": "error", "error": fmt.Sprintf("%s", ports) + " prot already exists"}
				c.ServeJSON()
				return
			}
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
		if node.Port != node_port {
			//如端口变动，就先删除之前在主机上面绑定端口的记录,再添加现在请求主机的绑定记录
			node.Mac.Delport(utils.Getmacports(node.Port))
			//检查该主机是否已有该端口存在
			if ok, ports := node.Mac.Addport(utils.Getmacports(node_port)); !ok {
				c.Data["json"] = map[string]interface{}{"message": "error", "error": fmt.Sprintf("%s", ports) + " prot already exists"}
				c.ServeJSON()
				return
			}
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
	beego.Info(&nodes)
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
	node := &mpro.Node{
		Id: uint(node_id),
	}
	if err := node.Read(); err != nil {
		beego.Error("动作:数据库操作,查询节点信息失败:", err)
	}
	if err := node.Mac.Read(); err != nil {
		beego.Error("动作:数据库操作,查询节点所属主机信息失败:", err)
	}
	if err := utils.Clidelcon(node.Mac.Adminurl, node.DocId); err != nil {
		c.Data["json"] = map[string]interface{}{"message": "error", "error": err.Error()}
		c.ServeJSON()
		return
	}
	if err := node.Delete(); err != nil {
		c.Data["json"] = map[string]interface{}{"message": "error", "error": err.Error()}
		c.ServeJSON()
		return
	} else {
		utils.Delnode(node.Mac.Id, node)
		beego.Info("删除node:", &node, "在主机:", &node.Mac)
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
			message <- "error"
			message <- err.Error()
		} else if err := utils.Gitpull(pro.Git, message); err != nil {
			message <- err.Error()
			message <- "error"
		} else {
			message <- "success"
		}
	}()
	beego.Info("websocket 接收到(node):", node)
	defer ws.Close()
	forcontr := false
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
				utils.Addnode(node.Mac.Id, node)
				beego.Debug("处理结束")
				forcontr = true
			}
			if mes == "error" {
				ws.Close()
				beego.Debug("处理结束")
				forcontr = true
			}
		} //end select
		if forcontr {
			break
		}

	} //end for
	beego.Info("message 循环已退出")
}

func (c *NodeController) Wsdeploy() {
	ws, err := websocket.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil, 1024, 1024)
	proid, interr := c.GetInt("proid")
	if interr != nil || proid == 0 {
		beego.Error("从URL地址获取项目ID出错,proid:", proid, ",error:", interr)
	}
	t := time.Now().UnixNano()
	md5path := utils.Md5(fmt.Sprintf("%d", t)) //为保持同一个项目部署的交叉执行,每一个生成不同的目录操作
	md5path = md5path[:10]
	var message = make(chan string, 2)
	var event = make(chan utils.Event, 2)
	var nodeupdate = make(chan string) //节点的容器ID更改后的容器ID
	defer ws.Close()
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(c.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup Websocket connection:", err)
		return
	}
	var exitstats = make(chan bool)
	var updatestats = make(chan bool)
	go func() { //循环检测节点容器状态
		var nodes []*mpro.Node
		if _, err := new(mpro.Node).Query().Filter("Pro__Id", proid).All(&nodes); err != nil {
			beego.Error("查询数据库出错", err)
		} else if len(nodes) > 0 {
			var tmp_nodes []*mpro.Node
			for _, node := range nodes {
				if node.DocId == "" {
					continue
				}
				if err := node.Mac.Read(); err != nil {
					beego.Error("查询节点所属主机出错", err)
				}
				tmp_nodes = append(tmp_nodes, node)
			}
		DONE:
			for {

				select {
				case code := <-exitstats:
					if code {
						beego.Debug("网页关闭状态检测")
						break DONE
					}
				case node_update_docid := <-nodeupdate:
					beego.Debug("容器id改变...")
					event <- utils.Event{
						Type:        utils.EVENT_UPDATE_NODE,
						Nodeid:      strings.Split(node_update_docid, "-")[1],
						Containerid: strings.Split(node_update_docid, "-")[0],
					}

					data, err := json.Marshal(event)
					if err != nil {
						beego.Error("Fail to marshal event:", err)
						continue
					}
					if err := ws.WriteMessage(websocket.TextMessage, data); err != nil {
						beego.Error("websocket 写出错:", err)
					}
				case isupdate := <-updatestats:
					if isupdate {
						for _, v := range tmp_nodes {
							beego.Debug(v)
						}
						beego.Debug("检测到更新容器")

						tmp_nodes = nil
						if _, err := new(mpro.Node).Query().Filter("Pro__Id", proid).All(&nodes); err != nil {
							beego.Error("查询数据库出错", err)
						} else if len(nodes) > 0 {
							for _, node := range nodes {
								if node.DocId == "" {
									continue
								}
								if err := node.Mac.Read(); err != nil {
									beego.Error("查询节点所属主机出错", err)
								}
								tmp_nodes = append(tmp_nodes, node)
							}
						}
						for _, v := range tmp_nodes {
							beego.Debug(v)
						}
					}
				default:
					events, err := utils.Cliinspectcon(tmp_nodes)
					if err != nil {
						mes := "检测节点状态出错"
						message <- mes
						//						var buf bytes.Buffer
						//						buf.WriteString(mes)
						//						if err := ws.WriteMessage(websocket.TextMessage, buf.Bytes()); err != nil {
						//							beego.Error("发送websocket出错:", err)
						//						}
						break DONE
					}
					beego.Debug("events.len:", len(events))

					for _, eve := range events {
						event <- *eve
						//						data, err := json.Marshal(event)
						//						if err != nil {
						//							beego.Error("Fail to marshal event:", err)
						//							continue
						//						}

						//						if ws.WriteMessage(websocket.TextMessage, data) != nil {
						//						}
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
		mess := string(tmp_ms[:len(tmp_ms)])
		beego.Info("message to string:", mess)
		node_id := strings.Split(mess, "-")[0]
		var node_ver []string
		for k, v := range strings.Split(mess, "-") {
			if k == 0 {
				continue
			}
			node_ver = append(node_ver, v)
		}
		var ms int
		if node_id != "" && node_id != "0" {
			ms, _ = strconv.Atoi(node_id)
		} else {
			beego.Error("节点ID:\"" + node_id + "\"解析出错")
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
			var tmp_node_ver []string
			for _, chil_ver := range strings.Split(node.CurVer, "-") {
				if chil_ver == "" {
					continue
				}
				for _, goto_chil_ver := range node_ver {
					if chil_ver == goto_chil_ver {
						tmp_node_ver = append(tmp_node_ver, goto_chil_ver)
					}
				}
			}
			beego.Debug("len(tmp_node_ver):", len(tmp_node_ver), "tmp_node_ver:", tmp_node_ver)
			beego.Debug("len(node_ver):", len(node_ver), "node_ver", node_ver)
			if len(tmp_node_ver) != len(node_ver) && node.CurVer != "" {

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
				var jsonfile = make(map[string]string)
				for _, v := range oldverconf {
					jsonfile[fmt.Sprintf("%s", v[0])] = fmt.Sprintf("%s", v[1])
				}
				if err := utils.Makejsonconf(node.Pro.Sign, node.Mac.Env.Sign, node.CurVer, jsonfile); err != nil {
					message <- "生成老版本配置失败, error:" + err.Error()
				} else {
					message <- "生成上个版本" + node.CurVer + "配置文件成功"
				}
			}

		}()
		go func() { //部署容器
			var comperr error
			comperr = nil
			contr := true
			//copy 项目到活动目录 .code
			//本次活动目录
			path := utils.EXECPATH
			gitpath := utils.Gittoname(node.Pro.Git)
			var pro_path = make([]string, len(gitpath))
			var tmp_pro_path = make([]string, len(gitpath))
			for k, v := range gitpath {
				pro_path[k] = path + utils.PD + "code" + utils.PD + v
				tmp_pro_path[k] = path + utils.PD + ".code" + utils.PD + md5path + utils.PD + v
			}
			beego.Debug("tmp_pro_path:", tmp_pro_path)
			beego.Debug("pro_path:", pro_path)

			if ok := utils.Gitchecver(node.Pro.Git, node_ver, message); !ok {
				comperr = errors.New("仓库切换tag 出错")
			}
			for k, v := range pro_path {
				if err := utils.Copypath(v, tmp_pro_path[k], ".git"); err != nil {
					message <- "copy 临时目录出错:" + err.Error()
					comperr = errors.New("copy 临时目录出错:" + err.Error())
				} else if err := utils.Compilecode(node.Pro.Compile, node.Pro.Compilever, tmp_pro_path[k], message); err != nil {
					message <- "编译代码失败, error:" + err.Error()
					comperr = errors.New("编译代码失败, error:" + err.Error())
				}
			}

			if comperr != nil {
				message <- comperr.Error()
			} else if err := utils.BuildImage(&node, node_ver, tmp_pro_path, node.Pro.Insfile, message); err != nil { //BuildImage
				message <- "镜像BUILD失败,error:" + err.Error()
				//build失败需要删除<none>镜像
				utils.DelNoneImage("", message)
			} else if err := utils.PushImages(node.Pro.Git, node_ver, message); err != nil {
				//上传镜像
				message <- "上传镜像失败,error:" + err.Error()
			} else if err := utils.Clipullimage(node.Mac.Adminurl, node.Pro.Git, utils.VerlisttoString(node_ver), message); err != nil {
				//客户端下载镜像
				message <- "客户端下载镜像失败,error:" + err.Error()
			} else if err := utils.Clidelcon(node.Mac.Adminurl, node.DocId); err != nil {
				message <- "Info:" + err.Error()
			} else if node_docid, createerr := utils.Clicreatecon(node.Mac.Adminurl, node.Port, utils.VerlisttoString(node_ver), node.Pro.Git, node.Mac.Env.Sign, node.Pro.Sharedpath, node.Pro.Dns); createerr != nil {
				message <- "客户端创建镜像失败,error:" + createerr.Error()
			} else if err := utils.Clistartcon(node.Mac.Adminurl, node_docid); err != nil {
				nodeupdate <- node_docid + "-" + strconv.FormatUint(uint64(node.Id), 10)
				node.DocId = node_docid
				node.CurVer = utils.VerlisttoString(node_ver)
				if err := node.Update(); err != nil {
					beego.Error("动作:数据库操作，添加节点版本失败：", err)
					//如果数据操作错误，就需要把刚刚创建 的容器删除了， 不然下次再次部署时，提示已有相同容器
					if err := utils.Clidelcon(node.Mac.Adminurl, node.DocId); err != nil {
						message <- "回滚事件删除容器error:" + err.Error()
					}
				} else {
					utils.Updatenode(node.Mac.Id, &node)
				}
				message <- "error:" + err.Error()
			} else {
				nodeupdate <- node_docid + "-" + strconv.FormatUint(uint64(node.Id), 10)
				message <- "容器已启动"
				node.DocId = node_docid
				node.CurVer = utils.VerlisttoString(node_ver)
				if err := node.Update(); err != nil {
					message <- "动作:数据库操作，添加节点版本失败：" + err.Error()
					beego.Error("动作:数据库操作，添加节点版本失败：", err)
					//如果数据操作错误，就需要把刚刚创建 的容器删除了， 不然下次再次部署时，提示已有相同容器
					if err := utils.Clidelcon(node.Mac.Adminurl, node.DocId); err != nil {
						message <- "回滚事件删除容器error:" + err.Error()
					}
				} else {
					utils.Updatenode(node.Mac.Id, &node)
					message <- "部署成功"
					message <- "success"
					contr = false
				}
			}
			md5pathindex := strings.LastIndex(tmp_pro_path[0], utils.PD)
			if err := utils.Deployover(tmp_pro_path[0][:md5pathindex]); err != nil {
				message <- "删除临时项目代码出错：" + err.Error() + ",请联系管理员手动删除"
				contr = true
			}
			if contr {
				message <- "error"
			}

		}()
		for {
			contron := false
			select {
			case data_event := <-event:
				data, err := json.Marshal(data_event)
				if err != nil {
					beego.Error("序列化EVENT TO json 出错:", err)
				}
				if ws.WriteMessage(websocket.TextMessage, data) != nil {
				}
			case mes := <-message:
				event <- utils.Event{
					Type:    utils.EVENT_MESSAGE,
					Message: mes,
				}

				if strings.Contains(mes, "error") {
					updatestats <- true
					beego.Debug("处理结束err")
					contron = true
				}
				if mes == "success" {
					updatestats <- true
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
	} //end for ws.ReadMessage()
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
		if err := utils.Clirestartcon(node.Mac.Adminurl, node.DocId); err != nil {
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
