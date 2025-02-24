package login_handler

import (
	"bilibiliaudio/ctx"
	"bilibiliaudio/utils"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func SetCookiesHandler(c *fiber.Ctx) error {
	ctx := utils.GetSrvCtx(c).(*ctx.SrvCtx)
	var body map[string]interface{}

	test := string(c.Body())
	_ = test
	err := json.Unmarshal([]byte(c.Body()), &body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid json",
		})
	}
	// 处理 cookies
	ctx.SetBiliCookies(body["cookies"].([]interface{}))

	return c.JSON(fiber.Map{
		"message": "cookies received",
	})
}
