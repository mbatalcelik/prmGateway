package main

import (
	"fmt"
	config "prmgateway/src/config"
	controllers "prmgateway/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	configs, _ := config.GetFromEnv()
	restApiInstance := fiber.New()
	controllers.RestApies(restApiInstance)
	staticContentConfig(restApiInstance)
	restApiInstance.Listen(fmt.Sprintf(":%s", configs.MyApp.Port))
}

func staticContentConfig(app *fiber.App) {
	app.Static("/src", "./src/public")
}
