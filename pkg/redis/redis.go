package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

var ctx = context.Background()

type DB struct {
	*redis.Client
}

func NewDB(addr string, password string, db int) (*DB, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("redis client ping result: %w", err)
	}

	return &DB{Client: client}, nil
}

func (db *DB) Set(key string, value interface{}, expiration time.Duration) error {
	return db.Client.Set(ctx, key, value, expiration).Err()
}

func (db *DB) Get(key string) (string, error) {
	value, err := db.Client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return "", nil
	}

	return value, err
}

func (db *DB) Delete(keys ...string) error {
	return db.Client.Del(ctx, keys...).Err()
}
