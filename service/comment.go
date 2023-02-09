package service

import (
	"Minimalist_TikTok/model"
	"Minimalist_TikTok/pkg/util"
	"Minimalist_TikTok/serializer"
	"fmt"
)

type CommentActionService struct {
	Token       string `json:"token"`
	VideoId     uint   `json:"video_id"`
	ActionType  string `json:"action_type"`
	CommentText string `json:"comment_text"`
	CommentId   uint   `json:"comment_id"`
}

type CommentListService struct {
	Token   string `form:"token",json:"token"`
	VideoID string `form:"video_id",json:"video_id"`
}

func (service *CommentActionService) CommentAction() serializer.CommentVResponse {
	claims, err := util.ParseToken(service.Token)
	if err != nil {
		util.LogrusObj.Info(err)
		return serializer.CommentVResponse{
			StatusCode: 1,
			StatusMsg:  "token解析错误",
		}
	}
	fmt.Println("user_id=", claims.Id)

	//comment := model.Comment{
	//	UserID: claims.Id,
	//	VideoID:service.VideoId,
	//	Content: service.CommentText,
	//	//ID: service.CommentId,????
	//}

	switch service.ActionType {
	case "1":
		{
			comment := model.Comment{
				UserID:  claims.Id,
				VideoID: service.VideoId,
				Content: service.CommentText,
				//ID: service.CommentId,????
			}
			//var user model.User
			//var video model.Video
			//model.DB.Model(&model.User{}).Where("id = ?",claims.Id).Find(&user)
			//model.DB.Model(&model.Video{}).Where("id = ?",service.VideoId).Find(&video)
			model.DB.Model(&model.Comment{}).Create(&comment)

			return serializer.CommentVResponse{
				StatusCode: 0,
				StatusMsg:  "评论成功",
				Comment:    comment,
			}
		}
	case "2":
		{

			comment := model.Comment{
				UserID:  claims.Id,
				VideoID: service.VideoId,
				Content: service.CommentText,
				ID:      service.CommentId,
			}

			model.DB.Model(&model.Comment{}).Where("id = ? ", service.CommentId).Delete(&comment)

			return serializer.CommentVResponse{
				StatusCode: 0,
				StatusMsg:  "控评成功",
			}
		}
	default:
		{
			return serializer.CommentVResponse{
				StatusCode: 1,
				StatusMsg:  "type不对",
			}
		}
	}
}

func (service *CommentListService) CommentList() serializer.CommentListResponse {
	var commentsList []model.Comment
	//claims框架
	_, err := util.ParseToken(service.Token)
	if err != nil {
		util.LogrusObj.Info(err)
		return serializer.CommentListResponse{
			StatusCode: 1,
			StatusMsg:  "token解析错误",
		}
	}

	model.DB.Model(&model.Comment{}).Preload("Commenter").Where("video_id = ?", service.VideoID).Find(&commentsList)

	//model.DB.Model(&model.Favourite{}).Where("user_id = ?",claim.Id).Find(&favouritevideos)
	//
	//var videosId []uint
	//for _, v := range favouritevideos {
	//	videosId=append(videosId,v.VideoID)
	//}
	//
	//model.DB.Model(&model.Video{}).Where("id = ?",videosId).Find(&favouritevideosList)
	//
	//
	////db.Where("name IN (?)", []string{"jinzhu", "jinzhu 2"})
	//
	//VideoLen := len(favouritevideosList)
	//for i := 0; i < VideoLen; i++ {
	//	favouritevideosList[i].PlayUrl = util.AndroidBeforeUrl + favouritevideosList[i].PlayUrl
	//	favouritevideosList[i].CoverUrl = util.AndroidBeforeUrl + favouritevideosList[i].CoverUrl
	//}
	//fmt.Println("data", favouritevideosList)
	return serializer.CommentListResponse{
		StatusCode:  0,
		StatusMsg:   "获取列表成功",
		CommentList: commentsList,
	}
}
