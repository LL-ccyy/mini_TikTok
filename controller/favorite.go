package controller

import (
	"Minimalist_TikTok/pkg/util"
	"Minimalist_TikTok/service"
	"github.com/gin-gonic/gin"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	var favouriteActionService service.FavouriteActionService
	if err := c.ShouldBind(&favouriteActionService); err == nil {
		res := favouriteActionService.FavouriteAction()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	var favouriteService service.FavouriteListService
	if err := c.ShouldBind(&favouriteService); err == nil {
		res := favouriteService.FavouriteList()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}
