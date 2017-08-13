/*
A package contains handlers for received HTTP requests.
*/
package handler

import (
	"fmt"
	"io"
	"net"
	"net/http"
)

type IfDataHandler struct{}

func (IfDataHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "" && r.Method != "GET" {
		http.Error(w, "These data are read only!", http.StatusForbidden)
		return
	}

	switch path := r.URL.Path; path {
	case "":
		ifaces, err := net.Interfaces()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		io.WriteString(w, fmt.Sprint(ifaces))
	default:
		iface, err := net.InterfaceByName(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		io.WriteString(w, fmt.Sprint(iface))
	}
}

func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/interfaces/", http.StripPrefix("/interfaces/", IfDataHandler{}))
	return mux
}
