package services

import (
	"gatewayH/internal/gateway/config"
	"gatewayH/internal/gateway/store"
	"gatewayH/pkg/idgenerator"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/gocraft/dbr/v2"
	"github.com/sirupsen/logrus"
	"strconv"
)

var (
	Logger *logrus.Logger
	Config *config.Config
)

// Services 处理 Restful API 的服务
type Services struct {
	Store    *store.Store
	Logger   *logrus.Logger
	Cfg      *config.Config
	IDWorker *idgenerator.IdWorker
}

// NewServices is used to create Services
func NewServices(
	config *config.Config,
	mysql *dbr.Connection,
	logger *logrus.Logger,
	redis *redis.Client,
	iw *idgenerator.IdWorker,
) *Services {
	Logger = logger
	Config = config
	return &Services{
		Store:    store.NewStore(config, mysql, logger, redis),
		Logger:   logger,
		Cfg:      config,
		IDWorker: iw,
	}
}

//limit 操作
func Limit(ctx *gin.Context) uint64 {
	limit := ctx.DefaultQuery("page_size", "20")
	num, _ := strconv.ParseUint(limit, 10, 64)
	return num
}

//offset 操作
func Offset(ctx *gin.Context) uint64 {
	offset := ctx.DefaultQuery("page", "1")
	page, _ := strconv.ParseUint(offset, 10, 64)
	if page < 1 {
		page = 1
	}
	return (page - 1) * Limit(ctx)
}
