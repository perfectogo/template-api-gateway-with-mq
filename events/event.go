package events

import (
	"github.com/perfectogo/template-api-gateway-with-mq/config"
	"github.com/perfectogo/template-api-gateway-with-mq/pkg/logger"
	"github.com/perfectogo/template-api-gateway-with-mq/pkg/rabbit"
)

// PubsubServer ...
type RabbitServer struct {
	cfg config.Config
	log logger.Logger
	RMQ *rabbit.RMQ
}

// New ...
func NewRabbitServer(cfg config.Config, log logger.Logger) (*RabbitServer, error) {
	rmq, err := rabbit.NewRMQ(cfg.RabbitURI, log)
	if err != nil {
		return nil, err
	}

	rmq.AddPublisher("v1.todo") // one publisher is enough for application service

	return &RabbitServer{
		cfg: cfg,
		log: log,
		RMQ: rmq,
	}, nil
}
