package model

type Favourite struct {
	UserID  uint  `json:"user_id"`
	VideoID uint  `json:"video_id"`
	User    User  `json:"user"`
	Video   Video `json:"video"`
}
