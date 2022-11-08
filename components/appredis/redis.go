package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

type RedisConnection interface {
	IsDisabled() bool
	Configure() error
	Name() string
	Get() interface{}
	Run() error
	Stop() <-chan bool
}

var (
	defaultRedisMaxActive = 0 // 0 is unlimited max active connection
	defaultRedisMaxIdle   = 10
)

type RedisDBOpt struct {
	RedisUri  string
	MaxActive int
	MaxIde    int
}

type redisDB struct {
	name   string
	client *redis.Client
	logger *logrus.Entry
	*RedisDBOpt
}

func NewRedisDB(name, uri string, logger *logrus.Logger) *redisDB {
	return &redisDB{
		name: name,
		RedisDBOpt: &RedisDBOpt{
			RedisUri:  uri,
			MaxActive: defaultRedisMaxActive,
			MaxIde:    defaultRedisMaxIdle,
		},
		logger: logger.WithField("service_name", "redisConnection"),
	}
}

func (r *redisDB) IsDisabled() bool {
	return r.RedisUri == ""
}

func (r *redisDB) Configure() error {
	if r.IsDisabled() {
		return nil
	}

	r.logger.Info("Connecting to Redis at ", r.RedisUri, "...")

	opt, err := redis.ParseURL(r.RedisUri)

	if err != nil {
		r.logger.Error("cannot parse Redis ", err.Error())
		return err
	}

	opt.PoolSize = r.MaxActive
	opt.MinIdleConns = r.MaxIde

	client := redis.NewClient(opt)
	// Ping to test Redis connection
	if err := client.Ping(context.Background()).Err(); err != nil {
		r.logger.Error("Cannot connect Redis. ", err.Error())
		return err
	}

	// Connect successfully, assign client to goRedisDB
	r.client = client
	return nil
}

func (r *redisDB) Name() string {
	return r.name
}

func (r *redisDB) Get() interface{} {
	return r.client
}

func (r *redisDB) Run() error {
	return r.Configure()
}

func (r *redisDB) Stop() <-chan bool {
	if r.client != nil {
		if err := r.client.Close(); err != nil {
			r.logger.Error("cannot close ", r.name)
		}
	}

	c := make(chan bool)
	go func() { c <- true }()
	return c
}
