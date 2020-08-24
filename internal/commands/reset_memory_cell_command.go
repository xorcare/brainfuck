package commands

import (
	"io"
)

type ResetMemoryCellCommand struct{}

func (c ResetMemoryCellCommand) Execute(_ io.Reader, _ io.Writer, conveyor Conveyor) error {
	conveyor.ResetCell()
	return nil
}

func (c ResetMemoryCellCommand) String() string {
	return "[-]"
}
