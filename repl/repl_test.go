package repl

import (
	"os"
	"strings"
	"testing"
	"testing/iotest"

	"github.com/sanka047/pokedex-go/cmd"
)

func TestReplHappyPath(t *testing.T) {
	in, err := os.Open("testdata/repl_input.txt")
	if err != nil {
		t.Fatal("Unable to open input fixture file.")
	}

	var out strings.Builder

	r, err := NewRepl([]cmd.Cmd{cmd.Exit{}}, in, &out)
	if err != nil {
		t.Fatal(err)
	}

	if r.IsActive() {
		t.Fatal("REPL did not start in the 'inactive' state.")
	}

	r.Start()
	if !r.IsActive() {
		t.Fatal("REPL should transition to an 'active' state")
	}

	for r.IsActive() {
		err = r.Next()
		if err != nil {
			t.Fatal(err)
		}
	}

	exp_out, err := os.Open("testdata/repl_output.txt")
	if err != nil {
		t.Fatal("Unable to open output fixture file.")
	}

	err = iotest.TestReader(exp_out, []byte(out.String()))
	if err != nil {
		t.Fatal(err)
	}
}

func TestReplNextOnInactive(t *testing.T) {
	in := strings.NewReader("help")
	var out strings.Builder

	r, err := NewRepl([]cmd.Cmd{}, in, &out)
	if err != nil {
		t.Fatal(err)
	}

	if r.IsActive() {
		t.Fatal("REPL did not start in the 'inactive' state.")
	}

	err = r.Next()
	if err == nil {
		t.Fatal("Expected error for calling Next without Start.")
	}
}
