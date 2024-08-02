package main

import (
	"context"
	"log"

	"github.com/Tomelin/fc-desafio-db/clean-architecture/configs"
	"github.com/Tomelin/fc-desafio-db/clean-architecture/internal/core/repository"
	"github.com/Tomelin/fc-desafio-db/clean-architecture/internal/core/service"
	handler_rest "github.com/Tomelin/fc-desafio-db/clean-architecture/internal/infra/handler/rest"
	"github.com/Tomelin/fc-desafio-db/clean-architecture/internal/infra/storage/database"
	"github.com/Tomelin/fc-desafio-db/clean-architecture/pkg/rest/httpserver"
)

// @title          manager orders
// @version        1.0.0
// @description    desafio fullcycle
// @contact.name   FullCycle
// @contact.url    www.FullCycle.com.br
// @contact.email  contato@FullCycle.com.br

// @schemes        http https
// @BasePath       /api
func main() {

	// Load configs
	config, err := configs.LoadConfig()
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}

	// Connect database
	pool, err := database.NewDBConnection(context.Background(), config.PathConfigFile, config.FileConfig.Filename, config.FileConfig.Extension)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
	defer pool.Close()

	err = pool.Ping()
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}

	ctx := context.Background()
	rest, err := httpserver.NewRestAPI(config.PathConfigFile, config.FileConfig.Filename, config.FileConfig.Extension)
	if err != nil {
		// log.Println(err.Error())
		panic(err)
	}

	orderRepo, err := repository.NewOrderRepository(ctx, pool)
	if err != nil {
		panic(err)
	}

	orderSvc := service.NewServiceOrder(orderRepo)
	if err != nil {
		panic(err)
	}

	handler_rest.NewOrderHandlerHttp(orderSvc, rest.RouterGroup)

	rest.Run(rest.Route.Handler())
}
