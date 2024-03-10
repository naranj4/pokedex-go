package cmd

import (
	"os"
	"strings"
	"testing"

	"github.com/sanka047/pokedex-go/testing/assert"
)

type StubCmd1 struct{}

func (s StubCmd1) Name() string                      { return "stub1" }
func (s StubCmd1) Aliases() []string                 { return []string{"s1"} }
func (s StubCmd1) Doc() string                       { return "Doc1" }
func (s StubCmd1) Cmd(args []string) (Result, error) { return Result{}, nil }

type StubCmd2 struct{}

func (s StubCmd2) Name() string                      { return "stub2" }
func (s StubCmd2) Aliases() []string                 { return []string{"s2"} }
func (s StubCmd2) Doc() string                       { return "Doc2" }
func (s StubCmd2) Cmd(args []string) (Result, error) { return Result{}, nil }

func TestHelpHappyPath(t *testing.T) {
	cmds := []Cmd{
		StubCmd1{},
		StubCmd2{},
	}
	h := NewHelp(cmds)

	assert.Equals(t, h.mesg, nil)

	dat, err := os.ReadFile("testdata/help-message.txt")
	assert.Ok(t, err)

	res, err := h.Cmd([]string{})
	assert.Ok(t, err)

	exp := Result{Mesg: strings.Trim(string(dat), "\n")}
	assert.Equals(t, res, exp)
	assert.NotEquals(t, h.mesg, nil)
}

func TestHelpExtraArgs(t *testing.T) {
	h := NewHelp([]Cmd{})
	_, err := h.Cmd([]string{"a", "b"})
	assert.Err(t, err)
}
