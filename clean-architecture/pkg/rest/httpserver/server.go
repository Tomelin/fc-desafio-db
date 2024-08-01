package httpserver

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/net/http2"
)

var cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "cpu_temperature_celsius",
	Help: "Current temperature of the CPU",
})

var routePath *gin.RouterGroup

func init() {
	prometheus.MustRegister(cpuTemp)
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func (s *RestAPI) Run(handler http.Handler) error {

	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", s.Config.Port),
		Handler: s.Route.Handler(),
	}

	http2.ConfigureServer(&srv, &http2.Server{})
	s.Route.Use(s.MiddlewareHeader)
	return srv.ListenAndServe()
}

func (s *RestAPI) RunTLS() error {
	return nil
}

func (s *RestAPI) MiddlewareHeader(c *gin.Context) {

	if c.GetHeader("blabla") == "teste" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid header blabla"})
		c.Writer.Flush()
		c.Abort()
		return
	}
	c.Next()
}

func (s *RestAPI) ValidateToken(c *gin.Context) {

	if c.GetHeader("Authorization") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid authorization token"})
		c.Writer.Flush()
		c.Abort()
		return
	}

	token := strings.Split(c.GetHeader("Authorization"), " ")[1]
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid authorization token"})
		c.Writer.Flush()
		c.Abort()
		return
	}
	c.Next()
}
