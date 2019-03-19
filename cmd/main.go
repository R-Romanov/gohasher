package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/R-Romanov/gohasher"
	"github.com/R-Romanov/gohasher/internal/hasherResponse"
	"net/http"
	"strconv"
)

func main() {
	port := flag.Int("port", 8080, "an int")
	flag.Parse()

	http.HandleFunc("/get-hash/", getHashHandler)
	http.HandleFunc("/hash-methods-list/", hashMethodsListHandler)

	portAddress := ":" + strconv.Itoa(*port)
	fmt.Printf("Start server on port %d", *port)

	err := http.ListenAndServe(portAddress, nil)

	if err != nil {
		panic("can not run http server")
	}

}

func getHashHandler(w http.ResponseWriter, r *http.Request) {
	response := hasherResponse.NewHasherResponse()

	datum, ok := r.URL.Query()["data"]

	if !ok || len(datum[0]) < 1 {
		response.SetError("url param 'data' is missing")
		echoResponse(w, response)
		return
	}
	data := datum[0]

	hashMethods, ok := r.URL.Query()["method"]

	if !ok || len(hashMethods[0]) < 1 {
		response.SetError("url param 'method' is missing")
		echoResponse(w, response)
		return
	}
	hashMethod := hashMethods[0]

	dataHashed, err := gohasher.HashString(data, hashMethod)

	if err != nil {
		response.SetError(err.Error())
	} else {
		response.SetHashedString(dataHashed)
	}

	echoResponse(w, response)
	return

}

func hashMethodsListHandler(w http.ResponseWriter, r *http.Request) {
	listString :=
		`The following encryption methods are available:
	- fnv1 : non-cryptographic hash function (int64)
`

	_, err := w.Write([]byte(listString))
	if err != nil {
		panic(err)
	}
}

func echoResponse(w http.ResponseWriter, response *hasherResponse.HasherResponse) {
	responseMarshaled, err := json.Marshal(*response)
	if err != nil {
		panic(err)
	}

	_, err = w.Write(responseMarshaled)
	if err != nil {
		panic(err)
	}
}