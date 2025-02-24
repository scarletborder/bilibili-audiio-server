package main

import (
	"bilibiliaudio/constant"
	"bilibiliaudio/ctx"
	"bilibiliaudio/handlers"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func init() {
	constant.LoadEnv()
}

func main() {
	// 接受一个.env路径参数
	if len(os.Args) > 1 {
		constant.LoadEnv(os.Args[1])
	}

	srvCtx := ctx.NewSrvCtx()

	app := fiber.New()
	handlers.RegisterRoutes(app, srvCtx)

	port := os.Getenv("LISTEN_PORT")
	if port == "" {
		port = "10764"
	}
	log.Printf("Listening on port %s", port)
	app.Listen(":" + port)
}
