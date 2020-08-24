package commands

import (
	"io"
)

type DecrementValueCommand struct{}

func (c DecrementValueCommand) Execute(_ io.Reader, _ io.Writer, conveyor Conveyor) error {
	conveyor.Decrement()
	return nil
}

func (c DecrementValueCommand) String() string {
	return "-"
}
