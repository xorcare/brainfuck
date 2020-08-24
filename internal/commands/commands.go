package commands

import (
	"io"
)

type Conveyor interface {
	Byte() byte
	Cell() []byte
	Decrement()
	Increment()
	Next() int
	Previous() int
}

type Command interface {
	Execute(stdin io.Reader, stdout io.Writer, conveyor Conveyor) error
	String() (character string)
}

func New(character string) Command {
	switch character {
	case "+":
		return IncrementValueCommand{}
	case ".":
		return WriteCommand{}
	case ",":
		return ReadCommand{}
	case ">":
		return IncrementPointerCommand{}
	case "<":
		return DecrementPointerCommand{}
	case "-":
		return DecrementValueCommand{}
	case "]":
		return LoopCloseCommand{}
	}

	return UnknownCommand(character)
}
