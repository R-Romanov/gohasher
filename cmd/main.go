package main

import (
	"flag"
	"fmt"
	"github.com/R-Romanov/gohasher/internal/server"
	"net/http"
	"strconv"
)

func main() {
	port := flag.Int("port", 8080, "port number")
	flag.Parse()

	portAddress := ":" + strconv.Itoa(*port)
	fmt.Printf("Listening on %d\n", *port)

	http.HandleFunc("/", server.HashRequestHandler)
	err := http.ListenAndServe(portAddress, nil)

	if err != nil {
		panic("can not run http server: " + err.Error())
	}
}
