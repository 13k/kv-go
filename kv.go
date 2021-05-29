// Package kv implements encoding and decoding of Valve's KeyValue format.
//
// https://developer.valvesoftware.com/wiki/KeyValues
package kv

import (
	"bytes"
	"encoding"
	"fmt"
	"strconv"
)

// KeyValue represents a node in a KeyValue tree.
type KeyValue interface {
	// Type returns the node's Type.
	Type() Type
	// SetType sets the node's Type and returns the receiver.
	SetType(Type) KeyValue
	// Key returns the node's Key.
	Key() string
	// SetKey sets the node's Key and returns the receiver.
	SetKey(key string) KeyValue
	// Value returns the node's Value.
	Value() string
	// AsString returns Value as string if Type is TypeString, otherwise returns an error.
	AsString() (string, error)
	// AsInt32 returns Value as int32 if Type is TypeInt32, otherwise returns an error.
	AsInt32() (int32, error)
	// AsInt64 returns Value as int64 if Type is TypeInt64, otherwise returns an error.
	AsInt64() (int64, error)
	// AsUint64 returns Value as uint64 if Type is TypeUint64, otherwise returns an error.
	AsUint64() (uint64, error)
	// AsFloat32 returns Value as float32 if Type is TypeFloat32, otherwise returns an error.
	AsFloat32() (float32, error)
	// AsColor returns Value as int32 if Type is TypeColor, otherwise returns an error.
	AsColor() (int32, error)
	// AsPointer returns Value as int32 if Type is TypePointer, otherwise returns an error.
	AsPointer() (int32, error)
	// SetValue sets the node's Value and returns the receiver.
	SetValue(value string) KeyValue
	// SetString sets Value to given string value if Type is TypeString, otherwise returns an error.
	SetString(string) error
	// SetInt32 sets Value to given int32 value if Type is TypeInt32, otherwise returns an error.
	SetInt32(int32) error
	// SetInt64 sets Value to given int64 value if Type is TypeInt64, otherwise returns an error.
	SetInt64(int64) error
	// SetUint64 sets Value to given uint64 value if Type is TypeUint64, otherwise returns an error.
	SetUint64(uint64) error
	// SetFloat32 sets Value to given float32 value if Type is TypeFloat32, otherwise returns an error.
	SetFloat32(float32) error
	// SetColor sets Value to given int32 value if Type is TypeColor, otherwise returns an error.
	SetColor(int32) error
	// SetPointer sets Value to given int32 value if Type is TypePointer, otherwise returns an error.
	SetPointer(int32) error
	// Parent returns the parent node.
	Parent() KeyValue
	// SetParent sets the node's parent node and returns the receiver.
	SetParent(KeyValue) KeyValue
	// Children returns all child nodes
	Children() []KeyValue
	// SetChildren sets the node's children and returns the receiver.
	SetChildren(...KeyValue) KeyValue
	// Child finds a child node with the given key.
	Child(key string) KeyValue
	// NewChild creates an empty child node and returns the child node.
	NewChild() KeyValue
	// AddChild adds a child node and returns the receiver.
	AddChild(KeyValue) KeyValue
	// AddObject adds an Object child node and returns the receiver.
	AddObject(key string) KeyValue
	// AddString adds a String child node and returns the receiver.
	AddString(key, value string) KeyValue
	// AddInt32 adds an Int32 child node and returns the receiver.
	AddInt32(key, value string) KeyValue
	// AddInt64 adds an Int64 child node and returns the receiver.
	AddInt64(key, value string) KeyValue
	// AddUint64 adds an Uint64 child node and returns the receiver.
	AddUint64(key, value string) KeyValue
	// AddFloat32 adds a Float32 child node and returns the receiver.
	AddFloat32(key, value string) KeyValue
	// AddColor adds a Color child node and returns the receiver.
	AddColor(key, value string) KeyValue
	// AddPointer adds a Pointer child node and returns the receiver.
	AddPointer(key, value string) KeyValue

	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
	encoding.TextMarshaler
	encoding.TextUnmarshaler
}

type keyValue struct {
	typ      Type
	key      string
	value    string
	parent   KeyValue
	children []KeyValue

	vInt32   *int32
	vFloat32 *float32
	vPointer *int32
	vColor   *int32
	vUint64  *uint64
	vInt64   *int64
}

// NewKeyValue creates a KeyValue node.
func NewKeyValue(t Type, key, value string, parent KeyValue) KeyValue {
	kv := &keyValue{
		typ:    t,
		key:    key,
		value:  value,
		parent: parent,
	}

	if parent != nil {
		parent.AddChild(kv)
	}

	return kv
}

// NewKeyValueEmpty creates an empty KeyValue node.
func NewKeyValueEmpty() KeyValue {
	return NewKeyValue(TypeInvalid, "", "", nil)
}

// NewKeyValueRoot creates a root KeyValue node.
func NewKeyValueRoot(key string) KeyValue {
	return NewKeyValueObject(key, nil)
}

