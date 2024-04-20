package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Tomelin/fc-desafio-db/internal/core/entity"
	"github.com/Tomelin/fc-desafio-db/internal/infra/requestHttp"
	"io"
)

type ServiceExchangeInterface interface {
	Request(ctx context.Context) (*string, error)
	Save(*entity.Exchange) error
}

type ServiceExchange struct {
	Server entity.ExchangeInterface
}

func NewServiceExchange(server entity.ExchangeInterface) (ServiceExchangeInterface, error) {
	return &ServiceExchange{Server: server}, nil
}

func (se *ServiceExchange) Save(dolar *entity.Exchange) error {
	response, err := entity.NewExchange(dolar.USDBRL)
	if err != nil {
		return err
	}
	se.Server.Create(response)
	return nil
}

func (se *ServiceExchange) Request(ctx context.Context) (*string, error) {

	baseUrl := fmt.Sprintf("https://economia.awesomeapi.com.br/json/last/USD-BRL")
	req, err := requestHttp.NewRequestHttp(ctx)
	if err != nil {
		return nil, err
	}

	response, err := req.Get(ctx, baseUrl)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var exchange entity.Exchange
	err = json.Unmarshal(body, &exchange)
	if err != nil {
		return nil, err
	}

	dolar := &exchange.USDBRL.Bid

	err = se.Save(&exchange)
	if err != nil {
		return nil, err
	}

	return dolar, nil
}
