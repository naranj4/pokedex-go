package cmd

import (
	"errors"
	"fmt"
)

type Exit struct {}

func NewExit() *Exit {
	return &Exit{}
}

func (c *Exit) Name() string {
	return "exit"
}

func (c *Exit) Doc() string {
	return "Quit the pokedex"
}

func (c *Exit) Cmd(args []string) (Result, error) {
	if len(args) != 0 {
		return Result{}, errors.New(fmt.Sprintf("%s: does not accept any arguments", c.Name()))
	}

	return Result{IsTerminal: true}, nil
}
