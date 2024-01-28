package cmd

import (
	"errors"
	"strings"
)

type Input struct {
	CmdName string
	Args    []string
}

func NewInput(text string) (Input, error) {
	words := strings.Fields(text)
	if len(words) == 0 {
		return Input{}, errors.New("No command")
	}
	return Input{CmdName: words[0], Args: words[1:]}, nil
}
