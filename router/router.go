package router

import (
 
    "github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
    app := fiber.New()
    api := app.Group("/api")
    auth := api.Group("/auth")

	// TODO

    return app
}