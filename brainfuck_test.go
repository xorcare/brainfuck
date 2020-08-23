package brainfuck

import (
	"bytes"
	"io"
	"testing"
)

func execute(t *testing.T, command string, stdin io.Reader, stdout io.Writer) []byte {
	stream := bytes.NewBufferString(command)
	memory := make([]byte, 16)
	if err := Execute(stream, stdin, stdout, memory); err != nil {
		t.Fatal(err)
	}

	return memory
}

func asset(t *testing.T, label string, want, got []byte) {
	memory := make([]byte, len(got))
	copy(memory, want)
	if bytes.Compare(memory, got) != 0 {
		t.Fatalf("%s: not equal\nwant: %q,\ngot: %q", label, memory, got)
	}
}

func TestExecute(t *testing.T) {
	tests := []struct {
		cmd    string
		stdin  []byte
		memory []byte
		stdout []byte
	}{
		{
			cmd:    "",
			memory: []byte{},
		},
		{
			cmd:    "++",
			memory: []byte{2},
		},
		{
			cmd:    "+*",
			memory: []byte{1},
		},
		{
			cmd:    "++>+++",
			memory: []byte{2, 3},
		},
		{
			cmd:    "++>+<++",
			memory: []byte{4, 1},
		},
		{
			cmd:    "++>+>++",
			memory: []byte{2, 1, 2},
		},
		{
			cmd:    "+++++++.",
			memory: []byte{7},
			stdout: []byte{7},
		},
		{
			cmd:    ",.",
			memory: []byte{65},
			stdin:  []byte{65},
			stdout: []byte{65},
		},
		{
			cmd:    "+>,.",
			memory: []byte{1, 65},
			stdin:  []byte{65},
			stdout: []byte{65},
		},
		{
			cmd:    "<<<<<",
			memory: []byte{},
			stdin:  []byte{},
			stdout: []byte{},
		},
		{
			cmd:    ">>>>>",
			memory: []byte{},
			stdin:  []byte{},
			stdout: []byte{},
		},
		{
			cmd:    "-",
			memory: []byte{255},
			stdin:  []byte{},
			stdout: []byte{},
		},
		{
			cmd:    "--",
			memory: []byte{254},
		},
		{
			cmd: "+-",
		},
	}

	for _, tt := range tests {
		t.Run(tt.cmd, func(t *testing.T) {
			stdout := &bytes.Buffer{}
			asset(t, "memory", tt.memory, execute(t, tt.cmd, bytes.NewBuffer(tt.stdin), stdout))
			asset(t, "stdout", tt.stdout, stdout.Bytes())
		})
	}
}
