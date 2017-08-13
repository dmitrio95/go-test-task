package main

import "flag"

type Config struct {
	Address string
}

var config = Config{
	Address: ":8080",
}

// Updates a configuration info and returns
// an actual Config object.
func GetConfig() Config {
	flag.StringVar(&config.Address, "address", config.Address, "TCP network address for the server to listen on")

	flag.Parse()

	return config
}
