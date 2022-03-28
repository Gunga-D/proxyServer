package app

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"time"
)

type Server struct {
	address        string
	handler        http.Handler
	maxHeaderBytes int
	readTimeout    int
	writeTimeout   int
	core           *http.Server
}

func NewServer(ip string, port string, handler http.Handler,
	maxHeaderBytes int, readTimeout int, writeTimeout int) *Server {
	proto := new(Server)

	proto.address = ip + ":" + port
	proto.handler = handler
	proto.maxHeaderBytes = maxHeaderBytes
	proto.readTimeout = readTimeout
	proto.writeTimeout = writeTimeout

	return proto
}

func (server *Server) Run() error {
	server.core = &http.Server{
		Addr:           server.address,
		Handler:        server.handler,
		MaxHeaderBytes: server.maxHeaderBytes,
		ReadTimeout:    time.Duration(server.readTimeout) * time.Second,
		WriteTimeout:   time.Duration(server.writeTimeout) * time.Second,
		TLSNextProto:   make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}

	log.Println("The server has been started!")

	return server.core.ListenAndServe()
}

func (server *Server) Shutdown(ctx context.Context) error {
	return server.core.Shutdown(ctx)
}
