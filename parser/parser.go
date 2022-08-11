package parser

import (
	"fmt"
	"regexp"

	"github.com/takushi-m/go-lisp/reader"
	"github.com/takushi-m/go-lisp/types"
)

type Parser struct {
	Reader *reader.Reader
}

func New(r *reader.Reader) *Parser {
	return &Parser{
		Reader: r,
	}
}

func (p *Parser) ParseForm() (*types.Node, error) {
	t, err := p.Reader.Peek()
	if err != nil {
		return nil, err
	}

	switch t {
	case reader.Token("("):
		return p.ParseList()
	default:
		return p.ParseAtom()
	}
}

func (p *Parser) ParseList() (*types.Node, error) {
	n := types.NewNode()

	_, err := p.Reader.Next()
	if err != nil {
		return nil, err
	}

	for {
		t, err := p.Reader.Peek()
		if err != nil {
			return nil, err
		}

		if t == reader.Token(")") {
			_, _ = p.Reader.Next()
			break
		}

		nn, err := p.ParseForm()
		if err != nil {
			return nil, err
		}
		n.Nodes = append(n.Nodes, nn)
	}

	return n, nil
}

var (
	numberRegExp = regexp.MustCompile("^[1-9][0-9]*$")
	boolRegExp   = regexp.MustCompile("^true|false$")
)

func (p *Parser) ParseAtom() (*types.Node, error) {
	n := types.NewNode()

	t, err := p.Reader.Next()
	if err != nil {
		return nil, err
	}

	switch {
	case numberRegExp.MatchString(string(t)):
		var i int64
		_, _ = fmt.Sscanf(string(t), "%d", &i)
		n.Type = types.TypeNumber
		n.Number = &i
	case boolRegExp.MatchString(string(t)):
		var v bool
		if string(t) == "true" {
			v = true
		}
		n.Type = types.TypeBool
		n.Bool = &v
	case string(t) == "nil":
		n.Type = types.TypeNil
	default:
		s := string(t)
		n.Type = types.TypeSymbol
		n.Symbol = &s
	}

	return n, nil
}
