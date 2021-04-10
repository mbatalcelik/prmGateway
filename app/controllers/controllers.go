package controllers

import (
	prm_controller "prmgateway/app/controllers/prm_controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func RestApies(app *fiber.App) {
	restApiesConfig(app)
	prm_controller.PrmApies(app)
}

func restApiesConfig(app *fiber.App) {
	app.Use(recover.New(), logger.New(), cors.New(), func(ctx *fiber.Ctx) error {
		if ctx.Is("json") {
			return ctx.Next()
		}
		return ctx.Status(fiber.StatusInternalServerError).SendString(" only json input")
	})
}
