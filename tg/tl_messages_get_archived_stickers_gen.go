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

// MessagesGetArchivedStickersRequest represents TL type `messages.getArchivedStickers#57f17692`.
// Get all archived stickers
//
// See https://core.telegram.org/method/messages.getArchivedStickers for reference.
type MessagesGetArchivedStickersRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Get mask stickers
	Masks bool
	// Offsets for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	OffsetID int64
	// Maximum number of results to return, see pagination¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Limit int
}

// MessagesGetArchivedStickersRequestTypeID is TL type id of MessagesGetArchivedStickersRequest.
const MessagesGetArchivedStickersRequestTypeID = 0x57f17692

func (g *MessagesGetArchivedStickersRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.Masks == false) {
		return false
	}
	if !(g.OffsetID == 0) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetArchivedStickersRequest) String() string {
	if g == nil {
		return "MessagesGetArchivedStickersRequest(nil)"
	}
	type Alias MessagesGetArchivedStickersRequest
	return fmt.Sprintf("MessagesGetArchivedStickersRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetArchivedStickersRequest from given interface.
func (g *MessagesGetArchivedStickersRequest) FillFrom(from interface {
	GetMasks() (value bool)
	GetOffsetID() (value int64)
	GetLimit() (value int)
}) {
	g.Masks = from.GetMasks()
	g.OffsetID = from.GetOffsetID()
	g.Limit = from.GetLimit()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetArchivedStickersRequest) TypeID() uint32 {
	return MessagesGetArchivedStickersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetArchivedStickersRequest) TypeName() string {
	return "messages.getArchivedStickers"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetArchivedStickersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getArchivedStickers",
		ID:   MessagesGetArchivedStickersRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Masks",
			SchemaName: "masks",
			Null:       !g.Flags.Has(0),
		},
		{
			Name:       "OffsetID",
			SchemaName: "offset_id",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetArchivedStickersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getArchivedStickers#57f17692 as nil")
	}
	b.PutID(MessagesGetArchivedStickersRequestTypeID)
	if !(g.Masks == false) {
		g.Flags.Set(0)
	}
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getArchivedStickers#57f17692: field flags: %w", err)
	}
	b.PutLong(g.OffsetID)
	b.PutInt(g.Limit)
	return nil
}

// SetMasks sets value of Masks conditional field.
func (g *MessagesGetArchivedStickersRequest) SetMasks(value bool) {
	if value {
		g.Flags.Set(0)
		g.Masks = true
	} else {
		g.Flags.Unset(0)
		g.Masks = false
	}
}

// GetMasks returns value of Masks conditional field.
func (g *MessagesGetArchivedStickersRequest) GetMasks() (value bool) {
	return g.Flags.Has(0)
}

// GetOffsetID returns value of OffsetID field.
func (g *MessagesGetArchivedStickersRequest) GetOffsetID() (value int64) {
	return g.OffsetID
}

// GetLimit returns value of Limit field.
func (g *MessagesGetArchivedStickersRequest) GetLimit() (value int) {
	return g.Limit
}

// Decode implements bin.Decoder.
func (g *MessagesGetArchivedStickersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getArchivedStickers#57f17692 to nil")
	}
	if err := b.ConsumeID(MessagesGetArchivedStickersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getArchivedStickers#57f17692: %w", err)
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.getArchivedStickers#57f17692: field flags: %w", err)
		}
	}
	g.Masks = g.Flags.Has(0)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getArchivedStickers#57f17692: field offset_id: %w", err)
		}
		g.OffsetID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getArchivedStickers#57f17692: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetArchivedStickersRequest.
var (
	_ bin.Encoder = &MessagesGetArchivedStickersRequest{}
	_ bin.Decoder = &MessagesGetArchivedStickersRequest{}
)

// MessagesGetArchivedStickers invokes method messages.getArchivedStickers#57f17692 returning error if any.
// Get all archived stickers
//
// See https://core.telegram.org/method/messages.getArchivedStickers for reference.
func (c *Client) MessagesGetArchivedStickers(ctx context.Context, request *MessagesGetArchivedStickersRequest) (*MessagesArchivedStickers, error) {
	var result MessagesArchivedStickers

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
