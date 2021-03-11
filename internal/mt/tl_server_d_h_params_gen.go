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

// ServerDHParamsFail represents TL type `server_DH_params_fail#79cb045d`.
type ServerDHParamsFail struct {
	// Nonce field of ServerDHParamsFail.
	Nonce bin.Int128
	// ServerNonce field of ServerDHParamsFail.
	ServerNonce bin.Int128
	// NewNonceHash field of ServerDHParamsFail.
	NewNonceHash bin.Int128
}

// ServerDHParamsFailTypeID is TL type id of ServerDHParamsFail.
const ServerDHParamsFailTypeID = 0x79cb045d

func (s *ServerDHParamsFail) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Nonce == bin.Int128{}) {
		return false
	}
	if !(s.ServerNonce == bin.Int128{}) {
		return false
	}
	if !(s.NewNonceHash == bin.Int128{}) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *ServerDHParamsFail) String() string {
	if s == nil {
		return "ServerDHParamsFail(nil)"
	}
	type Alias ServerDHParamsFail
	return fmt.Sprintf("ServerDHParamsFail%+v", Alias(*s))
}

// FillFrom fills ServerDHParamsFail from given interface.
func (s *ServerDHParamsFail) FillFrom(from interface {
	GetNonce() (value bin.Int128)
	GetServerNonce() (value bin.Int128)
	GetNewNonceHash() (value bin.Int128)
}) {
	s.Nonce = from.GetNonce()
	s.ServerNonce = from.GetServerNonce()
	s.NewNonceHash = from.GetNewNonceHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ServerDHParamsFail) TypeID() uint32 {
	return ServerDHParamsFailTypeID
}

// TypeName returns name of type in TL schema.
func (*ServerDHParamsFail) TypeName() string {
	return "server_DH_params_fail"
}

// TypeInfo returns info about TL type.
func (s *ServerDHParamsFail) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "server_DH_params_fail",
		ID:   ServerDHParamsFailTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Nonce",
			SchemaName: "nonce",
		},
		{
			Name:       "ServerNonce",
			SchemaName: "server_nonce",
		},
		{
			Name:       "NewNonceHash",
			SchemaName: "new_nonce_hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *ServerDHParamsFail) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode server_DH_params_fail#79cb045d as nil")
	}
	b.PutID(ServerDHParamsFailTypeID)
	b.PutInt128(s.Nonce)
	b.PutInt128(s.ServerNonce)
	b.PutInt128(s.NewNonceHash)
	return nil
}

// GetNonce returns value of Nonce field.
func (s *ServerDHParamsFail) GetNonce() (value bin.Int128) {
	return s.Nonce
}

// GetServerNonce returns value of ServerNonce field.
func (s *ServerDHParamsFail) GetServerNonce() (value bin.Int128) {
	return s.ServerNonce
}

// GetNewNonceHash returns value of NewNonceHash field.
func (s *ServerDHParamsFail) GetNewNonceHash() (value bin.Int128) {
	return s.NewNonceHash
}

// Decode implements bin.Decoder.
func (s *ServerDHParamsFail) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode server_DH_params_fail#79cb045d to nil")
	}
	if err := b.ConsumeID(ServerDHParamsFailTypeID); err != nil {
		return fmt.Errorf("unable to decode server_DH_params_fail#79cb045d: %w", err)
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode server_DH_params_fail#79cb045d: field nonce: %w", err)
		}
		s.Nonce = value
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode server_DH_params_fail#79cb045d: field server_nonce: %w", err)
		}
		s.ServerNonce = value
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode server_DH_params_fail#79cb045d: field new_nonce_hash: %w", err)
		}
		s.NewNonceHash = value
	}
	return nil
}

// construct implements constructor of ServerDHParamsClass.
func (s ServerDHParamsFail) construct() ServerDHParamsClass { return &s }

// Ensuring interfaces in compile-time for ServerDHParamsFail.
var (
	_ bin.Encoder = &ServerDHParamsFail{}
	_ bin.Decoder = &ServerDHParamsFail{}

	_ ServerDHParamsClass = &ServerDHParamsFail{}
)

// ServerDHParamsOk represents TL type `server_DH_params_ok#d0e8075c`.
type ServerDHParamsOk struct {
	// Nonce field of ServerDHParamsOk.
	Nonce bin.Int128
	// ServerNonce field of ServerDHParamsOk.
	ServerNonce bin.Int128
	// EncryptedAnswer field of ServerDHParamsOk.
	EncryptedAnswer []byte
}

