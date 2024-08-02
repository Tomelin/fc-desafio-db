package rest

import (
	// "log"
	"net/http"

	"github.com/Tomelin/fc-desafio-db/clean-architecture/internal/core/service"
	"github.com/gin-gonic/gin"
)

type OrderHandlerHttp struct {
	Service service.ServiceOrderInterface
}

type HandlerHttp interface {
	GetAll(c *gin.Context)
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

func (h *OrderHandlerHttp) GetAll(c *gin.Context){

	 
	_, err := h.FindAll()
	if err != nil {
		c.JSON(404, gin.H{"error": "not found"})
		c.Writer.Flush()
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "result"})
	// log.Println("finishing handler")
	return

}



