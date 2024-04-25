// SERVICE to get exchange between Dollar and Real
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
)

func main() {
	// LOAD CONFIGURATION FROM YAML
	// this code block load teh configuration from .config/config.yaml. you need have the .config/config.yaml in directory
	_, currentFile, _, _ := runtime.Caller(0)
	currentDirectoy := filepath.Dir(currentFile)
	pathConfig := fmt.Sprintf("%s/.config", currentDirectoy)
	fmt.Println("Current working directory:", currentDirectoy)

	// IF FILE_CONFIG variable exists, the code will load config.yaml from FILE_CONFIG path
	if os.Getenv("FILE_CONFIG") != "" {
		pathConfig = os.Getenv("FILE_CONFIG")
	}

	// Get configuration about database and webserver
	conf, err := configs.NewConfig(pathConfig)
	if err != nil {
		panic(err)
	}

	// CONNECT IN DATABASE acconding drive and host in config.yaml
	db, err := sql.Open(conf.Database.Driver, conf.Database.Host)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// CREATE TABLE  if not exists
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS exchange (id TEXT PRIMARY KEY, name TEXT, high TEXT, low TEXT, bid TEXT, timestamp TEXT )")
	if err != nil {
		panic(err)
	}
	_, err = statement.Exec()
	if err != nil {
		panic(err)
	}

	// Setted timeout for http and db
	// values about timeout are in config.yaml
	timeout := map[string]int{
		"db":   conf.Database.Timeout,
		"http": conf.Webserver.Timeout,
	}

	// CREATE A context for database
	// I don`t know if context need create when call repository or when call the service because service will call method of interface
	ctx := context.WithValue(context.Background(), "timeout", timeout)

	// call reposintory and return the entity.ExchangeInterface interface
	repositoryServer, err := repository.NewRepositoryExchange(ctx, db)
	if err != nil {
		panic(err)
	}

	// call reposintory and return the entity.ExchangeInterface interface
	serviceServer := service.NewServiceServer(repositoryServer)

	// HANDLER call a handler
	hand, err := handler.NewHandlerHttp(pathConfig, serviceServer)
	if err != nil {
		panic(err)
	}

	// start the webserver
	hand.Run(ctx)
}
