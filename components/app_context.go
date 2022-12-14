package component

import (
	redis "social-media-be/components/appredis"
	cloudprovider "social-media-be/components/cloudprovider"
	"social-media-be/pubsub"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	GetLogger(name string) *logrus.Entry
	GetRedisDBConnection() redis.RedisConnection
	GetCloudProvider() cloudprovider.CloudProvider
	GetPubSub() pubsub.PubSub
	SecretKey() string
}

type appCtx struct {
	db            *gorm.DB
	logger        *logrus.Logger
	redisDb       redis.RedisConnection
	cloudProvider cloudprovider.CloudProvider
	ps            pubsub.PubSub
	secretKey     string
}

func NewAppContext(db *gorm.DB, logger *logrus.Logger, redisDb redis.RedisConnection, cloudProvider cloudprovider.CloudProvider, ps pubsub.PubSub, secretKey string) *appCtx {
	return &appCtx{db: db, logger: logger, redisDb: redisDb, cloudProvider: cloudProvider, ps: ps, secretKey: secretKey}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) GetLogger(name string) *logrus.Entry {
	return ctx.logger.WithField("service_name", name)
}

func (ctx *appCtx) GetCloudProvider() cloudprovider.CloudProvider {
	return ctx.cloudProvider
}

func (ctx *appCtx) GetPubSub() pubsub.PubSub {
	return ctx.ps
}

func (ctx *appCtx) SecretKey() string {
	return ctx.secretKey
}

func (ctx *appCtx) GetRedisDBConnection() redis.RedisConnection {
	return ctx.redisDb
}
