package cmd

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/sanka047/pokedex-go/pkmn/pokeapi"
	"github.com/sanka047/pokedex-go/testing/assert"
)

type StubPokeAPI struct {}

func (s StubPokeAPI) GetPokemon(ctx context.Context, identifier string) (pokeapi.Pokemon, error) {
	return pokeapi.Pokemon{}, nil
}

func TestPokemonLookupHappyPath(t *testing.T) {
	pl := NewPokemonLookup(StubPokeAPI{})

	res, err := pl.Cmd([]string{"fake-pokemon"})
	assert.Ok(t, err)

	mesg, err := json.MarshalIndent(pokeapi.Pokemon{}, "", "  ")
	assert.Ok(t, err)

	exp := Result{Mesg: string(mesg)}
	assert.Equals(t, res, exp)
}
