package service

import (
	"Minimalist_TikTok/model"
	"Minimalist_TikTok/pkg/util"
	"Minimalist_TikTok/serializer"
	"fmt"
	"time"
)

type CommentActionService struct {
	Token       string `form:"token",json:"token"`
	VideoId     uint   `form:"video_id",json:"video_id"`
	ActionType  string `form:"action_type",json:"action_type"`
	CommentText string `form:"comment_text",json:"comment_text"`
	CommentId   uint   `form:"comment_id",json:"comment_id"`
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
			var commenter model.User
			err = model.DB.Model(&model.User{}).Where("id=?", claims.Id).Find(&commenter).Error
			if err != nil {
				util.LogrusObj.Info(err)
				return serializer.CommentVResponse{
					StatusCode: 1,
					StatusMsg:  "查找人失败",
				}
			}
			comment := model.Comment{
				CommenterID: claims.Id,
				Commenter:   commenter,
				VideoID:     service.VideoId,
				Content:     service.CommentText,
				CreateDate:  time.Now().Format("2006-01-02 15:04:05"),
				//ID: service.CommentId,????
			}
			//var user model.User
			//var video model.Video
			//model.DB.Model(&model.User{}).Where("id = ?",claims.Id).Find(&user)
			//model.DB.Model(&model.Video{}).Where("id = ?",service.VideoId).Find(&video)
			err = model.DB.Model(&model.Comment{}).Create(&comment).Error
			if err != nil {
				util.LogrusObj.Info(err)
				return serializer.CommentVResponse{
					StatusCode: 1,
					StatusMsg:  "添加评论数据库操作失败",
				}
			}
			return serializer.CommentVResponse{
				StatusCode: 0,
				StatusMsg:  "评论成功",
				Comment:    comment,
			}
		}
	case "2":
		{
			comment := model.Comment{
				CommenterID: claims.Id,
				VideoID:     service.VideoId,
				Content:     service.CommentText,
				ID:          service.CommentId,
			}

			err = model.DB.Model(&model.Comment{}).Where("id = ? ", service.CommentId).Delete(&comment).Error
			if err != nil {
				util.LogrusObj.Info(err)
				return serializer.CommentVResponse{
					StatusCode: 1,
					StatusMsg:  "删除评论数据库操作失败",
				}
			}
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

	err = model.DB.Model(&model.Comment{}).Preload("Commenter").Preload("Video").Preload("Video.Author").Where("video_id = ?", service.VideoID).Find(&commentsList).Error
	if err != nil {
		util.LogrusObj.Info(err)
		return serializer.CommentListResponse{
			StatusCode: 1,
			StatusMsg:  "评论数据库查找错误",
		}
	}

	return serializer.CommentListResponse{
		StatusCode:  0,
		StatusMsg:   "获取列表成功",
		CommentList: commentsList,
	}
}
