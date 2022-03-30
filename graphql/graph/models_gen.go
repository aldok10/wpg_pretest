// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph

import (
	"time"
)

type AccountIDInput struct {
	ID string `json:"id"`
}

type AccountInput struct {
	Name string `json:"name"`
}

type Order struct {
	ID         string            `json:"id"`
	CreatedAt  time.Time         `json:"createdAt"`
	TotalPrice float64           `json:"totalPrice"`
	Products   []*OrderedProduct `json:"products"`
}

type OrderInput struct {
	AccountID string               `json:"accountId"`
	Products  []*OrderProductInput `json:"products"`
}

type OrderProductInput struct {
	ID       string `json:"id"`
	Quantity int    `json:"quantity"`
}

type OrderedProduct struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

type PaginationInput struct {
	Skip *int `json:"skip"`
	Take *int `json:"take"`
}

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type ProductInput struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
