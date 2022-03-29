package main

import (
	"log"

	"github.com/Gunga-D/proxy-server/internal/app"
	serverConfig "github.com/Gunga-D/proxy-server/internal/infrastructure/config/server"
	"github.com/Gunga-D/proxy-server/internal/transport"
)

func main() {
	handler := transport.NewProxyHandler()

	serverCoreConfig, err := serverConfig.GetServerCoreConfig()
	if err != nil {
		log.Fatalf("Server core config initialization error: %s", err.Error())
	}
	server := app.NewServer(handler, *serverCoreConfig)

	if serverCoreConfig.TransferProtocol == "https" {
		serverTlsConfig, err := serverConfig.GetServerTLSConfig()
		if err != nil {
			log.Fatalf("Server tls config initialization error: %s", err.Error())
		}
		server.InitializeTLS(*serverTlsConfig)
	}

	err = server.Run()
	if err != nil {
		log.Fatalf("Server startup error: %s", err.Error())
	}
}
