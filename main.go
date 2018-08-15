package main

import (
	"fmt"
	"github.com/c12s/lunar-gateway/model/configs"
	"github.com/c12s/lunar-gateway/server"
	"log"
)

func main() {
	conf, err := configs.ConfigFile()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("conf")
	fmt.Println(conf)

	server := server.NewServer(conf)
	server.Start()
}
