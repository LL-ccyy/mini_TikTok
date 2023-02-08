package controller

import (
	"Minimalist_TikTok/pkg/util"
	"Minimalist_TikTok/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var userRegisterService service.UserService
	if err := c.ShouldBind(&userRegisterService); err == nil {
		res := userRegisterService.Register()
		fmt.Println("userRegisterService=", userRegisterService)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func Login(c *gin.Context) {
	var userLoginService service.UserService
	if err := c.ShouldBind(&userLoginService); err == nil {
		res := userLoginService.Login()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func UserInfo(c *gin.Context) {
	var userinfo service.SearchIDService
	//token := c.Query("token")
	//claim, _ := util.ParseToken(token)
	//fmt.Println("claim=", claim)
	//
	if err := c.ShouldBind(&userinfo); err == nil {
		res := userinfo.SearchById()
		fmt.Println("res=", res)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}
