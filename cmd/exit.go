package cmd

import (
	"errors"
	"fmt"
)

type Exit struct{}

func NewExit() Exit {
	return Exit{}
}

func (e Exit) Name() string {
	return "exit"
}

func (e Exit) Aliases() []string {
	return []string{"quit"}
}

func (e Exit) Doc() string {
	return "Quit the pokedex REPL."
}

func (e Exit) Cmd(args []string) (Result, error) {
	if len(args) != 0 {
		return Result{}, errors.New(fmt.Sprintf("%s: does not accept any arguments", e.Name()))
	}

	return Result{IsTerminal: true}, nil
}
