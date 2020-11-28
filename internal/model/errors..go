package model

import (
	"errors"
)

var (
	ErrNotValidPrice     = errors.New("not valid price")
	ErrNotValidQuantity  = errors.New("not valid quantity")
	ErrNotValidSymbol    = errors.New("not valid symbol")
	ErrNotValidSide      = errors.New("not valid side")
	ErrNotValidOrderType = errors.New("not valid order type")
)
