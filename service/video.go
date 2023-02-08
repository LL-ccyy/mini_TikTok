package service

import (
	"Minimalist_TikTok/model"
	"Minimalist_TikTok/pkg/util"
	"Minimalist_TikTok/serializer"
	"fmt"
)

type PublishService struct {
}

type PublishListService struct {
	Token string `form:"token",json:"token"`
	ID    string `form:"user_id",json:"user_id"`
}

//func UserInfo(c *gin.Context) {
//	var userinfo service.SearchIDService
//	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
//	fmt.Println("claim=", claim)
//	if err := c.ShouldBind(&userinfo); err == nil {
//		res := userinfo.SearchById(claim.Id)
//		c.JSON(200, res)
//	} else {
//		c.JSON(400, ErrorResponse(err))
//		util.LogrusObj.Info(err)
//	}
//}

func (service *PublishService) SearchByUid(uid uint) (model.User, error) {
	var user model.User
	err := model.DB.Model(&model.User{}).Where("ID = ?", uid).Find(&user).Error
	if err != nil {
		util.LogrusObj.Info(err)
		return user, err
		//return serializer.UserResponse{
		//	Response: serializer.Response{
		//		StatusCode: 0,
		//		StatusMsg:  "查找错误？不存在ID？",
		//		Error:      err.Error(),},
		//}
	}
	fmt.Println("user=", user)
	return user, nil
	//return serializer.UserResponse{
	//	Response: serializer.Response{
	//		StatusCode: 1,
	//		StatusMsg:  "查找成功"},
	//	User:user,
	//}
}

func Publish(User model.User, Uid uint, PlayUrl string, CoverUrl string, Title string) serializer.Response {
	//var video model.Video
	video := model.Video{
		Author:   User,
		Uid:      Uid,
		PlayUrl:  PlayUrl,
		CoverUrl: CoverUrl,
		Title:    Title,
	}
	err := model.DB.Model(&model.Video{}).Create(&video).Error
	if err != nil {
		util.LogrusObj.Info(err)
		return serializer.Response{
			StatusCode: 1,
			StatusMsg:  "投稿失败",
			Error:      err.Error(),
		}
	}
	return serializer.Response{
		StatusCode: 0,
		StatusMsg:  video.Title + " uploaded successfully",
	}
}

func (service *PublishListService) PublishList() serializer.FeedResponse {
	//var videos []model.Video
	fmt.Println("service.id=", service.ID)
	//fmt.Println("s.t=",service.Token)
	//model.DB.Model(&model.Video{}).Preload("User").Where("uid=", service.ID).Order("created_at DESC").Find(&videos)
	return serializer.FeedResponse{}
}
