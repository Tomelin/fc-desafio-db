package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"

	"github.com/Tomelin/fc-desafio-db/clean-architecture/internal/core/entity"
	"github.com/Tomelin/fc-desafio-db/clean-architecture/internal/core/repository"
	"github.com/Tomelin/fc-desafio-db/clean-architecture/internal/infra/storage/cache"
)

type ServiceOrderInterface interface {
	entity.OrderInterface
}

type ServiceOrder struct {
	Repository repository.OrderRepositoryInterface
	Cache      *redis.Client
}

func NewServiceOrder(repo repository.OrderRepositoryInterface, c cache.DesafioCache) ServiceOrderInterface {
	return &ServiceOrder{
		Repository: repo,
		Cache:      c,
	}
}

func (so *ServiceOrder) FindAll() ([]entity.OrderResponse, error) {

	var orders []entity.OrderResponse
	val, err := so.Cache.Get(context.Background(), "orders_all").Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			orders, err = so.Repository.FindAll()
			if err != nil {
				return nil, err
			}
			s, err := json.Marshal(orders)
			if err != nil {
				return nil, err
			}

			err = so.Cache.Set(context.Background(), "orders_all", s, 3*time.Second).Err()
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		err = json.Unmarshal([]byte(val), &orders)
		if err != nil {
			return nil, err
		}
	}

	return orders, nil
}

func (so *ServiceOrder) FindByFilter(in *string) ([]entity.OrderResponse, error) {
	var orders []entity.OrderResponse

	if in == nil{
		return nil,errors.New("filter cannot be empty")
	}

	val, err := so.Cache.Get(context.Background(), "orders_by_filter").Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			orders, err = so.Repository.FindByFilter(in)
			if err != nil {
				return nil, err
			}
			s, err := json.Marshal(orders)
			if err != nil {
				return nil, err
			}

			err = so.Cache.Set(context.Background(), "orders_by_filter", s, 3*time.Second).Err()
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		err = json.Unmarshal([]byte(val), &orders)
		if err != nil {
			return nil, err
		}
	}

	return orders, nil
}
func (so *ServiceOrder) FindByID(id *string) (*entity.OrderResponse, error) {

	var order *entity.OrderResponse

	if id == nil || len(*id) != 36{
		return nil,errors.New("filter cannot be empty")
	}
	
	val, err := so.Cache.Get(context.Background(), "orders_by_id").Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			order, err = so.Repository.FindByID(id)
			if err != nil {
				return nil, err
			}
			s, err := json.Marshal(order)
			if err != nil {
				return nil, err
			}

			err = so.Cache.Set(context.Background(), "orders_by_id", s, 3*time.Second).Err()
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		err = json.Unmarshal([]byte(val), &order)
		if err != nil {
			return nil, err
		}
	}

	return order, nil
}
func (so *ServiceOrder) Create(order *entity.Order) (*entity.OrderResponse, error) {
	return nil, nil
}
func (so *ServiceOrder) Delete(id *string) error {
	return nil
}
