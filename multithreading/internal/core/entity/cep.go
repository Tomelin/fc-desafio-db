package entity

import (
	"errors"
	"time"
)

type CEP struct {
	Number   string
	Origin   string
	Duration time.Time
}

func NewCEP(n, o string) *CEP {
	c := CEP{
		Number: n,
		Origin: o,
	}

	err := c.Validate()
	if err != nil {
		panic(err)
	}
	return &c
}

func (c *CEP) Validate() error {

	if len(c.Number) != 8 {
		errors.New("cep invalid")
	}
	return nil
}
