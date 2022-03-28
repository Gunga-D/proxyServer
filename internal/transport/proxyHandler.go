package transport

import (
	"log"
	"net/http"

	"github.com/Gunga-D/proxy-server/internal/transport/connectionTypes"
	"github.com/Gunga-D/proxy-server/package/additionalHttp"
)

type proxyHandler struct {
	transmitter *connectionTypes.Transmitter
	tunnel      *connectionTypes.Tunnel
}

func NewProxyHandler() *proxyHandler {
	proto := new(proxyHandler)

	proto.transmitter = connectionTypes.NewTransmitter()
	proto.tunnel = connectionTypes.NewTunnel()

	return proto
}

func (handler *proxyHandler) changeLocation(request *http.Request) {
	request.Header.Set("X-Forwarded-For", request.RemoteAddr)
	request.Header.Set("Accept-Language", "en-US")
}

func (handler *proxyHandler) normalizeProxyRequest(request *http.Request) {
	request.RequestURI = ""

	additionalHttp.DeleteHopByHopHeaders(request.Header)
}

func (handler *proxyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Println("The request by " + request.RemoteAddr + " and to " + request.URL.Host + " was passed")

	handler.changeLocation(request)
	handler.normalizeProxyRequest(request)

	if request.Method == http.MethodConnect {
		handler.tunnel.Forward(writer, request)
	} else {
		handler.transmitter.Forward(writer, request)
	}
}
