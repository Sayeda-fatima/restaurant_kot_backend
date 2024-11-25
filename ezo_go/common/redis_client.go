package common

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type RedisClient interface{
	Get(key string) (string, error)
	Set(key string, value interface{}, expiration time.Duration) error
	Delete(key string) error
}

type redisClient struct{
	Client *redis.Client
}

func NewRedisClient(host string, port int, password string, db int) RedisClient{

	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", host, port),
		Password: password,
		DB: db,
	})

	// test connection
	_, err := rdb.Ping(ctx).Result()
	if err != nil{
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}
	return &redisClient{rdb}
}

func (rc *redisClient) Get(key string) (string, error){

	return rc.Client.Get(ctx, key).Result()
}

func (rc *redisClient) Set(key string, value interface{}, expiration time.Duration) error{

	return rc.Client.Set(ctx, key, value, expiration).Err()
}

func (rc *redisClient) Delete(key string) error{

	return rc.Client.Del(ctx, key).Err()
}