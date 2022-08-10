package types

type TypeEnum int

const (
	TypeNone TypeEnum = iota
	TypeNumber
	TypeSymbol
)

type Node struct {
	Nodes []*Node

	Type   TypeEnum
	Number *int64
	Symbol *string
}

func NewNode() *Node {
	ns := make([]*Node, 0)
	return &Node{
		Nodes: ns,
	}
}
