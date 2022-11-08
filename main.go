package main

import (
	"os"
	component "social-media-be/components"
	redis "social-media-be/components/appredis"
	middleware "social-media-be/middlewares"
	module "social-media-be/modules"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	envErr := godotenv.Load(".env")

	if envErr != nil {
		logger.Error("Could not load .env file")
	}

	// Connect to db
	mySqlConnStr, ok := os.LookupEnv("MYSQL_CONNECTION")
	if !ok {
		logger.Error("Missing MySQL connection string.")
	}

	dsn := mySqlConnStr
	db, errCon := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db = db.Debug()

	if errCon != nil {
		logger.Error(errCon)
	}

	logger.Println("Connected:", db)

	redisUri, ok := os.LookupEnv("REDIS_HOST")
	if !ok {
		logger.Error("Missing Redis Host connection string.")
	}

	secretKey, ok := os.LookupEnv("SECRET_KEY")
	if !ok {
		logger.Error("Missing Secret Key string.")
	}

	redisDb := redis.NewRedisDB("redis", redisUri, logger)

	appCtx := component.NewAppContext(db, logger, redisDb, secretKey)

	router := gin.Default()

	router.Use(middleware.Recover(appCtx))

	module.MainRoute(router, appCtx)

	router.Run(":6000")
}
