package router

import (
    "hagavi-otp/handler"
    "database/sql"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
)

func New(db *sql.DB) *fiber.App {
    app := fiber.New()
    app.Use(cors.New())
    api := app.Group("/api")
    api.Post("/register", func (c *fiber.Ctx) error {
        return handler.Register(c, db)
    })
    api.Post("/login", func(c *fiber.Ctx) error {
        return handler.Login(c, db)
    })
	api.Post("/verify_otp", func(c *fiber.Ctx) error {
        return handler.VerifyOTP(c, db)
    })
    api.Post("/resend_otp", func(c *fiber.Ctx) error {
        return handler.ResendOTP(c, db)
    })
    return app
}