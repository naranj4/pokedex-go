package cmd

import "fmt"

type CmdFunc func(args []string) (Result, error)

type Registry interface {
	Register(name string, aliases []string, doc string, cmd CmdFunc)
	GetCmd(key string) (CmdFunc, error)
	ListCommands() []RegistryEntry
}

type RegistryEntry struct {
	Name    string
	Aliases []string
	Doc     string
	Cmd     CmdFunc
}

type registry struct {
	cmds    []RegistryEntry
	cmd_map map[string]RegistryEntry
}

func NewRegistry() Registry {
	return &registry{
		cmds:    make([]RegistryEntry, 0),
		cmd_map: make(map[string]RegistryEntry),
	}
}

// TODO Registration may fail, and if so it will panic with the offending
// alias/name that resulted in a collision
func (r registry) Register(name string, aliases []string, doc string, cmd CmdFunc) {
	entry := RegistryEntry{Name: name, Aliases: aliases, Doc: doc, Cmd: cmd}
	r.cmds = append(r.cmds, entry)
	err := r.insert(name, entry)
	if err != nil {
		panic(err)
	}
	for _, alias := range aliases {
		err = r.insert(alias, entry)
		if err != nil {
			panic(err)
		}
	}
}

func (r registry) insert(key string, entry RegistryEntry) error {
	if existing, ok := r.cmd_map[key]; ok {
		return fmt.Errorf(
			"Alias '%s' has already been registered. Existing:\n\"\"\"\n%v\n\"\"\"\nNew:\n\"\"\"\n%v\n\"\"\"",
			key,
			existing,
			entry,
		)
	}

	r.cmd_map[key] = entry
	return nil
}

func (r registry) GetCmd(key string) (CmdFunc, error) {
	if existing, ok := r.cmd_map[key]; ok {
		return existing.Cmd, nil
	}
	return nil, fmt.Errorf("Command not found with name or alias matching '%s'", key)
}

func (r registry) ListCommands() []RegistryEntry {
	cmd_lst := make([]RegistryEntry, len(r.cmds))
	copy(cmd_lst, r.cmds)
	return cmd_lst
}
