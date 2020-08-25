package interpreter

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xorcare/golden"

	"github.com/xorcare/brainfuck/internal/memory"
)

func TestPrepare(t *testing.T) {
	tests := []struct {
		have string
		want string
	}{
		{
			have: "+",
			want: "+",
		},
		{
			have: "+[-]",
			want: "+[-]",
		},
		{
			have: "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[-]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.",
			want: "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[-]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.",
		},
		{
			have: "+++++++++++++[->++>>>+++++>++>+<<<<<<]>>>>>++++++>--->>>>>>>>>>+++++++++++++++",
			want: "+++++++++++++[->++>>>+++++>++>+<<<<<<]>>>>>++++++>--->>>>>>>>>>+++++++++++++++",
		},
		{
			have: "++++++>>[-<<<<+>>>>]<<<<[->>>>+<<[-]<<]>>[<<<<<<<+<[-<+>>>>+<<[-]]>",
			want: "++++++>>[-<<<<+>>>>]<<<<[->>>>+<<[-]<<]>>[<<<<<<<+<[-<+>>>>+<<[-]]>",
		},
		{
			have: "++++++>>[-<<<<+>>>>]<<<<[->>>>+<<[-]<<]>>[<<<<<<<+<[-<+>>>>+<<[-]]>]",
			want: "++++++>>[-<<<<+>>>>]<<<<[->>>>+<<[-]<<]>>[<<<<<<<+<[-<+>>>>+<<[-]]>]",
		},
		{
			have: "->>>>>>>>>#<<<<<<<<<",
			want: "->>>>>>>>><<<<<<<<<",
		},
		{
			have: "->>>>>>>>>$<<<<<<<<<\n",
			want: "->>>>>>>>><<<<<<<<<",
		},
		{
			have: "[[[[[[[[[[[[[[[[]",
			want: "[[[[[[[[[[[[[[[[]",
		},
		{
			have: "[[[[[[[[[[[[[[[[]]]]]]]]]]]]]]",
			want: "[[[[[[[[[[[[[[[[]]]]]]]]]]]]]]",
		},
		{
			have: ",----------[----------------------.,----------]",
			want: ",----------[----------------------.,----------]",
		},
		{
			have: ",----------[----------------------.,----------]",
			want: ",----------[----------------------.,----------]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.have, func(t *testing.T) {
			got, err := Prepare(bytes.NewBufferString(tt.have))
			require.NoError(t, err)
			require.Equal(t, tt.want, got.String())
		})
	}
}

func BenchmarkMandelbrot(b *testing.B) {
	text := golden.Read(b)
	buffer := bytes.NewReader(text)
	b.Run("Prepare", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := Prepare(buffer)
			if err != nil {
				b.Fatal(err)
			}
			if _, err := buffer.Seek(0, 0); err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("Execute", func(b *testing.B) {
		commands, err := Prepare(bytes.NewBuffer(text))
		if err != nil {
			b.Fatal(err)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if err := commands.Execute(nil, ioutil.Discard, memory.New()); err != nil {
				b.Fatal(err)
			}
		}
	})
}
