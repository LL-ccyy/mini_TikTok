package controller

import (
	"Minimalist_TikTok/pkg/util"
	"Minimalist_TikTok/serializer"
	"Minimalist_TikTok/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strings"
)

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	var publishservice service.PublishService
	token := c.PostForm("token")
	claim, err := util.ParseToken(token)
	//claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusOK, serializer.Response{
			StatusCode: 1,
			StatusMsg:  "token解析错误",
		})
	}
	//if err := c.ShouldBind(&publishservice); err == nil {
	//	user := publishservice.SearchByUid(claim.Id)
	//	c.JSON(200, res)
	//} else {
	//	c.JSON(400, ErrorResponse(err))
	//	util.LogrusObj.Info(err)
	//}
	title := c.PostForm("title")
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, serializer.Response{
			StatusCode: 1,
			StatusMsg:  "获取视频数据失败",
			Error:      err.Error(),
		})
	}

	//获取文件名带后缀
	filename := filepath.Base(data.Filename)
	filenameWithSuffix := filepath.Ext(filename)
	filenameOnly := strings.TrimSuffix(filenameWithSuffix, filename)
	user, nil := publishservice.SearchByUid(claim.Id)
	finalName := fmt.Sprintf("%d_%s", user.ID, filename)
	saveFile := filepath.Join("./public/", finalName)

	ffm_ima, err := util.GetSnapshot(saveFile, filenameOnly, 1)
	if err != nil {
		fmt.Println("截取视频封面出错，", err)
	}

	res := service.Publish(user, claim.Id, saveFile, ffm_ima, title)

	c.JSON(http.StatusOK, res)

	//if err := c.SaveUploadedFile(data, saveFile); err != nil {
	//	c.JSON(http.StatusOK, Response{
	//		StatusCode: 1,
	//		StatusMsg:  err.Error(),
	//	})
	//	return
	//}
	//
	//if err != nil {
	//	c.JSON(http.StatusOK, Response{
	//		StatusCode: 1,
	//		StatusMsg:  err.Error(),
	//	})
	//	return
	//}
	//
	//
	//finalName := fmt.Sprintf("%d_%s", user.ID, filename)
	//
	//if err := c.SaveUploadedFile(data, saveFile); err != nil {
	//	c.JSON(http.StatusOK, model.Response{
	//		StatusCode: 1,
	//		StatusMsg:  err.Error(),
	//	})
	//	return
	//}
	//
	//filename := filepath.Base(data.Filename)
	//user := usersLoginInfo[token]
	//finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	//saveFile := filepath.Join("./public/", finalName)
	//if err := c.SaveUploadedFile(data, saveFile); err != nil {
	//	c.JSON(http.StatusOK, Response{
	//		StatusCode: 1,
	//		StatusMsg:  err.Error(),
	//	})
	//	return
	//}
	//
	//c.JSON(http.StatusOK, Response{
	//	StatusCode: 0,
	//	StatusMsg:  finalName + " uploaded successfully",
	//})
}

// PublishList all users have same publish video list
//func PublishList(c *gin.Context) {
//	c.JSON(http.StatusOK, VideoListResponse{
//		Response: Response{
//			StatusCode: 0,
//		},
//		VideoList: DemoVideos,
//	})
//}
