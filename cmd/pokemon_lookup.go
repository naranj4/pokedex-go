package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/sanka047/pokedex-go/pkmn"
)

type PokemonLookup struct {
	service pkmn.IPokeService
}

func NewPokemonLookup(s pkmn.IPokeService) *PokemonLookup {
	return &PokemonLookup{service: s}
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
	if len(args) != 2 {
		return Result{}, errors.New(
			fmt.Sprintf("%s: requires exactly 2 argument (version-group, name)", l.Name()),
		)
	}

	vg := pkmn.Name(args[0])
	name := pkmn.Name(args[1])
	pk, err := l.service.GetPokemon(context.Background(), vg, name)
	if err != nil {
		return Result{}, err
	}

	raw, err := json.MarshalIndent(pk, "", "  ")
	if err != nil {
		return Result{}, err
	}

	return Result{Mesg: string(raw)}, nil
}
