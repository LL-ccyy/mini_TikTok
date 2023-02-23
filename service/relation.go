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
		FollowID:   uint(service.ToUserID),
		FollowerID: claims.Id,
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

// A follow B
// A 粉丝 follower
// B 偶像 follow
// FollowList 关注列表
// FollowerList 粉丝列表
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
	err = model.DB.Model(&model.Follow{}).Preload("Follower").Preload("Follow").Where("follower_id = ?", service.UserID).Find(&follows).Error
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

	for j := 0; j < len(followslist); j++ {
		var count int
		model.DB.Model(&model.Follow{}).Where("follow_id = ? And follower_id = ? ", followslist[j].ID, claims.Id).Count(&count)
		followslist[j].IsFollow = (count == 1)
	}

	return serializer.FollowListResponse{
		StatusCode: 0,
		StatusMsg:  "查询关注列表成功",
		UserList:   followslist,
	}
}

func (service *FollowService) FollowerList() serializer.FollowListResponse {
	claims, err := util.ParseToken(service.Token)
	if err != nil {
		util.LogrusObj.Info(err)
		return serializer.FollowListResponse{
			StatusCode: 1,
			StatusMsg:  "token解析错误",
		}
	}

	var followers []model.Follow
	err = model.DB.Model(&model.Follow{}).Preload("Follower").Preload("Follow").Where("follow_id = ?", service.UserID).Find(&followers).Error
	if err != nil {
		util.LogrusObj.Info(err)
		return serializer.FollowListResponse{
			StatusCode: 1,
			StatusMsg:  "查询follow数据库错误",
		}
	}

	var followerslist []model.User
	for _, v := range followers {
		followerslist = append(followerslist, v.Follower)
	}

	for j := 0; j < len(followerslist); j++ {
		var count int
		model.DB.Model(&model.Follow{}).Where("follow_id = ? And follower_id = ? ", followerslist[j].ID, claims.Id).Count(&count)
		followerslist[j].IsFollow = (count == 1)
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

	var friends_half []model.Follow
	err = model.DB.Model(&model.Follow{}).Preload("Follower").Preload("Follow").Where("follow_id = ?", service.UserID).Find(&friends_half).Error
	if err != nil {
		util.LogrusObj.Info(err)
		return serializer.FollowListResponse{
			StatusCode: 1,
			StatusMsg:  "查询follow数据库错误",
		}
	}
	var friendslist_half_index []uint
	for _, v := range friends_half {
		friendslist_half_index = append(friendslist_half_index, v.FollowerID)
	}

	var half_friends []model.Follow
	err = model.DB.Model(&model.Follow{}).Preload("Follower").Preload("Follow").Where("follower_id = ?", service.UserID).Find(&half_friends).Error
	if err != nil {
		util.LogrusObj.Info(err)
		return serializer.FollowListResponse{
			StatusCode: 1,
			StatusMsg:  "查询follow数据库错误",
		}
	}
	var half_friendslist_index []uint
	for _, v := range half_friends {
		half_friendslist_index = append(half_friendslist_index, v.FollowID)
	}

	friendslistindex := IntersectArray(friendslist_half_index, half_friendslist_index)
	fmt.Println("friendslist_half_index", friendslist_half_index)
	fmt.Println("half_friendslist_index", half_friendslist_index)
	fmt.Println("friendslistindex", friendslistindex)

	var friendslist []model.User
	model.DB.Model(&model.User{}).Where("id in (?)", friendslistindex).Find(&friendslist)
	//
	//var followerslist []model.User
	//for _,v := range followers{
	//	followerslist = append(followerslist, v.Follow)
	//}
	fmt.Println(friendslist)

	return serializer.FollowListResponse{
		StatusCode: 0,
		StatusMsg:  "查询朋友列表成功",
		UserList:   friendslist,
	}
}

func IntersectArray(a []uint, b []uint) []uint {
	var inter []uint
	mp := make(map[uint]bool)

	for _, s := range a {
		if _, ok := mp[s]; !ok {
			mp[s] = true
		}
	}
	for _, s := range b {
		if _, ok := mp[s]; ok {
			inter = append(inter, s)
		}
	}
	return inter
}
