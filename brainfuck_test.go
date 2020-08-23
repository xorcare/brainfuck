package brainfuck

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xorcare/golden"

	"github.com/xorcare/brainfuck/internal/memory"
)

func execute(t *testing.T, command string, stdin io.Reader, stdout io.Writer) []byte {
	stream := bytes.NewBufferString(command)
	memory := memory.NewNano()
	if err := Execute(stream, stdin, stdout, memory); err != nil {
		t.Fatal(err)
	}

	return memory.Bytes()
}

func asset(t *testing.T, label string, want, got []byte) {
	t.Logf("%#v", want)
	t.Logf("%#v", string(want))
	want = append(want, make([]byte, len(got)-len(want))...)
	if bytes.Compare(want, got) != 0 {
		require.Equal(t, want, got, "%s: %q", label, want)
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
		{
			cmd:    "++.--",
			stdout: []byte{2},
		},
		{
			cmd:    "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.",
			memory: []byte{0x0, 0x0, 0x48, 0x64, 0x57, 0x21, 0xa},
			stdout: []byte("Hello World!\n"),
		},
		{
			cmd:    "++[>]",
			memory: []byte{2},
		},
		{
			cmd:    "++[+[+[-]+-]+-]",
			memory: []byte{0},
		},
		{
			cmd:    "+.-",
			stdout: []byte{1},
		},
		{
			cmd: "++[-]",
		},
		{
			cmd:    "++[-]+",
			memory: []byte{1},
		},
		{
			cmd:    "++[-]+++[[[[-]]]]+++[-]+++[-]++++++++",
			memory: []byte{8},
		},
		{
			cmd: "+++++[-><]",
		},
		{
			cmd:    "+++++[>+<-]",
			memory: []byte{0, 5},
		},
		{
			cmd:    ",.[-]",
			memory: []byte{0},
			stdin:  []byte("A"),
			stdout: []byte("A"),
		},
		{
			cmd:    ",----------[----------------------.,----------]",
			memory: []byte{0},
			stdin:  []byte("a\n"),
			stdout: []byte("A"),
		},
		{
			cmd:    ",>++++++[<-------->-],[<+>-]<.",
			memory: []byte{0xe},
			stdin:  []byte("4\n2\n"),
			stdout: []byte{0xe},
		},
	}

	for _, tt := range tests {
		t.Run(tt.cmd, func(t *testing.T) {
			stdout := &bytes.Buffer{}
			stdin := bytes.NewBuffer(tt.stdin)
			asset(t, "memory", tt.memory, execute(t, tt.cmd, stdin, stdout))
			asset(t, "stdout", tt.stdout, stdout.Bytes())

			t.Logf("stdin %q", stdin.Bytes())
			t.Logf("stdout %q", stdout.Bytes())
		})
	}

	t.Run("Fibonacci in Brainfuck", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		execute(t, string(golden.Read(t)), bytes.NewBuffer(nil), stdout)
		golden.Equal(t, stdout.Bytes()).FailNow()
	})

	t.Run("Factorial in Brainfuck", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		execute(t, string(golden.Read(t)), bytes.NewBuffer(nil), stdout)
		golden.Equal(t, stdout.Bytes()).FailNow()
	})

	t.Run("Oobrain", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		execute(t, string(golden.Read(t)), bytes.NewBuffer(nil), stdout)
		golden.Equal(t, stdout.Bytes()).FailNow()
	})
}
