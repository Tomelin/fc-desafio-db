package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Tomelin/fc-desafio-db/multithreading/internal/core/entity"
	"github.com/Tomelin/fc-desafio-db/multithreading/internal/infra/requestHttp"
	"io"
	"log"
	"time"
)

var channel chan entity.CEP = make(chan entity.CEP)

func main() {

	cep := "90230181"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	go func(ctx context.Context, cep string) {
		result, err := BasilAPI(ctx, cep)
		if err != nil {
			panic(err)
		}
		channel <- *result

	}(ctx, cep)

	go func(ctx context.Context, cep string) {
		result, err := ViaCEP(ctx, cep)
		if err != nil {
			panic(err)
		}

		channel <- *result
	}(ctx, cep)

	log.Println(<-channel)
}

func BasilAPI(ctx context.Context, cep string) (*entity.CEP, error) {
	//func BasilAPI(ctx context.Context, cep string) (*entity.CEP, error) {
	var address string = fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)

	request, err := requestHttp.NewRequestHttp(ctx)
	if err != nil {
		return nil, err
	}

	data, err := request.Get(ctx, address)
	if err != nil {
		return nil, err
	}

	defer data.Body.Close()
	b, err := io.ReadAll(data.Body)
	if err != nil {
		return nil, err
	}

	var result *entity.BrasilAPI
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}

	return entity.NewCEP(result.CEP, "brasilapi"), nil

}

func ViaCEP(ctx context.Context, cep string) (*entity.CEP, error) {
	var address string = fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)

	request, err := requestHttp.NewRequestHttp(ctx)
	if err != nil {
		return nil, err
	}

	data, err := request.Get(ctx, address)
	if err != nil {
		return nil, err
	}

	defer data.Body.Close()
	b, err := io.ReadAll(data.Body)
	if err != nil {
		return nil, err
	}

	var result *entity.ViaCEP
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}

	return entity.NewCEP(result.CEP, "viacep"), nil
}
