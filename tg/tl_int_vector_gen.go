// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// IntVector is a box for Vector<int>
type IntVector struct {
	// Elements of Vector<int>
	Elems []int
}

// IntVectorTypeID is TL type id of IntVector.
const IntVectorTypeID = bin.TypeVector

func (vec *IntVector) Zero() bool {
	if vec == nil {
		return true
	}
	if !(vec.Elems == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (vec *IntVector) String() string {
	if vec == nil {
		return "IntVector(nil)"
	}
	type Alias IntVector
	return fmt.Sprintf("IntVector%+v", Alias(*vec))
}

// FillFrom fills IntVector from given interface.
func (vec *IntVector) FillFrom(from interface {
	GetElems() (value []int)
}) {
	vec.Elems = from.GetElems()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*IntVector) TypeID() uint32 {
	return IntVectorTypeID
}

// TypeName returns name of type in TL schema.
func (*IntVector) TypeName() string {
	return ""
}

// TypeInfo returns info about TL type.
func (vec *IntVector) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "",
		ID:   IntVectorTypeID,
	}
	if vec == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Elems",
			SchemaName: "Elems",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (vec *IntVector) Encode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<int> as nil")
	}
	b.PutVectorHeader(len(vec.Elems))
	for _, v := range vec.Elems {
		b.PutInt(v)
	}
	return nil
}

// GetElems returns value of Elems field.
func (vec *IntVector) GetElems() (value []int) {
	return vec.Elems
}

// Decode implements bin.Decoder.
func (vec *IntVector) Decode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<int> to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode Vector<int>: field Elems: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode Vector<int>: field Elems: %w", err)
			}
			vec.Elems = append(vec.Elems, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for IntVector.
var (
	_ bin.Encoder = &IntVector{}
	_ bin.Decoder = &IntVector{}
)
