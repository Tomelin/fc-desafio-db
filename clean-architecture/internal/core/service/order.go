package service

import (
	"github.com/Tomelin/fc-desafio-db/clean-architecture/internal/core/entity"
	"github.com/Tomelin/fc-desafio-db/clean-architecture/internal/core/repository"
)

type ServiceOrderInterface interface {
	entity.OrderInterface
}

type ServiceOrder struct {
	Repository repository.OrderRepositoryInterface
	// cache
}

func NewServiceOrder(repo repository.OrderRepositoryInterface) ServiceOrderInterface {
	return &ServiceOrder{
		Repository: repo,
	}
}

func (so *ServiceOrder) FindAll() ([]entity.OrderResponse, error) {
	return nil, nil
}
func (so *ServiceOrder) FindByFilter(*string) ([]entity.OrderResponse, error) {
	return nil, nil
}
func (so *ServiceOrder) FindByID(id *string) (*entity.OrderResponse, error) {
	return nil, nil
}
func (so *ServiceOrder) Create(order *entity.Order) (*entity.OrderResponse, error) {
	return nil, nil
}
func (so *ServiceOrder) Delete(id *string) error {
	return nil
}
