/*
A package contains handlers for received HTTP requests.
*/
package handler

import (
	"io"
	"net"
	"net/http"

	"github.com/dmitrio95/go-test-task/response"
)

type IfDataHandler struct{}

func (IfDataHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "" && r.Method != "GET" {
		http.Error(w, "These data are read only!", http.StatusForbidden)
		return
	}

	fmt := response.NewResponseFormatter("text/html")
	var data response.ResponseData

	switch path := r.URL.Path; path {
	case "":
		ifaces, err := net.Interfaces()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data = response.NewResponseData(ifaces)
	default:
		iface, err := net.InterfaceByName(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		data = response.NewResponseData([]net.Interface{*iface})
	}

	resp, err := fmt.FormatResponse(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, resp)
}

func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/interfaces/", http.StripPrefix("/interfaces/", IfDataHandler{}))
	return mux
}
