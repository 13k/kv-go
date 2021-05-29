package kv

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	textIndent      = "  "
	textObjectStart = "{"
	textObjectEnd   = "}"
)

// TextEncoder writes text-encoded KeyValue nodes to an output stream.
type TextEncoder struct {
	w *bufio.Writer
}

// NewTextEncoder returns a new text encoder that writes to w.
func NewTextEncoder(w io.Writer) *TextEncoder {
	return &TextEncoder{w: bufio.NewWriter(w)}
}

// Encode writes the KeyValue text encoding of kv to the stream.
func (e *TextEncoder) Encode(kv KeyValue) error {
	return e.encode(kv, 0)
}

func (e *TextEncoder) encode(kv KeyValue, level int) error {
	switch kv.Type() {
	case TypeInvalid, TypeEnd:
		return fmt.Errorf("kv: cannot encode nodes of type %s", kv.Type())
	}

	indent := strings.Repeat(textIndent, level)
	qKey := strconv.Quote(kv.Key())

	if _, err := fmt.Fprintf(e.w, "%s%s ", indent, qKey); err != nil {
		return err
	}

	switch kv.Type() {
	case TypeObject:
		if err := e.encodeObject(kv, indent, level); err != nil {
			return err
		}
	default:
		qVal := strconv.Quote(kv.Value())

		if _, err := e.w.WriteString(qVal); err != nil {
			return err
		}
	}

	return e.w.Flush()
}

func (e *TextEncoder) encodeObject(kv KeyValue, indent string, level int) error {
	if _, err := fmt.Fprintf(e.w, "%s\n", textObjectStart); err != nil {
		return err
	}

	nextLevel := level + 1

	for _, c := range kv.Children() {
		if err := e.encode(c, nextLevel); err != nil {
			return err
		}

		if err := e.w.WriteByte('\n'); err != nil {
			return err
		}
	}

	if _, err := fmt.Fprintf(e.w, "%s%s\n", indent, textObjectEnd); err != nil {
		return err
	}

	return nil
}
