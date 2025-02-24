package playlist_handler

import (
	"bilibiliaudio/ctx"
	song_handler "bilibiliaudio/handlers/song"
	"bilibiliaudio/utils"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
)

func EasyViewHandler(c *fiber.Ctx) error {
	ctx := utils.GetSrvCtx(c).(*ctx.SrvCtx)
	cache := utils.GetQueryCache(c).(*cache.Cache)
	mlid := c.Query("mlid", "")

	if mlid == "" {
		return c.Status(fiber.StatusBadRequest).SendString("mlid is required")
	}
	mlid_int, err := strconv.Atoi(mlid)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("mlid is invalid")
	}

	var ret []map[string]interface{}
	page_idx := 0
	for {
		page, err := getPlaylistLogic(ctx, cache, mlid_int, page_idx)
		if err != nil || len(page) == 0 {
			break
		}

		for _, video_detail := range page {
			video, err := song_handler.GetSongByBvid(ctx, cache, video_detail.Bvid)
			if err != nil {
				continue
			}

			for _, part := range video["songs"].([]map[string]interface{}) {
				url, err := song_handler.GetSongUrlLogic(ctx, cache, video_detail.Bvid, part["cid"].(int))
				if err != nil {
					break
				}
				ret = append(ret, map[string]interface{}{
					"id":      fmt.Sprintf("%s_%d", video_detail.Bvid, part["cid"].(int)),
					"name":    part["name"],
					"album":   video["title"],
					"cover":   video["cover"],
					"url":     url,
					"artists": video["artists"],
				})
			}
		}

		page_idx++
	}

	return c.JSON(fiber.Map{
		"code": 0,
		"msg":  "success",
		"data": ret,
	})
}
