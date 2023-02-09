package model

type Follow struct {
	ID         uint `json:"id"`
	Follow     User `json:"follow"`
	FollowID   uint `json:"follow_id"`
	Follower   User `json:"follower"`
	FollowerID uint `json:"follower_id"`
}
