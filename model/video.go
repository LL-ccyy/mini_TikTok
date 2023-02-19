package model

import (
	"time"
)

type Video struct {
	ID            uint `json:"id",gorm:"primary_key"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `sql:"index"`
	Author        User       `json:"author"`
	AuthorID      uint       `json:"author_id;not null"`
	PlayUrl       string     `json:"play_url,omitempty"`
	CoverUrl      string     `json:"cover_url,omitempty"`
	FavoriteCount int64      `json:"favorite_count,omitempty",gorm:"default:0"`
	CommentCount  int64      `json:"comment_count,omitempty",gorm:"default:0"`
	IsFavorite    bool       `json:"is_favorite,omitempty",gorm:"default:False"`
	Title         string     `json:"title,omitempty"`
}
