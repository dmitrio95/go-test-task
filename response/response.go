package response

import (
	"fmt"
	"net"
)

// Contains MIME types of available ResponseFormatters
var (
	Formats = map[string](func() ResponseFormatter){
		"text/html":  newHTMLFormatter,
		"text/plain": newSimpleFormatter,
	}
	DefaultFormat = "text/html"
)

type ResponseEntryProperty struct {
	Name  string
	Value string
}

type ResponseEntry struct {
	Link string
	Name string
	Data []ResponseEntryProperty
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
	f, ok := Formats[mimetype]
	if ok {
		return f()
	} else {
		return nil
	}
}

func getInterfaceData(iface net.Interface) []ResponseEntryProperty {
	data := []ResponseEntryProperty{
		{"Index", fmt.Sprint(iface.Index)},
		{"Name", iface.Name},
		{"MTU", fmt.Sprint(iface.MTU)},
		{"Hardware Address", iface.HardwareAddr.String()},
		{"Flags", iface.Flags.String()},
	}

	addrs, err := iface.Addrs()
	if err != nil {
		data = append(data, ResponseEntryProperty{"Error on retrieving unicast addresses", err.Error()})
	} else {
		for _, a := range addrs {
			data = append(data, ResponseEntryProperty{"Network", a.Network()})
			data = append(data, ResponseEntryProperty{"Address", a.String()})
		}
	}

	addrs, err = iface.MulticastAddrs()
	if err != nil {
		data = append(data, ResponseEntryProperty{"Error on retrieving multicast addresses", err.Error()})
	} else {
		for _, a := range addrs {
			data = append(data, ResponseEntryProperty{"Network", a.Network()})
			data = append(data, ResponseEntryProperty{"Multicast Address", a.String()})
		}
	}

	return data
}

// Converts a slice of network interfaces to the
// representation that can be passed to the chosen
// ResponseFormatter.
func NewResponseData(ifaces []net.Interface) ResponseData {
	entries := make([]ResponseEntry, len(ifaces))

	for i, iface := range ifaces {
		entries[i] = ResponseEntry{iface.Name, iface.Name, getInterfaceData(iface)}
	}

	var title string
	if len(ifaces) == 1 {
		title = ifaces[0].Name
	} else {
		title = "Interfaces"
	}

	return ResponseData{title, entries}
}
