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

// SecureRequiredType represents TL type `secureRequiredType#829d99da`.
// Required type
//
// See https://core.telegram.org/constructor/secureRequiredType for reference.
type SecureRequiredType struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Native names
	NativeNames bool
	// Is a selfie required
	SelfieRequired bool
	// Is a translation required
	TranslationRequired bool
	// Secure value type
	Type SecureValueTypeClass
}

// SecureRequiredTypeTypeID is TL type id of SecureRequiredType.
const SecureRequiredTypeTypeID = 0x829d99da

func (s *SecureRequiredType) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.NativeNames == false) {
		return false
	}
	if !(s.SelfieRequired == false) {
		return false
	}
	if !(s.TranslationRequired == false) {
		return false
	}
	if !(s.Type == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SecureRequiredType) String() string {
	if s == nil {
		return "SecureRequiredType(nil)"
	}
	type Alias SecureRequiredType
	return fmt.Sprintf("SecureRequiredType%+v", Alias(*s))
}

// FillFrom fills SecureRequiredType from given interface.
func (s *SecureRequiredType) FillFrom(from interface {
	GetNativeNames() (value bool)
	GetSelfieRequired() (value bool)
	GetTranslationRequired() (value bool)
	GetType() (value SecureValueTypeClass)
}) {
	s.NativeNames = from.GetNativeNames()
	s.SelfieRequired = from.GetSelfieRequired()
	s.TranslationRequired = from.GetTranslationRequired()
	s.Type = from.GetType()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SecureRequiredType) TypeID() uint32 {
	return SecureRequiredTypeTypeID
}

// TypeName returns name of type in TL schema.
func (*SecureRequiredType) TypeName() string {
	return "secureRequiredType"
}

// TypeInfo returns info about TL type.
func (s *SecureRequiredType) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "secureRequiredType",
		ID:   SecureRequiredTypeTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "NativeNames",
			SchemaName: "native_names",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "SelfieRequired",
			SchemaName: "selfie_required",
			Null:       !s.Flags.Has(1),
		},
		{
			Name:       "TranslationRequired",
			SchemaName: "translation_required",
			Null:       !s.Flags.Has(2),
		},
		{
			Name:       "Type",
			SchemaName: "type",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SecureRequiredType) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureRequiredType#829d99da as nil")
	}
	b.PutID(SecureRequiredTypeTypeID)
	if !(s.NativeNames == false) {
		s.Flags.Set(0)
	}
	if !(s.SelfieRequired == false) {
		s.Flags.Set(1)
	}
	if !(s.TranslationRequired == false) {
		s.Flags.Set(2)
	}
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode secureRequiredType#829d99da: field flags: %w", err)
	}
	if s.Type == nil {
		return fmt.Errorf("unable to encode secureRequiredType#829d99da: field type is nil")
	}
	if err := s.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode secureRequiredType#829d99da: field type: %w", err)
	}
	return nil
}

// SetNativeNames sets value of NativeNames conditional field.
func (s *SecureRequiredType) SetNativeNames(value bool) {
	if value {
		s.Flags.Set(0)
		s.NativeNames = true
	} else {
		s.Flags.Unset(0)
		s.NativeNames = false
	}
}

// GetNativeNames returns value of NativeNames conditional field.
func (s *SecureRequiredType) GetNativeNames() (value bool) {
	return s.Flags.Has(0)
}

// SetSelfieRequired sets value of SelfieRequired conditional field.
func (s *SecureRequiredType) SetSelfieRequired(value bool) {
	if value {
		s.Flags.Set(1)
		s.SelfieRequired = true
	} else {
		s.Flags.Unset(1)
		s.SelfieRequired = false
	}
}

// GetSelfieRequired returns value of SelfieRequired conditional field.
func (s *SecureRequiredType) GetSelfieRequired() (value bool) {
	return s.Flags.Has(1)
}

// SetTranslationRequired sets value of TranslationRequired conditional field.
func (s *SecureRequiredType) SetTranslationRequired(value bool) {
	if value {
		s.Flags.Set(2)
		s.TranslationRequired = true
	} else {
		s.Flags.Unset(2)
		s.TranslationRequired = false
	}
}

// GetTranslationRequired returns value of TranslationRequired conditional field.
func (s *SecureRequiredType) GetTranslationRequired() (value bool) {
	return s.Flags.Has(2)
}

// GetType returns value of Type field.
func (s *SecureRequiredType) GetType() (value SecureValueTypeClass) {
	return s.Type
}

// Decode implements bin.Decoder.
func (s *SecureRequiredType) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureRequiredType#829d99da to nil")
	}
	if err := b.ConsumeID(SecureRequiredTypeTypeID); err != nil {
		return fmt.Errorf("unable to decode secureRequiredType#829d99da: %w", err)
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode secureRequiredType#829d99da: field flags: %w", err)
		}
	}
	s.NativeNames = s.Flags.Has(0)
	s.SelfieRequired = s.Flags.Has(1)
	s.TranslationRequired = s.Flags.Has(2)
	{
		value, err := DecodeSecureValueType(b)
		if err != nil {
			return fmt.Errorf("unable to decode secureRequiredType#829d99da: field type: %w", err)
		}
		s.Type = value
	}
	return nil
}

