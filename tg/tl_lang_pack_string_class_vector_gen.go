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

// LangPackStringClassVector is a box for Vector<LangPackString>
type LangPackStringClassVector struct {
	// Elements of Vector<LangPackString>
	Elems []LangPackStringClass
}

// LangPackStringClassVectorTypeID is TL type id of LangPackStringClassVector.
const LangPackStringClassVectorTypeID = bin.TypeVector

func (vec *LangPackStringClassVector) Zero() bool {
	if vec == nil {
		return true
	}
	if !(vec.Elems == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (vec *LangPackStringClassVector) String() string {
	if vec == nil {
		return "LangPackStringClassVector(nil)"
	}
	type Alias LangPackStringClassVector
	return fmt.Sprintf("LangPackStringClassVector%+v", Alias(*vec))
}

// FillFrom fills LangPackStringClassVector from given interface.
func (vec *LangPackStringClassVector) FillFrom(from interface {
	GetElems() (value []LangPackStringClass)
}) {
	vec.Elems = from.GetElems()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*LangPackStringClassVector) TypeID() uint32 {
	return LangPackStringClassVectorTypeID
}

// TypeName returns name of type in TL schema.
func (*LangPackStringClassVector) TypeName() string {
	return ""
}

// TypeInfo returns info about TL type.
func (vec *LangPackStringClassVector) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "",
		ID:   LangPackStringClassVectorTypeID,
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
func (vec *LangPackStringClassVector) Encode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<LangPackString> as nil")
	}
	b.PutVectorHeader(len(vec.Elems))
	for idx, v := range vec.Elems {
		if v == nil {
			return fmt.Errorf("unable to encode Vector<LangPackString>: field Elems element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode Vector<LangPackString>: field Elems element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetElems returns value of Elems field.
func (vec *LangPackStringClassVector) GetElems() (value []LangPackStringClass) {
	return vec.Elems
}

// MapElems returns field Elems wrapped in LangPackStringClassArray helper.
func (vec *LangPackStringClassVector) MapElems() (value LangPackStringClassArray) {
	return LangPackStringClassArray(vec.Elems)
}

// Decode implements bin.Decoder.
func (vec *LangPackStringClassVector) Decode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<LangPackString> to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode Vector<LangPackString>: field Elems: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeLangPackString(b)
			if err != nil {
				return fmt.Errorf("unable to decode Vector<LangPackString>: field Elems: %w", err)
			}
			vec.Elems = append(vec.Elems, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for LangPackStringClassVector.
var (
	_ bin.Encoder = &LangPackStringClassVector{}
	_ bin.Decoder = &LangPackStringClassVector{}
)
