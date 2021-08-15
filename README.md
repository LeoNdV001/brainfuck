# brainfuck
Golang brainfuck interpreter.

Made as an assessment, but still could use some improvements.

### Run in command line
```./brainfuck compiler parse examples/helloworld.bf```

#### Import as Library
```
interpreter := NewInterpreter()
interpreter.Parse([some bf string])

fmt.Println(interpreter.ParsedCommands)
```
