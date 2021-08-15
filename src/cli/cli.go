package cli

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

// BrainfuckCLI struct container.
type BrainfuckCLI struct {
	cli *cli.App
}

// Run starts the Capilla CLI with an array of arguments.
func Run(arguments []string) {
	var bf BrainfuckCLI

	bf.cli = cli.NewApp()
	bf.setInfo()
	bf.registerCLICommands()

	err := bf.cli.Run(arguments)
	handleError(err)
}

// setInfo sets the CLI information that is shown with help
func (bf BrainfuckCLI) setInfo() {
	bf.cli.Name = "Brainfuck CLI"
	bf.cli.Usage = "This is a tool that is used to compile bf files"
	bf.cli.Version = "0.0.1"
}

// run tries to rub the command and logs the output
func (bf BrainfuckCLI) run(_ *cli.Context, command *exec.Cmd) {
	var (
		out    bytes.Buffer
		stderr bytes.Buffer
	)

	command.Stdout = &out
	command.Stderr = &stderr

	err := command.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}

	if stderr.String() != "" {
		fmt.Printf("%s\n", stderr.String())
	}

	if out.String() != "" {
		fmt.Printf("%s\n", out.String())
	}
}

// handleError handles the errors
func handleError(err error) {
	if err != nil {
		println("A fatal error has occurred")
		println(err.Error())
		os.Exit(1)
	}
}
