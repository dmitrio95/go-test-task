/*
A package contains handlers for received HTTP requests.
*/
package handler

import (
	"io"
	"net"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/dmitrio95/go-test-task/response"
)

const ifacesURL = "/interfaces/"

type IfDataHandler struct{}

func (IfDataHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !checkGET(w, r) {
		return
	}

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

	writeResponse(w, r, data)
}

func HandleRootURL(w http.ResponseWriter, r *http.Request) {
	if !checkGET(w, r) {
		return
	}

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	data := response.ResponseData{"go-test-task", []response.ResponseEntry{{ifacesURL, "List of network interfaces", nil}}}

	writeResponse(w, r, data)
}

func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HandleRootURL)
	mux.Handle(ifacesURL, http.StripPrefix(ifacesURL, IfDataHandler{}))
	return mux
}

func chooseResponseFormat(acceptHeader string) (format string) {
	rawFormats := strings.Split(acceptHeader, ",")

	type fmtpriority struct {
		Format   string
		Priority float64
	}

	priority := make([]fmtpriority, len(rawFormats))

	for _, s := range rawFormats {
		s = strings.TrimSpace(s)
		splits := strings.Split(s, ";")
		if len(splits) < 2 {
			priority = append(priority, fmtpriority{splits[0], 1.0})
		} else if len(splits) == 2 {
			var qval = 1.0
			qvalsplit := strings.Split(strings.TrimSpace(splits[1]), "=")
			if len(qvalsplit) != 2 {
				// Probably an error. Ignore.
			} else {
				if qvalsplit[0] == "q" {
					qval1, err := strconv.ParseFloat(qvalsplit[1], 64)
					if err == nil {
						qval = qval1
					}
				}
			}
			priority = append(priority, fmtpriority{splits[0], qval})
		} else {
			// Something strange, probably an error.
			// Ignore this entry.
		}
	}

	// Sort priorities list descending by Priority value
	sort.Slice(priority, func(i, j int) bool { return priority[i].Priority > priority[j].Priority })

	for _, fmtprior := range priority {
		_, avail := response.Formats[fmtprior.Format]
		if avail {
			return fmtprior.Format
		}
	}

	return response.DefaultFormat
}

func checkGET(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "" && r.Method != "GET" {
		http.Error(w, "These data are read only!", http.StatusForbidden)
		return false
	}
	return true
}

// Writes a response by the given ResponseData structure
func writeResponse(w http.ResponseWriter, r *http.Request, data response.ResponseData) {
	acceptHeader := r.Header.Get("Accept")
	format := chooseResponseFormat(acceptHeader)
	fmt := response.NewResponseFormatter(format)

	resp, err := fmt.FormatResponse(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, resp)
}
