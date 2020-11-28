package server

import (
	"context"
	"errors"

	devchallenge "github.com/devchallenge/stock-service/internal/grpcapi"
	"github.com/devchallenge/stock-service/internal/model"
)

type Stock interface {
	CreateOrder(ctx context.Context, params model.CreateOrderParams) (model.Order, error)
}

type Server struct {
	stock Stock
}

var _ devchallenge.StockServer = &Server{}

func New(stock Stock) *Server {
	return &Server{
		stock: stock,
	}
}

const (
	MessageOK            = "OK"
	MessageNotValidPrice = "Not valid price"
)

func (s *Server) CreateOrder(ctx context.Context, req *devchallenge.CreateOrderRequest,
) (*devchallenge.CreateOrderResponse, error) {
	order, err := s.stock.CreateOrder(ctx, model.CreateOrderParams{
		Symbol:   req.Symbol,
		Side:     model.OrderSide(req.Side),
		Type:     model.OrderType(req.Type),
		Price:    req.Price,
		Quantity: req.Quantity,
	})

	if errors.Is(model.ErrNotValidPrice, err) {
		return &devchallenge.CreateOrderResponse{
			Status:  devchallenge.ResponseStatus_FAILED,
			Message: MessageNotValidPrice,
		}, nil
	}

	if err != nil {
		return &devchallenge.CreateOrderResponse{
			Status:  devchallenge.ResponseStatus_FAILED,
			Message: err.Error(),
		}, nil
	}

	return &devchallenge.CreateOrderResponse{
		Status:  devchallenge.ResponseStatus_SUCCESS,
		Message: MessageOK,
		Id:      order.ID,
		Time:    uint64(order.CreatedAt.Unix()),
	}, nil
}

func (s *Server) CancelOrder(ctx context.Context, req *devchallenge.CancelOrderRequest,
) (*devchallenge.CancelOrderResponse, error) {
	panic("implement me")
}

func (s *Server) GetOrderBook(ctx context.Context, req *devchallenge.GetOrderBookRequest,
) (*devchallenge.GetOrderBookResponse, error) {
	panic("implement me")
}

func (s *Server) GetTrades(ctx context.Context, req *devchallenge.GetTradesRequest,
) (*devchallenge.GetTradesResponse, error) {
	panic("implement me")
}

func (s *Server) GetStats(ctx context.Context, req *devchallenge.GetStatsRequest,
) (*devchallenge.GetStatsResponse, error) {
	panic("implement me")
}
