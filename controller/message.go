package controller

import (
	"Minimalist_TikTok/pkg/util"
	"Minimalist_TikTok/service"
	"github.com/gin-gonic/gin"
)

func MessageAction(c *gin.Context) {
	var messageActionService service.MessageActionService
	if err := c.ShouldBind(&messageActionService); err == nil {
		res := messageActionService.MessageAction()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

func MessageChat(c *gin.Context) {
	var messageListService service.MessageChatService
	if err := c.ShouldBind(&messageListService); err == nil {
		res := messageListService.MessageChat()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}
