package main

import (
	"fmt"
	"os"
	"project/handler"
	"project/repository/repositoryImpl"
	"project/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {

	db := initDatabase()

	subscribeRepository := repositoryImpl.NewSubscribeRepositoryImpl(db)
	subscribeService := service.NewSubscribeService(subscribeRepository)
	subscribeHandler := handler.NewSubscribeHandler(subscribeService)

	router := gin.New()
	router.Use(gin.Recovery())

	route := router.Group("/api/v1/")
	route.Use(cors.New(cors.Config{
		AllowMethods: []string{"GET", "PUT", "POST", "PATCH", "DELETE", "OPTIONS", "HEAD"},
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Content-Type", "x-iims-mock", "x-api-key"},
	}))

	route.GET("subscribes", subscribeHandler.getSubscribes)
	route.POST("subscribe", subscribeHandler.subscribe)
	route.POST("unsubscribe", subscribeHandler.unsubscribe)

	fmt.Println("Started port 3000")
	err := router.Run(":3000")
	if err != nil {
		fmt.Println(err)
	}
}

func initDatabase() *gorm.DB {

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Database is connected")

	return db
}
