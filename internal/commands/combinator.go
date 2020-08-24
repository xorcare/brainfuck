package commands

import (
	"io"
)

var _ Command = Combinator{}

type Combinator []Command

func (c Combinator) String() (str string) {
	for _, supervisor := range c {
		str += supervisor.String()
	}
	return
}

func (c Combinator) Execute(stdin io.Reader, stdout io.Writer, conveyor Conveyor) error {
	for i := range c {
		if err := c[i].Execute(stdin, stdout, conveyor); err != nil {
			return err
		}
	}
	return nil
}
