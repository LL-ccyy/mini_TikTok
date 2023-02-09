package service

import (
	"Minimalist_TikTok/model"
	"Minimalist_TikTok/pkg/util"
	"Minimalist_TikTok/serializer"
	"fmt"
	"github.com/jinzhu/gorm"
)

type RelationActionService struct {
	Token      string `form:"token",json:"token"`
	ToUserID   int64  `form:"to_user_id",json:"to_user_id"`
	ActionType string `form:"action_type",json:"action_type"`
}

type FollowService struct {
	Token  string `form:"token",json:"token"`
	UserID int64  `form:"user_id",json:"user_id"`
}

func (service *RelationActionService) RelationAction() serializer.Response {
	claims, err := util.ParseToken(service.Token)
	if err != nil {
		util.LogrusObj.Info(err)
		return serializer.Response{
			StatusCode: 1,
			StatusMsg:  "token解析错误",
		}
	}
	fmt.Println("from_user_id=", claims.Id)

	follow := model.Follow{
		FollowID:   claims.Id,
		FollowerID: uint(service.ToUserID),
	}
	switch service.ActionType {
	case "1":
		{
			err = model.DB.Model(&model.Follow{}).Create(&follow).Error
			if err != nil {
				util.LogrusObj.Info(err)
				return serializer.Response{
					StatusCode: 1,
					StatusMsg:  "添加关注数据库操作失败",
				}
			}

			err = model.DB.Model(&model.User{}).Where("id = ?", claims.Id).Update("follow_count", gorm.Expr("follow_count + ?", 1)).Error
			if err != nil {
				util.LogrusObj.Info(err)
				return serializer.Response{
					StatusCode: 1,
					StatusMsg:  "关注+1失败",
				}
			}

			err = model.DB.Model(&model.User{}).Where("id = ?", service.ToUserID).Update("follower_count", gorm.Expr("follower_count + ?", 1)).Error
			if err != nil {
				util.LogrusObj.Info(err)
				return serializer.Response{
					StatusCode: 1,
					StatusMsg:  "粉丝+1失败",
				}
			}

			return serializer.Response{
				StatusCode: 0,
				StatusMsg:  "关注成功",
			}
		}
	case "2":
		{
			err = model.DB.Model(&model.Follow{}).Where("follow_id = ? AND follower_id = ?", claims.Id, service.ToUserID).Delete(&follow).Error
			if err != nil {
				util.LogrusObj.Info(err)
				return serializer.Response{
					StatusCode: 1,
					StatusMsg:  "删除关注数据库操作失败",
				}
			}

			err = model.DB.Model(&model.User{}).Where("id = ?", claims.Id).Update("follow_count", gorm.Expr("follow_count - ?", 1)).Error
			if err != nil {
				util.LogrusObj.Info(err)
				return serializer.Response{
					StatusCode: 1,
					StatusMsg:  "关注-1失败",
				}
			}

			err = model.DB.Model(&model.User{}).Where("id = ?", service.ToUserID).Update("follower_count", gorm.Expr("follower_count - ?", 1)).Error
			if err != nil {
				util.LogrusObj.Info(err)
				return serializer.Response{
					StatusCode: 1,
					StatusMsg:  "粉丝-1失败",
				}
			}

			return serializer.Response{
				StatusCode: 0,
				StatusMsg:  "取关成功",
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

func (service *FollowService) FollowList() serializer.FollowListResponse {
	claims, err := util.ParseToken(service.Token)
	if err != nil {
		util.LogrusObj.Info(err)
		return serializer.FollowListResponse{
			StatusCode: 1,
			StatusMsg:  "token解析错误",
		}
	}
	fmt.Println("from_user_id=", claims.Id)

	var follows []model.Follow
	err = model.DB.Model(&model.Follow{}).Preload("Follow").Where("follow_id = ?", service.UserID).Find(&follows).Error
	if err != nil {
		util.LogrusObj.Info(err)
		return serializer.FollowListResponse{
			StatusCode: 1,
			StatusMsg:  "查询follow数据库错误",
		}
	}

	var followslist []model.User
	for _, v := range follows {
		followslist = append(followslist, v.Follow)
	}

	return serializer.FollowListResponse{
		StatusCode: 0,
		StatusMsg:  "查询关注列表成功",
		UserList:   followslist,
	}
}

func (service *FollowService) FollowerList() serializer.FollowListResponse {
	_, err := util.ParseToken(service.Token)
	if err != nil {
		util.LogrusObj.Info(err)
		return serializer.FollowListResponse{
			StatusCode: 1,
			StatusMsg:  "token解析错误",
		}
	}

	var followers []model.Follow
	err = model.DB.Model(&model.Follow{}).Preload("Follower").Where("follower_id = ?", service.UserID).Find(&followers).Error
	if err != nil {
		util.LogrusObj.Info(err)
		return serializer.FollowListResponse{
			StatusCode: 1,
			StatusMsg:  "查询follow数据库错误",
		}
	}

	var followerslist []model.User
	for _, v := range followers {
		followerslist = append(followerslist, v.Follow)
	}

	return serializer.FollowListResponse{
		StatusCode: 0,
		StatusMsg:  "查询关注列表成功",
		UserList:   followerslist,
	}
}

// 没想好
func (service *FollowService) FriendList() serializer.FollowListResponse {
	_, err := util.ParseToken(service.Token)
	if err != nil {
		util.LogrusObj.Info(err)
		return serializer.FollowListResponse{
			StatusCode: 1,
			StatusMsg:  "token解析错误",
		}
	}

	//var followers []model.Follow
	//err = model.DB.Model(&model.Follow{}).Preload("Follower").Where("follower_id = ?",service.UserID).Find(&followers).Error
	//if err != nil {
	//	util.LogrusObj.Info(err)
	//	return serializer.FollowListResponse{
	//		StatusCode: 1,
	//		StatusMsg:  "查询follow数据库错误",
	//	}
	//}
	//
	//var followerslist []model.User
	//for _,v := range followers{
	//	followerslist = append(followerslist, v.Follow)
	//}

	return serializer.FollowListResponse{
		StatusCode: 0,
		//StatusMsg:  "查询关注列表成功",
		//UserList: followerslist,
	}
}
