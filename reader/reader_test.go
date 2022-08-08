package reader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReader_Peek(t *testing.T) {
	cases := []struct {
		input string
		want  Token
	}{
		{
			"1",
			Token("1"),
		},
		{
			"7",
			Token("7"),
		},
		{
			"   7",
			Token("7"),
		},
		{
			"-123",
			Token("-123"),
		},
		{
			"+",
			Token("+"),
		},
		{
			"abc",
			Token("abc"),
		},
		{
			"   abc",
			Token("abc"),
		},
		{
			"abc5",
			Token("abc5"),
		},
		{
			"abc-def",
			Token("abc-def"),
		},
	}

	for _, c := range cases {
		r := New(c.input)
		tk, err := r.Next()
		assert.NoError(t, err)
		assert.Equal(t, c.want, tk)

		_, err = r.Peek()
		assert.Error(t, err)
	}
}
