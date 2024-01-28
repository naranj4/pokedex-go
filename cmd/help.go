package cmd

import (
	"errors"
	"fmt"
	"strings"
)

type Help struct {
	cmds []Cmd

	// internal: cached value for the help message
	mesg *string
}

func NewHelp(cmds []Cmd) *Help {
	return &Help{cmds: cmds}
}

func (c *Help) Name() string {
	return "help"
}

func (c *Help) Doc() string {
	return "Print this help message."
}

func (c *Help) Cmd(args []string) (Result, error) {
	if len(args) != 0 {
		return Result{}, errors.New(fmt.Sprintf("%s: does not accept any arguments", c.Name()))
	}

	if c.mesg == nil {
		mesg := c.generateHelpMessage()
		c.mesg = &mesg
	}

	return Result{Mesg: *c.mesg}, nil
}

func (c *Help) generateHelpMessage() string {
	var b strings.Builder
	b.WriteString("The Pokedex-CLI supports the following commands:\n")
	for _, cmd := range c.cmds {
		b.WriteString(document(cmd))
	}

	// make sure to include documentation for itself
	b.WriteString(document(c, ""))

	return b.String()
}

func document(cmd Cmd, args ...string) string {
	term := "\n"
	if len(args) > 0 {
		term = args[0]
	}
	return fmt.Sprintf("- `%s`: %s%s", cmd.Name(), cmd.Doc(), term)
}
