package model

type Comment struct {
	ID         uint   `json:"id"`
	User       User   `json:"user"`
	UserID     uint   `json:"user_id"`
	Video      Video  `json:"video"`
	VideoID    uint   `json:"video_id"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}
