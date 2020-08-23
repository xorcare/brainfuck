package brainfuck

import (
	"io"

	"github.com/xorcare/brainfuck/internal/interpreter"
)

func Execute(
	stream io.Reader,
	stdin io.Reader,
	stdout io.Writer,
	conveyor interpreter.Conveyor,
) error {
	pip, err := interpreter.Prepare(stream)
	if err != nil {
		return err
	}

	return pip.Execute(stdin, stdout, conveyor)
}
