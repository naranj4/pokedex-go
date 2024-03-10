package repl

import (
	"os"
	"strings"
	"testing"
	"testing/iotest"

	"github.com/sanka047/pokedex-go/cmd"
	"github.com/sanka047/pokedex-go/testing/assert"
)

func TestReplHappyPath(t *testing.T) {
	in, err := os.Open("testdata/repl_input.txt")
	assert.Ok(t, err)

	var out strings.Builder

	r, err := NewRepl([]cmd.Cmd{cmd.Exit{}}, in, &out)
	assert.Ok(t, err)

	assert.False(t, r.IsActive(), "REPL did not start in the 'inactive' state.")

	r.Start()
	assert.True(t, r.IsActive(), "REPL should transition to an 'active' state")

	for r.IsActive() {
		err = r.Next()
		assert.Ok(t, err)
	}

	exp_out, err := os.Open("testdata/repl_output.txt")
	assert.Ok(t, err)

	err = iotest.TestReader(exp_out, []byte(out.String()))
	assert.Ok(t, err)
}

func TestReplNextOnInactive(t *testing.T) {
	in := strings.NewReader("help")
	var out strings.Builder

	r, err := NewRepl([]cmd.Cmd{}, in, &out)
	assert.Ok(t, err)

	assert.False(t, r.IsActive(), "REPL did not start in the 'inactive' state.")

	err = r.Next()
	assert.Err(t, err)
}
