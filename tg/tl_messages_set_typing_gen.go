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

// MessagesSetTypingRequest represents TL type `messages.setTyping#58943ee2`.
// Sends a current user typing event (see SendMessageAction¹ for all event types) to a conversation partner or group.
//
// Links:
//  1) https://core.telegram.org/type/SendMessageAction
//
// See https://core.telegram.org/method/messages.setTyping for reference.
type MessagesSetTypingRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Target user or group
	Peer InputPeerClass
	// Thread ID¹
	//
	// Links:
	//  1) https://core.telegram.org/api/threads
	//
	// Use SetTopMsgID and GetTopMsgID helpers.
	TopMsgID int
	// Type of actionParameter added in Layer 17¹.
	//
	// Links:
	//  1) https://core.telegram.org/api/layers#layer-17
	Action SendMessageActionClass
}

// MessagesSetTypingRequestTypeID is TL type id of MessagesSetTypingRequest.
const MessagesSetTypingRequestTypeID = 0x58943ee2

func (s *MessagesSetTypingRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Peer == nil) {
		return false
	}
	if !(s.TopMsgID == 0) {
		return false
	}
	if !(s.Action == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSetTypingRequest) String() string {
	if s == nil {
		return "MessagesSetTypingRequest(nil)"
	}
	type Alias MessagesSetTypingRequest
	return fmt.Sprintf("MessagesSetTypingRequest%+v", Alias(*s))
}

// FillFrom fills MessagesSetTypingRequest from given interface.
func (s *MessagesSetTypingRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetTopMsgID() (value int, ok bool)
	GetAction() (value SendMessageActionClass)
}) {
	s.Peer = from.GetPeer()
	if val, ok := from.GetTopMsgID(); ok {
		s.TopMsgID = val
	}

	s.Action = from.GetAction()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSetTypingRequest) TypeID() uint32 {
	return MessagesSetTypingRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSetTypingRequest) TypeName() string {
	return "messages.setTyping"
}

// TypeInfo returns info about TL type.
func (s *MessagesSetTypingRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.setTyping",
		ID:   MessagesSetTypingRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "TopMsgID",
			SchemaName: "top_msg_id",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "Action",
			SchemaName: "action",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *MessagesSetTypingRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.setTyping#58943ee2 as nil")
	}
	b.PutID(MessagesSetTypingRequestTypeID)
	if !(s.TopMsgID == 0) {
		s.Flags.Set(0)
	}
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.setTyping#58943ee2: field flags: %w", err)
	}
	if s.Peer == nil {
		return fmt.Errorf("unable to encode messages.setTyping#58943ee2: field peer is nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.setTyping#58943ee2: field peer: %w", err)
	}
	if s.Flags.Has(0) {
		b.PutInt(s.TopMsgID)
	}
	if s.Action == nil {
		return fmt.Errorf("unable to encode messages.setTyping#58943ee2: field action is nil")
	}
	if err := s.Action.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.setTyping#58943ee2: field action: %w", err)
	}
	return nil
}

// GetPeer returns value of Peer field.
func (s *MessagesSetTypingRequest) GetPeer() (value InputPeerClass) {
	return s.Peer
}

// SetTopMsgID sets value of TopMsgID conditional field.
func (s *MessagesSetTypingRequest) SetTopMsgID(value int) {
	s.Flags.Set(0)
	s.TopMsgID = value
}

// GetTopMsgID returns value of TopMsgID conditional field and
// boolean which is true if field was set.
func (s *MessagesSetTypingRequest) GetTopMsgID() (value int, ok bool) {
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.TopMsgID, true
}

// GetAction returns value of Action field.
func (s *MessagesSetTypingRequest) GetAction() (value SendMessageActionClass) {
	return s.Action
}

// Decode implements bin.Decoder.
func (s *MessagesSetTypingRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.setTyping#58943ee2 to nil")
	}
	if err := b.ConsumeID(MessagesSetTypingRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.setTyping#58943ee2: %w", err)
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.setTyping#58943ee2: field flags: %w", err)
		}
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.setTyping#58943ee2: field peer: %w", err)
		}
		s.Peer = value
	}
	if s.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.setTyping#58943ee2: field top_msg_id: %w", err)
		}
		s.TopMsgID = value
	}
	{
		value, err := DecodeSendMessageAction(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.setTyping#58943ee2: field action: %w", err)
		}
		s.Action = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesSetTypingRequest.
var (
	_ bin.Encoder = &MessagesSetTypingRequest{}
	_ bin.Decoder = &MessagesSetTypingRequest{}
)

// MessagesSetTyping invokes method messages.setTyping#58943ee2 returning error if any.
// Sends a current user typing event (see SendMessageAction¹ for all event types) to a conversation partner or group.
//
// Links:
//  1) https://core.telegram.org/type/SendMessageAction
//
// Possible errors:
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this
//  400 CHAT_ID_INVALID: The provided chat id is invalid
//  403 CHAT_WRITE_FORBIDDEN: You can't write in this chat
//  400 INPUT_USER_DEACTIVATED: The specified user was deleted
//  400 MSG_ID_INVALID: Invalid message ID provided
//  400 PEER_ID_INVALID: The provided peer id is invalid
//  400 USER_BANNED_IN_CHANNEL: You're banned from sending messages in supergroups/channels
//  400 USER_IS_BLOCKED: You were blocked by this user
//  400 USER_IS_BOT: Bots can't send messages to other bots
//
// See https://core.telegram.org/method/messages.setTyping for reference.
// Can be used by bots.
func (c *Client) MessagesSetTyping(ctx context.Context, request *MessagesSetTypingRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
