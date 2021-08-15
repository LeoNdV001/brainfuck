package cli

import (
	"os"
	"testing"

	"github.com/urfave/cli"

	"github.com/stretchr/testify/suite"
)

// Test Suite which encapsulate the tests for the interpreter
type TestSuite struct {
	suite.Suite
	bf BrainfuckCLI
}

// SetupTest sets up often used objects
func (test *TestSuite) SetupTest() {
	test.bf.cli = cli.NewApp()
}

func (test *TestSuite) TearDownTest() {
	os.Clearenv()
}

// TestTestSuite Runs the testsuite
func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

// TestSetInfo tests the setInfo func
func (test *TestSuite) TestRun() {
	// application only registers func and doesn't return any. Check log output...
	Run([]string{"--help"})
}

// TestSetInfo tests the setInfo func
func (test *TestSuite) TestSetInfo() {
	os.Setenv("APPLICATION_NAME", "brainfuck")

	test.bf.setInfo()
	test.Equal("Brainfuck CLI", test.bf.cli.Name)
}

// TestRegisterCLICommands tests the registerCLICommands func
func (test *TestSuite) TestRegisterCLICommands() {
	test.bf.registerCLICommands()

	command := test.bf.cli.Command("compiler")

	test.NotNil(command)
	test.Equal(1, len(test.bf.cli.Commands))
	test.Greater(len(command.Subcommands), 0)

	for i := range command.Subcommands {
		cmd := command.Subcommands[i]
		test.NotNil(cmd.Name)

		run := []string{"brainfuck", "compiler", cmd.Name}

		if cmd.Name == "parse" {
			run = append(run, "examples/helloworld.bf")
		}

		err := test.bf.cli.Run(run)
		test.Nil(err)
	}
}