package joke

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

type jokesStore interface {
	GetRandomJokeParts() []string
}

// Connection handler
func handleConn(conn net.Conn, jokeStore jokesStore) {
	defer conn.Close() // close connection

	// show the jokes forever
	for {
		randJokeInParts := jokeStore.GetRandomJokeParts()
		time.Sleep(2 * time.Second)

		for i, jokePart := range randJokeInParts {
			respTime := time.Now().Format(time.RFC1123)
			time.Sleep(2 * time.Second)

			var out string
			if i == 0 {
				out += fmt.Sprintf("|---- %s", respTime)
			}
			out += fmt.Sprintf("\n\t - %s", jokePart)
			if i == len(randJokeInParts)-1 {
				out += "\n-----\n"
			}

			conn.Write([]byte(out))
		}
	}

}

type TCPServerOpts struct {
	Host       string
	Port       string
	JokesStore jokesStore
}

func NewTCPServerOpts(jokesStore jokesStore, host, port string) *TCPServerOpts {
	return &TCPServerOpts{
		Host:       host,
		Port:       port,
		JokesStore: jokesStore,
	}
}

func RunTCPServer(opts TCPServerOpts) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", opts.Host, opts.Port))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Printf("TCP server is running on host %s and port %s \n", opts.Host, opts.Port)

	// listen for connections in infinite loop
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue // ignore the connection
		}
		fmt.Printf("Handled connection: %s", conn.LocalAddr().String())
		go handleConn(conn, opts.JokesStore)
	}
}