// NewKeyValueObject creates a KeyValue node with TypeObject type.
func NewKeyValueObject(key string, parent KeyValue) KeyValue {
	return NewKeyValue(TypeObject, key, "", parent)
}

// NewKeyValueString creates a KeyValue node with TypeString type.
func NewKeyValueString(key, value string, parent KeyValue) KeyValue {
	return NewKeyValue(TypeString, key, value, parent)
}

// NewKeyValueInt32 creates a KeyValue node with TypeInt32 type.
func NewKeyValueInt32(key, value string, parent KeyValue) KeyValue {
	return NewKeyValue(TypeInt32, key, value, parent)
}

// NewKeyValueInt64 creates a KeyValue node with TypeInt64 type.
func NewKeyValueInt64(key, value string, parent KeyValue) KeyValue {
	return NewKeyValue(TypeInt64, key, value, parent)
}

// NewKeyValueUint64 creates a KeyValue node with TypeUint64 type.
func NewKeyValueUint64(key, value string, parent KeyValue) KeyValue {
	return NewKeyValue(TypeUint64, key, value, parent)
}

// NewKeyValueFloat32 creates a KeyValue node with TypeFloat32 type.
func NewKeyValueFloat32(key, value string, parent KeyValue) KeyValue {
	return NewKeyValue(TypeFloat32, key, value, parent)
}

// NewKeyValueColor creates a KeyValue node with TypeColor type.
func NewKeyValueColor(key, value string, parent KeyValue) KeyValue {
	return NewKeyValue(TypeColor, key, value, parent)
}

// NewKeyValuePointer creates a KeyValue node with TypePointer type.
func NewKeyValuePointer(key, value string, parent KeyValue) KeyValue {
	return NewKeyValue(TypePointer, key, value, parent)
}

func (kv *keyValue) resetValues() {
	kv.vInt32 = nil
	kv.vFloat32 = nil
	kv.vPointer = nil
	kv.vColor = nil
	kv.vUint64 = nil
	kv.vInt64 = nil
}

func (kv *keyValue) Type() Type { return kv.typ }
func (kv *keyValue) SetType(t Type) KeyValue {
	kv.resetValues()
	kv.typ = t

	return kv
}

func (kv *keyValue) Key() string { return kv.key }
func (kv *keyValue) SetKey(k string) KeyValue {
	kv.key = k
	return kv
}

func (kv *keyValue) Value() string { return kv.value }
func (kv *keyValue) SetValue(v string) KeyValue {
	kv.resetValues()
	kv.value = v

	return kv
}

func (kv *keyValue) AsString() (string, error) {
	if kv.typ != TypeString {
		return "", fmt.Errorf("kv: cannot convert Value of type %s to %s", kv.typ, TypeString)
	}

	return kv.value, nil
}

func (kv *keyValue) asInt32(p **int32) (int32, error) {
	if *p == nil {
		n, err := strconv.ParseInt(kv.value, 10, 32)

		if err != nil {
			return 0, err
		}

		n32 := int32(n)
		*p = &n32
	}

	return **p, nil
}

func (kv *keyValue) AsInt32() (int32, error) {
	if kv.typ != TypeInt32 {
		return 0, fmt.Errorf("kv: cannot convert Value of type %s to %s", kv.typ, TypeInt32)
	}

	return kv.asInt32(&kv.vInt32)
}

func (kv *keyValue) AsInt64() (int64, error) {
	if kv.typ != TypeInt64 {
		return 0, fmt.Errorf("kv: cannot convert Value of type %s to %s", kv.typ, TypeInt64)
	}

	if kv.vInt64 == nil {
		n, err := strconv.ParseInt(kv.value, 10, 64)

		if err != nil {
			return 0, err
		}

		kv.vInt64 = &n
	}

	return *kv.vInt64, nil
}

func (kv *keyValue) AsUint64() (uint64, error) {
	if kv.typ != TypeUint64 {
		return 0, fmt.Errorf("kv: cannot convert Value of type %s to %s", kv.typ, TypeUint64)
	}

	if kv.vUint64 == nil {
		n, err := strconv.ParseUint(kv.value, 10, 64)

		if err != nil {
			return 0, err
		}

		kv.vUint64 = &n
	}

	return *kv.vUint64, nil
}

func (kv *keyValue) AsFloat32() (float32, error) {
	if kv.typ != TypeFloat32 {
		return 0, fmt.Errorf("kv: cannot convert Value of type %s to %s", kv.typ, TypeFloat32)
	}

	if kv.vFloat32 == nil {
		n, err := strconv.ParseFloat(kv.value, 32)

		if err != nil {
			return 0, err
		}

		n32 := float32(n)
		kv.vFloat32 = &n32
	}

	return *kv.vFloat32, nil
}

func (kv *keyValue) AsColor() (int32, error) {
	if kv.typ != TypeColor {
		return 0, fmt.Errorf("kv: cannot convert Value of type %s to %s", kv.typ, TypeColor)
	}

	return kv.asInt32(&kv.vColor)
}

