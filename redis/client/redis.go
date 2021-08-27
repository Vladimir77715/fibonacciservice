package client

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

type RedisClientWrapper struct {
	redisContext *context.Context
	redisClient  *redis.Client
	redisMutex   sync.Mutex
}

type RedisEntity struct {
	Key   string
	Value string
}

func InitRedisClient(ctx *context.Context, flushData bool) (*RedisClientWrapper, error) {
	var (
		addr     = "redis:6379"
		password = ""
		db       = 0
	)
	if adr := os.Getenv("REDIS_ADDR"); adr != "" {
		log.Printf("REDIS_ADDR is: %v", adr)
		addr = adr
	}
	if pass := os.Getenv("REDIS_PASS"); pass != "" {
		log.Printf("REDIS_ADDR is: %v", pass)
		password = pass
	}
	if _db, err := strconv.Atoi(os.Getenv("REDIS_DB")); err == nil {
		log.Printf("REDIS_DB is: %v", _db)
		db = _db
	}
	log.Printf("Redis client trying to connect to %v with pass %v ", addr, password)
	var redisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db})
	rcw := RedisClientWrapper{
		redisClient: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       db}),
		redisContext: ctx,
	}
	if scmd := redisClient.Ping(*ctx); scmd.Err() != nil {
		return nil, scmd.Err()
	}
	if flushData {
		if scmd := rcw.redisClient.FlushAll(*ctx); scmd.Err() != nil {
			return nil, scmd.Err()
		}
	}
	return &RedisClientWrapper{
		redisClient:  redisClient,
		redisContext: ctx,
	}, nil
}

func (rcw *RedisClientWrapper) Set(key string, value interface{}, expire time.Duration) error {
	rcw.redisMutex.Lock()
	ctx, cancel := context.WithTimeout(*rcw.redisContext, time.Second*20)
	defer func() {
		cancel()
		rcw.redisMutex.Unlock()
	}()
	if err := rcw.redisClient.Set(ctx, key, value, expire).Err(); err != nil {
		return err
	}
	return nil
}
func (rcw *RedisClientWrapper) Get(key string) (*RedisEntity, error) {
	rcw.redisMutex.Lock()
	ctx, cancel := context.WithTimeout(*rcw.redisContext, time.Second*20)
	defer func() {
		cancel()
		rcw.redisMutex.Unlock()
	}()
	var scmd = rcw.redisClient.Get(ctx, key)
	if scmd.Err() != nil {
		return nil, scmd.Err()
	}
	return &RedisEntity{Key: key, Value: scmd.Val()}, nil
}
