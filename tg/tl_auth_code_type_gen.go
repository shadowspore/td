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

// AuthCodeTypeSMS represents TL type `auth.codeTypeSms#72a3158c`.
// Type of verification code that will be sent next if you call the resendCode method: SMS code
//
// See https://core.telegram.org/constructor/auth.codeTypeSms for reference.
type AuthCodeTypeSMS struct {
}

// AuthCodeTypeSMSTypeID is TL type id of AuthCodeTypeSMS.
const AuthCodeTypeSMSTypeID = 0x72a3158c

func (c *AuthCodeTypeSMS) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *AuthCodeTypeSMS) String() string {
	if c == nil {
		return "AuthCodeTypeSMS(nil)"
	}
	type Alias AuthCodeTypeSMS
	return fmt.Sprintf("AuthCodeTypeSMS%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthCodeTypeSMS) TypeID() uint32 {
	return AuthCodeTypeSMSTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthCodeTypeSMS) TypeName() string {
	return "auth.codeTypeSms"
}

// TypeInfo returns info about TL type.
func (c *AuthCodeTypeSMS) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "auth.codeTypeSms",
		ID:   AuthCodeTypeSMSTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *AuthCodeTypeSMS) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode auth.codeTypeSms#72a3158c as nil")
	}
	b.PutID(AuthCodeTypeSMSTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (c *AuthCodeTypeSMS) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode auth.codeTypeSms#72a3158c to nil")
	}
	if err := b.ConsumeID(AuthCodeTypeSMSTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.codeTypeSms#72a3158c: %w", err)
	}
	return nil
}

// construct implements constructor of AuthCodeTypeClass.
func (c AuthCodeTypeSMS) construct() AuthCodeTypeClass { return &c }

// Ensuring interfaces in compile-time for AuthCodeTypeSMS.
var (
	_ bin.Encoder = &AuthCodeTypeSMS{}
	_ bin.Decoder = &AuthCodeTypeSMS{}

	_ AuthCodeTypeClass = &AuthCodeTypeSMS{}
)

// AuthCodeTypeCall represents TL type `auth.codeTypeCall#741cd3e3`.
// Type of verification code that will be sent next if you call the resendCode method: SMS code
//
// See https://core.telegram.org/constructor/auth.codeTypeCall for reference.
type AuthCodeTypeCall struct {
}

// AuthCodeTypeCallTypeID is TL type id of AuthCodeTypeCall.
const AuthCodeTypeCallTypeID = 0x741cd3e3

func (c *AuthCodeTypeCall) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *AuthCodeTypeCall) String() string {
	if c == nil {
		return "AuthCodeTypeCall(nil)"
	}
	type Alias AuthCodeTypeCall
	return fmt.Sprintf("AuthCodeTypeCall%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthCodeTypeCall) TypeID() uint32 {
	return AuthCodeTypeCallTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthCodeTypeCall) TypeName() string {
	return "auth.codeTypeCall"
}

// TypeInfo returns info about TL type.
func (c *AuthCodeTypeCall) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "auth.codeTypeCall",
		ID:   AuthCodeTypeCallTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *AuthCodeTypeCall) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode auth.codeTypeCall#741cd3e3 as nil")
	}
	b.PutID(AuthCodeTypeCallTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (c *AuthCodeTypeCall) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode auth.codeTypeCall#741cd3e3 to nil")
	}
	if err := b.ConsumeID(AuthCodeTypeCallTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.codeTypeCall#741cd3e3: %w", err)
	}
	return nil
}

// construct implements constructor of AuthCodeTypeClass.
func (c AuthCodeTypeCall) construct() AuthCodeTypeClass { return &c }

// Ensuring interfaces in compile-time for AuthCodeTypeCall.
var (
	_ bin.Encoder = &AuthCodeTypeCall{}
	_ bin.Decoder = &AuthCodeTypeCall{}

	_ AuthCodeTypeClass = &AuthCodeTypeCall{}
)

// AuthCodeTypeFlashCall represents TL type `auth.codeTypeFlashCall#226ccefb`.
// Type of verification code that will be sent next if you call the resendCode method: SMS code
//
// See https://core.telegram.org/constructor/auth.codeTypeFlashCall for reference.
type AuthCodeTypeFlashCall struct {
}

// AuthCodeTypeFlashCallTypeID is TL type id of AuthCodeTypeFlashCall.
const AuthCodeTypeFlashCallTypeID = 0x226ccefb

