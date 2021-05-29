package kv

import (
	"io"

	"github.com/13k/kv-go/parser"
)

// TextDecoder reads and decodes text-encoded KeyValue nodes from an input stream.
type TextDecoder struct {
	p *parser.TextParser
}

// NewTextDecoder returns a new text decoder that reads from r.
func NewTextDecoder(r io.Reader) *TextDecoder {
	return &TextDecoder{p: parser.NewTextParser("", r)}
}

// Decode reads the next text-encoded KeyValue node from its input and stores it in the value
// pointed to by kv.
//
// The parser makes no assumptions regarding field types, so all fields are of type TypeString.
func (d *TextDecoder) Decode(kv KeyValue) error {
	root, err := d.p.Parse()

	if err != nil {
		return err
	}

	applyAST(kv, root)

	return nil
}

func applyAST(kv KeyValue, node *parser.Node) {
	kv.SetChildren()
	kv.SetKey(node.Key)

	switch node.Type {
	case parser.Object:
		kv.SetType(TypeObject)

		for _, nodeChild := range node.Children {
			applyAST(kv.NewChild(), nodeChild)
		}
	case parser.Field:
		kv.SetType(TypeString)
		kv.SetValue(node.Value)
	}
}
