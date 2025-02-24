package media_handler

import (
	"bilibiliaudio/ctx"
	song_handler "bilibiliaudio/handlers/song"
	"bilibiliaudio/utils"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
)

func BvidCidHandler(c *fiber.Ctx) error {
	ctx := utils.GetSrvCtx(c).(*ctx.SrvCtx)
	cache := utils.GetQueryCache(c).(*cache.Cache)
	bvidcid := c.Params("bvidcid", "")
	if bvidcid == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "need bvidcid",
		})
	}
	sp := strings.Split(bvidcid, "_")
	if len(sp) != 2 {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid bvidcid",
		})
	}
	bvid := sp[0]
	cid := sp[1]
	cid_int, err := strconv.Atoi(cid)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid bvidcid",
		})
	}
	url, err := song_handler.GetSongUrlLogic(ctx, cache, bvid, cid_int)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	bytes, err := getBiliM4s(url)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	setResponseHeader(c, bvidcid)
	return c.Send(bytes)
}
