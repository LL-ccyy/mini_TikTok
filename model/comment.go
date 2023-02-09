package model

type Comment struct {
	ID          uint   `json:"id"`
	Commenter   User   `json:"user"`
	CommenterID uint   `json:"user_id"`
	Video       Video  `json:"video"`
	VideoID     uint   `json:"video_id"`
	Content     string `json:"content,omitempty"`
	CreateDate  string `json:"create_date,omitempty"`
}
