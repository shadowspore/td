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

// PhoneDiscardCallRequest represents TL type `phone.discardCall#b2cbc1c0`.
// Refuse or end running call
//
// See https://core.telegram.org/method/phone.discardCall for reference.
type PhoneDiscardCallRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether this is a video call
	Video bool
	// The phone call
	Peer InputPhoneCall
	// Call duration
	Duration int
	// Why was the call discarded
	Reason PhoneCallDiscardReasonClass
	// Preferred libtgvoip relay ID
	ConnectionID int64
}

// PhoneDiscardCallRequestTypeID is TL type id of PhoneDiscardCallRequest.
const PhoneDiscardCallRequestTypeID = 0xb2cbc1c0

func (d *PhoneDiscardCallRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Flags.Zero()) {
		return false
	}
	if !(d.Video == false) {
		return false
	}
	if !(d.Peer.Zero()) {
		return false
	}
	if !(d.Duration == 0) {
		return false
	}
	if !(d.Reason == nil) {
		return false
	}
	if !(d.ConnectionID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *PhoneDiscardCallRequest) String() string {
	if d == nil {
		return "PhoneDiscardCallRequest(nil)"
	}
	type Alias PhoneDiscardCallRequest
	return fmt.Sprintf("PhoneDiscardCallRequest%+v", Alias(*d))
}

// FillFrom fills PhoneDiscardCallRequest from given interface.
func (d *PhoneDiscardCallRequest) FillFrom(from interface {
	GetVideo() (value bool)
	GetPeer() (value InputPhoneCall)
	GetDuration() (value int)
	GetReason() (value PhoneCallDiscardReasonClass)
	GetConnectionID() (value int64)
}) {
	d.Video = from.GetVideo()
	d.Peer = from.GetPeer()
	d.Duration = from.GetDuration()
	d.Reason = from.GetReason()
	d.ConnectionID = from.GetConnectionID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhoneDiscardCallRequest) TypeID() uint32 {
	return PhoneDiscardCallRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PhoneDiscardCallRequest) TypeName() string {
	return "phone.discardCall"
}

// TypeInfo returns info about TL type.
func (d *PhoneDiscardCallRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "phone.discardCall",
		ID:   PhoneDiscardCallRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Video",
			SchemaName: "video",
			Null:       !d.Flags.Has(0),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Duration",
			SchemaName: "duration",
		},
		{
			Name:       "Reason",
			SchemaName: "reason",
		},
		{
			Name:       "ConnectionID",
			SchemaName: "connection_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *PhoneDiscardCallRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode phone.discardCall#b2cbc1c0 as nil")
	}
	b.PutID(PhoneDiscardCallRequestTypeID)
	if !(d.Video == false) {
		d.Flags.Set(0)
	}
	if err := d.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.discardCall#b2cbc1c0: field flags: %w", err)
	}
	if err := d.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.discardCall#b2cbc1c0: field peer: %w", err)
	}
	b.PutInt(d.Duration)
	if d.Reason == nil {
		return fmt.Errorf("unable to encode phone.discardCall#b2cbc1c0: field reason is nil")
	}
	if err := d.Reason.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.discardCall#b2cbc1c0: field reason: %w", err)
	}
	b.PutLong(d.ConnectionID)
	return nil
}

// SetVideo sets value of Video conditional field.
func (d *PhoneDiscardCallRequest) SetVideo(value bool) {
	if value {
		d.Flags.Set(0)
		d.Video = true
	} else {
		d.Flags.Unset(0)
		d.Video = false
	}
}

// GetVideo returns value of Video conditional field.
func (d *PhoneDiscardCallRequest) GetVideo() (value bool) {
	return d.Flags.Has(0)
}

// GetPeer returns value of Peer field.
func (d *PhoneDiscardCallRequest) GetPeer() (value InputPhoneCall) {
	return d.Peer
}

// GetDuration returns value of Duration field.
func (d *PhoneDiscardCallRequest) GetDuration() (value int) {
	return d.Duration
}

// GetReason returns value of Reason field.
func (d *PhoneDiscardCallRequest) GetReason() (value PhoneCallDiscardReasonClass) {
	return d.Reason
}

// GetConnectionID returns value of ConnectionID field.
func (d *PhoneDiscardCallRequest) GetConnectionID() (value int64) {
	return d.ConnectionID
}

// Decode implements bin.Decoder.
func (d *PhoneDiscardCallRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode phone.discardCall#b2cbc1c0 to nil")
	}
	if err := b.ConsumeID(PhoneDiscardCallRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.discardCall#b2cbc1c0: %w", err)
	}
	{
		if err := d.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.discardCall#b2cbc1c0: field flags: %w", err)
		}
	}
	d.Video = d.Flags.Has(0)
	{
		if err := d.Peer.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.discardCall#b2cbc1c0: field peer: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phone.discardCall#b2cbc1c0: field duration: %w", err)
		}
		d.Duration = value
	}
	{
		value, err := DecodePhoneCallDiscardReason(b)
		if err != nil {
			return fmt.Errorf("unable to decode phone.discardCall#b2cbc1c0: field reason: %w", err)
		}
		d.Reason = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode phone.discardCall#b2cbc1c0: field connection_id: %w", err)
		}
		d.ConnectionID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for PhoneDiscardCallRequest.
var (
	_ bin.Encoder = &PhoneDiscardCallRequest{}
	_ bin.Decoder = &PhoneDiscardCallRequest{}
)

// PhoneDiscardCall invokes method phone.discardCall#b2cbc1c0 returning error if any.
// Refuse or end running call
//
// Possible errors:
//  400 CALL_ALREADY_ACCEPTED: The call was already accepted
//  400 CALL_PEER_INVALID: The provided call peer object is invalid
//
// See https://core.telegram.org/method/phone.discardCall for reference.
func (c *Client) PhoneDiscardCall(ctx context.Context, request *PhoneDiscardCallRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
