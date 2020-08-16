package config

import (
	"fmt"
	"github.com/go-redis/redis"
	. "github.com/itsursujit/fiber-boilerplate/app"
)

type CacheConfiguration struct {
	Cache_DSN string
	Cache_DB  int
}

var CacheConfig *CacheConfiguration //nolint:gochecknoglobals

func LoadCacheConfig() {
	loadDefaultCacheConfig()
	ViperConfig.Unmarshal(&CacheConfig)
	option, err := redis.ParseURL(fmt.Sprintf("redis://%s/%d", CacheConfig.Cache_DSN, CacheConfig.Cache_DB))
	if err != nil {
		panic(err)
	}
	RedisClient = redis.NewClient(option)
}

func loadDefaultCacheConfig() {
	ViperConfig.SetDefault("CACHE_DSN", "127.0.0.1:6379")
	ViperConfig.SetDefault("CACHE_DB", 0)
}
