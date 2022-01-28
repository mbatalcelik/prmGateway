package main

import (
	"fmt"
	controllers "prmgateway/app/controllers"
	config "prmgateway/pkg/config"
	middleware "prmgateway/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func main() {
	configs, _ := config.GetFromEnv()
	fiberInstance := fiber.New()
	middleware.SwaggerConfig(fiberInstance)
	controllers.RestApies(fiberInstance)
	middleware.StaticContentConfig(fiberInstance)
	fiberInstance.Listen(fmt.Sprintf(":%s", configs.MyApp.Port))
}
