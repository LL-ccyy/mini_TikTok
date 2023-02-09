package controller

import (
	"Minimalist_TikTok/pkg/util"
	"Minimalist_TikTok/service"
	"github.com/gin-gonic/gin"
)

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	var relationActionService service.RelationActionService
	if err := c.ShouldBind(&relationActionService); err == nil {
		res := relationActionService.RelationAction()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

// // FollowList all users have same follow list
func FollowList(c *gin.Context) {
	var followservice service.FollowService
	if err := c.ShouldBind(&followservice); err == nil {
		res := followservice.FollowList()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	var followerservice service.FollowService
	if err := c.ShouldBind(&followerservice); err == nil {
		res := followerservice.FollowerList()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}

// FriendList all users have same friend list
func FriendList(c *gin.Context) {
	var friendservice service.FollowService
	if err := c.ShouldBind(&friendservice); err == nil {
		res := friendservice.FriendList()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Info(err)
	}
}
