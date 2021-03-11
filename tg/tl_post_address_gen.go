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

// PostAddress represents TL type `postAddress#1e8caaeb`.
// Shipping address
//
// See https://core.telegram.org/constructor/postAddress for reference.
type PostAddress struct {
	// First line for the address
	StreetLine1 string
	// Second line for the address
	StreetLine2 string
	// City
	City string
	// State, if applicable (empty otherwise)
	State string
	// ISO 3166-1 alpha-2 country code
	CountryIso2 string
	// Address post code
	PostCode string
}

// PostAddressTypeID is TL type id of PostAddress.
const PostAddressTypeID = 0x1e8caaeb

func (p *PostAddress) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.StreetLine1 == "") {
		return false
	}
	if !(p.StreetLine2 == "") {
		return false
	}
	if !(p.City == "") {
		return false
	}
	if !(p.State == "") {
		return false
	}
	if !(p.CountryIso2 == "") {
		return false
	}
	if !(p.PostCode == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PostAddress) String() string {
	if p == nil {
		return "PostAddress(nil)"
	}
	type Alias PostAddress
	return fmt.Sprintf("PostAddress%+v", Alias(*p))
}

// FillFrom fills PostAddress from given interface.
func (p *PostAddress) FillFrom(from interface {
	GetStreetLine1() (value string)
	GetStreetLine2() (value string)
	GetCity() (value string)
	GetState() (value string)
	GetCountryIso2() (value string)
	GetPostCode() (value string)
}) {
	p.StreetLine1 = from.GetStreetLine1()
	p.StreetLine2 = from.GetStreetLine2()
	p.City = from.GetCity()
	p.State = from.GetState()
	p.CountryIso2 = from.GetCountryIso2()
	p.PostCode = from.GetPostCode()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PostAddress) TypeID() uint32 {
	return PostAddressTypeID
}

// TypeName returns name of type in TL schema.
func (*PostAddress) TypeName() string {
	return "postAddress"
}

// TypeInfo returns info about TL type.
func (p *PostAddress) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "postAddress",
		ID:   PostAddressTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "StreetLine1",
			SchemaName: "street_line1",
		},
		{
			Name:       "StreetLine2",
			SchemaName: "street_line2",
		},
		{
			Name:       "City",
			SchemaName: "city",
		},
		{
			Name:       "State",
			SchemaName: "state",
		},
		{
			Name:       "CountryIso2",
			SchemaName: "country_iso2",
		},
		{
			Name:       "PostCode",
			SchemaName: "post_code",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PostAddress) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode postAddress#1e8caaeb as nil")
	}
	b.PutID(PostAddressTypeID)
	b.PutString(p.StreetLine1)
	b.PutString(p.StreetLine2)
	b.PutString(p.City)
	b.PutString(p.State)
	b.PutString(p.CountryIso2)
	b.PutString(p.PostCode)
	return nil
}

// GetStreetLine1 returns value of StreetLine1 field.
func (p *PostAddress) GetStreetLine1() (value string) {
	return p.StreetLine1
}

// GetStreetLine2 returns value of StreetLine2 field.
func (p *PostAddress) GetStreetLine2() (value string) {
	return p.StreetLine2
}

// GetCity returns value of City field.
func (p *PostAddress) GetCity() (value string) {
	return p.City
}

// GetState returns value of State field.
func (p *PostAddress) GetState() (value string) {
	return p.State
}

// GetCountryIso2 returns value of CountryIso2 field.
func (p *PostAddress) GetCountryIso2() (value string) {
	return p.CountryIso2
}

// GetPostCode returns value of PostCode field.
func (p *PostAddress) GetPostCode() (value string) {
	return p.PostCode
}

// Decode implements bin.Decoder.
func (p *PostAddress) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode postAddress#1e8caaeb to nil")
	}
	if err := b.ConsumeID(PostAddressTypeID); err != nil {
		return fmt.Errorf("unable to decode postAddress#1e8caaeb: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode postAddress#1e8caaeb: field street_line1: %w", err)
		}
		p.StreetLine1 = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode postAddress#1e8caaeb: field street_line2: %w", err)
		}
		p.StreetLine2 = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode postAddress#1e8caaeb: field city: %w", err)
		}
		p.City = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode postAddress#1e8caaeb: field state: %w", err)
		}
		p.State = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode postAddress#1e8caaeb: field country_iso2: %w", err)
		}
		p.CountryIso2 = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode postAddress#1e8caaeb: field post_code: %w", err)
		}
		p.PostCode = value
	}
	return nil
}

// Ensuring interfaces in compile-time for PostAddress.
var (
	_ bin.Encoder = &PostAddress{}
	_ bin.Decoder = &PostAddress{}
)
