package main

import (
	"net/http"
	"os"

	"github.com/sanka047/pokedex-go/cmd"
	"github.com/sanka047/pokedex-go/pkmn/pokeapi"
	"github.com/sanka047/pokedex-go/repl"
)

func getCommands(pokeAPI *pokeapi.PokeAPI) []cmd.Cmd {
	cmds := []cmd.Cmd{
		cmd.NewExit(),
		cmd.NewPokemonLookup(pokeAPI),
		// TODO:
		// - regions
		// - region <name/id> (sets context)
		// - locations (within region)
		// - location <name/id> (sets context)
		// - areas (within location)
		// - area <name/id> (sets context)
		// - pokemon <name/id> (looks up pokemon)
	}
	return cmds
}

func main() {
	pk_api := pokeapi.NewPokeAPI(http.DefaultClient)
	r, err := repl.NewRepl(getCommands(pk_api), os.Stdin, os.Stdout)
	if err != nil {
		panic(err)
	}

	r.Start()
	for r.IsActive() {
		r.Next()
	}
}
