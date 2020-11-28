package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/benbjohnson/clock"
	devchallenge "github.com/devchallenge/stock-service/internal/grpcapi"
	"github.com/devchallenge/stock-service/internal/server"
	"github.com/devchallenge/stock-service/internal/stock"
	"github.com/devchallenge/stock-service/internal/util"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"
)

func ExecuteServer() error {
	config := &Config{}
	config.InitFlags()

	pflag.Parse()

	address := fmt.Sprintf("%s:%s", config.Host, config.Port)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen via address=%s: %w", address, err)
	}

	log.Printf("Listening on address=%s", address)

	s := stock.New(clock.New(), &util.UUIDGen{})
	serv := server.New(s)

	grpcServer := grpc.NewServer()
	devchallenge.RegisterStockServer(grpcServer, serv)

	return grpcServer.Serve(lis)
}
