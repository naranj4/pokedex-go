package cmd

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/sanka047/pokedex-go/pkmn"
	"github.com/sanka047/pokedex-go/testing/assert"
)

type StubPokeAPI struct{}

func (s StubPokeAPI) GetPokemon(
    ctx context.Context,
    vg_id pkmn.Identifier,
    pk_id pkmn.Identifier,
) (pkmn.Pokemon, error) {
	return pkmn.Pokemon{}, nil
}

func TestPokemonLookupHappyPath(t *testing.T) {
	pl := NewPokemonLookup(StubPokeAPI{})

	res, err := pl.Cmd([]string{"fake-vg", "fake-pokemon"})
	assert.Ok(t, err)

	mesg, err := json.MarshalIndent(pkmn.Pokemon{}, "", "  ")
	assert.Ok(t, err)

	exp := Result{Mesg: string(mesg)}
	assert.Equals(t, res, exp)
}
