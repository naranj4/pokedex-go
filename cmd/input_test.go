package cmd

import (
	"reflect"
	"testing"
)

func TestInputHappyPath(t *testing.T) {
	in, err := NewInput("command arg1 arg2")
	if err != nil {
		t.FailNow()
	}

	exp := Input{CmdName: "command", Args: []string{"arg1", "arg2"}}
	if !reflect.DeepEqual(in, exp) {
		t.Fatalf("Expected:\n\"\"\"\n%v\n\"\"\"\nReceived:\n\"\"\"\n%v\n\"\"\"", exp, in)
	}
}

func TestInputNoArgs(t *testing.T) {
	in, err := NewInput("command")
	if err != nil {
		t.FailNow()
	}

	exp := Input{CmdName: "command", Args: []string{}}
	if !reflect.DeepEqual(in, exp) {
		t.Fatalf("Expected:\n\"\"\"\n%v\n\"\"\"\nReceived:\n\"\"\"\n%v\n\"\"\"", exp, in)
	}
}

func TestInputNoCommand(t *testing.T) {
	_, err := NewInput("")
	if err == nil {
		t.FailNow()
	}
}
