package models

import "time"

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Tweet struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
}

type Follow struct {
	FollowerID string `json:"follower_id"`
	FolloweeID string `json:"followee_id"`
}
