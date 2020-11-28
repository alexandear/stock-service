package cmd

import (
	"github.com/spf13/pflag"
)

func ExecuteServer() error {
	_ = &Config{}

	pflag.Parse()

	println("Stock service")

	return nil
}
