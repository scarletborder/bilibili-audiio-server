package login_handler

import (
	"bilibiliaudio/ctx"
	"bilibiliaudio/utils"

	"github.com/CuteReimu/bilibili/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
	qc "github.com/skip2/go-qrcode"
)

func NewQrcodeLogin(c *fiber.Ctx) error {
	srvCtx := utils.GetSrvCtx(c).(*ctx.SrvCtx)
	cache := utils.GetLoginCache(c).(*cache.Cache)

	transaction_id := uuid.New().String()
	qrcode, err := srvCtx.B23_client.GetQRCode()
	if err != nil {
		return c.JSON(fiber.Map{
			"code": 1,
			"msg":  "获取二维码失败",
		})
	}

	SetToCache(cache, transaction_id, map[string]string{
		"url": qrcode.Url,
		"key": qrcode.QrcodeKey,
	})
	return c.JSON(fiber.Map{
		"code":           0,
		"msg":            "获取二维码成功",
		"transaction_id": transaction_id,
	})
}

func GetQrCodeImg(c *fiber.Ctx) error {
	tid := c.Query("tid", "")
	if tid == "" {
		return c.Status(400).SendString("transaction_id is required")
	}

	cache := utils.GetLoginCache(c).(*cache.Cache)
	qrcode, found := GetFromCache(cache, tid)
	if !found {
		return c.Status(404).SendString("transaction_id not found")
	}
	data := qrcode.(map[string]string)

	buf, err := qc.Encode(data["url"], qc.Medium, 256)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	c.Set("Content-Type", "image/png")
	return c.Send(buf)
}

func GetQrcodeLoginStatus(c *fiber.Ctx) error {
	tid := c.Query("tid", "")
	if tid == "" {
		return c.Status(400).SendString("transaction_id is required")
	}

	srvCtx := utils.GetSrvCtx(c).(*ctx.SrvCtx)
	cache := utils.GetLoginCache(c).(*cache.Cache)

	qrcode, found := GetFromCache(cache, tid)
	if !found {
		return c.Status(404).SendString("transaction_id not found")
	}
	data := qrcode.(map[string]string)

	result, err := srvCtx.B23_client.LoginWithQRCode(bilibili.LoginWithQRCodeParam{
		QrcodeKey: data["key"],
	})
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(result)
}
