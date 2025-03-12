package mysql

import (
	"time"

	log "github.com/ningzining/L-log"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	stdlog "log"
)

func NewClient(dsn string, opts ...Option) (*gorm.DB, error) {
	o := defaultOptions()
	for _, opt := range opts {
		opt(o)
	}

	// 连接mysql并且获取实例
	client, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(stdlog.New(o.Writer, "\r\n", stdlog.LstdFlags), logger.Config{
			SlowThreshold:             time.Millisecond * 200,
			Colorful:                  true,
			IgnoreRecordNotFoundError: false,
			ParameterizedQueries:      false,
			LogLevel:                  logger.LogLevel(log.Opts().Level.LevelForGorm()),
		}),
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
	sqlDB.SetMaxIdleConns(o.MaxIdleConnections)
	// 设置连接池的最大连接数
	sqlDB.SetMaxOpenConns(o.MaxOpenConnections)
	// 设置连接最大存活时间
	sqlDB.SetConnMaxLifetime(o.MaxConnectionLifeTime)

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("mysql ping异常: %s", err.Error())
		return nil, err
	}
	log.Info("mysql初始化成功", zap.String("dsn", dsn))

	return client, nil
}
