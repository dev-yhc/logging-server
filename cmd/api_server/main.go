package main

import (
	"log"
	"yhc/logging-server/app/api_server"
	"yhc/logging-server/app/config"
)

func main() {
	config, err := config.NewConfig()

	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	api_server.Run(config)
}
