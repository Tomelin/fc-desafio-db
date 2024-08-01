package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

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
	rows, err := so.Db.QueryContext(ctx, "SELECT id, name, description, stock, price, amount, category FROM order")
	if err != nil {
		return nil, err
	}

	var orders []entity.OrderResponse
	for rows.Next() {
		var id string
		var name string
		var description string
		var stock uint
		var price float32
		var amount uint
		var category entity.Category
		log.Println(rows.Next())
		log.Println(rows.Columns())
		err := rows.Scan(&id, &name, &description, &stock, &price, &amount, &category)
		log.Println(err, id, name, description, stock, price, amount, category)

		orders = append(orders, entity.OrderResponse{
			ID: id,
			Order: entity.Order{
				Name:        name,
				Description: description,
				Stock:       stock,
				Price:       price,
				Amount:      amount,
				Category:    category,
			},
		})
	}
	return orders, nil
}

func (so *OrderRepositoryDB) FindByFilter(*string) ([]entity.OrderResponse, error) {
	return nil, nil
}

func (so *OrderRepositoryDB) FindByID(id *string) (*entity.OrderResponse, error) {
	return nil, nil
}
func (so *OrderRepositoryDB) Create(order *entity.Order) (*entity.OrderResponse, error) {
	return nil, nil
}
func (so *OrderRepositoryDB) Delete(id *string) error {
	return nil
}
