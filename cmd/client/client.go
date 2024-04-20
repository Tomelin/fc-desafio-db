package main

import (
	"context"
	"fmt"
	"github.com/Tomelin/fc-desafio-db/internal/infra/requestHttp"
	"io"
	"log"
	"os"
	"time"
)

type HttpVersion uint

const (
	http1 HttpVersion = iota + 1
	http2
)

const (
	server   = "localhost"
	port     = "8080"
	timeout  = 300
	filePath = "/tmp/cotacao"
	fileName = "cotacao.txt"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*timeout)
	defer cancel()

	client, err := requestHttp.NewRequestHttp(ctx)
	if err != nil {
		panic(err)
	}

	response, err := client.Get(ctx, fmt.Sprintf("http://%s:%s/cotacao", server, port))
	if err != nil {
		panic(err)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	log.Println(string(data))
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		err = os.MkdirAll(filePath, 0777)
		if err != nil {
			panic(err)
		}
	}

	err = os.WriteFile(fmt.Sprintf("%s/%s", filePath, fileName), data, 0660)
	if err != nil {
		panic(err)
	}

}
