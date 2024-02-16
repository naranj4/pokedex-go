package main

import (
	"os"

	"github.com/sanka047/pokedex-go/cmd"
	"github.com/sanka047/pokedex-go/repl"
)

func getCommands() []cmd.Cmd {
	cmds := []cmd.Cmd{
		cmd.NewExit(),
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
	r, err := repl.NewRepl(getCommands(), os.Stdin, os.Stdout)
	if err != nil {
		panic(err)
	}

	r.Start()
	for r.IsActive() {
		r.Next()
	}
}
