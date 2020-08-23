package interpreter

import (
	"fmt"
	"io"
)

type Conveyor interface {
	Byte() byte
	Cell() []byte
	Decrement()
	Increment()
	Next() int
	Previous() int
}

type Command interface {
	Execute(stdin io.Reader, stdout io.Writer, memory Conveyor) error
	String() string
}

type Combinator []Command

func (p Combinator) String() (str string) {
	for _, supervisor := range p {
		str += supervisor.String()
	}
	return
}

func (p Combinator) Execute(stdin io.Reader, stdout io.Writer, memory Conveyor) error {
	for i := range p {
		if err := p[i].Execute(stdin, stdout, memory); err != nil {
			return err
		}
	}
	return nil
}

func Prepare(reader io.Reader) (Command, error) {
	var pip = make(Combinator, 0, 64)
	for {
		char := []byte{0}
		count, err := reader.Read(char)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if count == 0 {
			break
		}
		controlChar := string(char)
		switch controlChar {
		case "+", ".", ",", ">", "<", "-":
			pip = append(pip, command{kind: controlChar})
		case "[":
			pipeline, err := Prepare(reader)
			if err != nil {
				return nil, err
			}
			pip = append(pip, command{kind: "[", pipeline: pipeline})
		case "]":
			pip = append(pip, command{kind: controlChar})
			return pip, nil
		}
	}

	return pip, nil
}

type command struct {
	kind     string
	pipeline Command
}

func (c command) String() string {
	if c.kind == "[" {
		return "[" + c.pipeline.String()
	}

	return c.kind
}

func (c command) Execute(stdin io.Reader, stdout io.Writer, memory Conveyor) error {
	switch c.kind {
	case "+":
		memory.Increment()
	case ".":
		_, err := stdout.Write(memory.Cell())
		if err != nil {
			return err
		}
	case ",":
		_, err := stdin.Read(memory.Cell())
		if err != nil {
			return err
		}
	case ">":
		memory.Next()
	case "<":
		memory.Previous()
	case "-":
		memory.Decrement()
	case "[":
		for memory.Byte() > 0 {
			if err := c.pipeline.Execute(stdin, stdout, memory); err != nil {
				return err
			}
		}
	case "]":
	default:
		return fmt.Errorf("unknown control char: %q", c.kind)
	}

	return nil
}
