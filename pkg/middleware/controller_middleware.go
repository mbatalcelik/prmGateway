package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func RestApiesConfig(app *fiber.App) {
	app.Use(recoverConfig(), requestidConfig(), loggerConfig(), corsConfig(), limiterConfig(), func(ctx *fiber.Ctx) error {
		if ctx.Is("json") {
			return ctx.Next()
		}
		return ctx.Status(fiber.StatusInternalServerError).SendString(" only json input")
	})
}

func loggerConfig() fiber.Handler {
	return logger.New(logger.Config{
		Next:         nil,
		Format:       "[${time}] ${pid} ${locals:requestid} ${status} - ${latency} ${method} ${path}Request Body:\n${body}\n Response Body:\n${resBody}\n",
		TimeFormat:   "2006-01-02T15:04:05-0700",
		TimeZone:     "Local",
		TimeInterval: 1,
	})
}

func requestidConfig() fiber.Handler {
	return requestid.New()
}

func recoverConfig() fiber.Handler {
	return recover.New()
}

func corsConfig() fiber.Handler {
	return cors.New()
}

func limiterConfig() fiber.Handler {
	return limiter.New()
}
