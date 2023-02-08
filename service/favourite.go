package service

import (
	"Minimalist_TikTok/model"
	"Minimalist_TikTok/pkg/util"
	"Minimalist_TikTok/serializer"
	"fmt"
)

type FavouriteActionService struct {
	Token      string `json:"token"`
	VideoId    uint   `json:"video_id"`
	ActionType string `json:"action_type"`
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
			model.DB.Model(&model.Favourite{}).Create(&favourite)

			return serializer.Response{
				StatusCode: 0,
				StatusMsg:  "喜欢成功",
			}
		}
	case "2":
		{
			model.DB.Model(&model.Favourite{}).Where("user_id = ? AND video_id= ?", claims.Id, service.VideoId).Delete(&favourite)

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
	var favouritevideosList []model.Video
	claim, err := util.ParseToken(service.Token)
	if err != nil {
		util.LogrusObj.Info(err)
		return serializer.FeedResponse{
			StatusCode: 1,
			StatusMsg:  "token解析错误",
		}
	}

	model.DB.Model(&model.Favourite{}).Where("user_id = ?", claim.Id).Find(&favouritevideos)

	var videosId []uint
	for _, v := range favouritevideos {
		videosId = append(videosId, v.VideoID)
	}

	model.DB.Model(&model.Video{}).Where("id = ?", videosId).Find(&favouritevideosList)

	//db.Where("name IN (?)", []string{"jinzhu", "jinzhu 2"})

	VideoLen := len(favouritevideosList)
	for i := 0; i < VideoLen; i++ {
		favouritevideosList[i].PlayUrl = util.AndroidBeforeUrl + favouritevideosList[i].PlayUrl
		favouritevideosList[i].CoverUrl = util.AndroidBeforeUrl + favouritevideosList[i].CoverUrl
	}
	fmt.Println("data", favouritevideosList)
	return serializer.FeedResponse{
		StatusCode: 0,
		StatusMsg:  "获取列表成功",
		VideoList:  favouritevideosList,
	}
}
