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

// SecureSecretSettings represents TL type `secureSecretSettings#1527bcac`.
// Secure settings
//
// See https://core.telegram.org/constructor/secureSecretSettings for reference.
type SecureSecretSettings struct {
	// Secure KDF algo
	SecureAlgo SecurePasswordKdfAlgoClass
	// Secure secret
	SecureSecret []byte
	// Secret ID
	SecureSecretID int64
}

// SecureSecretSettingsTypeID is TL type id of SecureSecretSettings.
const SecureSecretSettingsTypeID = 0x1527bcac

func (s *SecureSecretSettings) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.SecureAlgo == nil) {
		return false
	}
	if !(s.SecureSecret == nil) {
		return false
	}
	if !(s.SecureSecretID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SecureSecretSettings) String() string {
	if s == nil {
		return "SecureSecretSettings(nil)"
	}
	type Alias SecureSecretSettings
	return fmt.Sprintf("SecureSecretSettings%+v", Alias(*s))
}

// FillFrom fills SecureSecretSettings from given interface.
func (s *SecureSecretSettings) FillFrom(from interface {
	GetSecureAlgo() (value SecurePasswordKdfAlgoClass)
	GetSecureSecret() (value []byte)
	GetSecureSecretID() (value int64)
}) {
	s.SecureAlgo = from.GetSecureAlgo()
	s.SecureSecret = from.GetSecureSecret()
	s.SecureSecretID = from.GetSecureSecretID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SecureSecretSettings) TypeID() uint32 {
	return SecureSecretSettingsTypeID
}

// TypeName returns name of type in TL schema.
func (*SecureSecretSettings) TypeName() string {
	return "secureSecretSettings"
}

// TypeInfo returns info about TL type.
func (s *SecureSecretSettings) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "secureSecretSettings",
		ID:   SecureSecretSettingsTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SecureAlgo",
			SchemaName: "secure_algo",
		},
		{
			Name:       "SecureSecret",
			SchemaName: "secure_secret",
		},
		{
			Name:       "SecureSecretID",
			SchemaName: "secure_secret_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SecureSecretSettings) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureSecretSettings#1527bcac as nil")
	}
	b.PutID(SecureSecretSettingsTypeID)
	if s.SecureAlgo == nil {
		return fmt.Errorf("unable to encode secureSecretSettings#1527bcac: field secure_algo is nil")
	}
	if err := s.SecureAlgo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode secureSecretSettings#1527bcac: field secure_algo: %w", err)
	}
	b.PutBytes(s.SecureSecret)
	b.PutLong(s.SecureSecretID)
	return nil
}

// GetSecureAlgo returns value of SecureAlgo field.
func (s *SecureSecretSettings) GetSecureAlgo() (value SecurePasswordKdfAlgoClass) {
	return s.SecureAlgo
}

// GetSecureSecret returns value of SecureSecret field.
func (s *SecureSecretSettings) GetSecureSecret() (value []byte) {
	return s.SecureSecret
}

// GetSecureSecretID returns value of SecureSecretID field.
func (s *SecureSecretSettings) GetSecureSecretID() (value int64) {
	return s.SecureSecretID
}

// Decode implements bin.Decoder.
func (s *SecureSecretSettings) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureSecretSettings#1527bcac to nil")
	}
	if err := b.ConsumeID(SecureSecretSettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode secureSecretSettings#1527bcac: %w", err)
	}
	{
		value, err := DecodeSecurePasswordKdfAlgo(b)
		if err != nil {
			return fmt.Errorf("unable to decode secureSecretSettings#1527bcac: field secure_algo: %w", err)
		}
		s.SecureAlgo = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode secureSecretSettings#1527bcac: field secure_secret: %w", err)
		}
		s.SecureSecret = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode secureSecretSettings#1527bcac: field secure_secret_id: %w", err)
		}
		s.SecureSecretID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for SecureSecretSettings.
var (
	_ bin.Encoder = &SecureSecretSettings{}
	_ bin.Decoder = &SecureSecretSettings{}
)
