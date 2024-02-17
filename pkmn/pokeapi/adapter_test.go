package pokeapi

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(f RoundTripFunc) *http.Client {
	return &http.Client{Transport: f}
}

func TestGetPokemonHappyPath(t *testing.T) {
	name := "fake-pokemon"
	client := NewTestClient(func(req *http.Request) *http.Response {
		equals(t, req.URL.String(), fmt.Sprintf("%s/pokemon/%s", BASE_API_PATH, Name(name)))
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString("{}")),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})

	api := NewPokeAPI(client)
	_, err := api.GetPokemon(context.Background(), Name(name))
	ok(t, err)
}

func TestGetPokemonContextDone(t *testing.T) {
	name := "fake-pokemon"
	client := NewTestClient(func(req *http.Request) *http.Response {
		equals(t, req.URL.String(), fmt.Sprintf("%s/pokemon/%s", BASE_API_PATH, Name(name)))
		return &http.Response{
			StatusCode: http.StatusOK,
			// Will fail JSON parsing if read
			Body: io.NopCloser(bytes.NewBufferString("")),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})

	// cancel immediately to ensure the API call fails
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	api := NewPokeAPI(client)
	_, err := api.GetPokemon(ctx, Name(name))
	notEquals(t, err, nil)
}

func equals(t *testing.T, act, exp any) {
	if act != exp {
		t.Fatalf("Expected:\n\"\"\"\n%v\n\"\"\"\nReceived:\n\"\"\"\n%v\n\"\"\"", exp, act)
	}
}

func notEquals(t *testing.T, act, not_exp any) {
	if act == not_exp {
		t.Fatalf("Not expected:\n\"\"\"\n%v\n\"\"\"\nReceived:\n\"\"\"\n%v\n\"\"\"", not_exp, act)
	}
}

func ok(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
