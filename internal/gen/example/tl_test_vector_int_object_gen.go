// Code generated by gotdgen, DO NOT EDIT.

package td

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
)

// TestVectorIntObject represents TL type `testVectorIntObject#f152999b`.
//
// See https://localhost:80/doc/constructor/testVectorIntObject for reference.
type TestVectorIntObject struct {
	// Vector of objects
	Value []TestInt
}

// TestVectorIntObjectTypeID is TL type id of TestVectorIntObject.
const TestVectorIntObjectTypeID = 0xf152999b

func (t *TestVectorIntObject) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Value == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *TestVectorIntObject) String() string {
	if t == nil {
		return "TestVectorIntObject(nil)"
	}
	type Alias TestVectorIntObject
	return fmt.Sprintf("TestVectorIntObject%+v", Alias(*t))
}

// FillFrom fills TestVectorIntObject from given interface.
func (t *TestVectorIntObject) FillFrom(from interface {
	GetValue() (value []TestInt)
}) {
	t.Value = from.GetValue()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TestVectorIntObject) TypeID() uint32 {
	return TestVectorIntObjectTypeID
}

// TypeName returns name of type in TL schema.
func (*TestVectorIntObject) TypeName() string {
	return "testVectorIntObject"
}

// TypeInfo returns info about TL type.
func (t *TestVectorIntObject) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "testVectorIntObject",
		ID:   TestVectorIntObjectTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Value",
			SchemaName: "value",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *TestVectorIntObject) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode testVectorIntObject#f152999b as nil")
	}
	b.PutID(TestVectorIntObjectTypeID)
	b.PutVectorHeader(len(t.Value))
	for idx, v := range t.Value {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode testVectorIntObject#f152999b: field value element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetValue returns value of Value field.
func (t *TestVectorIntObject) GetValue() (value []TestInt) {
	return t.Value
}

// Decode implements bin.Decoder.
func (t *TestVectorIntObject) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode testVectorIntObject#f152999b to nil")
	}
	if err := b.ConsumeID(TestVectorIntObjectTypeID); err != nil {
		return fmt.Errorf("unable to decode testVectorIntObject#f152999b: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode testVectorIntObject#f152999b: field value: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value TestInt
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode testVectorIntObject#f152999b: field value: %w", err)
			}
			t.Value = append(t.Value, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for TestVectorIntObject.
var (
	_ bin.Encoder = &TestVectorIntObject{}
	_ bin.Decoder = &TestVectorIntObject{}
)
