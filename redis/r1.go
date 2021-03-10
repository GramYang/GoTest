package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

const ADDR = "106.54.87.204:8095"
const PASSWORD = "cqdq54123"

var ctx = context.Background()

func main() {
	c1()
}

func c1() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     ADDR,
		Password: PASSWORD,
		DB:       0,
	})
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}
