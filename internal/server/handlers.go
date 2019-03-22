package server

import (
	"encoding/json"
	"fmt"
	"github.com/R-Romanov/gohasher/pkg/hasher"
	"net/http"
)

const (
	dataParamName = "data"
)

func HashRequestHandler(w http.ResponseWriter, r *http.Request) {
	response := NewResponse()

	dataParams, isParamPresent := r.URL.Query()[dataParamName]

	if !isParamPresent {
		response.SetError(fmt.Sprintf("url param \"%s\" is missing", dataParamName))
		echoResponse(w, response)
		return
	}

	if len(dataParams[0]) < 1 {
		response.SetError(fmt.Sprintf("url param \"%s\" is empty", dataParamName))
		echoResponse(w, response)
		return
	}

	data := dataParams[0]

	h, err := hasher.Hash(data)

	if err == nil {
		response.SetHash(h)
	} else {
		response.SetError(err.Error())
	}

	echoResponse(w, response)
	return
}

func echoResponse(w http.ResponseWriter, response *Response) {
	responseMarshaled, err := json.Marshal(*response)
	if err != nil {
		panic(err)
	}

	_, err = w.Write(responseMarshaled)
	if err != nil {
		panic(err)
	}
}
