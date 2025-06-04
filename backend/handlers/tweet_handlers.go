package handlers

import (
	"sort"
	"strconv"
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

// CreateTweet publica un tweet nuevo.
// POST /tweets
// Body: { "user_id": "usuario1", "text": "Hola mundo" }
func CreateTweet(c *fiber.Ctx) error {
	var req CreateTweetRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	if len(req.Text) == 0 || len(req.Text) > 280 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "tweet must be 1-280 characters"})
	}
	if _, ok := storage.Users[req.UserID]; !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user does not exist"})
	}
	tweet := models.Tweet{
		ID:        uuid.NewString(),
		UserID:    req.UserID,
		Text:      req.Text,
		Timestamp: time.Now(),
	}
	// Indexar el tweet por usuario
	storage.TweetsByUser[req.UserID] = append(storage.TweetsByUser[req.UserID], tweet)
	return c.Status(fiber.StatusCreated).JSON(tweet)
}

type FollowRequest struct {
	FollowerID string `json:"follower_id"`
	FolloweeID string `json:"followee_id"`
}

// FollowUser permite a un usuario seguir a otro.
// POST /follow
// Body: { "follower_id": "usuario1", "followee_id": "usuario2" }
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

// GetTimeline devuelve el timeline de un usuario, ordenado y paginado.
// GET /timeline?user_id=usuario1&page=1&size=10
func GetTimeline(c *fiber.Ctx) error {
	userID := c.Query("user_id")
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user_id is required"})
	}
	if _, ok := storage.Users[userID]; !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user does not exist"})
	}
	// Obtener a quién sigue el usuario
	following := map[string]bool{}
	for _, f := range storage.Follows {
		if f.FollowerID == userID {
			following[f.FolloweeID] = true
		}
	}
	following[userID] = true

	// Obtener los tweets de los usuarios seguidos usando el índice
	var timeline []models.Tweet
	for uid := range following {
		timeline = append(timeline, storage.TweetsByUser[uid]...)
	}

	// Ordenar por fecha descendente
	sort.Slice(timeline, func(i, j int) bool {
		return timeline[i].Timestamp.After(timeline[j].Timestamp)
	})

	// Paginación
	page := 1
	size := 10
	if p := c.Query("page"); p != "" {
		if v, err := strconv.Atoi(p); err == nil && v > 0 {
			page = v
		}
	}
	if s := c.Query("size"); s != "" {
		if v, err := strconv.Atoi(s); err == nil && v > 0 {
			size = v
		}
	}
	start := (page - 1) * size
	end := start + size
	if start > len(timeline) {
		start = len(timeline)
	}
	if end > len(timeline) {
		end = len(timeline)
	}
	paged := timeline[start:end]

	return c.JSON(fiber.Map{
		"page":     page,
		"size":     size,
		"total":    len(timeline),
		"timeline": paged,
	})
}

type CreateUserRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// CreateUser crea un usuario nuevo.
// POST /users
// Body: { "id": "usuario1", "name": "Nombre" }
func CreateUser(c *fiber.Ctx) error {
	var req CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	if req.ID == "" || req.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "id and name are required"})
	}
	if _, exists := storage.Users[req.ID]; exists {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "user already exists"})
	}
	user := models.User{
		ID:   req.ID,
		Name: req.Name,
	}
	storage.Users[req.ID] = user
	return c.Status(fiber.StatusCreated).JSON(user)
}
