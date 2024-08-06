package main

import (
	"context"
	"log"

	_ "github.com/Tomelin/fc-desafio-db/clean-architecture/docs/swagger"
	"github.com/Tomelin/fc-desafio-db/clean-architecture/configs"
	"github.com/Tomelin/fc-desafio-db/clean-architecture/internal/core/repository"
	"github.com/Tomelin/fc-desafio-db/clean-architecture/internal/core/service"
	grpc_order "github.com/Tomelin/fc-desafio-db/clean-architecture/internal/infra/handler/grpc"
	handler_rest "github.com/Tomelin/fc-desafio-db/clean-architecture/internal/infra/handler/rest"
	"github.com/Tomelin/fc-desafio-db/clean-architecture/internal/infra/storage/cache"
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

	// Connect cache
	cache, err := cache.NewCacheConnection(context.Background(), config.PathConfigFile, config.FileConfig.Filename, config.FileConfig.Extension)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}

	ctx := context.Background()
	rest, err := httpserver.NewRestAPI(config.PathConfigFile, config.FileConfig.Filename, config.FileConfig.Extension)
	if err != nil {
		panic(err)
	}

	orderRepo, err := repository.NewOrderRepository(ctx, pool)
	if err != nil {
		panic(err)
	}

	orderSvc := service.NewServiceOrder(orderRepo, cache)
	if err != nil {
		panic(err)
	}

	// BEGINER gRPC
	log.Println("loading gRPC")
	go grpc_order.NewOrderHandlerGrpc(orderSvc)
	// END gRPC

	// BEGINER GraphQL
	log.Println("loading GraphQL")
	go graphql("8082", orderSvc)
	// END GraphQL

	handler_rest.NewOrderHandlerHttp(orderSvc, rest.RouterGroup)
	log.Println("loading http with Gin")
	rest.Run(rest.Route.Handler())
}
