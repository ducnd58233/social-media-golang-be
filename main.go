package main

import (
	"os"
	component "social-media-be/components"
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
		logger.Fatalln("Could not load .env file")
	}

	// Connect to db
	mySqlConnStr, ok := os.LookupEnv("MYSQL_CONNECTION")
	if !ok {
		logger.Fatalln("Missing MySQL connection string.")
	}

	dsn := mySqlConnStr
	db, errCon := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db = db.Debug()
	

	if errCon != nil {
		logger.Fatalln(errCon)
	}

	logger.Println("Connected:", db)

	appCtx := component.NewAppContext(db, logger)
	
	router := gin.Default()

	router.Use(middleware.Recover(appCtx))

	module.MainRoute(router, appCtx)

	router.Run(":6000")
}