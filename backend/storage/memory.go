package storage

import (
	"tvitteruala-backend/models"
)

var (
	Users   = map[string]models.User{}
	Tweets  = []models.Tweet{}
	Follows = []models.Follow{}
)