// ServerDHParamsOkTypeID is TL type id of ServerDHParamsOk.
const ServerDHParamsOkTypeID = 0xd0e8075c

func (s *ServerDHParamsOk) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Nonce == bin.Int128{}) {
		return false
	}
	if !(s.ServerNonce == bin.Int128{}) {
		return false
	}
	if !(s.EncryptedAnswer == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *ServerDHParamsOk) String() string {
	if s == nil {
		return "ServerDHParamsOk(nil)"
	}
	type Alias ServerDHParamsOk
	return fmt.Sprintf("ServerDHParamsOk%+v", Alias(*s))
}

// FillFrom fills ServerDHParamsOk from given interface.
func (s *ServerDHParamsOk) FillFrom(from interface {
	GetNonce() (value bin.Int128)
	GetServerNonce() (value bin.Int128)
	GetEncryptedAnswer() (value []byte)
}) {
	s.Nonce = from.GetNonce()
	s.ServerNonce = from.GetServerNonce()
	s.EncryptedAnswer = from.GetEncryptedAnswer()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ServerDHParamsOk) TypeID() uint32 {
	return ServerDHParamsOkTypeID
}

// TypeName returns name of type in TL schema.
func (*ServerDHParamsOk) TypeName() string {
	return "server_DH_params_ok"
}

// TypeInfo returns info about TL type.
func (s *ServerDHParamsOk) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "server_DH_params_ok",
		ID:   ServerDHParamsOkTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Nonce",
			SchemaName: "nonce",
		},
		{
			Name:       "ServerNonce",
			SchemaName: "server_nonce",
		},
		{
			Name:       "EncryptedAnswer",
			SchemaName: "encrypted_answer",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *ServerDHParamsOk) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode server_DH_params_ok#d0e8075c as nil")
	}
	b.PutID(ServerDHParamsOkTypeID)
	b.PutInt128(s.Nonce)
	b.PutInt128(s.ServerNonce)
	b.PutBytes(s.EncryptedAnswer)
	return nil
}

// GetNonce returns value of Nonce field.
func (s *ServerDHParamsOk) GetNonce() (value bin.Int128) {
	return s.Nonce
}

// GetServerNonce returns value of ServerNonce field.
func (s *ServerDHParamsOk) GetServerNonce() (value bin.Int128) {
	return s.ServerNonce
}

// GetEncryptedAnswer returns value of EncryptedAnswer field.
func (s *ServerDHParamsOk) GetEncryptedAnswer() (value []byte) {
	return s.EncryptedAnswer
}

// Decode implements bin.Decoder.
func (s *ServerDHParamsOk) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode server_DH_params_ok#d0e8075c to nil")
	}
	if err := b.ConsumeID(ServerDHParamsOkTypeID); err != nil {
		return fmt.Errorf("unable to decode server_DH_params_ok#d0e8075c: %w", err)
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode server_DH_params_ok#d0e8075c: field nonce: %w", err)
		}
		s.Nonce = value
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode server_DH_params_ok#d0e8075c: field server_nonce: %w", err)
		}
		s.ServerNonce = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode server_DH_params_ok#d0e8075c: field encrypted_answer: %w", err)
		}
		s.EncryptedAnswer = value
	}
	return nil
}

// construct implements constructor of ServerDHParamsClass.
func (s ServerDHParamsOk) construct() ServerDHParamsClass { return &s }

// Ensuring interfaces in compile-time for ServerDHParamsOk.
var (
	_ bin.Encoder = &ServerDHParamsOk{}
	_ bin.Decoder = &ServerDHParamsOk{}

	_ ServerDHParamsClass = &ServerDHParamsOk{}
)

// ServerDHParamsClass represents Server_DH_Params generic type.
//
// Example:
//  g, err := mt.DecodeServerDHParams(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *mt.ServerDHParamsFail: // server_DH_params_fail#79cb045d
//  case *mt.ServerDHParamsOk: // server_DH_params_ok#d0e8075c
//  default: panic(v)
//  }
type ServerDHParamsClass interface {
	bin.Encoder
	bin.Decoder
	construct() ServerDHParamsClass

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

	// Nonce field of ServerDHParamsFail.
	GetNonce() (value bin.Int128)
	// ServerNonce field of ServerDHParamsFail.
	GetServerNonce() (value bin.Int128)
}

// DecodeServerDHParams implements binary de-serialization for ServerDHParamsClass.
func DecodeServerDHParams(buf *bin.Buffer) (ServerDHParamsClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ServerDHParamsFailTypeID:
		// Decoding server_DH_params_fail#79cb045d.
		v := ServerDHParamsFail{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ServerDHParamsClass: %w", err)
		}
		return &v, nil
	case ServerDHParamsOkTypeID:
		// Decoding server_DH_params_ok#d0e8075c.
		v := ServerDHParamsOk{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ServerDHParamsClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ServerDHParamsClass: %w", bin.NewUnexpectedID(id))
	}
}

