package main

import (
	"log"

	"github.com/Gunga-D/proxy-server/internal/app"
	"github.com/Gunga-D/proxy-server/internal/infrastructure/config/server"
	"github.com/Gunga-D/proxy-server/internal/transport"
)

func main() {
	serverCoreConfig, err := server.GetServerCoreConfig()
	if err != nil {
		log.Fatalf("Server core config initialization error: %s", err.Error())
	}

	serverTlsConfig, err := server.GetServerTLSConfig()
	if err != nil {
		log.Fatalf("Server tls config initialization error: %s", err.Error())
	}

	handler := transport.NewProxyHandler()

	server := app.NewServer(handler, *serverCoreConfig)
	server.InitializeTLS(*serverTlsConfig)
	err = server.Run()
	if err != nil {
		log.Fatalf("Server startup error: %s", err.Error())
	}
	log.Println("The server has been started!")
}
