package cmd

type HelloCommand struct{}

func (c *HelloCommand) Run(args []string) error {
	println("Hello, World!")
	return nil
}
