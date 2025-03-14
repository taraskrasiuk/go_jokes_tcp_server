package joke

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"golang.org/x/exp/rand"
)

// A simple store for jokes.
type JokesStore struct {
	jokes []jokeEntry
}

func NewJokesStore() *JokesStore {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	file, err := os.OpenFile(fmt.Sprintf("%s/jokes/jokes.json", wd), os.O_RDONLY, 04) // only read
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
	return &JokesStore{
		jokes,
	}
}

func (j JokesStore) GetRandomJokeParts() []string {
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
