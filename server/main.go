package main

import (
	"Minimalist_TikTok/server/cache"
	"Minimalist_TikTok/server/config"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
	"strings"
)

const month = 60 * 60 * 24 * 30 //30天过期

type SendMsg struct {
	Type    int    `json:"type"`
	Content string `json:"content"`
}

type ReplyMsg struct {
	From    string `json:"from"`
	Code    int    `json:"code"`
	Content string `json:"content"`
}

type Client struct {
	ID     string
	SendID string
	Socket *websocket.Conn
	Send   chan []byte
}

type Broadcast struct {
	Client  *Client
	Message []byte
	Type    int
}

type ClientManager struct {
	Clients    map[string]*Client
	Broadcast  chan *Broadcast
	Reply      chan *Client
	Register   chan *Client
	Unregister chan *Client
}

type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

var Manager = ClientManager{
	Clients:    make(map[string]*Client), // 参与连接的用户，出于性能的考虑，需要设置最大连接数
	Broadcast:  make(chan *Broadcast),
	Register:   make(chan *Client),
	Reply:      make(chan *Client),
	Unregister: make(chan *Client),
}

func CreateID(uid, toUid string) string {
	var byte_buf bytes.Buffer

	byte_buf.WriteString(uid)
	byte_buf.WriteString(".")
	byte_buf.WriteString(toUid)
	return byte_buf.String()
}
func WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { // CheckOrigin解决跨域问题
			return true
		}}).Upgrade(w, r, nil) // 升级成ws协议
	if err != nil {
		http.NotFound(w, r)
		return
	}
	var s string
	//for  {
	m, p, _ := conn.ReadMessage()
	//if e!=nil {
	//	break
	//}
	s = string(p)
	fmt.Println(m, string(p))
	//	defer conn.Close()
	//}
	arr := strings.Split(s, "+")
	fmt.Println("arr", arr)
	uid := arr[0]
	toUid := arr[1]
	content := arr[2]
	iden := arr[3]
	if iden == "0" {
		//创建一个用户实例
		client := &Client{
			ID:     CreateID(uid, toUid),
			SendID: CreateID(toUid, uid),
			Socket: conn,
			Send:   make(chan []byte),
		}
		// 用户注册到用户管理上
		Manager.Register <- client
		go client.Read(content, iden)
		go client.Write()
	} else if iden == "1" {
		client := &Client{
			ID:     CreateID(uid, toUid),
			SendID: CreateID(toUid, uid),
			Socket: conn,
			Send:   make(chan []byte),
		}
		// 用户注册到用户管理上
		Manager.Register <- client
		go client.Read(content, iden)
		go client.Write()
	} else {

	}
}

func (c *Client) Read(content string, iden string) {
	fmt.Println("reading")
	//ctx := context.Background()
	defer func() {
		Manager.Unregister <- c
		_ = c.Socket.Close()
	}()
	if iden == "0" {
		var send []byte = []byte(content)
		Manager.Broadcast <- &Broadcast{
			Client:  c,
			Message: send,
			//Message: []byte(sendMsg.Content), //发送过来的消息
		}
	} else if iden == "1" {
		timeT := 99999
		//}
		result, _ := cache.FindMany(config.MongoDBName, c.SendID, c.ID, int64(timeT), 10) //获取10条历史消息
		fmt.Println(c.SendID, c.ID)
		//fmt.Println("result",result)
		if len(result) > 10 {
			result = result[:10]
		} else if len(result) == 0 {
			replyMsg := &ReplyMsg{
				Code:    1,
				Content: "到底了",
			}
			msg, _ := json.Marshal(replyMsg)
			_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
		}
		replyMsg := ReplyMsg{
			From:    "guess",
			Content: strconv.Itoa(len(result)),
		}
		msg, _ := json.Marshal(replyMsg)
		_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
		for _, result := range result {
			replyMsg := ReplyMsg{
				From:    result.From,
				Content: result.Msg,
			}
			msg, _ := json.Marshal(replyMsg)
			_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
		}
	}
	//timeT,err := strconv.Atoi(sendMsg.Content)
	//if err != nil {

	//for {
	//c.Socket.PongHandler()
	//sendMsg := new(SendMsg)
	////c.Socket.ReadMessage() 不是json格式
	//err := c.Socket.ReadJSON(&sendMsg) //是json格式
	//if err != nil {
	//	fmt.Println("数据格式不正确", err)
	//	Manager.Unregister <- c
	//	_ = c.Socket.Close()
	//	break
	//}
	//if sendMsg.Type==1{
	//	如果传过来的type=1的话（接受消息），那么我们就可以先去redis上面查询一下当前有多少人进行了连接。
	//r1, _ := cache.RedisClient.Get(con, c.ID).Result()
	//r2, _ := cache.RedisClient.Get(con, c.SendID).Result()
	////if r1 > "3" && r2 == "" {
	//	//	1给2发，一直不回
	//	replyMsg := &ReplyMsg{
	//		Code:    1,
	//		Content: "别发了没回",
	//	}
	//	msg, _ := json.Marshal(replyMsg) //序列化
	//	_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
	//	continue
	//}else {
	//如果没有的话，就先记录到redis中进行缓存
	//cache.RedisClient.Incr(con, c.ID)
	//_, _ = cache.RedisClient.Expire(con, c.ID, time.Hour*24*30*3).Result()
	//	防止过快分手，建立连接三个月过期
	//}
	//广播消息
	//。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。
	//var send []byte = []byte(content)
	//Manager.Broadcast <- &Broadcast{
	//	Client:  c,
	//	Message :send,
	//	//Message: []byte(sendMsg.Content), //发送过来的消息
	//}
	//	。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。
	//timeT,err := strconv.Atoi(sendMsg.Content)
	//if err != nil {
	//	timeT = 99999
	//}
	//result,_ := cache.FindMany(config.MongoDBName,c.SendID,c.ID,int64(timeT),10)//获取10条历史消息
	//fmt.Println(c.SendID,c.ID)
	//if len(result)>10 {
	//	result=result[:10]
	//}else if len(result)==0 {
	//		replyMsg := &ReplyMsg{
	//			Code:    1,
	//			Content: "到底了",
	//		}
	//		msg,_ := json.Marshal(replyMsg)
	//		_=c.Socket.WriteMessage(websocket.TextMessage,msg)
	//		continue
	//}
	//for _,result := range result{
	//	replyMsg := ReplyMsg{
	//		From: result.From,
	//		Content: result.Msg,
	//	}
	//	msg,_ := json.Marshal(replyMsg)
	//	_=c.Socket.WriteMessage(websocket.TextMessage,msg)
	//	}
	//}
}

func (c *Client) Write() {
	defer func() {
		_ = c.Socket.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				_ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			replyMsg := ReplyMsg{
				Code:    0,
				Content: fmt.Sprintf("%s", string(message)),
			}
			msg, _ := json.Marshal(replyMsg)
			_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
		}
	}
}

func main() {
	config.Init()
	go Manager.Start()
	http.HandleFunc("/", WsHandler)
	http.ListenAndServe(":8888", nil)
}
