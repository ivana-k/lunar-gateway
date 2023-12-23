package main

import (
	"fmt"
	"gateway/config"
	"gateway/startup"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	noAuthConf, err := config.LoadNoAuthConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	gateway := startup.NewServer(conf, noAuthConf)
	gateway.Start()
}
