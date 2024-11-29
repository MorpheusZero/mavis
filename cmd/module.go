package cmd

import "fmt"

type Command interface {
	Run(args []string) error
}

type CommandModule struct {
	commands map[string]Command
}

func NewCommandModule() *CommandModule {
	return &CommandModule{
		commands: make(map[string]Command),
	}
}

func (m *CommandModule) Register(name string, command Command) {
	m.commands[name] = command
}

func (m *CommandModule) Get(name string) Command {
	return m.commands[name]
}

func (m *CommandModule) Run(name string, args []string) error {
	command := m.Get(name)
	if command == nil {
		return fmt.Errorf("command %s not found", name)
	}
	return command.Run(args)
}

func (m *CommandModule) List() []string {
	var names []string
	for name := range m.commands {
		names = append(names, name)
	}
	return names
}

func CommandHandler(args []string) error {
	module := NewCommandModule()
	module.Register("hello", &HelloCommand{})
	module.Register("server", &ServerCommand{})

	if len(args) < 2 {
		return fmt.Errorf("usage: %s <command> [args...]", args[0])
	} else {
		command := args[1]
		args = args[2:]
		err := module.Run(command, args)
		if err != nil {
			return err
		}
	}
	return nil
}
