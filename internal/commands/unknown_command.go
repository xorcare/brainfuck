package commands

import (
	"io"
)

type UnknownCommand string

func (c UnknownCommand) Execute(io.Reader, io.Writer, Conveyor) error {
	return nil
}

func (c UnknownCommand) String() string {
	return string(c)
}
