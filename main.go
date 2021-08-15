package main

import (
	"os"

	"github.com/LeoNdV001/brainfuck/src/cli"
)

// main bootstraps the application.
func main() {
	cli.Run(os.Args)
}
