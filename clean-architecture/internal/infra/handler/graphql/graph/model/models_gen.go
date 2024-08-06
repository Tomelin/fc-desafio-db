// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Order struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
	Amount      int     `json:"amount"`
	Category    *int    `json:"category,omitempty"`
}

type Query struct {
}