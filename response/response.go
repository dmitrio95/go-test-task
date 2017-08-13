package response

import (
	"net"
)

// Contains MIME types of available ResponseFormatters
var Formats = map[string]bool{
	"text/html": true,
}

type ResponseEntry struct {
	Link string
	Name string
	Data map[string]string
}

type ResponseData struct {
	Title   string
	Entries []ResponseEntry
}

type ResponseFormatter interface {
	FormatResponse(data ResponseData) (string, error)
}

// Returns new ResponseFormatter by its MIME type. If such
// a formatter does not exist, returns nil.
func NewResponseFormatter(mimetype string) ResponseFormatter {
	switch mimetype {
	case "text/html":
		return newHTMLFormatter()
	default:
		return nil
	}
}

// Converts a slice of network interfaces to the
// representation that can be passed to the chosen
// ResponseFormatter.
func NewResponseData(ifaces []net.Interface) ResponseData {
	entries := make([]ResponseEntry, len(ifaces))

	for i, iface := range ifaces {
		entries[i] = ResponseEntry{iface.Name, iface.Name, map[string]string{"key1": "value1", "key2": "value2"}}
	}

	var title string
	if len(ifaces) == 1 {
		title = ifaces[0].Name
	} else {
		title = "Interfaces"
	}

	return ResponseData{title, entries}
}
