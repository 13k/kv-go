package kv

import (
	"bufio"
	"encoding/binary"
	"io"
	"strconv"
)

const (
	binaryDelimString byte = 0x0
)

// BinaryDecoder reads and decodes binary-encoded KeyValue nodes from an input stream.
type BinaryDecoder struct {
	r *bufio.Reader
}

// NewBinaryDecoder returns a new binary decoder that reads from r.
func NewBinaryDecoder(r io.Reader) *BinaryDecoder {
	return &BinaryDecoder{r: bufio.NewReader(r)}
}

// Decode reads the next binary-encoded KeyValue node from its input and stores it in the value
// pointed to by kv.
func (d *BinaryDecoder) Decode(kv KeyValue) error {
	var (
		typ   Type
		value string
		err   error
	)

	typ, err = d.readType()

	if err != nil {
		return err
	}

	kv.SetType(typ)

	if typ == TypeEnd {
		return nil
	}

	key, err := d.readString()

	if err != nil {
		return err
	}

	kv.SetKey(key)

	switch typ {
	case TypeObject:
		var children []KeyValue

		if children, err = d.readObject(); err != nil {
			return err
		}

		kv.SetChildren(children...)
	case TypeString:
		if value, err = d.readString(); err != nil {
			return err
		}

		kv.SetValue(value)
	case TypeInt32, TypeColor, TypePointer:
		if value, err = d.readInt32String(); err != nil {
			return err
		}

		kv.SetValue(value)
	case TypeInt64:
		if value, err = d.readInt64String(); err != nil {
			return err
		}

		kv.SetValue(value)
	case TypeUint64:
		if value, err = d.readUint64String(); err != nil {
			return err
		}

		kv.SetValue(value)
	case TypeFloat32:
		if value, err = d.readFloat32String(); err != nil {
			return err
		}

		kv.SetValue(value)
	}

	return nil
}

func (d *BinaryDecoder) readObject() ([]KeyValue, error) {
	var kvs []KeyValue

	for {
		kv := NewKeyValueEmpty()

		if err := d.Decode(kv); err != nil {
			return nil, err
		}

		if kv.Type() == TypeEnd {
			break
		}

		kvs = append(kvs, kv)
	}

	return kvs, nil
}

func (d *BinaryDecoder) readType() (Type, error) {
	b, err := d.r.ReadByte()

	if err != nil {
		return 0, err
	}

	return TypeFromByte(b), nil
}

func (d *BinaryDecoder) readString() (string, error) {
	s, err := d.r.ReadString(binaryDelimString)

	if err != nil {
		return "", err
	}

	return s[:len(s)-1], nil
}

func (d *BinaryDecoder) readInt64String() (string, error) {
	var n int64

	if err := binary.Read(d.r, binary.LittleEndian, &n); err != nil {
		return "", err
	}

	return strconv.FormatInt(n, 10), nil
}

func (d *BinaryDecoder) readInt32String() (string, error) {
	var n int32

	if err := binary.Read(d.r, binary.LittleEndian, &n); err != nil {
		return "", err
	}

	return strconv.FormatInt(int64(n), 10), nil
}

func (d *BinaryDecoder) readUint64String() (string, error) {
	var n uint64

	if err := binary.Read(d.r, binary.LittleEndian, &n); err != nil {
		return "", err
	}

	return strconv.FormatUint(n, 10), nil
}

func (d *BinaryDecoder) readFloat32String() (string, error) {
	var n float32

	if err := binary.Read(d.r, binary.LittleEndian, &n); err != nil {
		return "", err
	}

	return strconv.FormatFloat(float64(n), 'f', -1, 32), nil
}
