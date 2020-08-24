package commands

import (
	"io"
)

type IncrementPointerCommand struct{}

func (c IncrementPointerCommand) Execute(_ io.Reader, _ io.Writer, conveyor Conveyor) error {
	conveyor.Next()
	return nil
}

func (c IncrementPointerCommand) String() string {
	return ">"
}
