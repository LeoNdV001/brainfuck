package cli

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/urfave/cli"

	"github.com/LeoNdV001/brainfuck/src/interpreter"
)

// registerCLICommands registers available commands
func (bf BrainfuckCLI) registerCLICommands() {
	cmd := cli.Command{
		Name:  "compiler",
		Usage: "Run and check for Brainfuck CLI",
		Subcommands: []cli.Command{
			bf.InfoCommand(),
			bf.ParseCommand(),
		},
	}

	bf.cli.Commands = append(bf.cli.Commands, cmd)
}

// InfoCommand returns information about the brainfuck package
func (bf BrainfuckCLI) InfoCommand() cli.Command {
	return cli.Command{
		Name:  "info",
		Usage: "Shows information about the bf package",
		Action: func(context *cli.Context) error {
			return nil
		},
	}
}

// ParseCommand parses a Brainfuck file and returns the parsed value
func (bf BrainfuckCLI) ParseCommand() cli.Command {
	return cli.Command{
		Name:  "parse",
		Usage: "Supply a file including the path",
		Action: func(context *cli.Context) error {
			file := context.Args().First()
			bf.run(context, handleParseCommand(file))

			return nil
		},
	}
}

// handleParseCommand handles file parsing
func handleParseCommand(file string) *exec.Cmd {
	bfCode, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Can't read the file. %s\n", err.Error())
		os.Exit(1)
	}

	interprtr := interpreter.NewInterpreter()

	err = interprtr.Parse(string(bfCode))
	if err != nil {
		fmt.Printf("Error parsing the file. %s\n", err.Error())
		os.Exit(1)
	}
	
	return exec.Command("echo", interprtr.ParsedCommands)
}