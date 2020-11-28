package cmd

import (
	"github.com/spf13/pflag"
)

const (
	defaultHost = "localhost"
	defaultPort = "9090"
)

type Config struct {
	Host string
	Port string
}

func (c *Config) InitFlags() {
	pflag.StringVar(&c.Host, "host", defaultHost, "host")
	pflag.StringVar(&c.Port, "port", defaultPort, "port")
}
