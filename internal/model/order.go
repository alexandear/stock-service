package model

import (
	"time"
)

type OrderSide int

const (
	OrderSideBuy OrderSide = iota + 1
	OrderSideSell
)

type OrderType int

const (
	OrderTypeLimit OrderType = iota + 1
	OrderTypeMarket
)

type Order struct {
	ID        uint64
	Symbol    string
	CreatedAt time.Time
	Side      OrderSide
	Type      OrderType
	Price     float64
	Quantity  float64
}

type CreateOrderParams struct {
	Symbol   string
	Side     OrderSide
	Type     OrderType
	Price    string
	Quantity string
}
