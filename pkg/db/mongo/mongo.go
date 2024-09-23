package mongo

import (
	"context"

	log "github.com/ningzining/L-log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
)

type Config struct {
	Uri string
}

func NewClient(cfg Config) (*mongo.Client, error) {
	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.Uri))
	if err != nil {
		log.Fatalf("mongo启动失败: %s", err.Error())
		return nil, err
	}
	if err := mongoClient.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatalf("mongo ping异常: %s", err.Error())
		return nil, err
	}

	log.Infof("mongo初始化成功", zap.String("uri", cfg.Uri))

	return mongoClient, nil
}
