package main

import (
	"github.com/c12s/lunar-gateway/server"
)

func main() {
	server := server.NewServer("localhost", "8080")
	server.Test()
	server.Start()
}
