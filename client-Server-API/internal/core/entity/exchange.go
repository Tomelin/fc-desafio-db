package entity

import (
	"context"
	"errors"
	"log/slog"
	"reflect"

	"github.com/google/uuid"
)

type ExchangeInterface interface {
	Get() ([]ResponseCurrency, error)
	Create(ctx context.Context, exchange *ResponseCurrency) (*ResponseCurrency, error)
	Delete(id *string) error
	Update(*Exchange) (*Exchange, error)
}

type ResponseCurrency struct {
	Id string `json:"id"`
	RequestCurrency
}

type ResponseBid string

type RequestCurrency struct {
	Code       string `json:"code"`
	CodeIn     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

type Exchange struct {
	USDBRL *RequestCurrency `json:"USDBRL"`
}

func NewExchange(e *RequestCurrency) (*ResponseCurrency, error) {

	err := e.Validate()
	if err != nil {
		return nil, err
	}

	return &ResponseCurrency{
		Id:              uuid.New().String(),
		RequestCurrency: *e,
	}, err
}

func (c *RequestCurrency) Validate() error {

	if reflect.ValueOf(c.Code).IsZero() {
		slog.Error("code cannot be empty")
		return errors.New("code cannot be empty")
	}

	if reflect.ValueOf(c.CodeIn).IsZero() {
		slog.Error("CodeIn cannot be empty")
		return errors.New("CodeIn cannot be empty")
	}
	if reflect.ValueOf(c.Name).IsZero() {
		slog.Error("Name cannot be empty")
		return errors.New("name cannot be empty")
	}
	if reflect.ValueOf(c.High).IsZero() {
		slog.Error("High cannot be empty")
		return errors.New("high cannot be empty")
	}
	if reflect.ValueOf(c.Low).IsZero() {
		slog.Error("Low cannot be empty")
		return errors.New("low cannot be empty")
	}
	if reflect.ValueOf(c.VarBid).IsZero() {
		slog.Error("VarBid cannot be empty")
		return errors.New("VarBid cannot be empty")
	}
	if reflect.ValueOf(c.PctChange).IsZero() {
		slog.Error("PctChange cannot be empty")
		return errors.New("PctChange cannot be empty")
	}
	if reflect.ValueOf(c.Bid).IsZero() {
		slog.Error("bid cannot be empty")
		return errors.New("bid cannot be empty")
	}
	if reflect.ValueOf(c.Ask).IsZero() {
		slog.Error("Ask cannot be empty")
		return errors.New("ask cannot be empty")
	}
	if reflect.ValueOf(c.Timestamp).IsZero() {
		slog.Error("Timestamp cannot be empty")
		return errors.New("timestamp cannot be empty")
	}
	if reflect.ValueOf(c.CreateDate).IsZero() {
		slog.Error("CreateDate cannot be empty")
		return errors.New("CreateDate cannot be empty")
	}

	return nil
}
