package pokeapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const BASE_API_PATH = "https://pokeapi.co/api/v2"

func Name(name string) string {
	return name
}

func Id(id int) string {
	return fmt.Sprint(id)
}

type IPokeAPI interface {
	GetPokemon(ctx context.Context, identifier string) (pokemon, error)
	GetVersionGroup(ctx context.Context, identifier string) (versionGroup, error)
}

// TODO: wrap this in another class to:
//  1. translate internal PokeAPI types to something less verbose (automatically filtered down by
//     version)
//  2. cache the translated types from above into a sqlite database
type PokeAPI struct {
	client *http.Client
}

func NewPokeAPI(c *http.Client) *PokeAPI {
	return &PokeAPI{client: c}
}

type result[T any] struct {
	resp T
	err  error
}

// Fetch a pokemon using the given identifier. For example: p.GetPokemon(Name("ditto")) OR
// p.GetPokemon(Id(132))
func (p *PokeAPI) GetPokemon(ctx context.Context, identifier string) (pokemon, error) {
	var pk pokemon

	n_ctx, cancel := context.WithTimeout(ctx, time.Duration(2000)*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(
		n_ctx,
		"GET",
		fmt.Sprintf("%s/pokemon/%s", BASE_API_PATH, identifier),
		nil,
	)
	if err != nil {
		return pk, err
	}

	body, err := p.call(req)
	if err != nil {
		return pk, err
	}

	err = json.Unmarshal(body, &pk)
	return pk, err
}

func (p *PokeAPI) GetVersionGroup(ctx context.Context, identifier string) (versionGroup, error) {
	var vg versionGroup

	n_ctx, cancel := context.WithTimeout(ctx, time.Duration(2000)*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(
		n_ctx,
		"GET",
		fmt.Sprintf("%s/version-group/%s", BASE_API_PATH, identifier),
		nil,
	)
	if err != nil {
		return vg, err
	}

	body, err := p.call(req)
	if err != nil {
		return vg, err
	}

	err = json.Unmarshal(body, &vg)
	return vg, err
}

func (p *PokeAPI) call(req *http.Request) ([]byte, error) {
	var body []byte
	res, err := p.client.Do(req)
	if err != nil {
		return body, fmt.Errorf("(%v) Failed to make request: %w", req.URL, err)
	}
	body, err = io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return body, fmt.Errorf(
			"(%v) Did not receive 200 OK response: %v (%v)",
			req.URL,
			res.Status,
			string(body),
		)
	}
	if err != nil {
		return body, fmt.Errorf("(%v) Failed to read response body: %w", req.URL, err)
	}
	return body, nil
}