// ServerDHParams boxes the ServerDHParamsClass providing a helper.
type ServerDHParamsBox struct {
	Server_DH_Params ServerDHParamsClass
}

// Decode implements bin.Decoder for ServerDHParamsBox.
func (b *ServerDHParamsBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ServerDHParamsBox to nil")
	}
	v, err := DecodeServerDHParams(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.Server_DH_Params = v
	return nil
}

// Encode implements bin.Encode for ServerDHParamsBox.
func (b *ServerDHParamsBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.Server_DH_Params == nil {
		return fmt.Errorf("unable to encode ServerDHParamsClass as nil")
	}
	return b.Server_DH_Params.Encode(buf)
}

// ServerDHParamsClassArray is adapter for slice of ServerDHParamsClass.
type ServerDHParamsClassArray []ServerDHParamsClass

// Sort sorts slice of ServerDHParamsClass.
func (s ServerDHParamsClassArray) Sort(less func(a, b ServerDHParamsClass) bool) ServerDHParamsClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ServerDHParamsClass.
func (s ServerDHParamsClassArray) SortStable(less func(a, b ServerDHParamsClass) bool) ServerDHParamsClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ServerDHParamsClass.
func (s ServerDHParamsClassArray) Retain(keep func(x ServerDHParamsClass) bool) ServerDHParamsClassArray {
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
func (s ServerDHParamsClassArray) First() (v ServerDHParamsClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ServerDHParamsClassArray) Last() (v ServerDHParamsClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ServerDHParamsClassArray) PopFirst() (v ServerDHParamsClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ServerDHParamsClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ServerDHParamsClassArray) Pop() (v ServerDHParamsClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsServerDHParamsFail returns copy with only ServerDHParamsFail constructors.
func (s ServerDHParamsClassArray) AsServerDHParamsFail() (to ServerDHParamsFailArray) {
	for _, elem := range s {
		value, ok := elem.(*ServerDHParamsFail)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsServerDHParamsOk returns copy with only ServerDHParamsOk constructors.
func (s ServerDHParamsClassArray) AsServerDHParamsOk() (to ServerDHParamsOkArray) {
	for _, elem := range s {
		value, ok := elem.(*ServerDHParamsOk)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// ServerDHParamsFailArray is adapter for slice of ServerDHParamsFail.
type ServerDHParamsFailArray []ServerDHParamsFail

// Sort sorts slice of ServerDHParamsFail.
func (s ServerDHParamsFailArray) Sort(less func(a, b ServerDHParamsFail) bool) ServerDHParamsFailArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ServerDHParamsFail.
func (s ServerDHParamsFailArray) SortStable(less func(a, b ServerDHParamsFail) bool) ServerDHParamsFailArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ServerDHParamsFail.
func (s ServerDHParamsFailArray) Retain(keep func(x ServerDHParamsFail) bool) ServerDHParamsFailArray {
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
func (s ServerDHParamsFailArray) First() (v ServerDHParamsFail, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ServerDHParamsFailArray) Last() (v ServerDHParamsFail, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ServerDHParamsFailArray) PopFirst() (v ServerDHParamsFail, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ServerDHParamsFail
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ServerDHParamsFailArray) Pop() (v ServerDHParamsFail, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// ServerDHParamsOkArray is adapter for slice of ServerDHParamsOk.
type ServerDHParamsOkArray []ServerDHParamsOk

// Sort sorts slice of ServerDHParamsOk.
func (s ServerDHParamsOkArray) Sort(less func(a, b ServerDHParamsOk) bool) ServerDHParamsOkArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ServerDHParamsOk.
func (s ServerDHParamsOkArray) SortStable(less func(a, b ServerDHParamsOk) bool) ServerDHParamsOkArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ServerDHParamsOk.
func (s ServerDHParamsOkArray) Retain(keep func(x ServerDHParamsOk) bool) ServerDHParamsOkArray {
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
func (s ServerDHParamsOkArray) First() (v ServerDHParamsOk, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ServerDHParamsOkArray) Last() (v ServerDHParamsOk, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ServerDHParamsOkArray) PopFirst() (v ServerDHParamsOk, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ServerDHParamsOk
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ServerDHParamsOkArray) Pop() (v ServerDHParamsOk, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
