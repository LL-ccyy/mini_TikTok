package controller

import (
	"Minimalist_TikTok/pkg/util"
	"Minimalist_TikTok/service"
	"github.com/gin-gonic/gin"
)

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	var commentActionService service.CommentActionService
	if err := c.ShouldBind(&commentActionService); err == nil {
		res := commentActionService.CommentAction()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

// CommentList all videos have same demo comment list
//func CommentList(c *gin.Context) {
//	var commentService service.CommentService
//	if err := c.ShouldBind(&commentService); err == nil {
//		res := commentService.CommentList()
//		c.JSON(200, res)
//	} else {
//		c.JSON(400, ErrorResponse(err))
//		util.LogrusObj.Info(err)
//	}
//}
