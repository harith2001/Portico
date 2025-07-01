package config

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func ExampleClient_connect_basic() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-16866.c91.us-east-1-3.ec2.redns.redis-cloud.com:16866",
		Username: "default",
		Password: "SdBAb8H95Yu4YX7cYAvlRW8oAN30ggYf",
		DB:       0,
	})

	rdb.Set(ctx, "foo", "bar", 0)
	result, err := rdb.Get(ctx, "foo").Result()

	if err != nil {
		panic(err)
	}

	fmt.Println(result) // >>> bar

}
