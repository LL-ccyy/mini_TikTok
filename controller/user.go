package controller

import (
	"Minimalist_TikTok/pkg/util"
	"Minimalist_TikTok/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var userRegisterService service.UserService //相当于创建了一个UserRegisterService对象，调用这个对象中的Register方法。
	//username := c.Query("username")
	//password := c.Query("password")
	if err := c.ShouldBind(&userRegisterService); err == nil {
		res := userRegisterService.Register()
		fmt.Println("userRegisterService=", userRegisterService)
		c.JSON(200, res)
	} else {
		fmt.Println(err)
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
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	fmt.Println("claim=", claim)
	if err := c.ShouldBind(&userinfo); err == nil {
		res := userinfo.SearchById(claim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}
