package main

import (
	"os"
	component "social-media-be/components"
	redis "social-media-be/components/appredis"

	// cloudinaryprovider "social-media-be/components/cloudprovider/cloudinary"
	awsprovider "social-media-be/components/cloudprovider/aws"
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

	// AWS
	s3Bucketname, ok := os.LookupEnv("S3_BUCKET_NAME")
	if !ok {
		logger.Error("Missing S3 Bucket Name string.")
	}

	s3Region, ok := os.LookupEnv("S3_REGION")
	if !ok {
		logger.Error("Missing S3 Region string.")
	}

	s3APIKey, ok := os.LookupEnv("S3_ACCESS_KEY")
	if !ok {
		logger.Error("Missing S3 API Key string.")
	}

	s3SecretKey, ok := os.LookupEnv("S3_SECRET_KEY")
	if !ok {
		logger.Error("Missing S3 Secret Key string.")
	}

	s3Domain, ok := os.LookupEnv("S3_DOMAIN")
	if !ok {
		logger.Error("Missing S3 Domain string.")
	}

	s3 := awsprovider.NewS3Provider(s3Bucketname, s3Region, s3APIKey, s3SecretKey, s3Domain, logger)

	// Cloudinary
	// cldName, ok := os.LookupEnv("CLOUDINARY_NAME")
	// if !ok {
	// 	logger.Error("Missing Cloudinary name string.")
	// }

	// cldApiKey, ok := os.LookupEnv("CLOUDINARY_API_KEY")
	// if !ok {
	// 	logger.Error("Missing Cloudinary API Key string.")
	// }

	// cldApiSecret, ok := os.LookupEnv("CLOUDINARY_API_SECRET")
	// if !ok {
	// 	logger.Error("Missing Cloudinary Secret string.")
	// }

	// cloudinary := cloudinaryprovider.NewCloudinaryProvider(cldName, cldApiKey, cldApiSecret, logger)

	appCtx := component.NewAppContext(db, logger, redisDb, s3, secretKey)

	router := gin.Default()

	router.Use(middleware.Recover(appCtx))

	module.MainRoute(router, appCtx)

	router.Run(":6000")
}
