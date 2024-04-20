package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Tomelin/fc-desafio-db/configs"
	"github.com/Tomelin/fc-desafio-db/internal/core/repository"
	"github.com/Tomelin/fc-desafio-db/internal/core/service"
	"github.com/Tomelin/fc-desafio-db/internal/infra/handler"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

func main() {
	// LOAD CONFIGURATION FROM YAML
	_, currentFile, _, _ := runtime.Caller(0)
	currentDirectoy := filepath.Dir(currentFile)
	pathConfig := fmt.Sprintf("%s/.config", currentDirectoy)
	fmt.Println("Current working directory:", currentDirectoy)

	if os.Getenv("FILE_CONFIG") != "" {
		pathConfig = os.Getenv("FILE_CONFIG")
	}

	conf, err := configs.NewConfig(pathConfig)
	if err != nil {
		panic(err)
	}

	// CONNECT IN DATABASE
	db, err := sql.Open(conf.Database.Driver, conf.Database.Host)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	ctxDb, cancelDb := context.WithTimeout(context.Background(), time.Millisecond*10)
	defer cancelDb()

	repositoryServer, err := repository.NewRepositoryExchange(ctxDb, db)
	if err != nil {
		panic(err)
	}

	serviceServer := service.NewServiceServer(repositoryServer)

	svcExchange, err := service.NewServiceExchange(serviceServer)
	if err != nil {
		panic(err)
	}

	// HANDLER
	hand, err := handler.NewHandlerHttp(pathConfig, svcExchange)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*300)
	defer cancel()
	hand.Run(ctx)
}
