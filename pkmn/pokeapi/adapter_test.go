package pokeapi

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/sanka047/pokedex-go/testing/assert"
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
		assert.Equals(t, req.URL.String(), fmt.Sprintf("%s/pokemon/%s", BASE_API_PATH, Name(name)))
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString("{}")),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})

	api := NewPokeAPI(client)
	_, err := api.GetPokemon(context.Background(), Name(name))
	assert.Ok(t, err)
}

func TestGetPokemonContextDone(t *testing.T) {
	name := "fake-pokemon"
	client := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equals(t, req.URL.String(), fmt.Sprintf("%s/pokemon/%s", BASE_API_PATH, Name(name)))
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
	assert.NotEquals(t, err, nil)
}
