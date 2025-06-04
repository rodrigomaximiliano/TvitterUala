package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"tvitteruala-backend/models"
	"tvitteruala-backend/storage"
)

type CreateTweetRequest struct {
	UserID string `json:"user_id"`
	Text   string `json:"text"`
}

func CreateTweet(c *fiber.Ctx) error {
	var req CreateTweetRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	if len(req.Text) == 0 || len(req.Text) > 280 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "tweet must be 1-280 characters"})
	}
	// Opcional: validar que el usuario exista
	if _, ok := storage.Users[req.UserID]; !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user does not exist"})
	}
	tweet := models.Tweet{
		ID:        uuid.NewString(),
		UserID:    req.UserID,
		Text:      req.Text,
		Timestamp: time.Now(),
	}
	storage.Tweets = append(storage.Tweets, tweet)
	return c.Status(fiber.StatusCreated).JSON(tweet)
}

type FollowRequest struct {
	FollowerID string `json:"follower_id"`
	FolloweeID string `json:"followee_id"`
}

func FollowUser(c *fiber.Ctx) error {
	var req FollowRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	if req.FollowerID == req.FolloweeID {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot follow yourself"})
	}
	// Validar que ambos usuarios existan
	if _, ok := storage.Users[req.FollowerID]; !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "follower does not exist"})
	}
	if _, ok := storage.Users[req.FolloweeID]; !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "followee does not exist"})
	}
	// Verificar si ya sigue
	for _, f := range storage.Follows {
		if f.FollowerID == req.FollowerID && f.FolloweeID == req.FolloweeID {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "already following"})
		}
	}
	storage.Follows = append(storage.Follows, models.Follow{
		FollowerID: req.FollowerID,
		FolloweeID: req.FolloweeID,
	})
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "followed"})
}

// ...aquí irán las funciones para manejar los endpoints...
