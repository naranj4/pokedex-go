package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/sanka047/pokedex-go/cmd"
)

const CLI_NAME = "pokedex"

func prompt() {
	fmt.Print(CLI_NAME, " > ")
}

func br() {
	fmt.Println()
}

func getCommands() []cmd.Cmd {
	cmds := []cmd.Cmd{
		cmd.NewExit(),
	}
	cmds = append(cmds, cmd.NewHelp(cmds))
	return cmds
}

func buildCommandMap(cmds []cmd.Cmd) map[string]cmd.Cmd {
	cmd_map := make(map[string]cmd.Cmd)
	for _, c := range cmds {
		cmd_map[c.Name()] = c
	}
	return cmd_map
}

type Input struct {
	CmdName string
	Args    []string
}

func cleanInput(text string) (Input, error) {
	words := strings.Fields(text)
	if len(words) == 0 {
		return Input{}, errors.New("No command")
	}
	return Input{CmdName: words[0], Args: words[1:]}, nil
}

func main() {
	fmt.Println("Welcome to Pokedex-CLI!")

	cmds := buildCommandMap(getCommands())
	status_code := 0
	scanner := bufio.NewScanner(os.Stdin)
	for {
		// main loop
		br()
		prompt()
		scanner.Scan()

		in, err := cleanInput(scanner.Text())
		if err != nil {
			continue
		}

		// Register commands in a registry or something instead of this
		if cmd, exists := cmds[in.CmdName]; exists {
			res, err := cmd.Cmd(in.Args)
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println(res.Mesg)
			if res.IsTerminal {
				os.Exit(status_code)
			}
		} else {
			fmt.Printf(
				"%s is not a valid command. Try 'help' for a list of valid commands.\n",
				in.CmdName,
			)
		}
	}
}
