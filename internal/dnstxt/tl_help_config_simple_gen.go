// Code generated by gotdgen, DO NOT EDIT.

package dnstxt

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

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
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
)

// HelpConfigSimple represents TL type `help.configSimple#5a592a6c`.
type HelpConfigSimple struct {
	// Date field of HelpConfigSimple.
	Date int
	// Expires field of HelpConfigSimple.
	Expires int
	// Rules field of HelpConfigSimple.
	Rules []AccessPointRule
}

// HelpConfigSimpleTypeID is TL type id of HelpConfigSimple.
const HelpConfigSimpleTypeID = 0x5a592a6c

func (c *HelpConfigSimple) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Date == 0) {
		return false
	}
	if !(c.Expires == 0) {
		return false
	}
	if !(c.Rules == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *HelpConfigSimple) String() string {
	if c == nil {
		return "HelpConfigSimple(nil)"
	}
	type Alias HelpConfigSimple
	return fmt.Sprintf("HelpConfigSimple%+v", Alias(*c))
}

// FillFrom fills HelpConfigSimple from given interface.
func (c *HelpConfigSimple) FillFrom(from interface {
	GetDate() (value int)
	GetExpires() (value int)
	GetRules() (value []AccessPointRule)
}) {
	c.Date = from.GetDate()
	c.Expires = from.GetExpires()
	c.Rules = from.GetRules()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpConfigSimple) TypeID() uint32 {
	return HelpConfigSimpleTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpConfigSimple) TypeName() string {
	return "help.configSimple"
}

// TypeInfo returns info about TL type.
func (c *HelpConfigSimple) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.configSimple",
		ID:   HelpConfigSimpleTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Date",
			SchemaName: "date",
		},
		{
			Name:       "Expires",
			SchemaName: "expires",
		},
		{
			Name:       "Rules",
			SchemaName: "rules",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *HelpConfigSimple) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode help.configSimple#5a592a6c as nil")
	}
	b.PutID(HelpConfigSimpleTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *HelpConfigSimple) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode help.configSimple#5a592a6c as nil")
	}
	b.PutInt(c.Date)
	b.PutInt(c.Expires)
	b.PutInt(len(c.Rules))
	for idx, v := range c.Rules {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare help.configSimple#5a592a6c: field rules element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetDate returns value of Date field.
func (c *HelpConfigSimple) GetDate() (value int) {
	return c.Date
}

// GetExpires returns value of Expires field.
func (c *HelpConfigSimple) GetExpires() (value int) {
	return c.Expires
}

// GetRules returns value of Rules field.
func (c *HelpConfigSimple) GetRules() (value []AccessPointRule) {
	return c.Rules
}

// Decode implements bin.Decoder.
func (c *HelpConfigSimple) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode help.configSimple#5a592a6c to nil")
	}
	if err := b.ConsumeID(HelpConfigSimpleTypeID); err != nil {
		return fmt.Errorf("unable to decode help.configSimple#5a592a6c: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *HelpConfigSimple) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode help.configSimple#5a592a6c to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode help.configSimple#5a592a6c: field date: %w", err)
		}
		c.Date = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode help.configSimple#5a592a6c: field expires: %w", err)
		}
		c.Expires = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode help.configSimple#5a592a6c: field rules: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value AccessPointRule
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare help.configSimple#5a592a6c: field rules: %w", err)
			}
			c.Rules = append(c.Rules, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpConfigSimple.
var (
	_ bin.Encoder     = &HelpConfigSimple{}
	_ bin.Decoder     = &HelpConfigSimple{}
	_ bin.BareEncoder = &HelpConfigSimple{}
	_ bin.BareDecoder = &HelpConfigSimple{}
)
