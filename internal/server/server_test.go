package server_test

import (
	"context"
	"testing"
	"time"

	"github.com/benbjohnson/clock"
	devchallenge "github.com/devchallenge/stock-service/internal/grpcapi"
	"github.com/devchallenge/stock-service/internal/orderbook"
	"github.com/devchallenge/stock-service/internal/server"
	"github.com/devchallenge/stock-service/internal/stock"
	"github.com/devchallenge/stock-service/internal/stock/mock"
	"github.com/golang/mock/gomock"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServer_CreateOrder(t *testing.T) {
	t.Run("ok limit", func(t *testing.T) {
		gctrl := gomock.NewController(t)
		defer gctrl.Finish()
		cm := clock.NewMock()
		um := mock.NewMockIDGen(gctrl)
		s := stock.New(cm, um, orderbook.New())
		serv := server.New(s)
		now := time.Now().UTC()
		cm.Set(now)
		id := uint64(1)
		um.EXPECT().New().Return(id).Times(1)

		resp, err := serv.CreateOrder(context.Background(), &devchallenge.CreateOrderRequest{
			Symbol:   fake.CharactersN(4),
			Side:     devchallenge.OrderSide_BUY,
			Type:     devchallenge.OrderType_LIMIT,
			Price:    "100.25",
			Quantity: "15.5",
		})

		require.NoError(t, err)
		assert.Equal(t, &devchallenge.CreateOrderResponse{
			Status:  devchallenge.ResponseStatus_SUCCESS,
			Message: "OK",
			Id:      "1",
			Time:    respTime(now),
		}, resp)
	})

	t.Run("ok market", func(t *testing.T) {
		gctrl := gomock.NewController(t)
		defer gctrl.Finish()
		cm := clock.NewMock()
		um := mock.NewMockIDGen(gctrl)
		s := stock.New(cm, um, orderbook.New())
		serv := server.New(s)
		now := time.Now().UTC()
		cm.Set(now)
		id := uint64(1)
		um.EXPECT().New().Return(id).Times(1)

		resp, err := serv.CreateOrder(context.Background(), &devchallenge.CreateOrderRequest{
			Symbol:   fake.CharactersN(4),
			Side:     devchallenge.OrderSide_BUY,
			Type:     devchallenge.OrderType_MARKET,
			Quantity: "15.5",
		})

		require.NoError(t, err)
		assert.Equal(t, &devchallenge.CreateOrderResponse{
			Status:  devchallenge.ResponseStatus_SUCCESS,
			Message: "OK",
			Id:      "1",
			Time:    respTime(now),
		}, resp)
	})

	t.Run("not valid price", func(t *testing.T) {
		gctrl := gomock.NewController(t)
		defer gctrl.Finish()
		s := &stock.Service{}
		serv := server.New(s)

		resp, err := serv.CreateOrder(context.Background(), &devchallenge.CreateOrderRequest{
			Symbol:   fake.CharactersN(4),
			Side:     devchallenge.OrderSide_BUY,
			Type:     devchallenge.OrderType_LIMIT,
			Price:    fake.Characters(),
			Quantity: "24.2",
		})

		require.NoError(t, err)
		assert.Equal(t, &devchallenge.CreateOrderResponse{
			Status:  devchallenge.ResponseStatus_FAILED,
			Message: "Not valid price",
		}, resp)
	})
}

func TestServer_CancelOrder(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		gctrl := gomock.NewController(t)
		defer gctrl.Finish()
		cm := clock.NewMock()
		um := mock.NewMockIDGen(gctrl)
		engine := orderbook.New()
		s := stock.New(cm, um, engine)
		serv := server.New(s)
		now := time.Now().UTC()
		cm.Set(now)
		id := uint64(100)
		symbol := fake.CharactersN(4)
		um.EXPECT().New().Return(id).Times(1)
		_, err := serv.CreateOrder(context.Background(), &devchallenge.CreateOrderRequest{
			Symbol:   symbol,
			Side:     devchallenge.OrderSide_BUY,
			Type:     devchallenge.OrderType_LIMIT,
			Price:    "100.25",
			Quantity: "15.5",
		})
		require.NoError(t, err)

		resp, err := serv.CancelOrder(context.Background(), &devchallenge.CancelOrderRequest{
			Symbol: symbol,
			Id:     "100",
		})

		require.NoError(t, err)
		assert.Equal(t, &devchallenge.CancelOrderResponse{
			Status:  devchallenge.ResponseStatus_SUCCESS,
			Message: "OK",
			Time:    respTime(now),
		}, resp)
	})

	t.Run("order not found", func(t *testing.T) {
		gctrl := gomock.NewController(t)
		defer gctrl.Finish()
		cm := clock.NewMock()
		s := stock.New(cm, nil, orderbook.New())
		serv := server.New(s)
		now := time.Now().UTC()
		cm.Set(now)

		resp, err := serv.CancelOrder(context.Background(), &devchallenge.CancelOrderRequest{
			Symbol: fake.CharactersN(4),
			Id:     "100500",
		})

		require.NoError(t, err)
		assert.Equal(t, &devchallenge.CancelOrderResponse{
			Status:  devchallenge.ResponseStatus_FAILED,
			Message: "Order not found",
			Time:    respTime(now),
		}, resp)
	})
}

