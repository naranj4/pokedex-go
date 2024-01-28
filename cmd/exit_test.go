package cmd

import "testing"

func TestExitHappyPath(t *testing.T) {
	e := NewExit()
	res, err := e.Cmd([]string{})
	if err != nil {
		t.FailNow()
	}

	exp := Result{IsTerminal: true}
	if res != exp {
		t.Fatalf("Expected:\n\"\"\"\n%v\n\"\"\"\nReceived:\n\"\"\"\n%v\n\"\"\"", exp, res)
	}
}

func TestExitExtraArgs(t *testing.T) {
	e := NewExit()
	_, err := e.Cmd([]string{"a", "b"})
	if err == nil {
		t.FailNow()
	}
}
