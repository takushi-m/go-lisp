package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

func (r Rep) Read() (string, bool, error) {
	if r.prompt != "" {
		_, err := r.writer.WriteString(r.prompt)
		if err != nil {
			return "", false, err
		}
		if err := r.writer.Flush(); err != nil {
			return "", false, err
		}
	}

	done := r.scanner.Scan()
	got := r.scanner.Text()
	err := r.scanner.Err()

	return got, !done && err == nil, err
}

func (r Rep) Eval(s string) string {
	return s
}

func (r Rep) Print(s string) error {
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
