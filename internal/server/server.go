package server

import (
	"context"
	"errors"
	"strconv"
	"time"

	devchallenge "github.com/devchallenge/stock-service/internal/grpcapi"
	"github.com/devchallenge/stock-service/internal/model"
)

type Stock interface {
	CreateOrder(ctx context.Context, params model.CreateOrderParams) (model.Order, error)
	CancelOrder(ctx context.Context, symbol, id string) (time.Time, error)
	OrderBook(symbol string, limit int) (bids, asks []model.Order, t time.Time)

	RunMatcher()
	Stop()
}

type Server struct {
	stock Stock
}

var _ devchallenge.StockServer = &Server{}

func New(stock Stock) *Server {
	s := &Server{
		stock: stock,
	}

	go s.stock.RunMatcher()

	return s
}

const (
	MessageOK            = "OK"
	MessageNotValidPrice = "Not valid price"
	MessageOrderNotFound = "Order not found"
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
		Id:      strconv.FormatUint(order.ID, 10),
		Time:    respTime(order.CreatedAt),
	}, nil
}

func (s *Server) CancelOrder(ctx context.Context, req *devchallenge.CancelOrderRequest,
) (*devchallenge.CancelOrderResponse, error) {
	t, err := s.stock.CancelOrder(ctx, req.Symbol, req.Id)

	if errors.Is(err, model.ErrOrderNotFound) {
		return &devchallenge.CancelOrderResponse{
			Status:  devchallenge.ResponseStatus_FAILED,
			Message: MessageOrderNotFound,
			Time:    respTime(t),
		}, nil
	}

	if err != nil {
		return &devchallenge.CancelOrderResponse{
			Status:  devchallenge.ResponseStatus_FAILED,
			Message: err.Error(),
			Time:    respTime(t),
		}, nil
	}

	return &devchallenge.CancelOrderResponse{
		Status:  devchallenge.ResponseStatus_SUCCESS,
		Message: MessageOK,
		Time:    respTime(t),
	}, nil
}

func (s *Server) GetOrderBook(ctx context.Context, req *devchallenge.GetOrderBookRequest,
) (*devchallenge.GetOrderBookResponse, error) {
	bids, asks, t := s.stock.OrderBook(req.Symbol, int(req.Limit))

	respBids := make([]*devchallenge.OrderBookEntity, 0, len(bids))
	for _, order := range bids {
		respBids = append(respBids, respOrderBookEntity(order))
	}

	respAsks := make([]*devchallenge.OrderBookEntity, 0, len(asks))
	for _, order := range asks {
		respAsks = append(respAsks, respOrderBookEntity(order))
	}

	return &devchallenge.GetOrderBookResponse{
		Bids: respBids,
		Asks: respAsks,
		Time: respTime(t),
	}, nil
}

func (s *Server) GetTrades(ctx context.Context, req *devchallenge.GetTradesRequest,
) (*devchallenge.GetTradesResponse, error) {
	panic("implement me")
}

func (s *Server) GetStats(ctx context.Context, req *devchallenge.GetStatsRequest,
) (*devchallenge.GetStatsResponse, error) {
	panic("implement me")
}

func (s *Server) Stop() {
	s.stock.Stop()
}

func respOrderBookEntity(order model.Order) *devchallenge.OrderBookEntity {
	return &devchallenge.OrderBookEntity{
		Price:    strconv.FormatFloat(order.Price, 'f', 2, 64),
		Quantity: strconv.FormatFloat(order.Quantity, 'f', 3, 64),
	}
}

func respTime(t time.Time) uint64 {
	return uint64(t.Unix())
}
