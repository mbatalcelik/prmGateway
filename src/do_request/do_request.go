package do_reqest

import (
	"prmgateway/src/config"

	"github.com/gofiber/fiber/v2"
)

func DoRequest(methodType string, path string) (int, string) {
	a := fiber.AcquireAgent()
	req := a.Request()
	resp := fiber.AcquireResponse()
	req.Header.SetMethod(methodType)
	configs, _ := config.GetFromEnv()
	req.SetRequestURI(configs.BusinessApp.Url + path)
	a.SetResponse(resp)

	if err := a.Parse(); err != nil {
		panic(err)
	}
	code, body, _ := a.Bytes()
	return code, string(body)
}
