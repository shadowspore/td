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

// MessagesUpdatePinnedMessageRequest represents TL type `messages.updatePinnedMessage#d2aaf7ec`.
// Pin a message
//
// See https://core.telegram.org/method/messages.updatePinnedMessage for reference.
type MessagesUpdatePinnedMessageRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Pin the message silently, without triggering a notification
	Silent bool
	// Whether the message should unpinned or pinned
	Unpin bool
	// Whether the message should only be pinned on the local side of a one-to-one chat
	PmOneside bool
	// The peer where to pin the message
	Peer InputPeerClass
	// The message to pin or unpin
	ID int
}

// MessagesUpdatePinnedMessageRequestTypeID is TL type id of MessagesUpdatePinnedMessageRequest.
const MessagesUpdatePinnedMessageRequestTypeID = 0xd2aaf7ec

func (u *MessagesUpdatePinnedMessageRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Flags.Zero()) {
		return false
	}
	if !(u.Silent == false) {
		return false
	}
	if !(u.Unpin == false) {
		return false
	}
	if !(u.PmOneside == false) {
		return false
	}
	if !(u.Peer == nil) {
		return false
	}
	if !(u.ID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *MessagesUpdatePinnedMessageRequest) String() string {
	if u == nil {
		return "MessagesUpdatePinnedMessageRequest(nil)"
	}
	type Alias MessagesUpdatePinnedMessageRequest
	return fmt.Sprintf("MessagesUpdatePinnedMessageRequest%+v", Alias(*u))
}

// FillFrom fills MessagesUpdatePinnedMessageRequest from given interface.
func (u *MessagesUpdatePinnedMessageRequest) FillFrom(from interface {
	GetSilent() (value bool)
	GetUnpin() (value bool)
	GetPmOneside() (value bool)
	GetPeer() (value InputPeerClass)
	GetID() (value int)
}) {
	u.Silent = from.GetSilent()
	u.Unpin = from.GetUnpin()
	u.PmOneside = from.GetPmOneside()
	u.Peer = from.GetPeer()
	u.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesUpdatePinnedMessageRequest) TypeID() uint32 {
	return MessagesUpdatePinnedMessageRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesUpdatePinnedMessageRequest) TypeName() string {
	return "messages.updatePinnedMessage"
}

// TypeInfo returns info about TL type.
func (u *MessagesUpdatePinnedMessageRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.updatePinnedMessage",
		ID:   MessagesUpdatePinnedMessageRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Silent",
			SchemaName: "silent",
			Null:       !u.Flags.Has(0),
		},
		{
			Name:       "Unpin",
			SchemaName: "unpin",
			Null:       !u.Flags.Has(1),
		},
		{
			Name:       "PmOneside",
			SchemaName: "pm_oneside",
			Null:       !u.Flags.Has(2),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *MessagesUpdatePinnedMessageRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode messages.updatePinnedMessage#d2aaf7ec as nil")
	}
	b.PutID(MessagesUpdatePinnedMessageRequestTypeID)
	if !(u.Silent == false) {
		u.Flags.Set(0)
	}
	if !(u.Unpin == false) {
		u.Flags.Set(1)
	}
	if !(u.PmOneside == false) {
		u.Flags.Set(2)
	}
	if err := u.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.updatePinnedMessage#d2aaf7ec: field flags: %w", err)
	}
	if u.Peer == nil {
		return fmt.Errorf("unable to encode messages.updatePinnedMessage#d2aaf7ec: field peer is nil")
	}
	if err := u.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.updatePinnedMessage#d2aaf7ec: field peer: %w", err)
	}
	b.PutInt(u.ID)
	return nil
}

// SetSilent sets value of Silent conditional field.
func (u *MessagesUpdatePinnedMessageRequest) SetSilent(value bool) {
	if value {
		u.Flags.Set(0)
		u.Silent = true
	} else {
		u.Flags.Unset(0)
		u.Silent = false
	}
}

// GetSilent returns value of Silent conditional field.
func (u *MessagesUpdatePinnedMessageRequest) GetSilent() (value bool) {
	return u.Flags.Has(0)
}

// SetUnpin sets value of Unpin conditional field.
func (u *MessagesUpdatePinnedMessageRequest) SetUnpin(value bool) {
	if value {
		u.Flags.Set(1)
		u.Unpin = true
	} else {
		u.Flags.Unset(1)
		u.Unpin = false
	}
}

// GetUnpin returns value of Unpin conditional field.
func (u *MessagesUpdatePinnedMessageRequest) GetUnpin() (value bool) {
	return u.Flags.Has(1)
}

// SetPmOneside sets value of PmOneside conditional field.
func (u *MessagesUpdatePinnedMessageRequest) SetPmOneside(value bool) {
	if value {
		u.Flags.Set(2)
		u.PmOneside = true
	} else {
		u.Flags.Unset(2)
		u.PmOneside = false
	}
}

// GetPmOneside returns value of PmOneside conditional field.
func (u *MessagesUpdatePinnedMessageRequest) GetPmOneside() (value bool) {
	return u.Flags.Has(2)
}

// GetPeer returns value of Peer field.
func (u *MessagesUpdatePinnedMessageRequest) GetPeer() (value InputPeerClass) {
	return u.Peer
}

// GetID returns value of ID field.
func (u *MessagesUpdatePinnedMessageRequest) GetID() (value int) {
	return u.ID
}

// Decode implements bin.Decoder.
func (u *MessagesUpdatePinnedMessageRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode messages.updatePinnedMessage#d2aaf7ec to nil")
	}
	if err := b.ConsumeID(MessagesUpdatePinnedMessageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.updatePinnedMessage#d2aaf7ec: %w", err)
	}
	{
		if err := u.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.updatePinnedMessage#d2aaf7ec: field flags: %w", err)
		}
	}
	u.Silent = u.Flags.Has(0)
	u.Unpin = u.Flags.Has(1)
	u.PmOneside = u.Flags.Has(2)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.updatePinnedMessage#d2aaf7ec: field peer: %w", err)
		}
		u.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.updatePinnedMessage#d2aaf7ec: field id: %w", err)
		}
		u.ID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesUpdatePinnedMessageRequest.
var (
	_ bin.Encoder = &MessagesUpdatePinnedMessageRequest{}
	_ bin.Decoder = &MessagesUpdatePinnedMessageRequest{}
)

// MessagesUpdatePinnedMessage invokes method messages.updatePinnedMessage#d2aaf7ec returning error if any.
// Pin a message
//
// Possible errors:
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this
//  400 CHAT_NOT_MODIFIED: The pinned message wasn't modified
//  403 CHAT_WRITE_FORBIDDEN: You can't write in this chat
//  400 MESSAGE_ID_INVALID: The provided message id is invalid
//  400 PIN_RESTRICTED: You can't pin messages
//
// See https://core.telegram.org/method/messages.updatePinnedMessage for reference.
// Can be used by bots.
func (c *Client) MessagesUpdatePinnedMessage(ctx context.Context, request *MessagesUpdatePinnedMessageRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
