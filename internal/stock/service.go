package stock

import (
	"context"
	"strconv"

	"github.com/benbjohnson/clock"
	"github.com/devchallenge/stock-service/internal/model"
)

type UUIDGen interface {
	New() string
}

//go:generate mockgen -source=$GOFILE -package mock -destination mock/interfaces.go

type Service struct {
	uuid  UUIDGen
	clock clock.Clock
}

func New(clock clock.Clock, uuid UUIDGen) *Service {
	return &Service{
		uuid:  uuid,
		clock: clock,
	}
}

func (s *Service) CreateOrder(ctx context.Context, params model.CreateOrderParams) (model.Order, error) {
	if params.Symbol == "" {
		return model.Order{}, model.ErrNotValidSymbol
	}

	switch params.Side {
	case model.OrderSideBuy, model.OrderSideSell:
	default:
		return model.Order{}, model.ErrNotValidSide
	}

	switch params.Type {
	case model.OrderTypeLimit, model.OrderTypeMarket:
	default:
		return model.Order{}, model.ErrNotValidOrderType
	}

	if params.Type == model.OrderTypeLimit {
		_, err := strconv.ParseFloat(params.Price, 64)
		if err != nil {
			return model.Order{}, model.ErrNotValidPrice
		}
	}

	_, err := strconv.ParseFloat(params.Quantity, 64)
	if err != nil {
		return model.Order{}, model.ErrNotValidQuantity
	}

	return model.Order{
		ID:        s.uuid.New(),
		CreatedAt: s.clock.Now(),
	}, nil
}
