package playlist_handler

import (
	"bilibiliaudio/ctx"
	"bilibiliaudio/model"
	"bilibiliaudio/utils"
	"fmt"
	"strconv"

	"github.com/CuteReimu/bilibili/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
)

type Video struct {
	Aid  int    `json:"aid"`
	Bvid string `json:"bvid"`
}

func GetPlaylist(c *fiber.Ctx) error {
	srvCtx := utils.GetSrvCtx(c).(*ctx.SrvCtx)
	cache := utils.GetQueryCache(c).(*cache.Cache)

	mlid, err := strconv.Atoi(c.Query("mlid", ""))
	if err != nil {
		return c.Status(400).SendString("mlid is required")
	}
	page, err := strconv.Atoi(c.Query("page", "0"))
	if err != nil {
		return c.Status(400).SendString("page is required")
	}

	ret, err := getPlaylistLogic(srvCtx, cache, mlid, page)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"code": 0,
		"msg":  "success",
		"data": ret,
	})

}

func getPlaylistLogic(srvCtx *ctx.SrvCtx, c *cache.Cache, mlid int, page int) (ret []model.SongDetail, err error) {
	key := fmt.Sprintf("playlist:%d:page:%d", mlid, page)
	data, found := getPlaylistFromCache(c, key)
	if found {
		ret = data.([]model.SongDetail)
		return
	}

	// get from sdk
	page += 1 // 兼容b站page
	playlist, err := srvCtx.B23_client.GetFavourList(bilibili.GetFavourListParam{
		MediaId: mlid,
		Ps:      20,
		Pn:      page,
	})

	if err != nil {
		return nil, err
	}

	for _, v := range playlist.Medias {
		ret = append(ret, model.SongDetail{
			Aid:   v.Id,
			Bvid:  v.Bvid,
			Title: v.Title,
			Cover: v.Cover,
			Artist: model.Artist{
				Id:   v.Upper.Mid,
				Name: v.Upper.Name,
			},
			HasPart: v.Page,
		})
	}

	setPlaylistToCache(c, key, ret)
	return
}

func getPlaylistFromCache(c *cache.Cache, key string) (interface{}, bool) {
	return c.Get(key)
}

func setPlaylistToCache(c *cache.Cache, key string, value interface{}) {
	c.Set(key, value, cache.DefaultExpiration)
}
