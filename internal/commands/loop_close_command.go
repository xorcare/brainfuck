package commands

import (
	"io"
)

type LoopCloseCommand struct{}

func (c LoopCloseCommand) Execute(io.Reader, io.Writer, Conveyor) error {
	return nil
}

func (c LoopCloseCommand) String() string {
	return "]"
}
