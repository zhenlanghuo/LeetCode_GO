package main

import "github.com/go-redis/redis"

func main() {
	redisCli := redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "",
		DB:       0, // use default DB
	})

	pubsub := redisCli.Subscribe()
	pubsub.Channel()

	redisCli.MGet()
}
