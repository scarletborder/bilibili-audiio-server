package playlist_handler

import (
	"bilibiliaudio/ctx"
	"bilibiliaudio/utils"
	"net/http"
	"strconv"

	"github.com/CuteReimu/bilibili/v2"
	"github.com/gofiber/fiber/v2"
)

// 列出指定用户 所有收藏夹

func ListAllHandler(c *fiber.Ctx) error {
	ctx := utils.GetSrvCtx(c).(*ctx.SrvCtx)
	// 获取用户id
	upMid := c.Query("up_mid")
	var (
		upMidInt int
		found    bool
		err      error
	)

	if upMid == "" {
		// 试图从cookies中获取用户id
		upMidInt, found = ExtractMidFromCookies(ctx.B23_client.GetCookies())
		if !found {
			return c.Status(400).SendString("up_mid is required or cookies is a loginned user state")
		}
	} else {
		upMidInt, err = strconv.Atoi(upMid)
		if err != nil {
			return c.Status(400).SendString("up_mid is invalid")
		}
	}

	res, err := ctx.B23_client.GetAllFavourFolderInfo(bilibili.GetAllFavourFolderInfoParam{
		UpMid: upMidInt,
	})

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	if res.List == nil {
		return c.JSON([]map[string]interface{}{})
	}

	ret := []map[string]interface{}{}
	for _, data := range res.List {
		ret = append(ret, map[string]interface{}{
			"id":    data.Id,
			"title": data.Title,
			"count": data.MediaCount,
		})
	}
	return c.JSON(ret)
}

func ExtractMidFromCookies(cookie []*http.Cookie) (int, bool) {
	for _, ck := range cookie {
		if ck.Name == "DedeUserID" {
			ckv, err := strconv.Atoi(ck.Value)
			if err != nil {
				return 0, false
			}
			return ckv, true
		}
	}
	return 0, false
}
