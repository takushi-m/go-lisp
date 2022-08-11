package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStep1(t *testing.T) {
	for _, c := range []struct {
		input  string
		output string
	}{
		{
			"(+ 1 2)",
			"(+ 1 2)",
		},
		{
			"()",
			"()",
		},
		{
			"( )",
			"()",
		},
		{
			"(nil)",
			"(nil)",
		},
		{
			"((3 4))",
			"((3 4))",
		},
		{
			"(+ 1 (+ 2 3))",
			"(+ 1 (+ 2 3))",
		},
		{
			"   ( +   1   (+   2 3   )   )",
			"(+ 1 (+ 2 3))",
		},
		{
			"(* 1 2)",
			"(* 1 2)",
		},
		{
			"(** 1 2)",
			"(** 1 2)",
		},
		{
			"(* -3 6)",
			"(* -3 6)",
		},
		{
			"(()())",
			"(() ())",
		},
		{
			"(1 2, 3,,,,),,",
			"(1 2 3)",
		},
		{
			"true",
			"true",
		},
		{
			"false",
			"false",
		},
		{
			"nil",
			"nil",
		},
	} {
		buf := &bytes.Buffer{}
		rep := New(strings.NewReader(c.input), buf, "")
		stop, err := rep.Run()

		assert.Equal(t, false, stop)
		assert.NoError(t, err)
		assert.Equal(t, c.output, buf.String())
	}
}

func TestStep0(t *testing.T) {
	cases := []string{
		"abcABC123",
		"!",
		"&",
		"+",
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
