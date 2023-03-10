package service

import (
	"Minimalist_TikTok/model"
	"Minimalist_TikTok/pkg/util"
	"Minimalist_TikTok/serializer"
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
)

type PublishService struct {
}

type PublishListService struct {
	Token string `form:"token",json:"token"`
	ID    string `form:"user_id",json:"user_id"`
}

type FeedService struct {
	LatestTime int    `form:"latest_time",json:"latest_time"`
	Token      string `form:"token",json:"token"`
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
		AuthorID: Uid,
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
	err = model.DB.Model(&model.User{}).Where("id = ?", Uid).Update("work_count", gorm.Expr("work_count + ?", 1)).Error
	if err != nil {
		util.LogrusObj.Info(err)
		return serializer.Response{
			StatusCode: 1,
			StatusMsg:  "work_count+1操作失败",
		}
	}
	return serializer.Response{
		StatusCode: 0,
		StatusMsg:  video.Title + " uploaded successfully",
	}
}

func (service *PublishListService) PublishList() serializer.FeedResponse {
	var videos []model.Video
	claim, err := util.ParseToken(service.Token)
	if err != nil {
		util.LogrusObj.Info(err)
		return serializer.FeedResponse{
			StatusCode: 1,
			StatusMsg:  "token解析错误",
		}
	}

	err = model.DB.Model(&model.Video{}).Preload("Author").Where("author_id=?", claim.Id).Order("created_at desc").Find(&videos).Error
	if err != nil {
		util.LogrusObj.Info(err)
		return serializer.FeedResponse{
			StatusCode: 1,
			StatusMsg:  "video数据库查询错误",
		}
	}

	VideoLen := len(videos)
	for i := 0; i < VideoLen; i++ {
		videos[i].PlayUrl = util.AndroidBeforeUrl + videos[i].PlayUrl
		videos[i].CoverUrl = util.AndroidBeforeUrl + videos[i].CoverUrl
	}
	//fmt.Println("data", videos)
	return serializer.FeedResponse{
		StatusCode: 0,
		StatusMsg:  "获取列表成功",
		VideoList:  videos,
	}
}

//
//func (service *FeedService) Feed() serializer.FeedResponse {
//	var videos []model.Video
//	timeLayout := "2006-01-02 15:04:05"
//	data, err := strconv.ParseInt(strconv.Itoa(service.LatestTime), 10, 64)
//	if err != nil {
//		util.LogrusObj.Info(err)
//		return serializer.FeedResponse{
//			StatusCode: 1,
//			StatusMsg:  "时间戳转化错误",
//		}
//	}
//	//登录时
//	if service.Token != "" {
//		//没有用到id信息
//		_, err := util.ParseToken(service.Token)
//		if err != nil {
//			util.LogrusObj.Info(err)
//			return serializer.FeedResponse{
//				StatusCode: 1,
//				StatusMsg:  "token解析错误",
//			}
//		}
//	}
//
//	err = model.DB.Model(&model.Video{}).Preload("Author").Where("created_at <= ?", time.Unix(data/1000, 0).Format(timeLayout)).Order("created_at desc").Limit(30).Find(&videos).Error
//	if err != nil {
//		util.LogrusObj.Info(err)
//		return serializer.FeedResponse{
//			StatusCode: 1,
//			StatusMsg:  "video数据库查询错误",
//		}
//	}
//
//	VideoLen := len(videos)
//	//fmt.Println("VideoLen", VideoLen)
//	for i := 0; i < VideoLen; i++ {
//		videos[i].PlayUrl = util.AndroidBeforeUrl + videos[i].PlayUrl
//		videos[i].CoverUrl = util.AndroidBeforeUrl + videos[i].CoverUrl
//	}
//
//	return serializer.FeedResponse{
//		StatusCode: 0,
//		StatusMsg:  "获取列表成功",
//		Nexttime:   int32(videos[VideoLen-1].CreatedAt.Unix()),
//		VideoList:  videos,
//	}
//}

func (service *FeedService) Feed() serializer.FeedResponse {
	var videos []model.Video
	timeLayout := "2006-01-02 15:04:05"
	data, err := strconv.ParseInt(strconv.Itoa(service.LatestTime), 10, 64)
	if err != nil {
		util.LogrusObj.Info(err)
		return serializer.FeedResponse{
			StatusCode: 1,
			StatusMsg:  "时间戳转化错误",
		}
	}
	err = model.DB.Model(&model.Video{}).Preload("Author").Where("created_at <= ?", time.Unix(data/1000, 0).Format(timeLayout)).Order("created_at desc").Limit(30).Find(&videos).Error
	if err != nil {
		util.LogrusObj.Info(err)
		return serializer.FeedResponse{
			StatusCode: 1,
			StatusMsg:  "video数据库查询错误",
		}
	}

	VideoLen := len(videos)
	for i := 0; i < VideoLen; i++ {
		videos[i].PlayUrl = util.AndroidBeforeUrl + videos[i].PlayUrl
		videos[i].CoverUrl = util.AndroidBeforeUrl + videos[i].CoverUrl
	}
	//登录时

	if service.Token != "" {
		fmt.Println("登录状态！")
		//没有用到id信息
		claim, err := util.ParseToken(service.Token)

		if err != nil {
			util.LogrusObj.Info(err)
			return serializer.FeedResponse{
				StatusCode: 1,
				StatusMsg:  "token解析错误",
			}
		}

		id := claim.Id
		for j := 0; j < VideoLen; j++ {
			var count int
			model.DB.Model(&model.Favourite{}).Where("user_id = ? And video_id = ? ", id, videos[j].ID).Count(&count)
			videos[j].IsFavorite = (count == 1)
		}
	}
	return serializer.FeedResponse{
		StatusCode: 0,
		StatusMsg:  "获取列表成功",
		Nexttime:   int32(videos[VideoLen-1].CreatedAt.Unix()),
		VideoList:  videos,
	}
}