func (c *AuthCodeTypeFlashCall) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *AuthCodeTypeFlashCall) String() string {
	if c == nil {
		return "AuthCodeTypeFlashCall(nil)"
	}
	type Alias AuthCodeTypeFlashCall
	return fmt.Sprintf("AuthCodeTypeFlashCall%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthCodeTypeFlashCall) TypeID() uint32 {
	return AuthCodeTypeFlashCallTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthCodeTypeFlashCall) TypeName() string {
	return "auth.codeTypeFlashCall"
}

// TypeInfo returns info about TL type.
func (c *AuthCodeTypeFlashCall) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "auth.codeTypeFlashCall",
		ID:   AuthCodeTypeFlashCallTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *AuthCodeTypeFlashCall) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode auth.codeTypeFlashCall#226ccefb as nil")
	}
	b.PutID(AuthCodeTypeFlashCallTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (c *AuthCodeTypeFlashCall) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode auth.codeTypeFlashCall#226ccefb to nil")
	}
	if err := b.ConsumeID(AuthCodeTypeFlashCallTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.codeTypeFlashCall#226ccefb: %w", err)
	}
	return nil
}

// construct implements constructor of AuthCodeTypeClass.
func (c AuthCodeTypeFlashCall) construct() AuthCodeTypeClass { return &c }

// Ensuring interfaces in compile-time for AuthCodeTypeFlashCall.
var (
	_ bin.Encoder = &AuthCodeTypeFlashCall{}
	_ bin.Decoder = &AuthCodeTypeFlashCall{}

	_ AuthCodeTypeClass = &AuthCodeTypeFlashCall{}
)

// AuthCodeTypeClass represents auth.CodeType generic type.
//
// See https://core.telegram.org/type/auth.CodeType for reference.
//
// Example:
//  g, err := tg.DecodeAuthCodeType(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.AuthCodeTypeSMS: // auth.codeTypeSms#72a3158c
//  case *tg.AuthCodeTypeCall: // auth.codeTypeCall#741cd3e3
//  case *tg.AuthCodeTypeFlashCall: // auth.codeTypeFlashCall#226ccefb
//  default: panic(v)
//  }
type AuthCodeTypeClass interface {
	bin.Encoder
	bin.Decoder
	construct() AuthCodeTypeClass

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

// DecodeAuthCodeType implements binary de-serialization for AuthCodeTypeClass.
func DecodeAuthCodeType(buf *bin.Buffer) (AuthCodeTypeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case AuthCodeTypeSMSTypeID:
		// Decoding auth.codeTypeSms#72a3158c.
		v := AuthCodeTypeSMS{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthCodeTypeClass: %w", err)
		}
		return &v, nil
	case AuthCodeTypeCallTypeID:
		// Decoding auth.codeTypeCall#741cd3e3.
		v := AuthCodeTypeCall{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthCodeTypeClass: %w", err)
		}
		return &v, nil
	case AuthCodeTypeFlashCallTypeID:
		// Decoding auth.codeTypeFlashCall#226ccefb.
		v := AuthCodeTypeFlashCall{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthCodeTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode AuthCodeTypeClass: %w", bin.NewUnexpectedID(id))
	}
}

// AuthCodeType boxes the AuthCodeTypeClass providing a helper.
type AuthCodeTypeBox struct {
	CodeType AuthCodeTypeClass
}

// Decode implements bin.Decoder for AuthCodeTypeBox.
func (b *AuthCodeTypeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode AuthCodeTypeBox to nil")
	}
	v, err := DecodeAuthCodeType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.CodeType = v
	return nil
}

// Encode implements bin.Encode for AuthCodeTypeBox.
func (b *AuthCodeTypeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.CodeType == nil {
		return fmt.Errorf("unable to encode AuthCodeTypeClass as nil")
	}
	return b.CodeType.Encode(buf)
}

// AuthCodeTypeClassArray is adapter for slice of AuthCodeTypeClass.
type AuthCodeTypeClassArray []AuthCodeTypeClass

// Sort sorts slice of AuthCodeTypeClass.
func (s AuthCodeTypeClassArray) Sort(less func(a, b AuthCodeTypeClass) bool) AuthCodeTypeClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of AuthCodeTypeClass.
func (s AuthCodeTypeClassArray) SortStable(less func(a, b AuthCodeTypeClass) bool) AuthCodeTypeClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of AuthCodeTypeClass.
func (s AuthCodeTypeClassArray) Retain(keep func(x AuthCodeTypeClass) bool) AuthCodeTypeClassArray {
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
func (s AuthCodeTypeClassArray) First() (v AuthCodeTypeClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s AuthCodeTypeClassArray) Last() (v AuthCodeTypeClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *AuthCodeTypeClassArray) PopFirst() (v AuthCodeTypeClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero AuthCodeTypeClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *AuthCodeTypeClassArray) Pop() (v AuthCodeTypeClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
