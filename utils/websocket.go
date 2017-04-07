package utils

import (
	"container/list"
	"encoding/json"
	mac "lvbu/models/machine"
	mpro "lvbu/models/project"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

type Subscriber struct {
	ID    string
	Type  EventType
	Proid string
	Sign  string
	Conn  *websocket.Conn // Only for WebSocket users; otherwise nil.
}

var (
	subscribers = list.New()
	hosts       = make(map[string][]*mac.Machine)
	nodes       = make(map[uint][]*mpro.Node)
)

func Join(id string, ws *websocket.Conn) {
	sub := Subscriber{ID: id, Conn: ws}
	beego.Debug("Join list.len():", subscribers.Len())
	if !isUserExist(subscribers, sub.ID) {
		subscribers.PushBack(sub) // Add user to the end of list.
		// Publish a JOIN event.
		beego.Info("新加入")
	} else {
		beego.Info("已加入")
	}
}

func Addhost(envsign string, host *mac.Machine) {
	hostlist := hosts[envsign]
	if hostlist == nil {
		hosts[envsign] = []*mac.Machine{host}
	} else {
		hosts[envsign] = append(hosts[envsign], host)
	}
}

func Updatehost(envsign string, host *mac.Machine) {
	hostlist := hosts[envsign]
	if hostlist == nil {
		hosts[envsign] = []*mac.Machine{host}
	} else {
		for _, v := range hosts[envsign] {
			if v.Id == host.Id {
				v.Adminurl = host.Adminurl
				v.Status = host.Status
				break
			}
		}
	}
}

func Delhost(envsign string, host *mac.Machine) {
	hostlist := hosts[envsign]
	if hostlist == nil {
		hosts[envsign] = []*mac.Machine{host}
	} else {
		for _, v := range hosts[envsign] {
			if v.Id == host.Id {
				v = nil
				break
			}
		}
	}
}

func Addnode(macid uint, node *mpro.Node) {
	nodelist := nodes[macid]
	if nodelist == nil {
		nodes[macid] = []*mpro.Node{node}
	} else {
		nodes[macid] = append(nodes[macid], node)
	}
}

func Updatenode(macid uint, node *mpro.Node) {
	nodelist := nodes[macid]
	if nodelist == nil {
		nodes[macid] = []*mpro.Node{node}
	} else {
		for _, v := range nodes[macid] {
			if v.Id == node.Id {
				v.Pro = node.Pro
				v.DocId = node.DocId
				break
			}
		}
	}
}

func Delnode(macid uint, node *mpro.Node) {
	nodelist := nodes[macid]
	if nodelist == nil {
		nodes[macid] = []*mpro.Node{node}
	} else {
		for _, v := range nodes[macid] {
			if v.Id == node.Id {
				v = nil
				break
			}
		}
	}
}
func Leave(id string) {
	for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
		if sub.Value.(Subscriber).ID == id {
			subscribers.Remove(sub)
			// Clone connection.
			ws := sub.Value.(Subscriber).Conn
			if ws != nil {
				ws.Close()
				beego.Error("WebSocket closed:", id)
			}
			break
		}
	}
	beego.Debug("list,len():", subscribers.Len())
}

func UpdateSubSign(id, sign string) {
	for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
		if sub.Value.(Subscriber).ID == id {
			subscribers.Remove(sub)
			subb := Subscriber{ID: id, Sign: sign, Conn: sub.Value.(Subscriber).Conn}
			subscribers.PushBack(subb)
			break
		}
	}
}

func ProDetection() {
	time.Sleep(time.Second * 5)
	inithostsandnodes()
	for {
		for _, v := range []string{"DE", "QE", "OE"} {
			if len(hosts[v]) == 0 {
				beego.Info("环境Id:", v, "没有主机")
				time.Sleep(time.Second * 1)
				break
			}
			pros := make(map[uint]map[string]int) //map[Pro_id]map["all"|"running"] len(nodes) | len(node_running)
			for _, mac := range hosts[v] {
				if mac == nil {
					continue
				}
				nodelist := nodes[mac.Id]
				if len(nodelist) == 0 {
					//					beego.Info("主机Id", mac.Id, "没有节点")
					continue
				}
				containers, err := Clilistcons(mac.Adminurl)
				if err != nil {
					beego.Info("查询主机Id", mac.Id, "出错", err)
					continue
				}
				for _, container := range containers {
					for _, node := range nodelist {
						mm, ok := pros[node.Pro.Id]
						if !ok {
							mm = make(map[string]int)
						}

						if node.DocId == "" {
							//beego.Debug("节点Id:", node.Id, "未初始化")
							mm["all"] = mm["all"] + 1
							continue
						}
						if container.ID == node.DocId {
							mm["all"] = mm["all"] + 1
							if container.State == "running" {
								mm["running"] = mm["running"] + 1
							} else if strings.Contains(container.Status, "Up") {
								mm["running"] = mm["running"] + 1
							}
						}
						pros[node.Pro.Id] = mm
						//						node_pros[node.Pro.Id] = m
					} //end for nodes
				} // end for containers
			} // end for macs
			var events []Event
			for key, value := range pros {

				proid := strconv.FormatUint(uint64(key), 10)
				all := value["all"]
				run := value["running"]
				event := Event{
					Type:        EVENT_PRO_STAT,
					Envsign:     v,
					Proid:       proid,
					Proallnodes: all,
					Prorunnodes: run,
				}
				events = append(events, event)

			}

			for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
				//				if sub.Value.(Subscriber).Sign != v {
				//					continue
				//				}
				ws := sub.Value.(Subscriber).Conn
				if ws != nil {
					data, err := json.Marshal(events)
					if err != nil {
						beego.Error("Fail to marshal event:", err)
						return
					}
					//beego.Debug("ws send :", websocket.TextMessage, sub.Value.(Subscriber).ID)
					if ws.WriteMessage(websocket.TextMessage, data) != nil {
						// User disconnected.
						Leave(sub.Value.(Subscriber).ID)
						beego.Debug("ws send err, always leave id:", sub.Value.(Subscriber).ID)
					}
				} else {
					Leave(sub.Value.(Subscriber).ID)
					beego.Debug("ws nil, always leave id:", sub.Value.(Subscriber).ID)
				}
			}
			time.Sleep(time.Second * 2)
		}

	}

}

func inithostsandnodes() {
	for _, v := range []string{"DE", "QE", "OE"} {
		macs := new(mac.Machine).GetMacforenv(v)
		hosts[v] = macs
		for _, mac := range macs {
			nodelist := new(mpro.Node).GetNodeformac(mac.Id)
			if len(nodelist) == 0 {
				//					beego.Info("主机Id", mac.Id, "没有节点")
				continue
			}
			nodes[mac.Id] = nodelist
		}
	}
}

func isUserExist(subscribers *list.List, id string) bool {
	for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
		if sub.Value.(Subscriber).ID == id {
			return true
		}
	}
	return false
}
