package controller

import (
	"Minimalist_TikTok/config"
	"Minimalist_TikTok/serializer"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type User struct {
	gorm.Model
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Password      string `json:"password,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

type Message struct {
	Id         int64  `json:"id,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
}

type MessageSendEvent struct {
	UserId     int64  `json:"user_id,omitempty"`
	ToUserId   int64  `json:"to_user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

type MessagePushEvent struct {
	FromUserId int64  `json:"user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

func ErrorResponse(err error) serializer.ErrorResponse {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := config.T(fmt.Sprintf("Field.%s", e.Field))
			tag := config.T(fmt.Sprintf("Tag.Valid.%s", e.Tag))
			return serializer.ErrorResponse{
				StatusCode: 40001,
				StatusMsg:  fmt.Sprintf("%s%s", field, tag),
				Error:      fmt.Sprint(err),
			}
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.ErrorResponse{
			StatusCode: 40001,
			StatusMsg:  "JSON类型不匹配",
			Error:      fmt.Sprint(err),
		}
	}
	return serializer.ErrorResponse{
		StatusCode: 40001,
		StatusMsg:  "参数错误",
		Error:      fmt.Sprint(err),
	}
}
