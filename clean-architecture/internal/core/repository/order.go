package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"
	"fmt"

	"github.com/Tomelin/fc-desafio-db/clean-architecture/internal/core/entity"
)

type OrderRepositoryInterface entity.OrderInterface

type OrderRepositoryDB struct {
	Db  *sql.DB
	Ctx context.Context
}

func NewOrderRepository(ctx context.Context, db *sql.DB) (entity.OrderInterface, error) {
	if db == nil {
		return nil, errors.New("db pool cannot be empty")
	}

	return &OrderRepositoryDB{
		Db:  db,
		Ctx: ctx,
	}, nil
}

func (so *OrderRepositoryDB) FindAll() ([]entity.OrderResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	rows, err := so.Db.QueryContext(ctx, "SELECT id, name, description, stock, price, amount FROM \"order\";")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []entity.OrderResponse
	for rows.Next() {
		var id string
		var name string
		var description string
		var stock uint
		var price float32
		var amount uint
		// var category entity.Category
		err = rows.Scan(&id, &name, &description, &stock, &price, &amount)
		if err != nil {
			return nil, err
		}

		orders = append(orders, entity.OrderResponse{
			ID: id,
			Order: entity.Order{
				Name:        name,
				Description: description,
				Stock:       stock,
				Price:       price,
				Amount:      amount,
				// Category:    category,
			},
		})
	}
	return orders, nil
}

func (so *OrderRepositoryDB) FindByFilter(in *string) ([]entity.OrderResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	query := fmt.Sprintf("SELECT id, name, description, stock, price, amount FROM \"order\"  WHERE name ILIKE '%%%s%%' or description ILIKE '%%%s%%';", *in, *in)

	rows, err := so.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []entity.OrderResponse
	for rows.Next() {
		var id string
		var name string
		var description string
		var stock uint
		var price float32
		var amount uint
		// var category entity.Category
		err = rows.Scan(&id, &name, &description, &stock, &price, &amount)

		if err != nil {
			return nil, err
		}

		orders = append(orders, entity.OrderResponse{
			ID: id,
			Order: entity.Order{
				Name:        name,
				Description: description,
				Stock:       stock,
				Price:       price,
				Amount:      amount,
				// Category:    category,
			},
		})
	}
	return orders, nil
}

func (so *OrderRepositoryDB) FindByID(id *string) (*entity.OrderResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	query := fmt.Sprintf("SELECT id, name, description, stock, price, amount FROM \"order\"  WHERE id = '%s';", *id)

	rows, err := so.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var order entity.OrderResponse
  if rows.Next() {
    err := rows.Scan(&order.ID, &order.Name, &order.Description, &order.Stock, &order.Price, &order.Amount)
    if err != nil {
      return nil, err
    }
    return &order, nil
  }
	return &order, nil
}

func (so *OrderRepositoryDB) Create(order *entity.Order) (*entity.OrderResponse, error) {
	return nil, nil
}
func (so *OrderRepositoryDB) Delete(id *string) error {
	return nil
}
