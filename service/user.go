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
	UserName string `form:"username" json:"username" binding:"required,min=3,max=15" example:"FanOne"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16" example:"FanOne666"`
}

// 查询用户ID的服务
type SearchIDService struct {
	Id uint `json:"id"`
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
	user.UserName = service.UserName
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
	token, _ := util.GenerateToken(user.ID, user.UserName, 0)
	return &serializer.UserLoginResponse{
		Response: serializer.Response{StatusCode: 0},
		UserId:   int64(user.ID),
		Token:    token,
	}
}

// Login 用户登陆函数
func (service *UserService) Login() *serializer.UserLoginResponse {
	var user model.User
	if err := model.DB.Where("user_name=?", service.UserName).First(&user).Error; err != nil {
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
		fmt.Println("user.PasswordDigest=1?", user.PasswordDigest)
		fmt.Println("password=", service.Password)
		return &serializer.UserLoginResponse{
			Response: serializer.Response{
				StatusCode: 1,
				StatusMsg:  "解密后密码错误",
			},
		}
	}
	token, err := util.GenerateToken(user.ID, service.UserName, 0)
	fmt.Println("登录的token", token)
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

func (service *SearchIDService) SearchById(id uint) serializer.SearchIDResponse {
	var user model.User
	err := model.DB.Where("Id = ?", id).Find(&user).Error
	if err != nil {
		util.LogrusObj.Info(err)
		return serializer.SearchIDResponse{
			StatusCode: 0,
			StatusMsg:  "错误",
			Error:      err.Error(),
		}
	}
	return serializer.SearchIDResponse{
		StatusCode: 1,
		StatusMsg:  "查询成功",
		Data:       user,
	}
}
