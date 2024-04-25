package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Tomelin/fc-desafio-db/internal/core/entity"
	"github.com/Tomelin/fc-desafio-db/internal/infra/requestHttp"
	"io"
	"strings"
	"time"
)

type ServiceServer struct {
	Repository entity.ExchangeInterface
}

func NewServiceServer(repo entity.ExchangeInterface) entity.ExchangeInterface {
	return &ServiceServer{
		Repository: repo,
	}
}

func (se *ServiceServer) Get() ([]entity.ResponseCurrency, error) {

	return se.Repository.Get()
}
func (se *ServiceServer) Create(ctx context.Context, exchange *entity.ResponseCurrency) (*entity.ResponseCurrency, error) {
	result, err := se.Request(ctx)
	if err != nil {
		return nil, err
	}

	response, err := entity.NewExchange(result.USDBRL)
	if err != nil {
		return nil, err
	}

	if _, err := se.Repository.Create(ctx, response); err != nil {
		return nil, err
	}

	return response, nil
}
func (se *ServiceServer) Delete(id *string) error { return nil }

func (se *ServiceServer) Update(*entity.Exchange) (*entity.Exchange, error) { return nil, nil }

func (se *ServiceServer) Request(ctx context.Context) (*entity.Exchange, error) {

	timeout := ctx.Value("timeout").(map[string]int)
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*time.Duration(timeout["http"]))
	defer cancel()

	baseUrl := fmt.Sprintf("https://economia.awesomeapi.com.br/json/last/USD-BRL")
	req, err := requestHttp.NewRequestHttp(ctx)
	if err != nil {
		return nil, err
	}

	response, err := req.Get(ctx, baseUrl)
	if err != nil {
		if strings.Contains(err.Error(), "context deadline exceeded") {
			return nil, errors.New("context deadline exceeded to request http")
		}
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

	return &exchange, nil
}
