package brainfuck

import (
	"io"
)

type Memo interface {
	Cursor() int
	Next() int
	Previous() int
	Bytes()
}

type Memory struct {
}

func Execute(stream io.Reader, stdin io.Reader, stdout io.Writer, memory []byte) error {
	var register uint = 0
	for {
		op := []byte{0}
		count, err := stream.Read(op)
		if err != nil && err != io.EOF {
			return err
		}
		if count == 0 {
			break
		}
		switch string(op) {
		case "+":
			memory[register]++
		case ".":
			if _, err := stdout.Write(memory[register:register]); err != nil {
				return err
			}
		case ",":
			if _, err := stdin.Read(memory[register : register+1]); err != nil {
				return err
			}
		case ">":
			register++
		case "<":
			register--
		case "-":
			memory[register]--
		case "[":

		case "]":

		}
	}

	return nil
}