// construct implements constructor of SecureRequiredTypeClass.
func (s SecureRequiredType) construct() SecureRequiredTypeClass { return &s }

// Ensuring interfaces in compile-time for SecureRequiredType.
var (
	_ bin.Encoder = &SecureRequiredType{}
	_ bin.Decoder = &SecureRequiredType{}

	_ SecureRequiredTypeClass = &SecureRequiredType{}
)

// SecureRequiredTypeOneOf represents TL type `secureRequiredTypeOneOf#27477b4`.
// One of
//
// See https://core.telegram.org/constructor/secureRequiredTypeOneOf for reference.
type SecureRequiredTypeOneOf struct {
	// Secure required value types
	Types []SecureRequiredTypeClass
}

// SecureRequiredTypeOneOfTypeID is TL type id of SecureRequiredTypeOneOf.
const SecureRequiredTypeOneOfTypeID = 0x27477b4

func (s *SecureRequiredTypeOneOf) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Types == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SecureRequiredTypeOneOf) String() string {
	if s == nil {
		return "SecureRequiredTypeOneOf(nil)"
	}
	type Alias SecureRequiredTypeOneOf
	return fmt.Sprintf("SecureRequiredTypeOneOf%+v", Alias(*s))
}

// FillFrom fills SecureRequiredTypeOneOf from given interface.
func (s *SecureRequiredTypeOneOf) FillFrom(from interface {
	GetTypes() (value []SecureRequiredTypeClass)
}) {
	s.Types = from.GetTypes()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SecureRequiredTypeOneOf) TypeID() uint32 {
	return SecureRequiredTypeOneOfTypeID
}

// TypeName returns name of type in TL schema.
func (*SecureRequiredTypeOneOf) TypeName() string {
	return "secureRequiredTypeOneOf"
}

// TypeInfo returns info about TL type.
func (s *SecureRequiredTypeOneOf) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "secureRequiredTypeOneOf",
		ID:   SecureRequiredTypeOneOfTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Types",
			SchemaName: "types",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SecureRequiredTypeOneOf) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureRequiredTypeOneOf#27477b4 as nil")
	}
	b.PutID(SecureRequiredTypeOneOfTypeID)
	b.PutVectorHeader(len(s.Types))
	for idx, v := range s.Types {
		if v == nil {
			return fmt.Errorf("unable to encode secureRequiredTypeOneOf#27477b4: field types element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode secureRequiredTypeOneOf#27477b4: field types element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetTypes returns value of Types field.
func (s *SecureRequiredTypeOneOf) GetTypes() (value []SecureRequiredTypeClass) {
	return s.Types
}

// MapTypes returns field Types wrapped in SecureRequiredTypeClassArray helper.
func (s *SecureRequiredTypeOneOf) MapTypes() (value SecureRequiredTypeClassArray) {
	return SecureRequiredTypeClassArray(s.Types)
}

// Decode implements bin.Decoder.
func (s *SecureRequiredTypeOneOf) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureRequiredTypeOneOf#27477b4 to nil")
	}
	if err := b.ConsumeID(SecureRequiredTypeOneOfTypeID); err != nil {
		return fmt.Errorf("unable to decode secureRequiredTypeOneOf#27477b4: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode secureRequiredTypeOneOf#27477b4: field types: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeSecureRequiredType(b)
			if err != nil {
				return fmt.Errorf("unable to decode secureRequiredTypeOneOf#27477b4: field types: %w", err)
			}
			s.Types = append(s.Types, value)
		}
	}
	return nil
}

// construct implements constructor of SecureRequiredTypeClass.
func (s SecureRequiredTypeOneOf) construct() SecureRequiredTypeClass { return &s }

// Ensuring interfaces in compile-time for SecureRequiredTypeOneOf.
var (
	_ bin.Encoder = &SecureRequiredTypeOneOf{}
	_ bin.Decoder = &SecureRequiredTypeOneOf{}

	_ SecureRequiredTypeClass = &SecureRequiredTypeOneOf{}
)

// SecureRequiredTypeClass represents SecureRequiredType generic type.
//
// See https://core.telegram.org/type/SecureRequiredType for reference.
//
// Example:
//  g, err := tg.DecodeSecureRequiredType(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.SecureRequiredType: // secureRequiredType#829d99da
//  case *tg.SecureRequiredTypeOneOf: // secureRequiredTypeOneOf#27477b4
//  default: panic(v)
//  }
type SecureRequiredTypeClass interface {
	bin.Encoder
	bin.Decoder
	construct() SecureRequiredTypeClass

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
}

// DecodeSecureRequiredType implements binary de-serialization for SecureRequiredTypeClass.
func DecodeSecureRequiredType(buf *bin.Buffer) (SecureRequiredTypeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case SecureRequiredTypeTypeID:
		// Decoding secureRequiredType#829d99da.
		v := SecureRequiredType{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SecureRequiredTypeClass: %w", err)
		}
		return &v, nil
	case SecureRequiredTypeOneOfTypeID:
		// Decoding secureRequiredTypeOneOf#27477b4.
		v := SecureRequiredTypeOneOf{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SecureRequiredTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode SecureRequiredTypeClass: %w", bin.NewUnexpectedID(id))
	}
}

// SecureRequiredType boxes the SecureRequiredTypeClass providing a helper.
type SecureRequiredTypeBox struct {
	SecureRequiredType SecureRequiredTypeClass
}

// Decode implements bin.Decoder for SecureRequiredTypeBox.
func (b *SecureRequiredTypeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode SecureRequiredTypeBox to nil")
	}
	v, err := DecodeSecureRequiredType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.SecureRequiredType = v
	return nil
}

