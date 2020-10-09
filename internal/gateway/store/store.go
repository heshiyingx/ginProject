package store

import (
	"gatewayH/internal/gateway/config"
	"github.com/go-redis/redis/v8"
	"github.com/gocraft/dbr/v2"
	"github.com/sirupsen/logrus"
)

// Store 负责数据层读取
type Store struct {
	Config  *config.Config
	Connect *dbr.Connection
	Logger  *logrus.Logger
	Redis   *redis.Client
}

// NewStore is used to create store
func NewStore(config *config.Config, connect *dbr.Connection, logger *logrus.Logger, redis *redis.Client) *Store {
	return &Store{
		Config:  config,
		Connect: connect,
		Logger:  logger,
		Redis:   redis,
	}
}
