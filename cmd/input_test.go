package cmd

import (
	"testing"

	"github.com/sanka047/pokedex-go/testing/assert"
)

func TestInputHappyPath(t *testing.T) {
	in, err := NewInput("command arg1 arg2")
	assert.Ok(t, err)

	exp := Input{CmdName: "command", Args: []string{"arg1", "arg2"}}
	assert.DeepEquals(t, in, exp)
}

func TestInputNoArgs(t *testing.T) {
	in, err := NewInput("command")
	assert.Ok(t, err)

	exp := Input{CmdName: "command", Args: []string{}}
	assert.DeepEquals(t, in, exp)
}

func TestInputNoCommand(t *testing.T) {
	_, err := NewInput("")
	assert.Err(t, err)
}
