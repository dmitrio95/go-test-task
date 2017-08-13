package main

import (
	"net/http"

	"github.com/dmitrio95/go-test-task/handler"
)

func main() {
	config := GetConfig()
	server := http.Server{Addr: config.Address, Handler: handler.NewHandler()}
	server.ListenAndServe()
}
