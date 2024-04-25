package main

import (
	"context"
	"fmt"
	"github.com/Tomelin/fc-desafio-db/internal/infra/requestHttp"
	"io"
	"log"
	"os"
	"strings"
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
	filePath = "./"
	fileName = "cotacao.txt"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*300)
	defer cancel()
	client, err := requestHttp.NewRequestHttp(ctx)
	if err != nil {
		panic(err)
	}

	// create a request from server and send the context
	response, err := client.Get(ctx, fmt.Sprintf("http://%s:%s/cotacao", server, port))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// check if there is a context deadline exceeded i body
	if strings.Contains(string(data), "context deadline exceeded") {
		log.Fatalf("the server returned this error: %s", string(data))
	}

	// if request ok, will be create the directory and file accoding to const
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
