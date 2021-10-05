package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalfian/go-kafka/app/comment"
)

func InitRouter(app *fiber.App) {
	commentProvider := comment.ProvideHandler()

	app.Get("/ping", func(c *fiber.Ctx) error {
		c.Status(200).JSON(&fiber.Map{
			"message": "Pong!",
		})
		return nil
	})

	api := app.Group("/api/v1")

	// Comment
	api.Post("/comment", commentProvider.AddComment)
}
