package httpserver

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type RestAPIConfig struct {
	Port           string `mapstructure:"port"`
	SSLEnabled     bool   `mapstructure:"ssl_enabled"`
	Host           string `mapstructure:"host"`
	Version        string `mapstructure:"version"`
	Name           string `mapstructure:"name"`
	CertificateCrt string `mapstructure:"certificate_crt"`
	CertificateKey string `mapstructure:"certificate_key"`
}

type RestAPI struct {
	Config *RestAPIConfig
	Route  *gin.Engine
	*gin.RouterGroup
}

func NewRestAPI(pathConfigFile, nameConfigFile, nameFileExtension string) (*RestAPI, error) {

	viper.AddConfigPath(pathConfigFile)
	viper.SetConfigName(nameConfigFile)
	viper.SetConfigType(nameFileExtension)
	viper.AutomaticEnv()

	v, ok := viper.Get("webserver").(map[string]interface{})
	if !ok {
		return nil, errors.New("error, not found webserver key at config file")
	}

	// var stringPort string = "8443"
	// if reflect.TypeOf(v["port"]).Key().Kind() == reflect.String {
	// 	log.Println("if ")
	// 	log.Println(reflect.TypeOf(v["port"]).Key().Kind())
	// 	stringPort = "8443"
	// 	// stringPort = strconv.Itoa(v["port"]).(int)
	// } else if reflect.TypeOf(v["port"]).Key().Kind() == reflect.Int {
	// 	log.Println("else ")
	// 	log.Println(reflect.TypeOf(v["port"]).Key().Kind())
	// 	stringPort = "8443"
	// 	// stringPort = v["port"].(string)
	// }

	rest := RestAPIConfig{
		Port:       "8443",
		Host:       v["host"].(string),
		Version:    v["version"].(string),
		Name:       v["name"].(string),
		SSLEnabled: false,
		// CertificateCrt: v["certificate_crt"].(string),
		// CertificateKey: v["certificate_key"].(string),
	}

	rest.Validate()
	r, g := newRestAPI(&rest)

	return &RestAPI{
		Config:      &rest,
		Route:       r,
		RouterGroup: g,
	}, nil
}

func (c *RestAPIConfig) Validate() {
	if c.Port == "" {
		c.Port = "8843"
	}

	if c.Host == "" {
		c.Host = "0.0.0.0"
	}

	if c.Name == "" {
		c.Name = "api"
	}

	if c.Version == "" {
		c.Version = "v1"
	}

	return
}

func newRestAPI(config *RestAPIConfig) (*gin.Engine, *gin.RouterGroup) {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.UseH2C = true

	routerGroupPath := fmt.Sprintf("/%s", config.Name)
	routerPath := router.Group(routerGroupPath)

	router.GET("/metrics", prometheusHandler())

	// Set swagger
	routerPath.GET("/docs/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))
	routerPath.GET("/docs/swagger", func(c *gin.Context) {
		c.Redirect(301, fmt.Sprintf("%s/docs/swagger/index.html", routerGroupPath))
	})

	routerPath.GET("/docs", func(c *gin.Context) {
		c.Redirect(301, fmt.Sprintf("%s/docs/swagger/index.html", routerGroupPath))
	})

	routerPath.GET("/", func(c *gin.Context) {
		c.Redirect(301, fmt.Sprintf("%s/docs/swagger/index.html", routerGroupPath))
	})

	router.Use(setHeader)

	return router, routerPath
}

func setHeader(c *gin.Context) {

	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.Next()
}
