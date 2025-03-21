package storage

import (
	"context"
	// "demo/internal/pkg/config"
	// "fmt"
	"github.com/go-redis/redis/v8"
	// "log"
)

var RdRedis *redis.Client
var RdRedisContext = context.Background()

func init() {
	// RdRedis = redis.NewClient(&redis.Options{
	// 	Addr:     config.RdRedis.Addr,
	// 	Password: config.RdRedis.Password,
	// 	DB:       0,
	// })

	// _, err := RdRedis.Ping(RdRedisContext).Result()

	// if err != nil {
	// log.Fatalf("Could not connect to Redis: %v", err)
	// }
	// fmt.Println("Connected to Redis")
}
