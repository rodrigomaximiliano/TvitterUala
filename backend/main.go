package main

import (
	"tvitteruala-backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Ruta de prueba
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("¡TvitterUala backend funcionando!")
	})

	app.Post("/users", handlers.CreateUser)
	app.Post("/tweets", handlers.CreateTweet)
	app.Post("/follow", handlers.FollowUser)
	app.Get("/timeline", handlers.GetTimeline)

	app.Listen(":3000")
}
