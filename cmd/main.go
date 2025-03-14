package main

import (
	"os"
	"strings"
	jokes "taraskrasiuk/go_jokes_tcp_server/jokes"
)

// Default:
// host = localhost
// port = 8080
type config struct {
	host string
	port string
}

// Just use the arguments.
func NewConfig() *config {
	args := os.Args[1:]

	var host, port string = "localhost", "8080"

	for _, arg := range args {
		argSpl := strings.Split(arg, "=")

		if argSpl[0] == "--host" {
			host = argSpl[1]
		}

		if argSpl[0] == "--port" {
			port = argSpl[1]
		}
	}

	return &config{host, port}
}

func main() {
	cfg := NewConfig()

	jokesStore := jokes.NewJokesStore()
	serverOpts := jokes.NewTCPServerOpts(jokesStore, cfg.host, cfg.port)

	jokes.RunTCPServer(*serverOpts)
}
