package main

import (
	"github.com/remy-z/gobank/api"
)

func main() {
	server := api.NewAPIServer(":3000")
	server.Run()
}
