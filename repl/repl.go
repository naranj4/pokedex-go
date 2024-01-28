package repl

import (
	"bufio"
	"errors"
	"fmt"
	"io"

	"github.com/sanka047/pokedex-go/cmd"
)

const CLI_NAME = "pokedex"

type Repl struct {
	cmds    map[string]cmd.Cmd
	scanner *bufio.Scanner
	w       *bufio.Writer
	active bool
}

func NewRepl(cmds []cmd.Cmd, r io.Reader, w io.Writer) *Repl {
	return &Repl{
		cmds:    buildCommandMap(cmds),
		scanner: bufio.NewScanner(r),
		w:       bufio.NewWriter(w),
		active: false, // must call Start to start using the REPL
	}
}

func (r *Repl) Start() {
	r.active = true
	r.println("Welcome to Pokedex-CLI!")
}

func (r *Repl) IsActive() bool {
	return r.active
}

func (r *Repl) Next() error {
	if !r.active {
		return errors.New("Must call 'Start' before attempting to use the REPL")
	}

	r.br()
	r.prompt()
	in, err := r.readInput()
	if err != nil {
		return nil
	}

	cmd, exists := r.cmds[in.CmdName]
	if !exists {
		r.printf(
			"%s is not a valid command. Try 'help' for a list of valid commands.\n",
			in.CmdName,
		)
		return nil
	}

	res, err := cmd.Cmd(in.Args)
	if err != nil {
		r.println(err)
		return nil
	}

	r.println(res.Mesg)
	if res.IsTerminal {
		r.active = false
	}
	return nil
}

func (r *Repl) readInput() (cmd.Input, error) {
	r.scanner.Scan()
	return cmd.NewInput(r.scanner.Text())
}

func (r *Repl) prompt() {
	r.print(CLI_NAME, " > ")
}

func (r *Repl) br() {
	r.println()
}

func (r *Repl) printf(format string, a ...any) (int, error) {
	n, err := r.w.WriteString(fmt.Sprintf(format, a...))
	if err != nil {
		return n, err
	}
	return n, errors.Join(err, r.w.Flush())
}

func (r *Repl) println(a ...any) (int, error) {
	n, err := r.w.WriteString(fmt.Sprintln(a...))
	if err != nil {
		return n, err
	}
	return n, errors.Join(err, r.w.Flush())
}

func (r *Repl) print(a ...any) (int, error) {
	n, err := r.w.WriteString(fmt.Sprint(a...))
	if err != nil {
		return n, err
	}
	return n, errors.Join(err, r.w.Flush())
}

func buildCommandMap(cmds []cmd.Cmd) map[string]cmd.Cmd {
	cmd_map := make(map[string]cmd.Cmd)
	for _, c := range cmds {
		cmd_map[c.Name()] = c
	}

	// always add the 'help' command to document options
	h := cmd.NewHelp(cmds)
	cmd_map[h.Name()] = h

	return cmd_map
}
