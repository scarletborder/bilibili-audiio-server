package utils

import "github.com/gofiber/fiber/v2"

// fiber ctx
func GetSrvCtx(c *fiber.Ctx) interface{} {
	return c.Locals("srvCtx")
}

func GetQueryCache(c *fiber.Ctx) interface{} {
	return c.Locals("queryCache")
}

func GetLoginCache(c *fiber.Ctx) interface{} {
	return c.Locals("loginCache")
}
