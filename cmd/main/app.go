package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"tz/iternal/heartbeat"
	"tz/iternal/rate-limiter"
	"tz/iternal/user"
	"tz/pkg/configuration"
	"tz/pkg/logging"
)

func main() {
	config := configuration.GetConfig()
	log := logging.GetLogger()

	log.Info("Create router")
	router := gin.New()
	router.Use(rate_limiter.RateLimiter())

	registerHandlers(router, log)
	start(router, config, log)
}

func registerHandlers(router *gin.Engine, log *logging.Logger) {
	v1 := router.Group("api/v1")
	{
		log.Info("Init user module")
		userHandler := user.NewHandler()
		userHandler.Register(v1)

		log.Info("Init rate limiter")
		rateService := rate_limiter.NewService()
		rateHandler := rate_limiter.NewHandler(rateService)
		rateHandler.Register(v1)
	}
	v2 := router.Group("api/v2")
	{
		log.Info("Init heartbeat module")
		heartbeatModule := heartbeat.NewHandler()
		heartbeatModule.Register(v2)
	}
}

func start(router *gin.Engine, config configuration.ApplicationConfig, log *logging.Logger) {
	log.Info("Start application")
	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", config.ServerPort),
		Handler:           router,
		WriteTimeout:      time.Duration(config.ServerWriteTimeout) * time.Second,
		ReadHeaderTimeout: time.Duration(config.ServerReadTimeout) * time.Second,
	}

	log.Info(fmt.Sprintf("Server has been started http://localhost:%d/", config.ServerPort))
	log.Fatalln(server.ListenAndServe())
}
