package commands

import (
	"io"
)

type DecrementPointerCommand struct{}

func (c DecrementPointerCommand) Execute(_ io.Reader, _ io.Writer, conveyor Conveyor) error {
	conveyor.Previous()
	return nil
}

func (c DecrementPointerCommand) String() string {
	return "<"
}