func TestServer_GetOrderBook(t *testing.T) {
	gctrl := gomock.NewController(t)
	defer gctrl.Finish()
	cm := clock.NewMock()
	um := mock.NewMockIDGen(gctrl)
	engine := orderbook.New()
	s := stock.New(cm, um, engine)
	serv := server.New(s)
	now := time.Now().UTC()
	cm.Set(now)
	symbol := fake.CharactersN(4)
	um.EXPECT().New().Return(uint64(1)).Times(1)
	um.EXPECT().New().Return(uint64(2)).Times(1)
	um.EXPECT().New().Return(uint64(3)).Times(1)
	um.EXPECT().New().Return(uint64(4)).Times(1)
	_, err := serv.CreateOrder(context.Background(), &devchallenge.CreateOrderRequest{
		Symbol:   symbol,
		Side:     devchallenge.OrderSide_BUY,
		Type:     devchallenge.OrderType_LIMIT,
		Price:    "1.01",
		Quantity: "2.002",
	})
	require.NoError(t, err)
	_, err = serv.CreateOrder(context.Background(), &devchallenge.CreateOrderRequest{
		Symbol:   symbol,
		Side:     devchallenge.OrderSide_BUY,
		Type:     devchallenge.OrderType_LIMIT,
		Price:    "3.03",
		Quantity: "4.004",
	})
	require.NoError(t, err)
	_, err = serv.CreateOrder(context.Background(), &devchallenge.CreateOrderRequest{
		Symbol:   symbol,
		Side:     devchallenge.OrderSide_BUY,
		Type:     devchallenge.OrderType_LIMIT,
		Price:    "5.05",
		Quantity: "6.006",
	})
	require.NoError(t, err)
	_, err = serv.CreateOrder(context.Background(), &devchallenge.CreateOrderRequest{
		Symbol:   symbol,
		Side:     devchallenge.OrderSide_SELL,
		Type:     devchallenge.OrderType_MARKET,
		Price:    "1.11",
		Quantity: "5",
	})
	require.NoError(t, err)
	time.Sleep(1 * time.Second)

	resp, err := serv.GetOrderBook(context.Background(), &devchallenge.GetOrderBookRequest{
		Symbol: symbol,
		Limit:  2,
	})

	require.NoError(t, err)
	assert.Equal(t, &devchallenge.GetOrderBookResponse{
		Bids: []*devchallenge.OrderBookEntity{
			{
				Price:    "5.05",
				Quantity: "6.006",
			},
			{
				Price:    "3.03",
				Quantity: "4.004",
			},
		},
		Asks: []*devchallenge.OrderBookEntity{},
		Time: respTime(now),
	}, resp)
}

func respTime(t time.Time) uint64 {
	return uint64(t.Unix())
}
