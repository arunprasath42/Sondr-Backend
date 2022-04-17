package main

import (
	"web-api/route"
	"web-api/src/repository"
	"web-api/utils/database"
	"web-api/utils/middleware"
	"web-api/utils/validator"

	config "web-api/config"
	"web-api/migration"
	logger "web-api/utils/logging"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//var log *zerolog.Logger = applogger.GetInstance()

func main() {
	config.LoadConfig()
	router := gin.Default()
	logger.SetupLogger(router)
	database.GetInstancemysql()
	migration.Migration()
	repository.MySqlInit()
	validator.Init()
	router.Use(middleware.TracingMiddleware())
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	router.Use(cors.New(corsConfig))
	route.SetupRoutes(router)

}
