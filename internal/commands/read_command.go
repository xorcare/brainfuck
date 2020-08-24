package commands

import (
	"io"
)

type ReadCommand struct{}

func (c ReadCommand) Execute(stdin io.Reader, _ io.Writer, conveyor Conveyor) error {
	if _, err := stdin.Read(conveyor.Cell()); err != nil {
		return err
	}

	return nil
}

func (c ReadCommand) String() string {
	return ","
}
