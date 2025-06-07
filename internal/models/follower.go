package models

import "github.com/jumaevkova04/posts/internal/helpers"

type Follower struct {
	ID          string `json:"id" db:"id"`
	FollowerID  string `json:"follower_id" db:"follower_id"`
	FollowingID string `json:"following_id" db:"following_id"`
	CreatedAt   string `json:"created_at" db:"created_at"`
}

func NewFollower(followerID string, followingID string) *Follower {
	return &Follower{
		ID:          helpers.NewUUID(),
		FollowerID:  followerID,
		FollowingID: followingID,
		CreatedAt:   helpers.TimeNowWithRFC3339Format(),
	}
}
