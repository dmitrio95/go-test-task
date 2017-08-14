package main

import (
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dmitrio95/go-test-task/handler"
)

const ifacesURL = "/interfaces/"

var badURLs = []string{
	"/asdf/",
	"/asdf/lo/",
	"/interfaces/asdf/",
}

func newEmptyReader() io.Reader {
	return strings.NewReader("")
}

// Currently service offers read-only data, so any POST
// request should fail.
func TestPOST(t *testing.T) {
	h := handler.NewHandler()

	req := httptest.NewRequest("POST", ifacesURL, newEmptyReader())
	resp := httptest.NewRecorder()

	h.ServeHTTP(resp, req)

	if resp.Code == http.StatusOK {
		t.Error("Succesful POST request")
	}
}

func TestBadURLs(t *testing.T) {
	h := handler.NewHandler()

	for _, target := range badURLs {
		req := httptest.NewRequest("GET", target, newEmptyReader())
		resp := httptest.NewRecorder()

		h.ServeHTTP(resp, req)

		if resp.Code == http.StatusOK {
			t.Error("Succesful GET request on bad url:", target)
		}
	}
}

// Basic test ensuring that interfaces list request
// finishes succesfully (or not succesfully, if
// net.Interfaces() call on current system returns a error)
func TestInterfacesListResponse(t *testing.T) {
	h := handler.NewHandler()

	req := httptest.NewRequest("GET", ifacesURL, newEmptyReader())
	resp := httptest.NewRecorder()

	h.ServeHTTP(resp, req)

	_, err := net.Interfaces()
	if err == nil {
		if resp.Code != http.StatusOK {
			t.Error("GET request on", ifacesURL, "failed")
		}
	} else {
		if resp.Code == http.StatusOK {
			t.Error("GET request on", ifacesURL, ": server has not reported a error:", err.Error())
		}
	}
}
