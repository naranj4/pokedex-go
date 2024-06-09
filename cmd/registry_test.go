package cmd

import (
	"testing"

	"github.com/sanka047/pokedex-go/testing/assert"
)

func TestOne(t *testing.T) {
	reg := NewRegistry()
	// registering works in general (with aliases)
	reg.Register(
		"commandone",
		[]string{"cmdone", "one"},
		"imagine using vim",
		func(args []string) (Result, error) { return Result{}, nil },
	)
	// registering works in general (without aliases)
	reg.Register(
		"commandtwo",
		[]string{},
		"imagine using nano",
		func(args []string) (Result, error) { return Result{}, nil },
	)
	// registering fails on alias collision
	assert.Panic(t, func() {
		reg.Register(
			"commandwon",
			[]string{"cmdone", "won", "one"},
			"imagine using nvim",
		    func(args []string) (Result, error) { return Result{}, nil },
		)
	})
	// registering fails on name collision
	assert.Panic(t, func() {
		reg.Register(
			"commandone",
			[]string{},
			"imagine using neovim",
		    func(args []string) (Result, error) { return Result{}, nil },
		)
	})
}
