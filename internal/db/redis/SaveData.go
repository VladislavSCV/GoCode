package redis

import (
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func SetData(key string, value interface{}) (error) {
	err := rdb.Set(ctx, key, value, time.Hour * 24).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("Saved Data!")
	return nil
}

func GetData(key string) (interface{}, error) {
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		fmt.Printf("%v does not exist \n", key)
	} else if err != nil {
		panic(err)
	} else {
		// fmt.Println(key, val)
	}
	return val, nil
}
