package main

import (
	"github.com/perfectogo/template-api-gateway-with-mq/api"
	"github.com/perfectogo/template-api-gateway-with-mq/config"
	"github.com/perfectogo/template-api-gateway-with-mq/events"
	"github.com/perfectogo/template-api-gateway-with-mq/pkg/logger"
)

func main() {
	//
	cfg := config.Load()

	//
	log := logger.New(cfg.App, cfg.LogLevel)

	//++ docs.SwaggerInfo.Host = cfg.ServiceHost + cfg.ServicePort
	//- docs.SwaggerInfo.BasePath = cfg.BasePath
	// docs.SwaggerInfo.Schemes = []string{cfg.ServiceScheme}

	//
	rabbitServer, err := events.NewRabbitServer(cfg, log)
	if err != nil {
		log.Panic("error on the event server", logger.Error(err))
	}

	//
	router, err := api.New(cfg, log, rabbitServer.RMQ)
	if err != nil {
		log.Panic("error on the api server", logger.Error(err))
	}

	//
	err = router.Run(cfg.HTTPPort) // this method will block the calling goroutine indefinitely unless an error happens
	if err != nil {
		panic(err)
	}

	log.Panic("api server has finished")
	return

}
