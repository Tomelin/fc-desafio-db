package entity

import (
	"errors"
	"github.com/google/uuid"
	"reflect"
)

type OrderInterface interface {
	Create(*Order) (*OrderResponse, error)
	FindAll() ([]OrderResponse, error)
	FindByFilter(*string) ([]OrderResponse, error)
	FindByID(*string) (*OrderResponse, error)
	Delete(*string) error
}

type Category uint

const (
	Software Category = iota + 1
	HomeAutomation
	BigData
	Unknow
)

type Order struct {
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description"`
	Stock       uint     `json:"stock" binding:"required"`
	Price       float32  `json:"price" binding:"required"`
	Amount      uint     `json:"amount" binding:"required"`
	Category    Category `json:"category" binding:"required"`
}

type OrderResponse struct {
	ID string `json:"id" binding:"required"`
	Order
}

func (o *Order) Validate() error {

	if reflect.ValueOf(o).IsNil() {
		return errors.New("order cannot be nil")
	}

	if o.Name == "" {
		return errors.New("order name cannot be blank")
	}

	if o.Amount == 0 {
		return errors.New("order cannot processed because the amount is zero")
	}

	if o.Amount > o.Stock {
		return errors.New("order cannot processed because amount greater than stock")
	}

	if o.Stock > 0 {
		if o.Price <= 0 {
			return errors.New("price must be greater than zero")
		}
	}

	if o.Category.GetCategoryName() == "Unknow" {
		return errors.New("category not found")
	}
	return nil
}

func (o *OrderResponse) Validate() error {
	err := o.Order.Validate()
	if err != nil {
		return err
	}

	_, err = uuid.Parse(o.ID)
	return err
}

func (c *Category) GetCategoryName() string {

	switch *c {
	case 1:
		return "Software"
	case 2:
		return "HomeAutomation"
	case 3:
		return "BigData"
	case 4:
		return "Unknow"
	default:
		return "Unknow"
	}
}
