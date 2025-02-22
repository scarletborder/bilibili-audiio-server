package main

import (
	"bilibiliaudio/constant"
	"bilibiliaudio/ctx"
	"bilibiliaudio/handlers"
	"os"

	"github.com/gofiber/fiber/v2"
)

func init() {
	constant.LoadEnv()
}

func main() {
	srvCtx := ctx.NewSrvCtx()

	app := fiber.New()
	handlers.RegisterRoutes(app, srvCtx)

	port := os.Getenv("LISTEN_PORT")
	if port == "" {
		port = "3000"
	}
	app.Listen(":" + port)
}
