package connectionTypes

import (
	"io"
	"net"
	"net/http"
	"time"
)

type Tunnel struct {
}

func NewTunnel() *Tunnel {
	result := new(Tunnel)

	return result
}

func (tunnel *Tunnel) transfer(destination io.WriteCloser, source io.ReadCloser) {
	defer destination.Close()
	defer source.Close()

	io.Copy(destination, source)
}

func (tunnel *Tunnel) Forward(writer http.ResponseWriter, request *http.Request) {
	destinationConnection, err := net.DialTimeout("tcp", request.Host, 10*time.Second)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)

	hijacker, isSupported := writer.(http.Hijacker)
	if !isSupported {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	clientConnection, _, err := hijacker.Hijack()
	if err != nil {
		writer.WriteHeader(http.StatusServiceUnavailable)
	}

	go tunnel.transfer(destinationConnection, clientConnection)
	go tunnel.transfer(clientConnection, destinationConnection)
}
