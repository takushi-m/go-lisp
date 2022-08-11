package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/takushi-m/go-lisp/parser"
	"github.com/takushi-m/go-lisp/printer"
	"github.com/takushi-m/go-lisp/reader"
	"github.com/takushi-m/go-lisp/types"
)

type Rep struct {
	scanner *bufio.Scanner
	writer  *bufio.Writer
	prompt  string
}

func New(r io.Reader, w io.Writer, prompt string) Rep {
	return Rep{
		scanner: bufio.NewScanner(r),
		writer:  bufio.NewWriter(w),
		prompt:  prompt,
	}
}

func (r Rep) Read() (*types.Node, bool, error) {
	if r.prompt != "" {
		_, err := r.writer.WriteString(r.prompt)
		if err != nil {
			return nil, false, err
		}
		if err := r.writer.Flush(); err != nil {
			return nil, false, err
		}
	}

	done := r.scanner.Scan()
	got := r.scanner.Text()
	err := r.scanner.Err()

	p := parser.New(reader.New(got))
	n, err := p.ParseForm()
	if err != nil {
		return nil, false, err
	}

	return n, !done && err == nil, err
}

func (r Rep) Eval(n *types.Node) *types.Node {
	return n
}

func (r Rep) Print(n *types.Node) error {
	s := printer.Print(n)

	_, err := r.writer.WriteString(s)
	if err != nil {
		return err
	}

	return r.writer.Flush()
}

func (r Rep) Run() (bool, error) {
	got, stop, err := r.Read()
	if err != nil {
		return false, err
	}

	if stop {
		return true, nil
	}

	got = r.Eval(got)

	err = r.Print(got)
	if err != nil {
		return false, err
	}

	return false, nil
}

func main() {
	rep := New(os.Stdin, os.Stdout, "=> ")

	for {
		stop, err := rep.Run()
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		if stop {
			break
		}
		fmt.Print("\n")
	}
}
