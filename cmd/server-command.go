package cmd

import "github.com/morpheuszero/mavis/core"

type ServerCommand struct{}

func (c *ServerCommand) Run(args []string) error {
	server := core.NewServer()
	return server.Run()
}
