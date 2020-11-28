package server

import (
	"context"
	"testing"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	devchallenge "github.com/devchallenge/stock-service/internal/grpcapi"
	"github.com/devchallenge/stock-service/internal/stock"
	"github.com/devchallenge/stock-service/internal/stock/mock"
)

func TestServer_CreateOrder(t *testing.T) {
	t.Run("ok limit", func(t *testing.T) {
		gctrl := gomock.NewController(t)
		defer gctrl.Finish()
		cm := clock.NewMock()
		um := mock.NewMockUUIDGen(gctrl)
		s := stock.New(cm, um)
		serv := &Server{stock: s}
		now := time.Now().UTC()
		cm.Set(now)
		id := uuid.New().String()
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
			Id:      id,
			Time:    respTime(now),
		}, resp)
	})

	t.Run("ok market", func(t *testing.T) {
		gctrl := gomock.NewController(t)
		defer gctrl.Finish()
		cm := clock.NewMock()
		um := mock.NewMockUUIDGen(gctrl)
		s := stock.New(cm, um)
		serv := &Server{stock: s}
		now := time.Now().UTC()
		cm.Set(now)
		id := uuid.New().String()
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
			Id:      id,
			Time:    respTime(now),
		}, resp)
	})

	t.Run("not valid price", func(t *testing.T) {
		gctrl := gomock.NewController(t)
		defer gctrl.Finish()
		s := stock.New(nil, nil)
		serv := &Server{stock: s}

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

func respTime(t time.Time) uint64 {
	return uint64(t.Unix())
}
