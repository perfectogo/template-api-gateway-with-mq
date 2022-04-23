package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	v1 "github.com/perfectogo/template-api-gateway-with-mq/api/handlers/v1"
	"github.com/perfectogo/template-api-gateway-with-mq/config"
	"github.com/perfectogo/template-api-gateway-with-mq/pkg/logger"
	"github.com/perfectogo/template-api-gateway-with-mq/pkg/rabbit"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func New(cfg config.Config, log logger.Logger, rmq *rabbit.RMQ) (*gin.Engine, error) {
	if cfg.Environment != "development" {
		gin.SetMode(gin.ReleaseMode)
	}
	log.Info("Api server")
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = append(config.AllowHeaders, "*")

	router.Use(cors.New(config))

	handlerV1 := v1.NewHandler(cfg, log, rmq)

	router.GET("/ping", handlerV1.Ping)

	rV1 := router.Group("/v1")
	{
		endpointsV1(rV1, handlerV1)
	}

	url := ginSwagger.URL("/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router, nil
}
