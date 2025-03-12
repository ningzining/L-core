package mysql

import (
	"gorm.io/gorm/logger"
	"io"
	"os"
	"time"
)

type Options struct {
	Writer io.Writer

	Level                 logger.LogLevel
	MaxIdleConnections    int
	MaxOpenConnections    int
	MaxConnectionLifeTime time.Duration
}

type Option func(option *Options)

func defaultOptions() *Options {
	return &Options{
		Writer:                os.Stdout,
		Level:                 logger.Info,
		MaxIdleConnections:    10,
		MaxOpenConnections:    20,
		MaxConnectionLifeTime: time.Second * 60,
	}
}

func WithLevel(level logger.LogLevel) Option {
	return func(option *Options) {
		option.Level = level
	}
}

func WithWriter(writer io.Writer) Option {
	return func(option *Options) {
		option.Writer = writer
	}
}

func WithMaxIdleConnections(maxIdleConnections int) Option {
	return func(opt *Options) {
		opt.MaxIdleConnections = maxIdleConnections
	}
}

func WithMaxOpenConnections(maxOpenConnections int) Option {
	return func(opt *Options) {
		opt.MaxOpenConnections = maxOpenConnections
	}
}

func WithMaxConnectionLifeTime(maxConnectionLifeTime time.Duration) Option {
	return func(opt *Options) {
		opt.MaxConnectionLifeTime = maxConnectionLifeTime
	}
}
