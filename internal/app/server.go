package app

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"time"

	"github.com/Gunga-D/proxy-server/internal/infrastructure/config/server"
)

type Server struct {
	tlsConfig  *server.TLSServerConfigEntity
	coreConfig *server.CoreServerConfigEntity
	handler    *http.Handler
	core       *http.Server
}

func NewServer(handler http.Handler, coreConfig server.CoreServerConfigEntity) *Server {
	proto := new(Server)

	proto.handler = &handler
	proto.coreConfig = &coreConfig

	return proto
}

func (server *Server) InitializeTLS(tlsConfig server.TLSServerConfigEntity) {
	server.tlsConfig = &tlsConfig
}

func (server *Server) Run() error {
	server.core = &http.Server{
		Addr:           server.coreConfig.Ip + ":" + server.coreConfig.Port,
		Handler:        *server.handler,
		MaxHeaderBytes: server.coreConfig.MaxHeaderBytes,
		ReadTimeout:    time.Duration(server.coreConfig.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(server.coreConfig.WriteTimeout) * time.Second,
		TLSNextProto:   make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}

	log.Println("Server has been started!")

	if server.coreConfig.TransferProtocol == "https" {
		return server.core.ListenAndServeTLS(server.tlsConfig.PemPath, server.tlsConfig.KeyPath)
	}

	return server.core.ListenAndServe()
}

func (server *Server) Shutdown(ctx context.Context) error {
	return server.core.Shutdown(ctx)
}
