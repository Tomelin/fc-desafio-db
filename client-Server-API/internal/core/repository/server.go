package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Tomelin/fc-desafio-db/internal/core/entity"
)

//
//type ExchangeRepositoryInterface interface {
//	Create(exchange *entity.Exchange) error
//	List() ([]entity.Exchange, error)
//	Delete(id *string) error
//	Update(exchange *entity.Exchange) error
//}

type ExchangeRepository struct {
	DB  *sql.DB
	Ctx context.Context
}

func NewRepositoryExchange(ctx context.Context, db *sql.DB) (entity.ExchangeInterface, error) {
	return &ExchangeRepository{Ctx: ctx, DB: db}, nil
}

func (e *ExchangeRepository) Create(ex *entity.ResponseCurrency) (*entity.ResponseCurrency, error) {

	query := fmt.Sprintf("INSERT INTO currency(id,code,codeIn, name, high, low, varBid,pctChange,bid,ask,timestamp,createDate) VALUES(%s,%s,%s,%s,%s,%s,%s,%s,%s);", ex.Id, ex.Code, ex.CodeIn, ex.Name, ex.High, ex.Low, ex.VarBid, ex.PctChange, ex.Bid, ex.Ask, ex.Timestamp, ex.CreateDate)

	request, err := e.DB.QueryContext(e.Ctx, query)
	if err != nil {
		return nil, err
	}

	defer request.Close()
	var result *entity.ResponseCurrency
	for request.Next() {
		err = request.Scan(&result.Id, &result.Code, &result.CodeIn, &result.Name, &result.High, &result.Low, &result.VarBid, &result.PctChange, &result.Bid, &result.Ask, &result.Timestamp, &result.CreateDate)
		if err != nil {

			return nil, err
		}
	}

	return result, nil
}

func (e *ExchangeRepository) Get() ([]entity.Exchange, error) { return nil, nil }

func (e *ExchangeRepository) Delete(id *string) error { return nil }

func (e *ExchangeRepository) Update(exchange *entity.Exchange) (*entity.Exchange, error) {
	return nil, nil
}
