package mysql

import (
	"time"

	log "github.com/ningzining/L-log"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Dsn                   string
	MaxIdleConnections    int
	MaxOpenConnections    int
	MaxConnectionLifeTime time.Duration
	Logger                logger.Interface
}

func (c *Config) complete() {
	if c.MaxIdleConnections <= 0 {
		c.MaxIdleConnections = 100
	}
	if c.MaxOpenConnections <= 0 {
		c.MaxOpenConnections = 100
	}
	if c.MaxConnectionLifeTime <= 0 {
		c.MaxConnectionLifeTime = 30 * time.Minute
	}
}

func NewMysqlClient(cfg Config) (*gorm.DB, error) {
	cfg.complete()
	// 连接mysql并且获取实例
	client, err := gorm.Open(mysql.Open(cfg.Dsn), &gorm.Config{
		Logger: cfg.Logger,
	})
	if err != nil {
		log.Fatalf("mysql连接异常: %s", err.Error())
		return nil, err
	}

	sqlDB, err := client.DB()
	if err != nil {
		log.Fatalf("mysql获取实例异常: %s", err.Error())
		return nil, err
	}
	// 设置连接池的空闲连接数
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConnections)
	// 设置连接池的最大连接数
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConnections)
	// 设置连接最大存活时间
	sqlDB.SetConnMaxLifetime(cfg.MaxConnectionLifeTime)

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("mysql ping失败: %s", err.Error())
		return nil, err
	}
	log.Infof("mysql初始化成功", zap.String("dsn", cfg.Dsn))

	return client, nil
}
