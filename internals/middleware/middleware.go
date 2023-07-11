package middleware

import "github.com/redis/go-redis/v9"

var redisClient *redis.Client

func NewClient(r *redis.Client) {
	redisClient = r
}
