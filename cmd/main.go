package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/R-Romanov/gohasher"
	"github.com/R-Romanov/gohasher/internal/server"
	"net/http"
	"strconv"
)

func main() {
	port := flag.Int("port", 8080, "port number")
	flag.Parse()

	http.HandleFunc("/", algoListHandler)
	http.HandleFunc("/get-hash/", getHashHandler)

	portAddress := ":" + strconv.Itoa(*port)
	fmt.Printf("Start server on port %d", *port)

	err := http.ListenAndServe(portAddress, nil)

	if err != nil {
		panic("can not run http server")
	}
}

func getHashHandler(w http.ResponseWriter, r *http.Request) {
	response := server.NewResponse()

	datum, ok := r.URL.Query()["data"]

	if !ok || len(datum[0]) < 1 {
		response.SetError("url param 'data' is missing")
		echoResponse(w, response)
		return
	}
	data := datum[0]

	algoParam, ok := r.URL.Query()["algo"]

	if !ok || len(algoParam[0]) < 1 {
		response.SetError("url param 'method' is missing")
		echoResponse(w, response)
		return
	}
	hashMethod := algoParam[0]

	dataHashed, err := gohasher.HashString(data, hashMethod)

	if err != nil {
		response.SetError(err.Error())
	} else {
		response.SetHashedString(dataHashed)
	}

	echoResponse(w, response)
	return

}

func algoListHandler(w http.ResponseWriter, r *http.Request) {
	listString := "Algorithms available:" +
		"\n- fnv1 : non-cryptographic 64-bit hash function (default)"

	_, err := w.Write([]byte(listString))
	if err != nil {
		panic(err)
	}
}

func echoResponse(w http.ResponseWriter, response *server.Response) {
	responseMarshaled, err := json.Marshal(*response)
	if err != nil {
		panic(err)
	}

	_, err = w.Write(responseMarshaled)
	if err != nil {
		panic(err)
	}
}
