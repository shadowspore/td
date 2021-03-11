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

// AuthExportedAuthorization represents TL type `auth.exportedAuthorization#df969c2d`.
// Data for copying of authorization between data centres.
//
// See https://core.telegram.org/constructor/auth.exportedAuthorization for reference.
type AuthExportedAuthorization struct {
	// current user identifier
	ID int
	// authorizes key
	Bytes []byte
}

// AuthExportedAuthorizationTypeID is TL type id of AuthExportedAuthorization.
const AuthExportedAuthorizationTypeID = 0xdf969c2d

func (e *AuthExportedAuthorization) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.ID == 0) {
		return false
	}
	if !(e.Bytes == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *AuthExportedAuthorization) String() string {
	if e == nil {
		return "AuthExportedAuthorization(nil)"
	}
	type Alias AuthExportedAuthorization
	return fmt.Sprintf("AuthExportedAuthorization%+v", Alias(*e))
}

// FillFrom fills AuthExportedAuthorization from given interface.
func (e *AuthExportedAuthorization) FillFrom(from interface {
	GetID() (value int)
	GetBytes() (value []byte)
}) {
	e.ID = from.GetID()
	e.Bytes = from.GetBytes()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthExportedAuthorization) TypeID() uint32 {
	return AuthExportedAuthorizationTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthExportedAuthorization) TypeName() string {
	return "auth.exportedAuthorization"
}

// TypeInfo returns info about TL type.
func (e *AuthExportedAuthorization) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "auth.exportedAuthorization",
		ID:   AuthExportedAuthorizationTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Bytes",
			SchemaName: "bytes",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *AuthExportedAuthorization) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode auth.exportedAuthorization#df969c2d as nil")
	}
	b.PutID(AuthExportedAuthorizationTypeID)
	b.PutInt(e.ID)
	b.PutBytes(e.Bytes)
	return nil
}

// GetID returns value of ID field.
func (e *AuthExportedAuthorization) GetID() (value int) {
	return e.ID
}

// GetBytes returns value of Bytes field.
func (e *AuthExportedAuthorization) GetBytes() (value []byte) {
	return e.Bytes
}

// Decode implements bin.Decoder.
func (e *AuthExportedAuthorization) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode auth.exportedAuthorization#df969c2d to nil")
	}
	if err := b.ConsumeID(AuthExportedAuthorizationTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.exportedAuthorization#df969c2d: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode auth.exportedAuthorization#df969c2d: field id: %w", err)
		}
		e.ID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode auth.exportedAuthorization#df969c2d: field bytes: %w", err)
		}
		e.Bytes = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AuthExportedAuthorization.
var (
	_ bin.Encoder = &AuthExportedAuthorization{}
	_ bin.Decoder = &AuthExportedAuthorization{}
)
