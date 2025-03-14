package main

import (
	"fmt"
	"os"
	"strings"
	"taraskrasiuk/go_jokes_tcp_server/joke"
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

	fmt.Println(args)
	for _, arg := range args {
		argSpl := strings.Split(arg, "=")

		fmt.Println(argSpl)

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

	joke.RunTCPServer(cfg.host, cfg.port)
}
