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

// MessageRangeVector is a box for Vector<MessageRange>
type MessageRangeVector struct {
	// Elements of Vector<MessageRange>
	Elems []MessageRange
}

// MessageRangeVectorTypeID is TL type id of MessageRangeVector.
const MessageRangeVectorTypeID = bin.TypeVector

func (vec *MessageRangeVector) Zero() bool {
	if vec == nil {
		return true
	}
	if !(vec.Elems == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (vec *MessageRangeVector) String() string {
	if vec == nil {
		return "MessageRangeVector(nil)"
	}
	type Alias MessageRangeVector
	return fmt.Sprintf("MessageRangeVector%+v", Alias(*vec))
}

// FillFrom fills MessageRangeVector from given interface.
func (vec *MessageRangeVector) FillFrom(from interface {
	GetElems() (value []MessageRange)
}) {
	vec.Elems = from.GetElems()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageRangeVector) TypeID() uint32 {
	return MessageRangeVectorTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageRangeVector) TypeName() string {
	return ""
}

// TypeInfo returns info about TL type.
func (vec *MessageRangeVector) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "",
		ID:   MessageRangeVectorTypeID,
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
func (vec *MessageRangeVector) Encode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<MessageRange> as nil")
	}
	b.PutVectorHeader(len(vec.Elems))
	for idx, v := range vec.Elems {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode Vector<MessageRange>: field Elems element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetElems returns value of Elems field.
func (vec *MessageRangeVector) GetElems() (value []MessageRange) {
	return vec.Elems
}

// Decode implements bin.Decoder.
func (vec *MessageRangeVector) Decode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<MessageRange> to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode Vector<MessageRange>: field Elems: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value MessageRange
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode Vector<MessageRange>: field Elems: %w", err)
			}
			vec.Elems = append(vec.Elems, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessageRangeVector.
var (
	_ bin.Encoder = &MessageRangeVector{}
	_ bin.Decoder = &MessageRangeVector{}
)
