package brainfuck

import (
	"io"
)

type Memory interface {
	Byte() byte
	Cell() []byte
	Decrement()
	Increment()
	Next() int
	Previous() int
}

func Execute(stream io.Reader, stdin io.Reader, stdout io.Writer, memory Memory) error {
	for {
		controlChar := []byte{0}
		count, err := stream.Read(controlChar)
		if err != nil && err != io.EOF {
			return err
		}
		if count == 0 {
			break
		}
		switch string(controlChar) {
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

		case "]":

		}
	}

	return nil
}
