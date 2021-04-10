package main

import (
	"fmt"
	config "prmgateway/src/config"
	controllers "prmgateway/src/controllers"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func main() {
	configs, _ := config.GetFromEnv()
	fiberInstance := fiber.New()
	swaggerConfig(fiberInstance)
	controllers.RestApies(fiberInstance)
	staticContentConfig(fiberInstance)
	fiberInstance.Listen(fmt.Sprintf(":%s", configs.MyApp.Port))
}

func staticContentConfig(app *fiber.App) {
	app.Static("/src", "./src/public")
}

func swaggerConfig(app *fiber.App) {
	app.Get("/swagger/*", swagger.Handler) // default
}
