package service

import (
	"Minimalist_TikTok/model"
	"Minimalist_TikTok/pkg/util"
	"Minimalist_TikTok/serializer"
	"fmt"
	"github.com/gorilla/websocket"
	"strconv"
	"strings"
	"time"
)

type MessageActionService struct {
	Token      string `form:"token",json:"token"`
	ToUserID   string `form:"to_user_id",json:"to_user_id"`
	ActionType string `form:"action_type",json:"action_type"`
	Content    string `form:"content",json:"content"`
}

type MessageChatService struct {
	Token    string `form:"token",json:"token"`
	ToUserID string `form:"to_user_id",json:"to_user_id"`
}

func (service *MessageActionService) MessageAction() serializer.Response {
	token := service.Token
	claims, err := util.ParseToken(token)
	if err != nil {
		util.LogrusObj.Info(err)
		return serializer.Response{
			StatusCode: 1,
			StatusMsg:  "token错误",
		}
	}
	toUid := service.ToUserID
	uid := strconv.Itoa(int(claims.Id))
	content := service.Content
	action_type := service.ActionType
	dl := websocket.Dialer{}
	conn, _, err := dl.Dial("ws://localhost:8888", nil)
	if err != nil {
		util.LogrusObj.Info(err)
		return serializer.Response{
			StatusCode: 1,
			StatusMsg:  "lianjie错误",
		}
	}
	if action_type == "1" {
		send := uid + "+" + toUid + "+" + content + "+" + "0"
		conn.WriteMessage(websocket.TextMessage, []byte(send))
		return serializer.Response{
			StatusCode: 0,
			StatusMsg:  "ok",
		}
	} else {
		return serializer.Response{
			StatusCode: 1,
			StatusMsg:  "err",
		}
	}
}

func (service *MessageChatService) MessageChat() serializer.ChatRecordResponse {
	token := service.Token
	claims, err := util.ParseToken(token)
	if err != nil {
		util.LogrusObj.Info(err)
		return serializer.ChatRecordResponse{
			StatusCode: 1,
			StatusMsg:  "token错误",
		}
	}
	toUid := service.ToUserID
	uid := strconv.Itoa(int(claims.Id))
	var from_user model.User
	var to_user model.User
	model.DB.Model(&model.User{}).Where("id=?", uid).Find(&from_user)
	model.DB.Model(&model.User{}).Where("id=?", toUid).Find(&to_user)
	dl := websocket.Dialer{}
	conn, _, err := dl.Dial("ws://localhost:8888", nil)
	if err != nil {
		util.LogrusObj.Info(err)
		return serializer.ChatRecordResponse{
			StatusCode: 1,
			StatusMsg:  "lianjie错误",
		}
	}
	send := uid + "+" + toUid + "+" + "ccyy" + "+" + "1"
	//fmt.Println(send)
	conn.WriteMessage(websocket.TextMessage, []byte(send))
	//连接服务器提示
	_, _, _ = conn.ReadMessage()
	//s=string(p)
	_, s, _ := conn.ReadMessage()
	//s=string(p)
	fmt.Println("一共n条消息,n=", string(s))
	arr := strings.Split(string(s), "\"")
	fmt.Println("arr", arr[9])
	num, err := strconv.Atoi(arr[9])
	from_user_id, _ := strconv.Atoi(uid)
	to_user_id, _ := strconv.Atoi(toUid)
	var message model.Message
	var messagelist []model.Message
	for i := 0; i < num; i++ {
		_, s, _ := conn.ReadMessage()
		fmt.Println("收到的", string(s))
		arr := strings.Split(string(s), "\"")
		from := arr[3]
		fmt.Println("msg:", arr[9][1:len(arr[9])-1])
		msg := arr[9][1 : len(arr[9])-1]
		sep := strings.Split(string(msg), " ")
		tim, _ := strconv.Atoi(sep[2])
		t := time.Unix(int64(tim), 0)
		message.Id = int64(i + 1)
		message.CreateTime = t.Format("03:04 PM")
		//message.CreateTime=sep[2]
		message.Content = sep[0]
		if from == "you" {
			message.FromUser = to_user
			message.ToUser = from_user
			message.FromUserID = uint(to_user_id)
			message.ToUserID = uint(from_user_id)
		} else {
			message.FromUser = from_user
			message.ToUser = to_user
			message.FromUserID = uint(from_user_id)
			message.ToUserID = uint(to_user_id)
		}
		messagelist = append(messagelist, message)
	}
	fmt.Println()
	return serializer.ChatRecordResponse{
		StatusCode:  0,
		StatusMsg:   "",
		MessageList: messagelist,
	}
}
