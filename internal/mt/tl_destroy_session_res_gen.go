// Code generated by gotdgen, DO NOT EDIT.

package mt

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

// DestroySessionOk represents TL type `destroy_session_ok#e22045fc`.
type DestroySessionOk struct {
	// SessionID field of DestroySessionOk.
	SessionID int64
}

// DestroySessionOkTypeID is TL type id of DestroySessionOk.
const DestroySessionOkTypeID = 0xe22045fc

func (d *DestroySessionOk) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.SessionID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DestroySessionOk) String() string {
	if d == nil {
		return "DestroySessionOk(nil)"
	}
	type Alias DestroySessionOk
	return fmt.Sprintf("DestroySessionOk%+v", Alias(*d))
}

// FillFrom fills DestroySessionOk from given interface.
func (d *DestroySessionOk) FillFrom(from interface {
	GetSessionID() (value int64)
}) {
	d.SessionID = from.GetSessionID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DestroySessionOk) TypeID() uint32 {
	return DestroySessionOkTypeID
}

// TypeName returns name of type in TL schema.
func (*DestroySessionOk) TypeName() string {
	return "destroy_session_ok"
}

// TypeInfo returns info about TL type.
func (d *DestroySessionOk) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "destroy_session_ok",
		ID:   DestroySessionOkTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SessionID",
			SchemaName: "session_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DestroySessionOk) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode destroy_session_ok#e22045fc as nil")
	}
	b.PutID(DestroySessionOkTypeID)
	b.PutLong(d.SessionID)
	return nil
}

// GetSessionID returns value of SessionID field.
func (d *DestroySessionOk) GetSessionID() (value int64) {
	return d.SessionID
}

// Decode implements bin.Decoder.
func (d *DestroySessionOk) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode destroy_session_ok#e22045fc to nil")
	}
	if err := b.ConsumeID(DestroySessionOkTypeID); err != nil {
		return fmt.Errorf("unable to decode destroy_session_ok#e22045fc: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode destroy_session_ok#e22045fc: field session_id: %w", err)
		}
		d.SessionID = value
	}
	return nil
}

// construct implements constructor of DestroySessionResClass.
func (d DestroySessionOk) construct() DestroySessionResClass { return &d }

// Ensuring interfaces in compile-time for DestroySessionOk.
var (
	_ bin.Encoder = &DestroySessionOk{}
	_ bin.Decoder = &DestroySessionOk{}

	_ DestroySessionResClass = &DestroySessionOk{}
)

// DestroySessionNone represents TL type `destroy_session_none#62d350c9`.
type DestroySessionNone struct {
	// SessionID field of DestroySessionNone.
	SessionID int64
}

// DestroySessionNoneTypeID is TL type id of DestroySessionNone.
const DestroySessionNoneTypeID = 0x62d350c9

func (d *DestroySessionNone) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.SessionID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DestroySessionNone) String() string {
	if d == nil {
		return "DestroySessionNone(nil)"
	}
	type Alias DestroySessionNone
	return fmt.Sprintf("DestroySessionNone%+v", Alias(*d))
}

// FillFrom fills DestroySessionNone from given interface.
func (d *DestroySessionNone) FillFrom(from interface {
	GetSessionID() (value int64)
}) {
	d.SessionID = from.GetSessionID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DestroySessionNone) TypeID() uint32 {
	return DestroySessionNoneTypeID
}

// TypeName returns name of type in TL schema.
func (*DestroySessionNone) TypeName() string {
	return "destroy_session_none"
}

// TypeInfo returns info about TL type.
func (d *DestroySessionNone) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "destroy_session_none",
		ID:   DestroySessionNoneTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SessionID",
			SchemaName: "session_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DestroySessionNone) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode destroy_session_none#62d350c9 as nil")
	}
	b.PutID(DestroySessionNoneTypeID)
	b.PutLong(d.SessionID)
	return nil
}

// GetSessionID returns value of SessionID field.
func (d *DestroySessionNone) GetSessionID() (value int64) {
	return d.SessionID
}

// Decode implements bin.Decoder.
func (d *DestroySessionNone) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode destroy_session_none#62d350c9 to nil")
	}
	if err := b.ConsumeID(DestroySessionNoneTypeID); err != nil {
		return fmt.Errorf("unable to decode destroy_session_none#62d350c9: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode destroy_session_none#62d350c9: field session_id: %w", err)
		}
		d.SessionID = value
	}
	return nil
}

// construct implements constructor of DestroySessionResClass.
func (d DestroySessionNone) construct() DestroySessionResClass { return &d }

// Ensuring interfaces in compile-time for DestroySessionNone.
var (
	_ bin.Encoder = &DestroySessionNone{}
	_ bin.Decoder = &DestroySessionNone{}

	_ DestroySessionResClass = &DestroySessionNone{}
)

