package login_handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
)

func NewLoginCache() *cache.Cache {
	return cache.New(3*time.Minute, 2*time.Minute)
}

// 获取登录事务
func GetFromCache(c *cache.Cache, key string) (interface{}, bool) {
	// 专有cache不用再拼接字符了
	return c.Get(key)
}

// 设置登录事务
func SetToCache(c *cache.Cache, key string, value interface{}) {
	c.Set(key, value, cache.DefaultExpiration)
}

func CacheMiddware(cache *cache.Cache) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("loginCache", cache)
		return c.Next()
	}
}
