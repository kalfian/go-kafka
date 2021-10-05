package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalfian/go-kafka/routes"
)

type App struct {
}

func main() {

	app := fiber.New()
	routes.InitRouter(app)

	app.Listen(":3000")
}
