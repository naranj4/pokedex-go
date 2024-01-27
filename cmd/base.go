package cmd

type Cmd interface {
	Name() string
	Doc() string
	Cmd(args []string) (Result, error)
}

type Result struct {
	Mesg string
	IsTerminal bool
}
