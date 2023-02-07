package model

import "github.com/jinzhu/gorm"

type Video struct {
	gorm.Model
	Author        User   `json:"author"`
	Uid           uint   `json:"uid;not null"`
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty",gorm:"default:0"`
	CommentCount  int64  `json:"comment_count,omitempty",gorm:"default:0"`
	IsFavorite    bool   `json:"is_favorite,omitempty",gorm:"default:False"`
	Title         string `json:"title,omitempty"`
}
