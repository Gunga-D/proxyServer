package main

import (
	"log"

	"github.com/Gunga-D/proxy-server/internal/app"
	"github.com/Gunga-D/proxy-server/internal/infrastructure/config"
	"github.com/Gunga-D/proxy-server/internal/transport"
)

func main() {
	serverConfig, err := config.GetServerConfig()
	if err != nil {
		log.Fatalf("Сonfig initialization error: %s", err.Error())
	}

	handler := transport.NewProxyHandler()

	server := app.NewServer(serverConfig.Ip, serverConfig.Port, handler,
		serverConfig.MaxHeaderBytes, serverConfig.ReadTimeout, serverConfig.WriteTimeout)
	err = server.Run()
	if err != nil {
		log.Fatalf("Server startup error: %s", err.Error())
	}
}
