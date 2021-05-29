package kv

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
)

// BinaryEncoder writes binary-encoded KeyValue nodes to an output stream.
type BinaryEncoder struct {
	w *bufio.Writer
}

// NewBinaryEncoder returns a new binary encoder that writes to w.
func NewBinaryEncoder(w io.Writer) *BinaryEncoder {
	return &BinaryEncoder{w: bufio.NewWriter(w)}
}

// Encode writes the KeyValue binary encoding of kv to the stream.
func (e *BinaryEncoder) Encode(kv KeyValue) error {
	switch kv.Type() {
	case TypeInvalid, TypeEnd, TypeWString:
		return fmt.Errorf("kv: cannot encode node of type %s", kv.Type())
	}

	var err error

	if err = e.writeType(kv.Type()); err != nil {
		return err
	}

	if err = e.writeString(kv.Key()); err != nil {
		return err
	}

	switch kv.Type() {
	case TypeObject:
		err = e.encodeObject(kv)
	case TypeString:
		err = e.encodeString(kv)
	case TypeInt32:
		err = e.encodeInt32(kv)
	case TypeInt64:
		err = e.encodeInt64(kv)
	case TypeUint64:
		err = e.encodeUint64(kv)
	case TypeFloat32:
		err = e.encodeFloat32(kv)
	case TypeColor:
		err = e.encodeColor(kv)
	case TypePointer:
		err = e.encodePointer(kv)
	}

	if err != nil {
		return err
	}

	return e.w.Flush()
}

func (e *BinaryEncoder) writeType(t Type) error {
	return e.w.WriteByte(t.Byte())
}

func (e *BinaryEncoder) writeString(s string) error {
	if _, err := e.w.WriteString(s); err != nil {
		return err
	}

	return e.w.WriteByte(binaryDelimString)
}

func (e *BinaryEncoder) writeInt32(n int32) error {
	return binary.Write(e.w, binary.LittleEndian, n)
}

func (e *BinaryEncoder) writeInt64(n int64) error {
	return binary.Write(e.w, binary.LittleEndian, n)
}

func (e *BinaryEncoder) writeUint64(n uint64) error {
	return binary.Write(e.w, binary.LittleEndian, n)
}

func (e *BinaryEncoder) writeFloat32(n float32) error {
	return binary.Write(e.w, binary.LittleEndian, n)
}

func (e *BinaryEncoder) encodeObject(kv KeyValue) error {
	for _, c := range kv.Children() {
		if err := e.Encode(c); err != nil {
			return err
		}
	}

	if err := e.writeType(TypeEnd); err != nil {
		return err
	}

	return nil
}

func (e *BinaryEncoder) encodeString(kv KeyValue) error {
	s, err := kv.AsString()

	if err != nil {
		return err
	}

	return e.writeString(s)
}

func (e *BinaryEncoder) encodeInt32(kv KeyValue) error {
	n, err := kv.AsInt32()

	if err != nil {
		return err
	}

	return e.writeInt32(n)
}

func (e *BinaryEncoder) encodeInt64(kv KeyValue) error {
	n, err := kv.AsInt64()

	if err != nil {
		return err
	}

	return e.writeInt64(n)
}

func (e *BinaryEncoder) encodeUint64(kv KeyValue) error {
	n, err := kv.AsUint64()

	if err != nil {
		return err
	}

	return e.writeUint64(n)
}

func (e *BinaryEncoder) encodeFloat32(kv KeyValue) error {
	n, err := kv.AsFloat32()

	if err != nil {
		return err
	}

	return e.writeFloat32(n)
}

func (e *BinaryEncoder) encodeColor(kv KeyValue) error {
	n, err := kv.AsColor()

	if err != nil {
		return err
	}

	return e.writeInt32(n)
}

func (e *BinaryEncoder) encodePointer(kv KeyValue) error {
	n, err := kv.AsPointer()

	if err != nil {
		return err
	}

	return e.writeInt32(n)
}
