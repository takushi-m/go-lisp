package printer

import (
	"fmt"
	"strings"

	"github.com/takushi-m/go-lisp/types"
)

func Print(n *types.Node) string {
	if n == nil {
		return ""
	}

	if n.Type != types.TypeNone {
		switch n.Type {
		case types.TypeNumber:
			return fmt.Sprintf("%d", *n.Number)
		case types.TypeSymbol:
			return *n.Symbol
		default:
			return ""
		}
	}

	ss := make([]string, len(n.Nodes))
	for i, node := range n.Nodes {
		ss[i] = Print(node)
	}
	return "(" + strings.Join(ss, " ") + ")"
}
