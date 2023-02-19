package common

import (
	"context"
	"github.com/redis/go-redis/v9"
	"strconv"
)

var ctx = context.Background()
var rdb *redis.Client

func init() {
	db, _ := strconv.Atoi(config.Redis.Database)
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Host,
		Password: "",
		DB:       db,
	})
}

func GetRedisClient() *redis.Client {
	return rdb
}
