package db

import (
	settings "github.com/ewenliu/gin-example/7.Gorm/env"
	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func GetRedisClient() *redis.Client{
	return rdb
}

func init()  {
	opt, err := redis.ParseURL(settings.Env["Redis"].(string))
	if err != nil {
		panic(err)
	}

	rdb = redis.NewClient(opt)
}
