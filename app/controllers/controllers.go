package controllers

import (
	prm_controller "prmgateway/app/controllers/prm_controller"
	controller_middleware "prmgateway/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func RestApies(app *fiber.App) {
	controller_middleware.RestApiesConfig(app)
	prm_controller.PrmApies(app)
}
