package model

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName       string `json:"name,omitempty",gorm:"column:user_name"`
	PasswordDigest string `json:"password,omitempty"`
	FollowCount    int64  `json:"follow_count,omitempty",gorm:"default:0"`
	FollowerCount  int64  `json:"follower_count,omitempty",gorm:"default:0"`
	IsFollow       bool   `json:"is_follow,omitempty",gorm:"default:false"`
}

// 加密
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// 验证密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	fmt.Println("user.PasswordDigest=", user.PasswordDigest)
	fmt.Println(bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password)))
	return err == nil
}
