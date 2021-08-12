package interpreter

import "fmt"

/*
> 	Increment the pointer.
< 	Decrement the pointer.
+ 	Increment the byte at the pointer.
- 	Decrement the byte at the pointer.
. 	Output the byte at the pointer.
, 	Input a byte and store it in the byte at the pointer.
[ 	Jump forward past the matching ] if the byte at the pointer is zero.
] 	Jump backward to the matching [ unless the byte at the pointer is zero.
*/

// Interpreter contains the commands to run
type Interpreter struct {
	size     int
	commands []byte
	pointer  int
}

// Loop contains helper properties for handling loops
type Loop struct {
	startPos  int
	endPos    int
	skipClose int
	isOpen    bool
}

// NewInterpreter initializes a new Interpreter
func NewInterpreter() *Interpreter {
	/*
		A Brainfuck program has an implicit byte pointer, called "the pointer",
		which is free to move around within an array of 30000 bytes, initially all set to zero.
		The pointer itself is initialized to point to the beginning of this array.
	*/
	byteSize := 30000

	return &Interpreter{
		size:     byteSize,
		commands: make([]byte, byteSize),
		pointer:  0,
	}
}

// Next handles the character ">" (increment the pointer).
func (interpreter *Interpreter) Next() {
	if interpreter.pointer == interpreter.size-1 {
		interpreter.pointer = 0
	} else {
		interpreter.pointer++
	}
}

// Previous handles the character "<" (decrement the pointer).
func (interpreter *Interpreter) Previous() {
	if interpreter.pointer == 0 {
		interpreter.pointer = interpreter.size - 1
	} else {
		interpreter.pointer--
	}
}

// Increment handles the character "+" (increment the byte at the pointer).
func (interpreter *Interpreter) Increment() {
	interpreter.commands[interpreter.pointer]++
}

// Decrement handles the character "-" (decrement the byte at the pointer).
func (interpreter *Interpreter) Decrement() {
	interpreter.commands[interpreter.pointer]--
}

// Output handles the character "." (output the byte at the pointer).
func (interpreter *Interpreter) Output() {
	// %c returns the character represented by the corresponding Unicode code point
	fmt.Printf("%c", interpreter.commands[interpreter.pointer])
}

// Input handles the character "," (input a byte and store it in the byte at the pointer).
func (interpreter *Interpreter) Input() {
	// ignore error handling for now
	_, _ = fmt.Scanf("%c", &interpreter.commands[interpreter.pointer])
}

// Run runs the interpreter
func (interpreter *Interpreter) Run() {
	//@TODO: read from file

}

// Parse runs the interpreter
func (interpreter *Interpreter) Parse(commands string) error {
	loop := Loop{
		startPos:  -1,
		endPos:    -1,
		skipClose: 0,
		isOpen:    false,
	}

	// loop over the characters within the
	for i := range commands {
		command := commands[i]

		if loop.isOpen {
			interpreter.ParseLoop(&loop, i, commands, string(command))
			continue
		}

		switch string(command) {
		case ">":
			interpreter.Next()
		case "<":
			interpreter.Previous()
		case "+":
			interpreter.Increment()
		case "-":
			interpreter.Decrement()
		case ".":
			interpreter.Output()
		case ",":
			interpreter.Input()
		case "[":
			interpreter.OpenLoop(&loop, i)
		case "]":
			// loops are closed in parse loops
			continue
		default:
			continue
		}
	}

	return nil
}

// OpenLoop opens the loop
func (interpreter *Interpreter) OpenLoop(loop *Loop, pos int) {
	loop.isOpen = true
	loop.startPos = pos + 1
}

// ParseLoop parses the loop
func (interpreter *Interpreter) ParseLoop(loop *Loop, pos int, commands, currentCommand string) {
	switch currentCommand {
	case "[":
		loop.skipClose++
	case "]":
		if loop.skipClose > 0 {
			return
		}

		// close loop
		loop.isOpen = false
		loop.endPos = pos

		// ignore loop and reset if it is closed right away
		if loop.startPos == loop.endPos {
			loop.startPos = -1
			loop.endPos = -1

			return
		}

		// define and loop over commands within loop
		loopCommands := commands[loop.startPos:loop.endPos]

		for interpreter.commands[interpreter.pointer] > 0 {
			err := interpreter.Parse(loopCommands)

			// print error
			if err != nil {
				fmt.Printf("%v", err.Error())
				return
			}
		}
	}
}
