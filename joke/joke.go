package joke

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"time"
)

// A simple store for jokes.
type JokeStore struct {
	jokes []jokeEntry
}

func NewJokeStore() *JokeStore {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	file, err := os.OpenFile(fmt.Sprintf("%s/joke/jokes.json", wd), os.O_RDONLY, 04) // only read
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	// Close the file.
	defer file.Close()

	jsonData, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	var jokes []jokeEntry
	if err := json.Unmarshal(jsonData, &jokes); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return &JokeStore{
		jokes,
	}
}

func (j JokeStore) GetRandomJokeParts() []string {
	time.Sleep(2 * time.Second)
	randIdx := rand.Intn(len(j.jokes))
	jokeEntry := j.jokes[randIdx]
	parts := []string{jokeEntry.Part1, jokeEntry.Part2}
	return parts
}

// A joke entry. Right now only the 'part1' and 'part2' fields are needed for unmarshalling.
type jokeEntry struct {
	link   string
	score  int
	Part1  string `json:"part1"`
	mature bool
	author string
	Part2  string `json:"part2"`
}

func (j *jokeEntry) getFullJoke() string {
	return fmt.Sprintf("%s\n%s\n", j.Part1, j.Part2)
}

// Connection handler
func handleConn(conn net.Conn, jokeStore *JokeStore) {
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

func RunTCPServer(host, port string) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Printf("TCP server is running on host %s and port %s \n", host, port)

	jokeStore := NewJokeStore()

	// listen for connections in infinite loop
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue // ignore the connection
		}
		fmt.Printf("Handled connection: %s", conn.LocalAddr().String())
		go handleConn(conn, jokeStore)
	}
}
