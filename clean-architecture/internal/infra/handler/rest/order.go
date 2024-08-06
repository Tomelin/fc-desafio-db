package rest

import (
	"log"
	"net/http"

	"github.com/Tomelin/fc-desafio-db/clean-architecture/internal/core/service"
	_ "github.com/Tomelin/fc-desafio-db/clean-architecture/internal/core/entity"
	"github.com/gin-gonic/gin"
)

type OrderHandlerHttp struct {
	Service service.ServiceOrderInterface
}

type HandlerHttp interface {
	FindAll(c *gin.Context)
	FindByID(c *gin.Context)
	FindByFilter(c *gin.Context)
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
	routerGroup.GET("/order/:id", h.FindByID)
	routerGroup.GET("/order/search", h.FindByFilter)
}

// OrderFind    godoc
// @Summary     find all orders
// @Tags         order
// @Accept       json
// @Produce     json
// @Description get all orders
// @Success     200 {object} []entity.Order
// @Failure     404 {object} string
// @Failure     500 {object} string
// @Router      /order [get]
func (svc *OrderHandlerHttp) FindAll(c *gin.Context) {

	result, err := svc.Service.FindAll()
	if err != nil {
		log.Println(err)
		c.JSON(404, gin.H{"error": err.Error()})
		c.Writer.Flush()
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
	// log.Println("finishing handler")
	return

}

// OrderFind    godoc
// @Summary     find order by ID
// @Tags         order
// @Accept       json
// @Produce     json
// @Description get order by ID
// @Param        id path string true "found"
// @Success     200 {object} entity.Order
// @Failure     404 {object} string
// @Failure     500 {object} string
// @Router      /order/{id} [get]
func (svc *OrderHandlerHttp) FindByID(c *gin.Context) {

	id := c.Param("id")
	if len(id) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "ID to get not found"})
		return
	}
	result, err := svc.Service.FindByID(&id)
	if err != nil {
		log.Println(err)
		c.JSON(404, gin.H{"error": err.Error()})
		c.Writer.Flush()
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
	// log.Println("finishing handler")
	return

}

// OrderFind    godoc
// @Summary     find order by filter
// @Tags         order
// @Accept       json
// @Produce     json
// @Param       filter        query string false "filter field"
// @Description get order by filter
// @Success     200 {object} []entity.Order
// @Failure     404 {object} string
// @Failure     500 {object} string
// @Router      /order/search [get]
func (svc *OrderHandlerHttp) FindByFilter(c *gin.Context) {
	filter :=  c.Query("filter")
	log.Println(filter)
	if len(filter) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "ID to get not found"})
		return
	}

	result, err := svc.Service.FindByFilter(&filter)
	if err != nil {
		log.Println(err)
		c.JSON(404, gin.H{"error": err.Error()})
		c.Writer.Flush()
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
	// log.Println("finishing handler")
	return

}
