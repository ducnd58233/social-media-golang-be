package component

import (
	redis "social-media-be/components/appredis"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	GetLogger(name string) *logrus.Entry
	GetRedisDBConnection() redis.RedisConnection
	SecretKey() string
}

type appCtx struct {
	db        *gorm.DB
	logger    *logrus.Logger
	redisDb   redis.RedisConnection
	secretKey string
}

func NewAppContext(db *gorm.DB, logger *logrus.Logger, redisDb redis.RedisConnection, secretKey string) *appCtx {
	return &appCtx{db: db, logger: logger, redisDb: redisDb, secretKey: secretKey}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) GetLogger(name string) *logrus.Entry {
	return ctx.logger.WithField("service_name", name)
}

func (ctx *appCtx) SecretKey() string {
	return ctx.secretKey
}

func (ctx *appCtx) GetRedisDBConnection() redis.RedisConnection {
	return ctx.redisDb
}