// Encode implements bin.Encode for SecureRequiredTypeBox.
func (b *SecureRequiredTypeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.SecureRequiredType == nil {
		return fmt.Errorf("unable to encode SecureRequiredTypeClass as nil")
	}
	return b.SecureRequiredType.Encode(buf)
}

// SecureRequiredTypeClassArray is adapter for slice of SecureRequiredTypeClass.
type SecureRequiredTypeClassArray []SecureRequiredTypeClass

// Sort sorts slice of SecureRequiredTypeClass.
func (s SecureRequiredTypeClassArray) Sort(less func(a, b SecureRequiredTypeClass) bool) SecureRequiredTypeClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of SecureRequiredTypeClass.
func (s SecureRequiredTypeClassArray) SortStable(less func(a, b SecureRequiredTypeClass) bool) SecureRequiredTypeClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of SecureRequiredTypeClass.
func (s SecureRequiredTypeClassArray) Retain(keep func(x SecureRequiredTypeClass) bool) SecureRequiredTypeClassArray {
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
func (s SecureRequiredTypeClassArray) First() (v SecureRequiredTypeClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s SecureRequiredTypeClassArray) Last() (v SecureRequiredTypeClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *SecureRequiredTypeClassArray) PopFirst() (v SecureRequiredTypeClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero SecureRequiredTypeClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *SecureRequiredTypeClassArray) Pop() (v SecureRequiredTypeClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsSecureRequiredType returns copy with only SecureRequiredType constructors.
func (s SecureRequiredTypeClassArray) AsSecureRequiredType() (to SecureRequiredTypeArray) {
	for _, elem := range s {
		value, ok := elem.(*SecureRequiredType)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsSecureRequiredTypeOneOf returns copy with only SecureRequiredTypeOneOf constructors.
func (s SecureRequiredTypeClassArray) AsSecureRequiredTypeOneOf() (to SecureRequiredTypeOneOfArray) {
	for _, elem := range s {
		value, ok := elem.(*SecureRequiredTypeOneOf)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// SecureRequiredTypeArray is adapter for slice of SecureRequiredType.
type SecureRequiredTypeArray []SecureRequiredType

// Sort sorts slice of SecureRequiredType.
func (s SecureRequiredTypeArray) Sort(less func(a, b SecureRequiredType) bool) SecureRequiredTypeArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of SecureRequiredType.
func (s SecureRequiredTypeArray) SortStable(less func(a, b SecureRequiredType) bool) SecureRequiredTypeArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of SecureRequiredType.
func (s SecureRequiredTypeArray) Retain(keep func(x SecureRequiredType) bool) SecureRequiredTypeArray {
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
func (s SecureRequiredTypeArray) First() (v SecureRequiredType, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s SecureRequiredTypeArray) Last() (v SecureRequiredType, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *SecureRequiredTypeArray) PopFirst() (v SecureRequiredType, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero SecureRequiredType
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *SecureRequiredTypeArray) Pop() (v SecureRequiredType, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SecureRequiredTypeOneOfArray is adapter for slice of SecureRequiredTypeOneOf.
type SecureRequiredTypeOneOfArray []SecureRequiredTypeOneOf

// Sort sorts slice of SecureRequiredTypeOneOf.
func (s SecureRequiredTypeOneOfArray) Sort(less func(a, b SecureRequiredTypeOneOf) bool) SecureRequiredTypeOneOfArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of SecureRequiredTypeOneOf.
func (s SecureRequiredTypeOneOfArray) SortStable(less func(a, b SecureRequiredTypeOneOf) bool) SecureRequiredTypeOneOfArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of SecureRequiredTypeOneOf.
func (s SecureRequiredTypeOneOfArray) Retain(keep func(x SecureRequiredTypeOneOf) bool) SecureRequiredTypeOneOfArray {
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
func (s SecureRequiredTypeOneOfArray) First() (v SecureRequiredTypeOneOf, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s SecureRequiredTypeOneOfArray) Last() (v SecureRequiredTypeOneOf, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *SecureRequiredTypeOneOfArray) PopFirst() (v SecureRequiredTypeOneOf, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero SecureRequiredTypeOneOf
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *SecureRequiredTypeOneOfArray) Pop() (v SecureRequiredTypeOneOf, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
