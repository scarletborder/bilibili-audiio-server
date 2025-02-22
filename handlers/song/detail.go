package song_handler

import (
	"bilibiliaudio/ctx"
	"bilibiliaudio/utils"
	"fmt"

	"github.com/CuteReimu/bilibili/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
)

// GET /song/detail

func SongDetail(c *fiber.Ctx) error {
	bvid := c.Query("bvid", "")
	if bvid == "" {
		return c.Status(400).SendString("bvid is required")
	}

	srvCtx := utils.GetSrvCtx(c).(*ctx.SrvCtx)
	cache := utils.GetQueryCache(c).(*cache.Cache)

	song, err := GetSongByBvid(srvCtx, cache, bvid)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(song)
}

func GetSongByBvid(srvCtx *ctx.SrvCtx, cache *cache.Cache, bvid string) (map[string]interface{}, error) {
	data, found := getDetailFromCache(cache, bvid)
	if found {
		return data.(map[string]interface{}), nil
	}

	ret := map[string]interface{}{}

	info, err := srvCtx.B23_client.GetVideoInfo(bilibili.VideoParam{
		Bvid: bvid,
	})

	if err != nil {
		return nil, err
	}

	page, err := srvCtx.B23_client.GetVideoPageList(bilibili.VideoParam{
		Bvid: bvid,
	})

	if err != nil {
		return nil, err
	}

	ret["title"] = info.Title
	ret["aid"] = info.Aid
	ret["bvid"] = info.Bvid
	ret["desc"] = info.Desc
	ret["cover"] = info.Pic
	ret["artist"] = map[string]interface{}{
		"id":   info.Owner.Mid,
		"name": info.Owner.Name,
	}
	ret["songs"] = [](map[string]interface{}){}

	for _, v := range page {
		ret["songs"] = append(ret["songs"].([]map[string]interface{}), map[string]interface{}{
			"cid":  v.Cid,
			"name": v.Part,
		})
	}

	setDetailToCache(cache, bvid, ret)
	return ret, nil
}

func getDetailFromCache(cache *cache.Cache, bvid string) (interface{}, bool) {
	key := fmt.Sprintf("bvid:%s", bvid)
	data, found := cache.Get(key)
	if !found {
		return nil, false
	}
	return data, true
}

func setDetailToCache(c *cache.Cache, bvid string, value interface{}) {
	key := fmt.Sprintf("bvid:%s", bvid)
	c.Set(key, value, cache.DefaultExpiration)
}
