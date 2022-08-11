package reader

import (
	"fmt"
)

type Reader struct {
	s      string
	tokens []Token
	pos    int
}

type Token string

func New(s string) *Reader {
	return &Reader{
		s:      s,
		tokens: tokenize(s),
		pos:    0,
	}
}

func (r *Reader) Next() (Token, error) {
	if r.pos < len(r.tokens) {
		pos := r.pos
		r.pos += 1
		return r.tokens[pos], nil
	}

	return Token(""), fmt.Errorf("Reader.Next: all token consumed")
}

func (r *Reader) Peek() (Token, error) {
	if r.pos < len(r.tokens) {
		return r.tokens[r.pos], nil
	}
	return Token(""), fmt.Errorf("Reader.Peek: all toekn consumed")
}
