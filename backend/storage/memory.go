package storage

import (
	"tvitteruala-backend/models"
)

var (
	Users = map[string]models.User{}
	// Tweets indexados por user_id
	TweetsByUser = map[string][]models.Tweet{}
	Follows      = []models.Follow{}
)
