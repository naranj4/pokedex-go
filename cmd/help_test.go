package cmd

import (
	"os"
	"strings"
	"testing"
)

type StubCmd1 struct {}
func (s StubCmd1) Name() string { return "stub1" }
func (s StubCmd1) Aliases() []string { return []string{"s1"} }
func (s StubCmd1) Doc() string { return "Doc1" }
func (s StubCmd1) Cmd(args []string) (Result, error) { return Result{}, nil }

type StubCmd2 struct {}
func (s StubCmd2) Name() string { return "stub2" }
func (s StubCmd2) Aliases() []string { return []string{"s2"} }
func (s StubCmd2) Doc() string { return "Doc2" }
func (s StubCmd2) Cmd(args []string) (Result, error) { return Result{}, nil }

func TestHelpHappyPath(t *testing.T) {
	cmds := []Cmd{
		StubCmd1{},
		StubCmd2{},
	}
	h := NewHelp(cmds)

	if h.mesg != nil {
		t.Fatal("There is an existing cached value in h.mesg")
	}

	dat, err := os.ReadFile("testdata/help-message.txt")
	if err != nil {
		t.Fatal("Failed to read data from fixture: ", err)
	}

	res, err := h.Cmd([]string{})
	if err != nil {
		t.Fatal(err)
	}

	exp := Result{Mesg: strings.Trim(string(dat), "\n")}
	if res != exp {
		t.Fatalf("Expected:\n\"\"\"\n%v\n\"\"\"\nReceived:\n\"\"\"\n%v\n\"\"\"", exp, res)
	}

	if h.mesg == nil {
		t.Fatal("Generated doc was not cached.")
	}
}

func TestHelpExtraArgs(t *testing.T) {
	h := NewHelp([]Cmd{})
	_, err := h.Cmd([]string{"a", "b"})
	if err == nil {
		t.FailNow()
	}
}
