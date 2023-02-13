package main

import (
	"Minimalist_TikTok/pkg/util"
	"Minimalist_TikTok/server/cache"
	"Minimalist_TikTok/server/config"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

func (manager *ClientManager) Start() {
	for {
		fmt.Println("----监听管道通信----")
		select {
		case conn := <-Manager.Register:
			fmt.Println("有新连接：%s", conn.ID)
			Manager.Clients[conn.ID] = conn //把连接放在用户管理上
			replyMsg := &ReplyMsg{
				Code:    0,
				Content: "已经连接到服务器了",
			}
			msg, _ := json.Marshal(replyMsg)
			_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)
		case conn := <-Manager.Unregister:
			fmt.Println("连接失败：%s", conn.ID)
			if _, ok := Manager.Clients[conn.ID]; ok {
				replyMsg := &ReplyMsg{
					Code:    0,
					Content: "连接中断",
				}
				msg, _ := json.Marshal(replyMsg)
				_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)
				close(conn.Send)
				delete(Manager.Clients, conn.ID)
			}
		case broadcast := <-Manager.Broadcast: //1->2
			message := broadcast.Message
			sendId := broadcast.Client.SendID //2->1
			flag := false                     //默认对方是不在线的
			for id, conn := range Manager.Clients {
				if id != sendId {
					continue
				}
				select {
				case conn.Send <- message:
					flag = true
				default:
					close(conn.Send)
					delete(Manager.Clients, conn.ID)
				}
			}
			id := broadcast.Client.ID //1->2
			if flag {
				replyMsg := &ReplyMsg{
					Code:    0,
					Content: "对方在线应答",
				}
				msg, _ := json.Marshal(replyMsg)
				_ = broadcast.Client.Socket.WriteMessage(websocket.TextMessage, msg)
				//cache.RedisInsert(cache.RedisDbName,id,string(message),1,int64(3*month))
				err := cache.InsertMsg(config.MongoDBName, id, string(message), 1, int64(3*month))
				if err != nil {
					util.LogrusObj.Info(err)
					fmt.Println("插入mongo错误", err)
				}
			} else {
				fmt.Println("对方不在线")
				reply := &ReplyMsg{
					Code:    0,
					Content: "对方不在线",
				}
				msg, _ := json.Marshal(reply)
				_ = broadcast.Client.Socket.WriteMessage(websocket.TextMessage, msg)
				//cache.RedisInsert(cache.RedisDbName,id,string(message),0,int64(3*month))
				err := cache.InsertMsg(config.MongoDBName, id, string(message), 1, int64(3*month))
				if err != nil {
					util.LogrusObj.Info(err)
					fmt.Println("插入mongo错误", err)
				}
			}
		}
	}
}
