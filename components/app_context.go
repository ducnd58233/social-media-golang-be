package component

import (
	"gorm.io/gorm"
	"github.com/sirupsen/logrus"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	CreateLogger(name string) *logrus.Entry
}

type appCtx struct {
	db *gorm.DB
	logger *logrus.Logger 
}

func NewAppContext(db *gorm.DB, logger *logrus.Logger) *appCtx {
	return &appCtx{db: db, logger: logger}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) CreateLogger(name string) *logrus.Entry {
	return ctx.logger.WithField("service_name", name)
}