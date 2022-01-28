package middleware

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func StaticContentConfig(app *fiber.App) {
	app.Static("/src", "./resource")
}

func SwaggerConfig(app *fiber.App) {
	app.Get("/swagger/*", swagger.Handler) // default
}
