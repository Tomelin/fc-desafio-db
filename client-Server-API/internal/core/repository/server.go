package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Tomelin/fc-desafio-db/internal/core/entity"
	"log"
	"time"
)

type ExchangeRepository struct {
	DB  *sql.DB
	Ctx context.Context
}

func NewRepositoryExchange(ctx context.Context, db *sql.DB) (entity.ExchangeInterface, error) {
	return &ExchangeRepository{Ctx: ctx, DB: db}, nil
}

func (e *ExchangeRepository) Create(ctx context.Context, ex *entity.ResponseCurrency) (*entity.ResponseCurrency, error) {

	statement, err := e.DB.Prepare("INSERT INTO exchange(id, name, high, low, bid,timestamp) VALUES (?,?,?,?,?,?)")
	if err != nil {
		return nil, err
	}

	timeout := ctx.Value("timeout").(map[string]int)
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*time.Duration(timeout["db"]))
	defer cancel()
	//ctxx, cacnel := context.WithTimeout(ctx, time.Millisecond*1)
	result, err := statement.ExecContext(ctx, ex.Id, ex.Name, ex.High, ex.Low, ex.Bid, ex.Timestamp)
	if err != nil {
		if err.Error() == "context deadline exceeded" {
			return nil, fmt.Errorf("context deadline exceeded to exec SQL query")
		}
		return nil, err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return nil, errors.New("error to insert exchange")
	}

	return ex, nil
}

func (e *ExchangeRepository) Get() ([]entity.ResponseCurrency, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	rows, err := e.DB.QueryContext(ctx, "SELECT id, name, high, low, bid, timestamp FROM exchange")
	if err != nil {
		return nil, err
	}

	var exchanges []entity.ResponseCurrency
	for rows.Next() {
		var id string
		var name string
		var high string
		var low string
		var bid string
		var timestamp string
		log.Println(rows.Next())
		log.Println(rows.Columns())
		err := rows.Scan(&id, &name, &high, &low, &bid, &timestamp)
		log.Println(err, id, name, high, low, bid, timestamp)

		exchanges = append(exchanges, entity.ResponseCurrency{
			Id: id,
			RequestCurrency: entity.RequestCurrency{
				Name:      name,
				High:      high,
				Low:       low,
				Bid:       bid,
				Timestamp: timestamp,
			},
		})
	}
	return exchanges, nil
}

func (e *ExchangeRepository) Delete(id *string) error { return nil }

func (e *ExchangeRepository) Update(exchange *entity.Exchange) (*entity.Exchange, error) {
	return nil, nil
}
