package handlers

import (
	"bilibiliaudio/ctx"
	login_handler "bilibiliaudio/handlers/login"
	playlist_handler "bilibiliaudio/handlers/playlist"
	song_handler "bilibiliaudio/handlers/song"
	sys_handler "bilibiliaudio/handlers/sys"
	user_handler "bilibiliaudio/handlers/user"
	"bilibiliaudio/storage"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, ctx *ctx.SrvCtx) {
	app.Use(SrvCtxMiddleWare(ctx))
	login_cache := login_handler.NewLoginCache() // login cache
	query_cache := storage.NewCache()            // general query cache

	sysGroup := app.Group("/sys")
	sysGroup.Get("/info", sys_handler.SysInfo)

	loginGroup := app.Group("/login")
	loginGroup.Use(login_handler.CacheMiddware(login_cache))
	loginGroup.Get("/qrcode", login_handler.NewQrcodeLogin)
	loginGroup.Get("/qrcode_img", login_handler.GetQrCodeImg)
	loginGroup.Get("/qrcode_status", login_handler.GetQrcodeLoginStatus)

	songGroup := app.Group("/song")
	songGroup.Use(QueryCacheMiddleWare(query_cache))
	songGroup.Get("/detail", song_handler.SongDetail)
	songGroup.Get("/url", song_handler.GetSongUrl)

	userGroup := app.Group("/user")
	userGroup.Use(QueryCacheMiddleWare(query_cache))
	userGroup.Get("/cookies", user_handler.GetBiliCookies)

	playlistGroup := app.Group("/playlist")
	playlistGroup.Use(QueryCacheMiddleWare(query_cache))
	playlistGroup.Get("/detail", playlist_handler.GetPlaylist)

}
