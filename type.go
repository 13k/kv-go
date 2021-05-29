package kv

//go:generate stringer -type Type -trimprefix Type

// Type represents a KeyValue's node type.
type Type int8

// KeyValue types.
const (
	TypeInvalid Type = iota - 1 // -1
	TypeObject                  // 0x00
	TypeString                  // 0x01
	TypeInt32                   // 0x02
	TypeFloat32                 // 0x03
	TypePointer                 // 0x04
	TypeWString                 // 0x05
	TypeColor                   // 0x06
	TypeUint64                  // 0x07
	TypeEnd                     // 0x08
	_                           // skip
	TypeInt64                   // 0x0a
)

// TypeFromByte converts a byte from binary format to a Type.
//
// Returns TypeInvalid if the given byte is not a valid Type.
func TypeFromByte(b byte) Type {
	t := Type(b)

	switch t {
	case TypeObject,
		TypeString,
		TypeInt32,
		TypeFloat32,
		TypePointer,
		TypeWString,
		TypeColor,
		TypeUint64,
		TypeEnd,
		TypeInt64:
		return t
	default:
		return TypeInvalid
	}
}

// Byte returns the corresponding byte in binary format.
func (t Type) Byte() byte {
	switch t {
	case TypeObject,
		TypeString,
		TypeInt32,
		TypeFloat32,
		TypePointer,
		TypeWString,
		TypeColor,
		TypeUint64,
		TypeEnd,
		TypeInt64:
		return byte(t)
	default:
		return 0
	}
}
