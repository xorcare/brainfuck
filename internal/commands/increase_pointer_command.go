package commands

import (
	"io"
	"strings"
)

type IncreasePointerCommand struct {
	Delta int
}

func (c IncreasePointerCommand) Execute(_ io.Reader, _ io.Writer, conveyor Conveyor) error {
	conveyor.AddPosition(c.Delta)
	return nil
}

func (c IncreasePointerCommand) String() string {
	if c.Delta > 0 {
		return strings.Repeat(">", c.Delta)
	}
	if c.Delta <= 0 {
		return strings.Repeat("<", c.Delta*-1)
	}
	return ""
}
