package additionalHttp

import "net/http"

var HopByHopHeaders = []string{
	"Connection",
	"Keep-Alive",
	"Proxy-Authenticate",
	"Proxy-Authorization",
	"TE",
	"Trailers",
	"Transfer-Encoding",
	"Upgrade",
}

func DeleteHopByHopHeaders(source http.Header) {
	for _, hopByHopHeader := range HopByHopHeaders {
		source.Del(hopByHopHeader)
	}
}

func CopyHeader(destination http.Header, source http.Header) {
	for key, value := range source {
		for _, matrixValue := range value {
			destination.Add(key, matrixValue)
		}
	}
}

func CopyHeadersWithoutHopByHop(destination http.Header, source http.Header) {
	DeleteHopByHopHeaders(source)

	CopyHeader(destination, source)
}
