package handler

import (
	"context"
	"fmt"
	"github.com/Tomelin/fc-desafio-db/internal/core/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlerHttpInterface interface {
	Run(ctx context.Context) error
}

type HandlerHttp struct {
	router  *gin.Engine
	config  *ConfigWebserver
	Service service.ServiceExchangeInterface
}

func NewHandlerHttp(f string, service service.ServiceExchangeInterface) (HandlerHttpInterface, error) {

	config, err := NewConfig(f)
	if err != nil {
		return nil, err
	}

	return &HandlerHttp{
		Service: service,
		router:  gin.Default(),
		config:  config,
	}, nil
}

func (h *HandlerHttp) Run(ctx context.Context) error {
	h.cotacao(ctx)
	h.router.UseH2C = h.config.EnabledHttp2
	return h.router.Run(fmt.Sprintf("%s:%s", h.config.Listen, h.config.Port))
}

func (h *HandlerHttp) cotacao(ctx context.Context) error {
	h.router.GET("/cotacao", func(c *gin.Context) {
		exchange, err := h.Service.Request(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error: ": err.Error()})
			c.Abort()
			return
		}
		c.Writer.Header().Set("Content-Type", "application/json")

		c.JSON(http.StatusOK, gin.H{"dolar": exchange})
	})
	return nil
}
