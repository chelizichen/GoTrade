package component_cache

import (
	"com_sgrid_gotrade/src/framework/config"
	"fmt"

	"github.com/go-redis/redis"
)

var Cache *redis.Client

func LoadCache(ctx *config.SgridConf) error {
	redis_addr := ctx.GetString("redis_addr")
	redis_pass := ctx.GetString("redis_pass")
	Cache = redis.NewClient(&redis.Options{
		Addr:     redis_addr,
		Password: redis_pass,
		DB:       0,
	})
	pong, err := Cache.Ping().Result()
	if err != nil {
		return fmt.Errorf("Error")
	}
	fmt.Println("conn success! ", pong)
	return nil
}
