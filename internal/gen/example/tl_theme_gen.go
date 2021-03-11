// Code generated by gotdgen, DO NOT EDIT.

package td

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

// Theme represents TL type `theme#28f1114`.
//
// See https://localhost:80/doc/constructor/theme for reference.
type Theme struct {
	// Name field of Theme.
	Name string
}

// ThemeTypeID is TL type id of Theme.
const ThemeTypeID = 0x28f1114

func (t *Theme) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Name == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *Theme) String() string {
	if t == nil {
		return "Theme(nil)"
	}
	type Alias Theme
	return fmt.Sprintf("Theme%+v", Alias(*t))
}

// FillFrom fills Theme from given interface.
func (t *Theme) FillFrom(from interface {
	GetName() (value string)
}) {
	t.Name = from.GetName()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Theme) TypeID() uint32 {
	return ThemeTypeID
}

// TypeName returns name of type in TL schema.
func (*Theme) TypeName() string {
	return "theme"
}

// TypeInfo returns info about TL type.
func (t *Theme) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "theme",
		ID:   ThemeTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Name",
			SchemaName: "name",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *Theme) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode theme#28f1114 as nil")
	}
	b.PutID(ThemeTypeID)
	b.PutString(t.Name)
	return nil
}

// GetName returns value of Name field.
func (t *Theme) GetName() (value string) {
	return t.Name
}

// Decode implements bin.Decoder.
func (t *Theme) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode theme#28f1114 to nil")
	}
	if err := b.ConsumeID(ThemeTypeID); err != nil {
		return fmt.Errorf("unable to decode theme#28f1114: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode theme#28f1114: field name: %w", err)
		}
		t.Name = value
	}
	return nil
}

// Ensuring interfaces in compile-time for Theme.
var (
	_ bin.Encoder = &Theme{}
	_ bin.Decoder = &Theme{}
)
