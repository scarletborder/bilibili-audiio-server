package song_handler

import (
	"bilibiliaudio/ctx"
	"bilibiliaudio/utils"
	"fmt"
	"strconv"

	"github.com/CuteReimu/bilibili/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
)

func GetSongUrlLogic(srvCtx *ctx.SrvCtx, cache *cache.Cache, bvid string, cid string) (string, error) {
	data, found := getUrlFromCache(cache, bvid, cid)
	if found {
		return data, nil
	}

	i_cid, err := strconv.Atoi(cid)
	if err != nil {
		return "", err
	}

	_res, err := srvCtx.B23_client.GetVideoStream(bilibili.GetVideoStreamParam{
		Bvid:  bvid,
		Fnval: 16,
		Cid:   i_cid,
	})

	if err != nil {
		return "", err
	}

	setUrlToCache(cache, bvid, cid, _res.Dash.Audio[0].BaseUrl)

	return _res.Dash.Audio[0].BaseUrl, nil
}

func GetSongUrl(c *fiber.Ctx) error {
	srvCtx := utils.GetSrvCtx(c).(*ctx.SrvCtx)
	cache := utils.GetQueryCache(c).(*cache.Cache)

	bvid := c.Query("bvid", "")
	if bvid == "" {
		return c.Status(400).SendString("bvid is required")
	}
	cid := c.Query("cid", "")
	if cid == "" {
		return c.Status(400).SendString("cid is required")
	}

	url, err := GetSongUrlLogic(srvCtx, cache, bvid, cid)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"code": 0,
		"msg":  "success",
		"url":  url,
	})

}

func getUrlFromCache(c *cache.Cache, bvid string, cid string) (string, bool) {
	key := fmt.Sprintf("bvid:%s:cid:%s", bvid, cid)
	data, found := c.Get(key)
	if found {
		return data.(string), true
	}
	return "", false
}

func setUrlToCache(c *cache.Cache, bvid string, cid string, url string) {
	key := fmt.Sprintf("bvid:%s:cid:%s", bvid, cid)
	c.Set(key, url, cache.DefaultExpiration)
}
