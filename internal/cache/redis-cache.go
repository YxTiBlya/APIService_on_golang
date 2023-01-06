package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/yxtiblya/internal/cfg"
)

var ctx = context.Background()

// returning redis client for next work
func getClient() *redis.Client {
	config := cfg.GetConfig()
	return redis.NewClient(&redis.Options{
		Addr:     config.RedisHost + ":" + config.RedisPort,
		Password: config.RedisPass,
		DB:       config.RedisDb,
	})
}

// redis `SET key value [expiration]` command.
func Set(key string, value interface{}) error {
	rdb := getClient()

	json, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = rdb.Set(ctx, key, json, time.Second*cfg.GetConfig().RedisExp).Err()
	if err != nil {
		return err
	}
	return nil
}

// redis `GET key` command. It returns redis.Nil error when key does not exist.
func Get(key string) ([]map[string]interface{}, error) {
	rdb := getClient()

	value, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var ans []map[string]interface{}
	if err = json.Unmarshal([]byte(value), &ans); err != nil {
		return nil, err
	}

	return ans, nil
}

func Del(keys ...string) {
	rdb := getClient()
	rdb.Del(ctx, keys...)
}
