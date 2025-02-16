package api_server

import (
	"fmt"
	"yhc/logging-server/app/config"

	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func Run(config *config.Config) {
	fmt.Println(config)
	// HTTP Server -.
	handler := gin.Default()

	// add router

	handler.Run()

	// Waiting signal -.
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		fmt.Println("app - Run - signal: " + s.String())
	}

	// graceful shut down
}
