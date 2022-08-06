package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStep0(t *testing.T) {
	cases := []string{
		"abcABC123",
		"hello mal world",
		"\"[]{}\"'* ;:()",
		"hello world abcdefghijklmnopqrstuvwxyz ABCDEFGHIJKLMNOPQRSTUVWXYZ 0123456789 (;:() []{}\"'* ;:() []{}\"'* ;:() []{}\"'*)",
		"!",
		"&",
		"+",
		",",
		"-",
		"/",
		"<",
		"=",
		">",
		"?",
		"@",
		"^",
		"_",
		"`",
		"~",
		"#",
		"$",
		"%",
		".",
		"|",
	}

	for _, s := range cases {
		buf := &bytes.Buffer{}
		rep := New(strings.NewReader(s), buf, "")
		stop, err := rep.Run()

		assert.Equal(t, false, stop)
		assert.NoError(t, err)
		assert.Equal(t, s, buf.String())
	}
}
