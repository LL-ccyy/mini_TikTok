package model

type Message struct {
	Id         int64  `json:"id,omitempty"`
	FromUser   User   `json:"from_user"`
	FromUserID uint   `json:"from_user_id"`
	ToUser     User   `json:"to_user"`
	ToUserID   uint   `json:"to_user_id"`
	Content    string `json:"content,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
}
