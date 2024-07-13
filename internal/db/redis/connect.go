package redis

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var (
	ctx = context.Background()

	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "VLADISLAV@31415926", // no password set
		DB:       0,                    // use default DB
	})
)

func Connect() {
	_, err := rdb.Ping(ctx).Result()
	if err == nil {
		log.Println("Successfully connected!")
	}
}

func ClearData() (bool) {
	err := rdb.FlushDB(context.Background()).Err()
	if err != nil {
		return false
	}
	return true
}