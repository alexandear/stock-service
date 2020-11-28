package stock

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/devchallenge/stock-service/internal/model"
	"github.com/devchallenge/stock-service/internal/orderbook"
	"github.com/shopspring/decimal"
)

type IDGen interface {
	New() uint64
}

//go:generate mockgen -source=$GOFILE -package mock -destination mock/interfaces.go

const (
	maxOrders = 10000
)

type Service struct {
	idGen  IDGen
	clock  clock.Clock
	engine *orderbook.Engine

	stop chan struct{}

	allOrders map[uint64]model.Order

	mu          sync.Mutex
	matchOrders map[uint64]model.Order
}

func New(clock clock.Clock, idGen IDGen, engine *orderbook.Engine) *Service {
	s := &Service{
		idGen:  idGen,
		clock:  clock,
		engine: engine,
	}
	s.stop = make(chan struct{})
	s.allOrders = make(map[uint64]model.Order, maxOrders)
	s.matchOrders = make(map[uint64]model.Order, maxOrders)

	return s
}

func (s *Service) Stop() {
	log.Print("stopping signal")
	close(s.stop)
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

	var price float64

	if params.Type == model.OrderTypeLimit {
		p, err := strconv.ParseFloat(params.Price, 64)
		if err != nil {
			return model.Order{}, model.ErrNotValidPrice
		}

		price = p
	}

	quantity, err := strconv.ParseFloat(params.Quantity, 64)
	if err != nil {
		return model.Order{}, model.ErrNotValidQuantity
	}

	id := s.idGen.New()
	order := model.Order{
		ID:        id,
		Symbol:    params.Symbol,
		CreatedAt: s.clock.Now(),
		Side:      params.Side,
		Type:      params.Type,
		Price:     price,
		Quantity:  quantity,
	}

	s.allOrders[id] = order

	s.mu.Lock()
	s.matchOrders[id] = order
	s.mu.Unlock()

	return order, nil
}

func (s *Service) RunMatcher() {
	log.Print("running matcher")

	for {
		s.mu.Lock()
		for id, order := range s.matchOrders {
			if err := s.Match(order); err != nil {
				log.Printf("failed to match order=%v: %v", order, err)
			}

			delete(s.matchOrders, id)
		}
		s.mu.Unlock()
	}
}

func (s *Service) Match(order model.Order) error {
	_, _, err := s.engine.Match(newEngineOrder(order))
	if err != nil {
		return fmt.Errorf("failed to engine match: %w", err)
	}

	return nil
}

func (s *Service) OrderBook(symbol string, limit int) (bids, asks []model.Order, t time.Time) {
	buys := s.engine.OrderBookBuys()
	bids = filterOrders(symbol, limit, s.orders(buys))

	sells := s.engine.OrderBookSells()
	asks = filterOrders(symbol, limit, s.orders(sells))

	return bids, asks, s.clock.Now()
}

func (s *Service) orders(orders []*orderbook.Order) []model.Order {
	res := make([]model.Order, 0, len(orders))

	for _, order := range orders {
		res = append(res, s.newModelOrder(order))
	}

	return res
}

func filterOrders(symbol string, limit int, orders []model.Order) []model.Order {
	res := make([]model.Order, 0, limit)

	for i, order := range orders {
		if !strings.EqualFold(order.Symbol, symbol) {
			continue
		}

		if i == limit {
			break
		}

		res = append(res, order)
	}

	return res
}

func (s *Service) CancelOrder(ctx context.Context, symbol, id string) (time.Time, error) {
	now := s.clock.Now()

	orderID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return now, model.ErrNotValidID
	}

	if _, ok := s.allOrders[orderID]; !ok {
		return now, model.ErrOrderNotFound
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.matchOrders, orderID)

	return now, nil
}

func newEngineOrder(order model.Order) *orderbook.Order {
	price := decimal.NewFromFloat(order.Price)
	volume := decimal.NewFromFloat(order.Quantity)

	side := orderbook.SideBuy
	if order.Side == model.OrderSideSell {
		side = orderbook.SideSell
	}

	kind := orderbook.KindMarket
	locked := volume

	if order.Type == model.OrderTypeLimit {
		kind = orderbook.KindLimit
		locked = orderbook.CalculateLocked(volume, price, side)
	}

	return &orderbook.Order{
		ID:       order.ID,
		Side:     side,
		Kind:     kind,
		Price:    price,
		Volume:   volume,
		Locked:   locked,
		Received: decimal.Decimal{},
	}
}

func (s *Service) newModelOrder(order *orderbook.Order) model.Order {
	o := s.allOrders[order.ID]

	return model.Order{
		ID:        order.ID,
		Symbol:    o.Symbol,
		CreatedAt: o.CreatedAt,
		Side:      o.Side,
		Type:      o.Type,
		Price:     o.Price,
		Quantity:  o.Quantity,
	}
}
