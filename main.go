package main

import (
	"fmt"
	"gateway/config"
	"gateway/startup"
)

var path = "config.yml"
var noAuthPath = "no_auth_config.yml"

func main() {
	conf, err := config.LoadConfig(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	noAuthConf, err := config.LoadConfig(noAuthPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	gateway := startup.NewServer(conf, noAuthConf)
	gateway.Start()
}
