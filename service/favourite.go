package service

import (
	"Minimalist_TikTok/model"
	"Minimalist_TikTok/pkg/util"
	"Minimalist_TikTok/serializer"
	"fmt"
	"github.com/jinzhu/gorm"
)

type FavouriteActionService struct {
	Token      string `form:"token",json:"token"`
	VideoId    uint   `form:"video_id",json:"video_id"`
	ActionType string `form:"action_type",json:"action_type"`
}

type FavouriteListService struct {
	Token string `form:"token",json:"token"`
	ID    string `form:"user_id",json:"user_id"`
}

func (service *FavouriteActionService) FavouriteAction() serializer.Response {
	claims, err := util.ParseToken(service.Token)
	if err != nil {
		util.LogrusObj.Info(err)
		return serializer.Response{
			StatusCode: 1,
			StatusMsg:  "token解析错误",
		}
	}
	fmt.Println("user_id=", claims.Id)

	favourite := model.Favourite{
		UserID:  claims.Id,
		VideoID: service.VideoId,
	}
	switch service.ActionType {
	case "1":
		{
			//var user model.User
			//var video model.Video
			//model.DB.Model(&model.User{}).Where("id = ?",claims.Id).Find(&user)
			//model.DB.Model(&model.Video{}).Where("id = ?",service.VideoId).Find(&video)
			err = model.DB.Model(&model.Favourite{}).Create(&favourite).Error
			if err != nil {
				util.LogrusObj.Info(err)
				return serializer.Response{
					StatusCode: 1,
					StatusMsg:  "添加喜欢数据库操作失败",
				}
			}
			err = model.DB.Model(&model.Video{}).Where("id = ?", service.VideoId).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
			if err != nil {
				util.LogrusObj.Info(err)
				return serializer.Response{
					StatusCode: 1,
					StatusMsg:  "视频喜欢+1操作失败",
				}
			}
			err = model.DB.Model(&model.User{}).Where("id = ?", claims.Id).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
			if err != nil {
				util.LogrusObj.Info(err)
				return serializer.Response{
					StatusCode: 1,
					StatusMsg:  "用户喜欢+1操作失败",
				}
			}
			return serializer.Response{
				StatusCode: 0,
				StatusMsg:  "喜欢成功",
			}
		}
	case "2":
		{
			err = model.DB.Model(&model.Favourite{}).Where("user_id = ? AND video_id= ?", claims.Id, service.VideoId).Delete(&favourite).Error
			if err != nil {
				util.LogrusObj.Info(err)
				return serializer.Response{
					StatusCode: 1,
					StatusMsg:  "删除喜欢数据库操作失败",
				}
			}
			err = model.DB.Model(&model.Video{}).Where("id = ?", service.VideoId).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error
			if err != nil {
				util.LogrusObj.Info(err)
				return serializer.Response{
					StatusCode: 1,
					StatusMsg:  "视频喜欢-1操作失败",
				}
			}
			err = model.DB.Model(&model.User{}).Where("id = ?", claims.Id).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error
			if err != nil {
				util.LogrusObj.Info(err)
				return serializer.Response{
					StatusCode: 1,
					StatusMsg:  "用户喜欢-1操作失败",
				}
			}
			return serializer.Response{
				StatusCode: 0,
				StatusMsg:  "不喜欢成功",
			}
		}
	default:
		{
			return serializer.Response{
				StatusCode: 1,
				StatusMsg:  "type不对",
			}
		}
	}
}

func (service *FavouriteListService) FavouriteList() serializer.FeedResponse {
	var favouritevideos []model.Favourite
	//var favouritevideosList []model.Video
	claim, err := util.ParseToken(service.Token)
	if err != nil {
		util.LogrusObj.Info(err)
		return serializer.FeedResponse{
			StatusCode: 1,
			StatusMsg:  "token解析错误",
		}
	}

	err = model.DB.Model(&model.Favourite{}).Preload("Video").Where("user_id = ?", claim.Id).Find(&favouritevideos).Error
	if err != nil {
		util.LogrusObj.Info(err)
		return serializer.FeedResponse{
			StatusCode: 1,
			StatusMsg:  "video数据库查询错误",
		}
	}

	//fmt.Println("favouritevideos???????=",favouritevideos)
	//
	//var videosId []uint
	//for _, v := range favouritevideos {
	//	videosId = append(videosId, v.VideoID)
	//}

	var videos []model.Video
	for _, v := range favouritevideos {
		videos = append(videos, v.Video)
	}
	VideoLen1 := len(videos)
	for i := 0; i < VideoLen1; i++ {
		videos[i].PlayUrl = util.AndroidBeforeUrl + videos[i].PlayUrl
		videos[i].CoverUrl = util.AndroidBeforeUrl + videos[i].CoverUrl
	}

	//model.DB.Model(&model.Video{}).Where("id = ?", videosId).Find(&favouritevideosList)

	//db.Where("name IN (?)", []string{"jinzhu", "jinzhu 2"})

	//VideoLen := len(favouritevideosList)
	//for i := 0; i < VideoLen; i++ {
	//	favouritevideosList[i].PlayUrl = util.AndroidBeforeUrl + favouritevideosList[i].PlayUrl
	//	favouritevideosList[i].CoverUrl = util.AndroidBeforeUrl + favouritevideosList[i].CoverUrl
	//}
	//fmt.Println("data", favouritevideosList)
	return serializer.FeedResponse{
		StatusCode: 0,
		StatusMsg:  "获取列表成功",
		VideoList:  videos,
	}
}
