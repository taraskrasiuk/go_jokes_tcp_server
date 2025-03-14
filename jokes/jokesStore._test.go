package joke

import "testing"

func TestJokesStore(t *testing.T) {
	jokesStore := NewJokesStore()

	// it should return a random joke, with 2 parts.
	randJoke := jokesStore.GetRandomJokeParts()

	if len(randJoke) != 2 {
		t.Errorf("the random joke should contain the 2 strings.")
	}
}
