package main

import (
	"github.com/morpheuszero/mavis/cmd"
	"os"
)

func main() {
	err := cmd.CommandHandler(os.Args)
	if err != nil {
		panic(err)
	}
}
