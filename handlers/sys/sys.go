package sys_handler

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func SysInfo(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"version": os.Getenv("VERSION"),
	})
}
