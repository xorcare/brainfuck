package commands

import (
	"io"
)

type WriteCommand struct{}

func (c WriteCommand) Execute(_ io.Reader, stdout io.Writer, memory Conveyor) error {
	if _, err := stdout.Write(memory.Cell()); err != nil {
		return err
	}

	return nil
}

func (c WriteCommand) String() string {
	return "."
}
