package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	log "github.com/ningzining/L-log"
	"go.uber.org/zap"
)

type Config struct {
	Host     string
	Password string
	DB       int
}

func NewRedisClient(cfg Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:        cfg.Host,
		Password:    cfg.Password,
		DB:          cfg.DB,
		DialTimeout: time.Minute,
		ReadTimeout: time.Minute,
		IdleTimeout: time.Minute,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("redis连接异常: %s", err.Error())
		return nil, err
	}
	log.Infof("redis初始化成功", zap.String("host", cfg.Host), zap.Int("db", cfg.DB))

	return client, nil
}
