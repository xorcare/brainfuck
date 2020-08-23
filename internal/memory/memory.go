package memory

type Memory struct {
	cursor uint32
	memory []byte
}

func (m *Memory) Cell() []byte {
	return m.memory[m.cursor : m.cursor+1]
}

func (m *Memory) Byte() byte {
	return m.memory[m.cursor]
}

func (m *Memory) Increment() {
	m.memory[m.cursor]++
}

func (m *Memory) Decrement() {
	m.memory[m.cursor]--
}

func (m *Memory) Next() int {
	m.cursor++
	if m.cursor >= uint32(len(m.memory)) {
		m.memory = append(m.memory, make([]byte, 64)...)
	}
	return int(m.cursor)
}

func (m *Memory) Previous() int {
	m.cursor--
	return int(m.cursor)
}

func (m *Memory) Bytes() []byte {
	return m.memory
}

func New() *Memory {
	return &Memory{
		// In the classic distribution, the array has 30,000 cells, and the
		// pointer begins at the leftmost cell. Even more cells are needed
		// to store things like the millionth Fibonacci number, and the easiest
		// way to make the language Turing complete is to make the array
		// unlimited on the right.
		// See https://en.wikipedia.org/wiki/Brainfuck
		memory: make([]byte, 30000),
	}
}

func NewNano() *Memory {
	return &Memory{
		memory: make([]byte, 16),
	}
}
