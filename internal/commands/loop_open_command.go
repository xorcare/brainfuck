package commands

import (
	"io"
)

type LoopOpenCommand struct {
	Combinator Combinator
}

func (c LoopOpenCommand) Execute(stdin io.Reader, stdout io.Writer, conveyor Conveyor) error {
	for conveyor.Byte() > 0 {
		if err := c.Combinator.Execute(stdin, stdout, conveyor); err != nil {
			return err
		}
	}
	return nil
}

func (c LoopOpenCommand) String() string {
	return "[" + c.Combinator.String()
}
