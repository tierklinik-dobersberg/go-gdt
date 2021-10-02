package gdt

import (
	"fmt"
	"sync"
)

// FieldType describes the type of a record field.
type FieldType string

// FieldDesc describes fields specified in GDT 2.10
type FieldDesc struct {
	ID        FieldID
	Type      FieldType
	Name      string
	Length    int
	MaxLength int
}

// ParseFieldFunc parses the encoded value according to desc.
type ParseFieldFunc func(value []byte, desc *FieldDesc) (interface{}, error)

// TypeRegistry keeps track of registered and known types.
type TypeRegistry struct {
	rl    sync.RWMutex
	types map[FieldType]ParseFieldFunc
}

// Register registeres a new field type ft using fn to parse and validate
// the type.
func (tr *TypeRegistry) Register(ft FieldType, fn ParseFieldFunc) error {
	tr.rl.Lock()
	defer tr.rl.Unlock()

	if _, ok := tr.types[ft]; ok {
		return fmt.Errorf("type %q already registered", ft)
	}

	if tr.types == nil {
		tr.types = make(map[FieldType]ParseFieldFunc)
	}

	tr.types[ft] = fn
	return nil
}

// ParseValue parses the data encoded in value using the type definition
// of desc.
func (tr *TypeRegistry) ParseValue(value []byte, desc *FieldDesc) (interface{}, error) {
	tr.rl.RLock()
	defer tr.rl.RUnlock()

	fn, ok := tr.types[desc.Type]
	if !ok {
		return nil, ErrUnknownType
	}
	return fn(value, desc)
}

// DefaultRegistry is the default type registry.
var DefaultRegistry = new(TypeRegistry)

// Register registeres a new field type at the DefaultRegistry.
func Register(ft FieldType, fn ParseFieldFunc) error {
	return DefaultRegistry.Register(ft, fn)
}

// MustRegister is like Register but panics in case of an error.
func MustRegister(ft FieldType, fn ParseFieldFunc) {
	if err := Register(ft, fn); err != nil {
		panic(err)
	}
}
