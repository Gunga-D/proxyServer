package connectionTypes

import (
	"io"
	"net/http"

	"github.com/Gunga-D/proxy-server/package/additionalHttp"
)

type Transmitter struct {
	core *http.Client
}

func NewTransmitter() *Transmitter {
	result := new(Transmitter)
	result.core = &http.Client{}

	return result
}

func (transmitter *Transmitter) transfer(destination http.ResponseWriter, source io.ReadCloser) {
	io.Copy(destination, source)
}

func (transmitter *Transmitter) Forward(writer http.ResponseWriter, request *http.Request) {
	response, err := transmitter.core.Do(request)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	additionalHttp.CopyHeader(writer.Header(), response.Header)
	additionalHttp.DeleteHopByHopHeaders(writer.Header())

	transmitter.transfer(writer, response.Body)
}
