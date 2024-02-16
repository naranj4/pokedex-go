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
	active  bool
}

func NewRepl(cmds []cmd.Cmd, r io.Reader, w io.Writer) (*Repl, error) {
	cmd_map, err := buildCommandMap(cmds)
	if err != nil {
		return nil, err
	}

	return &Repl{
		cmds:    cmd_map,
		scanner: bufio.NewScanner(r),
		w:       bufio.NewWriter(w),
		active:  false, // must call Start to start using the REPL
	}, nil
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

func insertWithoutReplacement(cmd_map *map[string]cmd.Cmd, alias string, c cmd.Cmd) error {
	if existing, ok := (*cmd_map)[alias]; ok {
		return fmt.Errorf(
			"Alias '%s' has already been registered. Existing:\n\"\"\"\n%v\n\"\"\"\nNew:\n\"\"\"\n%v\n\"\"\"",
			alias,
			existing,
			c,
		)
	}

	(*cmd_map)[alias] = c
	return nil
}

func buildCommandMap(cmds []cmd.Cmd) (map[string]cmd.Cmd, error) {
	// always add the 'help' command to document options
	cmds = append(cmds, cmd.NewHelp(cmds))

	cmd_map := make(map[string]cmd.Cmd)
	for _, c := range cmds {
		err := insertWithoutReplacement(&cmd_map, c.Name(), c)
		if err != nil {
			return nil, err
		}

		for _, a := range c.Aliases() {
			err := insertWithoutReplacement(&cmd_map, a, c)
			if err != nil {
				return nil, err
			}
		}
	}

	return cmd_map, nil
}
