package serializer

import "Minimalist_TikTok/model"

// Response 基础序列化器
// 投稿接口、点赞操作、关注操作、发送消息
type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
	Error      string `json:"error"`
}

// 注册、登录
type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User model.User `json:"user"`
}

type ErrorResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
	Error      string `json:"error"`
}

type SearchIDResponse struct {
	StatusCode int32      `json:"status_code"`
	StatusMsg  string     `json:"status_msg,omitempty"`
	Error      string     `json:"error"`
	User       model.User `json:"user"`
}

// 视频流接口、发布列表、喜欢列表
type FeedResponse struct {
	StatusCode int32         `json:"status_code"`
	StatusMsg  string        `json:"status_msg,omitempty"`
	Nexttime   int32         `json:"nexttime"`
	VideoList  []model.Video `json:"video_list"`
}

// 评论操作
type CommentVResponse struct {
	StatusCode int32         `json:"status_code"`
	StatusMsg  string        `json:"status_msg,omitempty"`
	Comment    model.Comment `json:"comment"`
}

// 评论列表
type CommentListResponse struct {
	StatusCode  int32           `json:"status_code"`
	StatusMsg   string          `json:"status_msg,omitempty"`
	CommentList []model.Comment `json:"comment_list"`
}

// 关注、粉丝、好友列表
type FollowListResponse struct {
	StatusCode int32        `json:"status_code"`
	StatusMsg  string       `json:"status_msg,omitempty"`
	UserList   []model.User `json:"user_list"`
}

type ChatRecordResponse struct {
	StatusCode  int32           `json:"status_code"`
	StatusMsg   string          `json:"status_msg,omitempty"`
	MessageList []model.Message `json:"message_list"`
}
