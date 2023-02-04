package service

import (
	"Minimalist_TikTok/model"
	"Minimalist_TikTok/pkg/util"
	"Minimalist_TikTok/serializer"
	"fmt"
	"github.com/jinzhu/gorm"
)

// UserService 用户注册服务
type UserService struct {
	UserName string `form:"name" json:"name" binding:"required,min=3,max=15" example:"FanOne"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16" example:"FanOne666"`
}

func (service *UserService) Register() *serializer.UserLoginResponse {
	var user model.User
	var count int64
	model.DB.Model(&model.User{}).Where("name=?", service.UserName).First(&user).Count(&count)
	// 表单验证
	fmt.Println("count=", count)
	if count == 1 {
		return &serializer.UserLoginResponse{
			Response: serializer.Response{
				StatusCode: 1,
				StatusMsg:  "User already exist"},
		}
	}
	user.Name = service.UserName
	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		util.LogrusObj.Info(err)
		return &serializer.UserLoginResponse{
			Response: serializer.Response{
				StatusCode: 1,
				StatusMsg:  "加密出错"},
		}
	}
	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		util.LogrusObj.Info(err)
		return &serializer.UserLoginResponse{
			Response: serializer.Response{
				StatusCode: 1,
				StatusMsg:  "创建出错"},
		}
	}
	token, _ := util.GenerateToken(user.ID, user.Name, 0)
	return &serializer.UserLoginResponse{
		Response: serializer.Response{StatusCode: 0},
		UserId:   int64(user.ID),
		Token:    token,
	}
}

// Login 用户登陆函数
func (service *UserService) Login() *serializer.UserLoginResponse {
	var user model.User
	if err := model.DB.Where("name=?", service.UserName).First(&user).Error; err != nil {
		// 如果查询不到，返回相应的错误
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			util.LogrusObj.Info(err)
			return &serializer.UserLoginResponse{
				Response: serializer.Response{
					StatusCode: 1,
					StatusMsg:  "User doesn't exist",
				},
			}
		}
		util.LogrusObj.Info(err)
		return &serializer.UserLoginResponse{
			Response: serializer.Response{
				StatusCode: 1,
				StatusMsg:  "查询失败",
			},
		}
	}
	//解密后密码错误
	if !user.CheckPassword(service.Password) {
		return &serializer.UserLoginResponse{
			Response: serializer.Response{
				StatusCode: 1,
				StatusMsg:  "解密后密码错误",
			},
		}
	}
	token, err := util.GenerateToken(user.ID, service.UserName, 0)
	//加密出错
	if err != nil {
		util.LogrusObj.Info(err)
		return &serializer.UserLoginResponse{
			Response: serializer.Response{
				StatusCode: 1,
				StatusMsg:  "加密出错",
			},
		}
	}
	return &serializer.UserLoginResponse{
		Response: serializer.Response{
			StatusCode: 0,
			StatusMsg:  "登录成功",
		},
		UserId: int64(user.ID),
		Token:  token,
	}
}