func (kv *keyValue) AsPointer() (int32, error) {
	if kv.typ != TypePointer {
		return 0, fmt.Errorf("kv: cannot convert Value of type %s to %s", kv.typ, TypePointer)
	}

	return kv.asInt32(&kv.vPointer)
}

func (kv *keyValue) SetString(v string) error {
	if kv.typ != TypeString {
		return fmt.Errorf("cannot set Value of type %s with value of type %s", kv.typ, TypeString)
	}

	kv.value = v

	return nil
}

func (kv *keyValue) SetInt32(v int32) error {
	if kv.typ != TypeInt32 {
		return fmt.Errorf("cannot set Value of type %s with value of type %s", kv.typ, TypeInt32)
	}

	kv.vInt32 = &v
	kv.value = strconv.FormatInt(int64(v), 10)

	return nil
}

func (kv *keyValue) SetInt64(v int64) error {
	if kv.typ != TypeInt64 {
		return fmt.Errorf("cannot set Value of type %s with value of type %s", kv.typ, TypeInt64)
	}

	kv.vInt64 = &v
	kv.value = strconv.FormatInt(v, 10)

	return nil
}

func (kv *keyValue) SetUint64(v uint64) error {
	if kv.typ != TypeUint64 {
		return fmt.Errorf("cannot set Value of type %s with value of type %s", kv.typ, TypeUint64)
	}

	kv.vUint64 = &v
	kv.value = strconv.FormatUint(v, 10)

	return nil
}

func (kv *keyValue) SetFloat32(v float32) error {
	if kv.typ != TypeFloat32 {
		return fmt.Errorf("cannot set Value of type %s with value of type %s", kv.typ, TypeFloat32)
	}

	kv.vFloat32 = &v
	kv.value = strconv.FormatFloat(float64(v), 'f', -1, 32)

	return nil
}

func (kv *keyValue) SetColor(v int32) error {
	if kv.typ != TypeColor {
		return fmt.Errorf("cannot set Value of type %s with value of type %s", kv.typ, TypeColor)
	}

	kv.vColor = &v
	kv.value = strconv.FormatInt(int64(v), 10)

	return nil
}

func (kv *keyValue) SetPointer(v int32) error {
	if kv.typ != TypePointer {
		return fmt.Errorf("cannot set Value of type %s with value of type %s", kv.typ, TypePointer)
	}

	kv.vPointer = &v
	kv.value = strconv.FormatInt(int64(v), 10)

	return nil
}

func (kv *keyValue) Parent() KeyValue { return kv.parent }
func (kv *keyValue) SetParent(p KeyValue) KeyValue {
	kv.parent = p
	return kv
}

func (kv *keyValue) Children() []KeyValue { return kv.children }
func (kv *keyValue) SetChildren(children ...KeyValue) KeyValue {
	for _, c := range children {
		c.SetParent(kv)
	}

	kv.children = children

	return kv
}

func (kv *keyValue) Child(key string) KeyValue {
	for _, c := range kv.children {
		if c.Key() == key {
			return c
		}
	}

	return nil
}

func (kv *keyValue) NewChild() KeyValue {
	return NewKeyValue(TypeInvalid, "", "", kv)
}

func (kv *keyValue) AddChild(c KeyValue) KeyValue {
	c.SetParent(kv)

	kv.children = append(kv.children, c)

	return kv
}

func (kv *keyValue) AddObject(key string) KeyValue {
	NewKeyValueObject(key, kv)
	return kv
}

func (kv *keyValue) AddString(key, value string) KeyValue {
	NewKeyValueString(key, value, kv)
	return kv
}

func (kv *keyValue) AddInt32(key, value string) KeyValue {
	NewKeyValueInt32(key, value, kv)
	return kv
}

func (kv *keyValue) AddInt64(key, value string) KeyValue {
	NewKeyValueInt64(key, value, kv)
	return kv
}

func (kv *keyValue) AddUint64(key, value string) KeyValue {
	NewKeyValueUint64(key, value, kv)
	return kv
}

func (kv *keyValue) AddFloat32(key, value string) KeyValue {
	NewKeyValueFloat32(key, value, kv)
	return kv
}

func (kv *keyValue) AddColor(key, value string) KeyValue {
	NewKeyValueColor(key, value, kv)
	return kv
}

func (kv *keyValue) AddPointer(key, value string) KeyValue {
	NewKeyValuePointer(key, value, kv)
	return kv
}

func (kv *keyValue) MarshalText() ([]byte, error) {
	b := &bytes.Buffer{}

	if err := NewTextEncoder(b).Encode(kv); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func (kv *keyValue) UnmarshalText(data []byte) error {
	return NewTextDecoder(bytes.NewReader(data)).Decode(kv)
}

func (kv *keyValue) MarshalBinary() ([]byte, error) {
	b := &bytes.Buffer{}

	if err := NewBinaryEncoder(b).Encode(kv); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func (kv *keyValue) UnmarshalBinary(data []byte) error {
	return NewBinaryDecoder(bytes.NewReader(data)).Decode(kv)
}
