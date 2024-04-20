package service

import (
	"github.com/Tomelin/fc-desafio-db/internal/core/entity"
)

type ServiceServer struct {
	Repository entity.ExchangeInterface
}

func NewServiceServer(repo entity.ExchangeInterface) entity.ExchangeInterface {
	return &ServiceServer{
		Repository: repo,
	}
}

func (se *ServiceServer) Get() ([]entity.Exchange, error) {
	return nil, nil
}
func (se *ServiceServer) Create(exchange *entity.ResponseCurrency) (*entity.ResponseCurrency, error) {
	return se.Repository.Create(exchange)

}
func (se *ServiceServer) Delete(id *string) error { return nil }

func (se *ServiceServer) Update(*entity.Exchange) (*entity.Exchange, error) { return nil, nil }
