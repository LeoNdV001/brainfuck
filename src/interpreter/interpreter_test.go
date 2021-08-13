package interpreter

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

// Test Suite which encapsulate the tests for the interpreter
type TestSuite struct {
	suite.Suite
	interpreter *Interpreter
	helloWorld string
}

// SetupTest sets up often used objects
func (test *TestSuite) SetupTest() {
	test.interpreter = NewInterpreter()
	test.helloWorld = "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."
}

// TestTestSuite Runs the testsuite
func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

// TestParse tests the Parse func
func (test *TestSuite) TestParse() {
	err := test.interpreter.Parse(test.helloWorld)
	test.Nil(err)
	test.Equal("Hello World!\n", test.interpreter.parsedCommands)
}