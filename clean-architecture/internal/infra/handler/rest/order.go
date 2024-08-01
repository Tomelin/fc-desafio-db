package rest

import (
	"net/http"

	"github.com/Tomelin/fc-desafio-db/clean-architecture/internal/core/service"
	"github.com/gin-gonic/gin"
)

type OrderHandlerHttp struct {
	Service service.ServiceOrderInterface
}

type HandlerHttp interface {
	FindAll(c *gin.Context)
}

func NewOrderHandlerHttp(svc service.ServiceOrderInterface, routerGroup *gin.RouterGroup) HandlerHttp {

	order := &OrderHandlerHttp{
		Service: svc,
	}

	order.handlers(routerGroup)
	return order
}

func (h *OrderHandlerHttp) handlers(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/order", h.FindAll)
}

func (h *OrderHandlerHttp) FindAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "invalid header blabla"})
	return
}
