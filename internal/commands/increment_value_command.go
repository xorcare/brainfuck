package commands

import (
	"io"
)

type IncrementValueCommand struct{}

func (c IncrementValueCommand) Execute(_ io.Reader, _ io.Writer, conveyor Conveyor) error {
	conveyor.Increment()
	return nil
}

func (c IncrementValueCommand) String() string {
	return "+"
}
