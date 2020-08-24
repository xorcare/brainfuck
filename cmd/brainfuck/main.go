package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/xorcare/brainfuck"
	"github.com/xorcare/brainfuck/internal/memory"
)

func main() {
	if !terminal.IsTerminal(0) {
		execute(bufio.NewReader(os.Stdin), bytes.NewBuffer(nil))
		return
	}

	for {
		code := ""
		data := ""
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter code: \n$ ")
		code, _ = reader.ReadString('\n')

		if strings.Contains(code, ",") {
			fmt.Print("Enter data: \n> ")
			data, _ = reader.ReadString('\n')
		}

		execute(bytes.NewBufferString(code), bytes.NewBufferString(data))
	}
}

func execute(stream, stdin io.Reader) {
	conveyor := memory.New()

	writer := bufio.NewWriter(os.Stdout)
	err := brainfuck.Execute(
		bufio.NewReader(stream),
		bufio.NewReader(stdin),
		writer,
		conveyor,
	)
	if err := writer.Flush(); err != nil {
		log.Println(err)
	}
	fmt.Println(fmt.Sprintf("\nmemory dump => %q\n", conveyor.Byte()))
	if err != nil {
		log.Fatal(err)
	}
}
