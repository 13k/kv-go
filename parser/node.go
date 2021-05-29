package parser

import (
	"fmt"
	"strconv"
	"strings"
)

// NodeType represents a Node's type.
type NodeType uint8

// Node types.
const (
	Object NodeType = iota
	Field
)

func (t NodeType) String() string {
	switch t {
	case Object:
		return "Object"
	case Field:
		return "Field"
	default:
		return fmt.Sprintf("NodeType(%d)", t)
	}
}

// Node is an AST node.
type Node struct {
	Parent   *Node
	Children []*Node
	Type     NodeType
	Key      string
	Value    string
}

func (n *Node) addChild(child *Node) *Node {
	child.Parent = n

	n.Children = append(n.Children, child)

	return child
}

func (n *Node) String() string {
	return n.toString(0)
}

func (n *Node) toString(level int) string {
	var b strings.Builder

	indent := strings.Repeat("  ", level)

	b.WriteString(fmt.Sprintf("%s[%s] %s ", indent, n.Type, strconv.Quote(n.Key)))

	switch n.Type {
	case Object:
		b.WriteRune(tokObjectStart)
		b.WriteRune('\n')

		for _, c := range n.Children {
			b.WriteString(c.toString(level + 1))
		}

		b.WriteString(indent)
		b.WriteRune(tokObjectEnd)
		b.WriteRune('\n')
	case Field:
		b.WriteString(strconv.Quote(n.Value))
		b.WriteRune('\n')
	}

	return b.String()
}
