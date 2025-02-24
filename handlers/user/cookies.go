package user_handler

import (
	"bilibiliaudio/ctx"
	"bilibiliaudio/utils"

	"github.com/gofiber/fiber/v2"
)

func GetBiliCookies(c *fiber.Ctx) error {
	ctx := utils.GetSrvCtx(c).(*ctx.SrvCtx)
	t := ctx.B23_client.GetCookiesString()
	_ = t
	return c.JSON(fiber.Map{
		"code":    0,
		"cookies": ctx.B23_client.GetCookies(),
	})
}
