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

// StatsLoadAsyncGraphRequest represents TL type `stats.loadAsyncGraph#621d5fa0`.
// Load channel statistics graph¹ asynchronously
//
// Links:
//  1) https://core.telegram.org/api/stats
//
// See https://core.telegram.org/method/stats.loadAsyncGraph for reference.
type StatsLoadAsyncGraphRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Graph token from statsGraphAsync¹ constructor
	//
	// Links:
	//  1) https://core.telegram.org/constructor/statsGraphAsync
	Token string
	// Zoom value, if required
	//
	// Use SetX and GetX helpers.
	X int64
}

// StatsLoadAsyncGraphRequestTypeID is TL type id of StatsLoadAsyncGraphRequest.
const StatsLoadAsyncGraphRequestTypeID = 0x621d5fa0

func (l *StatsLoadAsyncGraphRequest) Zero() bool {
	if l == nil {
		return true
	}
	if !(l.Flags.Zero()) {
		return false
	}
	if !(l.Token == "") {
		return false
	}
	if !(l.X == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (l *StatsLoadAsyncGraphRequest) String() string {
	if l == nil {
		return "StatsLoadAsyncGraphRequest(nil)"
	}
	type Alias StatsLoadAsyncGraphRequest
	return fmt.Sprintf("StatsLoadAsyncGraphRequest%+v", Alias(*l))
}

// FillFrom fills StatsLoadAsyncGraphRequest from given interface.
func (l *StatsLoadAsyncGraphRequest) FillFrom(from interface {
	GetToken() (value string)
	GetX() (value int64, ok bool)
}) {
	l.Token = from.GetToken()
	if val, ok := from.GetX(); ok {
		l.X = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StatsLoadAsyncGraphRequest) TypeID() uint32 {
	return StatsLoadAsyncGraphRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*StatsLoadAsyncGraphRequest) TypeName() string {
	return "stats.loadAsyncGraph"
}

// TypeInfo returns info about TL type.
func (l *StatsLoadAsyncGraphRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stats.loadAsyncGraph",
		ID:   StatsLoadAsyncGraphRequestTypeID,
	}
	if l == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Token",
			SchemaName: "token",
		},
		{
			Name:       "X",
			SchemaName: "x",
			Null:       !l.Flags.Has(0),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (l *StatsLoadAsyncGraphRequest) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode stats.loadAsyncGraph#621d5fa0 as nil")
	}
	b.PutID(StatsLoadAsyncGraphRequestTypeID)
	if !(l.X == 0) {
		l.Flags.Set(0)
	}
	if err := l.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.loadAsyncGraph#621d5fa0: field flags: %w", err)
	}
	b.PutString(l.Token)
	if l.Flags.Has(0) {
		b.PutLong(l.X)
	}
	return nil
}

// GetToken returns value of Token field.
func (l *StatsLoadAsyncGraphRequest) GetToken() (value string) {
	return l.Token
}

// SetX sets value of X conditional field.
func (l *StatsLoadAsyncGraphRequest) SetX(value int64) {
	l.Flags.Set(0)
	l.X = value
}

// GetX returns value of X conditional field and
// boolean which is true if field was set.
func (l *StatsLoadAsyncGraphRequest) GetX() (value int64, ok bool) {
	if !l.Flags.Has(0) {
		return value, false
	}
	return l.X, true
}

// Decode implements bin.Decoder.
func (l *StatsLoadAsyncGraphRequest) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode stats.loadAsyncGraph#621d5fa0 to nil")
	}
	if err := b.ConsumeID(StatsLoadAsyncGraphRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stats.loadAsyncGraph#621d5fa0: %w", err)
	}
	{
		if err := l.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stats.loadAsyncGraph#621d5fa0: field flags: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode stats.loadAsyncGraph#621d5fa0: field token: %w", err)
		}
		l.Token = value
	}
	if l.Flags.Has(0) {
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode stats.loadAsyncGraph#621d5fa0: field x: %w", err)
		}
		l.X = value
	}
	return nil
}

// Ensuring interfaces in compile-time for StatsLoadAsyncGraphRequest.
var (
	_ bin.Encoder = &StatsLoadAsyncGraphRequest{}
	_ bin.Decoder = &StatsLoadAsyncGraphRequest{}
)

// StatsLoadAsyncGraph invokes method stats.loadAsyncGraph#621d5fa0 returning error if any.
// Load channel statistics graph¹ asynchronously
//
// Links:
//  1) https://core.telegram.org/api/stats
//
// Possible errors:
//  400 GRAPH_INVALID_RELOAD: Invalid graph token provided, please reload the stats and provide the updated token
//  400 GRAPH_OUTDATED_RELOAD: The graph is outdated, please get a new async token using stats.getBroadcastStats
//
// See https://core.telegram.org/method/stats.loadAsyncGraph for reference.
func (c *Client) StatsLoadAsyncGraph(ctx context.Context, request *StatsLoadAsyncGraphRequest) (StatsGraphClass, error) {
	var result StatsGraphBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.StatsGraph, nil
}
