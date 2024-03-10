package cmd

import (
	"testing"

	"github.com/sanka047/pokedex-go/testing/assert"
)

func TestExitHappyPath(t *testing.T) {
	e := NewExit()
	res, err := e.Cmd([]string{})
	assert.Ok(t, err)

	exp := Result{IsTerminal: true}
	assert.Equals(t, res, exp)
}

func TestExitExtraArgs(t *testing.T) {
	e := NewExit()
	_, err := e.Cmd([]string{"a", "b"})
	assert.Err(t, err)
}
