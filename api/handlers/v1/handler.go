package v1

import (
	"fmt"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/perfectogo/template-api-gateway-with-mq/config"
	"github.com/perfectogo/template-api-gateway-with-mq/models"
	"github.com/perfectogo/template-api-gateway-with-mq/pkg/logger"
	"github.com/perfectogo/template-api-gateway-with-mq/pkg/rabbit"
)

// Handler ...
type Handler struct {
	cfg config.Config
	log logger.Logger
	rmq *rabbit.RMQ
}

// NewHandler ...
func NewHandler(cfg config.Config, log logger.Logger, rmq *rabbit.RMQ) *Handler {
	return &Handler{
		cfg: cfg,
		log: log,
		rmq: rmq,
	}
}

//Responses
func (h *Handler) handleSuccessResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, models.SuccessModel{
		Code:    fmt.Sprint(code),
		Message: message,
		Data:    data,
	})
}

func (h *Handler) handleErrorResponse(c *gin.Context, code int, message string, err interface{}) {
	h.log.Error(message, logger.Int("code", code), logger.Any("error", err))
	c.JSON(code, models.ErrorModel{
		Code:    fmt.Sprint(code),
		Message: message,
		Error:   err,
	})
}

// func (h *Handler) parseOffsetQueryParam(c *gin.Context) (int, error) {
// 	return strconv.Atoi(c.DefaultQuery("offset", h.cfg.DefaultOffset))
// }

// func (h *Handler) parseLimitQueryParam(c *gin.Context) (int, error) {
// 	return strconv.Atoi(c.DefaultQuery("limit", h.cfg.DefaultLimit))
// }

//Proxy
func (h *Handler) makeProxy(c *gin.Context, proxyURL string, endpoint string) (err error) {
	fmt.Println(proxyURL)
	fmt.Println(endpoint)

	// parse the url
	url, err := url.Parse(proxyURL)
	if err != nil {
		h.log.Error("error in parse addr: %v", logger.Error(err))
		return err
	}

	// create the  reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(url)

	// Update the headers to allow for SSL redirection
	c.Request.URL.Host = url.Host
	c.Request.URL.Scheme = url.Scheme
	c.Request.URL.Path = endpoint
	c.Request.Header.Set("X-Forwarded-Host", c.Request.Header.Get("Host"))
	c.Request.Host = url.Host

	// Note that ServeHttp is non blocking and uses a go routine under the hood
	proxy.ServeHTTP(c.Writer, c.Request)

	return nil
}
