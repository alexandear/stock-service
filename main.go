package main

import (
	"log"

	"github.com/devchallenge/stock-service/cmd"
)

func main() {
	if err := cmd.ExecuteServer(); err != nil {
		log.Fatal(err)
	}
}