// DestroySessionResClass represents DestroySessionRes generic type.
//
// Example:
//  g, err := mt.DecodeDestroySessionRes(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *mt.DestroySessionOk: // destroy_session_ok#e22045fc
//  case *mt.DestroySessionNone: // destroy_session_none#62d350c9
//  default: panic(v)
//  }
type DestroySessionResClass interface {
	bin.Encoder
	bin.Decoder
	construct() DestroySessionResClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	// SessionID field of DestroySessionOk.
	GetSessionID() (value int64)
}

// DecodeDestroySessionRes implements binary de-serialization for DestroySessionResClass.
func DecodeDestroySessionRes(buf *bin.Buffer) (DestroySessionResClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case DestroySessionOkTypeID:
		// Decoding destroy_session_ok#e22045fc.
		v := DestroySessionOk{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DestroySessionResClass: %w", err)
		}
		return &v, nil
	case DestroySessionNoneTypeID:
		// Decoding destroy_session_none#62d350c9.
		v := DestroySessionNone{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DestroySessionResClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode DestroySessionResClass: %w", bin.NewUnexpectedID(id))
	}
}

// DestroySessionRes boxes the DestroySessionResClass providing a helper.
type DestroySessionResBox struct {
	DestroySessionRes DestroySessionResClass
}

// Decode implements bin.Decoder for DestroySessionResBox.
func (b *DestroySessionResBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode DestroySessionResBox to nil")
	}
	v, err := DecodeDestroySessionRes(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.DestroySessionRes = v
	return nil
}

// Encode implements bin.Encode for DestroySessionResBox.
func (b *DestroySessionResBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.DestroySessionRes == nil {
		return fmt.Errorf("unable to encode DestroySessionResClass as nil")
	}
	return b.DestroySessionRes.Encode(buf)
}

// DestroySessionResClassArray is adapter for slice of DestroySessionResClass.
type DestroySessionResClassArray []DestroySessionResClass

// Sort sorts slice of DestroySessionResClass.
func (s DestroySessionResClassArray) Sort(less func(a, b DestroySessionResClass) bool) DestroySessionResClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of DestroySessionResClass.
func (s DestroySessionResClassArray) SortStable(less func(a, b DestroySessionResClass) bool) DestroySessionResClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of DestroySessionResClass.
func (s DestroySessionResClassArray) Retain(keep func(x DestroySessionResClass) bool) DestroySessionResClassArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s DestroySessionResClassArray) First() (v DestroySessionResClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s DestroySessionResClassArray) Last() (v DestroySessionResClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *DestroySessionResClassArray) PopFirst() (v DestroySessionResClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero DestroySessionResClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *DestroySessionResClassArray) Pop() (v DestroySessionResClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsDestroySessionOk returns copy with only DestroySessionOk constructors.
func (s DestroySessionResClassArray) AsDestroySessionOk() (to DestroySessionOkArray) {
	for _, elem := range s {
		value, ok := elem.(*DestroySessionOk)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsDestroySessionNone returns copy with only DestroySessionNone constructors.
func (s DestroySessionResClassArray) AsDestroySessionNone() (to DestroySessionNoneArray) {
	for _, elem := range s {
		value, ok := elem.(*DestroySessionNone)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// DestroySessionOkArray is adapter for slice of DestroySessionOk.
type DestroySessionOkArray []DestroySessionOk

// Sort sorts slice of DestroySessionOk.
func (s DestroySessionOkArray) Sort(less func(a, b DestroySessionOk) bool) DestroySessionOkArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of DestroySessionOk.
func (s DestroySessionOkArray) SortStable(less func(a, b DestroySessionOk) bool) DestroySessionOkArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of DestroySessionOk.
func (s DestroySessionOkArray) Retain(keep func(x DestroySessionOk) bool) DestroySessionOkArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s DestroySessionOkArray) First() (v DestroySessionOk, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s DestroySessionOkArray) Last() (v DestroySessionOk, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *DestroySessionOkArray) PopFirst() (v DestroySessionOk, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero DestroySessionOk
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *DestroySessionOkArray) Pop() (v DestroySessionOk, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// DestroySessionNoneArray is adapter for slice of DestroySessionNone.
type DestroySessionNoneArray []DestroySessionNone

// Sort sorts slice of DestroySessionNone.
func (s DestroySessionNoneArray) Sort(less func(a, b DestroySessionNone) bool) DestroySessionNoneArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of DestroySessionNone.
func (s DestroySessionNoneArray) SortStable(less func(a, b DestroySessionNone) bool) DestroySessionNoneArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of DestroySessionNone.
func (s DestroySessionNoneArray) Retain(keep func(x DestroySessionNone) bool) DestroySessionNoneArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s DestroySessionNoneArray) First() (v DestroySessionNone, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s DestroySessionNoneArray) Last() (v DestroySessionNone, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *DestroySessionNoneArray) PopFirst() (v DestroySessionNone, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero DestroySessionNone
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *DestroySessionNoneArray) Pop() (v DestroySessionNone, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
