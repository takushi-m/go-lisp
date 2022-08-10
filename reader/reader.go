package reader

import (
	"fmt"
	"regexp"

	"github.com/takushi-m/go-lisp/types"
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

func ReadForm(r *Reader) (*types.Node, error) {
	t, err := r.Peek()
	if err != nil {
		return nil, err
	}

	switch t {
	case Token("("):
		return ReadList(r)
	default:
		return ReadAtom(r)
	}
}

func ReadList(r *Reader) (*types.Node, error) {
	n := types.NewNode()

	_, err := r.Next()
	if err != nil {
		return nil, err
	}

	for {
		t, err := r.Peek()
		if err != nil {
			return nil, err
		}

		if t == Token(")") {
			_, _ = r.Next()
			break
		}

		nn, err := ReadForm(r)
		if err != nil {
			return nil, err
		}
		n.Nodes = append(n.Nodes, nn)
	}

	return n, nil
}

var (
	numberReg = regexp.MustCompile("^[1-9][0-9]*$")
)

func ReadAtom(r *Reader) (*types.Node, error) {
	n := types.NewNode()

	t, err := r.Next()
	if err != nil {
		return nil, err
	}

	switch {
	case numberReg.MatchString(string(t)):
		var i int64
		_, _ = fmt.Sscanf(string(t), "%d", &i)
		n.Type = types.TypeNumber
		n.Number = &i
	default:
		s := string(t)
		n.Type = types.TypeSymbol
		n.Symbol = &s
	}

	return n, nil
}
