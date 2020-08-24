package interpreter

import (
	"io"

	commands "github.com/xorcare/brainfuck/internal/commands"
)

type Conveyor = commands.Conveyor

func Prepare(reader io.Reader) (commands.Combinator, error) {
	var combinator = make(commands.Combinator, 0, 64)
	for {
		single := []byte{0}
		count, err := reader.Read(single)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if count == 0 {
			break
		}
		character := string(single)
		switch character {
		case "+", ".", ",", ">", "<", "-":
			combinator = append(combinator, commands.New(character))
		case "[":
			prepare, err := Prepare(reader)
			if err != nil {
				return nil, err
			}
			combinator = append(combinator, commands.LoopOpenCommand{
				Combinator: prepare,
			})
		case "]":
			combinator = append(combinator, commands.New(character))
			return combinator, nil
		}
	}

	return combinator, nil
}
