package storage

import (
	"time"

	"github.com/patrickmn/go-cache"
)

func NewCache() *cache.Cache {
	return cache.New(5*time.Minute, 3*time.Minute)
}
