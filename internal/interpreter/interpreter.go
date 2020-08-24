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

func Optimize(combinator commands.Combinator) (optimized commands.Combinator) {
	optimized = make(commands.Combinator, 0, len(combinator))

	for i := 0; i < len(combinator); i++ {
		command := combinator[i]
		switch command.(type) {
		case commands.LoopOpenCommand:
			if command.String() == "[-]" {
				optimized = append(optimized, commands.ResetMemoryCellCommand{})
				break
			}
			optimized = append(optimized, command)
		case commands.IncrementPointerCommand:
			count := 1
			for i = i + 1; i < len(combinator); i++ {
				command := combinator[i]
				if _, ok := command.(commands.IncrementPointerCommand); !ok {
					i--
					break
				}
				count++
			}
			optimized = append(optimized, commands.IncreasePointerCommand{Delta: count})
		case commands.DecrementPointerCommand:
			count := -1
			for i = i + 1; i < len(combinator); i++ {
				command := combinator[i]
				if _, ok := command.(commands.DecrementPointerCommand); !ok {
					i--
					break
				}
				count--
			}
			optimized = append(optimized, commands.IncreasePointerCommand{Delta: count})
		case commands.UnknownCommand:
			// skip an unknown command.
		default:
			optimized = append(optimized, command)
		}
	}

	if len(combinator) != len(optimized) {
		return Optimize(optimized)
	}

	return optimized
}
