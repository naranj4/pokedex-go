package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/sanka047/pokedex-go/pkmn/pokeapi"
)

type PokemonLookup struct {
	pokeAPI pokeapi.IPokeAPI
}

func NewPokemonLookup(p pokeapi.IPokeAPI) *PokemonLookup {
	return &PokemonLookup{pokeAPI: p}
}

func (l *PokemonLookup) Name() string {
	return "pokemon"
}

func (l *PokemonLookup) Aliases() []string {
	return []string{"pk"}
}

func (l *PokemonLookup) Doc() string {
	return "Lookup a Pokemon by name."
}

func (l *PokemonLookup) Cmd(args []string) (Result, error) {
	// TODO: Format this in a more sane way
	if len(args) != 1 {
		return Result{}, errors.New(fmt.Sprintf("%s: requires exactly 1 argument (name)", l.Name()))
	}

	pk, err := l.pokeAPI.GetPokemon(context.Background(), pokeapi.Name(args[0]))
	if err != nil {
		return Result{}, err
	}

	raw, err := json.MarshalIndent(pk, "", "  ")
	if err != nil {
		return Result{}, err
	}

	return Result{Mesg: string(raw)}, nil
}
