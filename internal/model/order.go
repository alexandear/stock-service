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
	ID        string
	CreatedAt time.Time
}

type CreateOrderParams struct {
	Symbol   string
	Side     OrderSide
	Type     OrderType
	Price    string
	Quantity string
}
