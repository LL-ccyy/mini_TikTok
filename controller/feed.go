package controller

import (
	"Minimalist_TikTok/pkg/util"
	"Minimalist_TikTok/service"
	"github.com/gin-gonic/gin"
)

func Feed(c *gin.Context) {
	var feedservice service.FeedService
	if err := c.ShouldBind(&feedservice); err == nil {
		res := feedservice.Feed()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}
