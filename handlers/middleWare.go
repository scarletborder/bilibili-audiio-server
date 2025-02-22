package handlers

import (
	"bilibiliaudio/ctx"

	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
)

func SrvCtxMiddleWare(ctx *ctx.SrvCtx) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("srvCtx", ctx)
		return c.Next()
	}
}

func QueryCacheMiddleWare(cache *cache.Cache) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("queryCache", cache)
		return c.Next()
	}
}
