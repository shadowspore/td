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

// PhoneInviteToGroupCallRequest represents TL type `phone.inviteToGroupCall#7b393160`.
//
// See https://core.telegram.org/method/phone.inviteToGroupCall for reference.
type PhoneInviteToGroupCallRequest struct {
	// Call field of PhoneInviteToGroupCallRequest.
	Call InputGroupCall
	// Users field of PhoneInviteToGroupCallRequest.
	Users []InputUserClass
}

// PhoneInviteToGroupCallRequestTypeID is TL type id of PhoneInviteToGroupCallRequest.
const PhoneInviteToGroupCallRequestTypeID = 0x7b393160

func (i *PhoneInviteToGroupCallRequest) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Call.Zero()) {
		return false
	}
	if !(i.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *PhoneInviteToGroupCallRequest) String() string {
	if i == nil {
		return "PhoneInviteToGroupCallRequest(nil)"
	}
	type Alias PhoneInviteToGroupCallRequest
	return fmt.Sprintf("PhoneInviteToGroupCallRequest%+v", Alias(*i))
}

// FillFrom fills PhoneInviteToGroupCallRequest from given interface.
func (i *PhoneInviteToGroupCallRequest) FillFrom(from interface {
	GetCall() (value InputGroupCall)
	GetUsers() (value []InputUserClass)
}) {
	i.Call = from.GetCall()
	i.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhoneInviteToGroupCallRequest) TypeID() uint32 {
	return PhoneInviteToGroupCallRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PhoneInviteToGroupCallRequest) TypeName() string {
	return "phone.inviteToGroupCall"
}

// TypeInfo returns info about TL type.
func (i *PhoneInviteToGroupCallRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "phone.inviteToGroupCall",
		ID:   PhoneInviteToGroupCallRequestTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Call",
			SchemaName: "call",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *PhoneInviteToGroupCallRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode phone.inviteToGroupCall#7b393160 as nil")
	}
	b.PutID(PhoneInviteToGroupCallRequestTypeID)
	if err := i.Call.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.inviteToGroupCall#7b393160: field call: %w", err)
	}
	b.PutVectorHeader(len(i.Users))
	for idx, v := range i.Users {
		if v == nil {
			return fmt.Errorf("unable to encode phone.inviteToGroupCall#7b393160: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode phone.inviteToGroupCall#7b393160: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetCall returns value of Call field.
func (i *PhoneInviteToGroupCallRequest) GetCall() (value InputGroupCall) {
	return i.Call
}

// GetUsers returns value of Users field.
func (i *PhoneInviteToGroupCallRequest) GetUsers() (value []InputUserClass) {
	return i.Users
}

// MapUsers returns field Users wrapped in InputUserClassArray helper.
func (i *PhoneInviteToGroupCallRequest) MapUsers() (value InputUserClassArray) {
	return InputUserClassArray(i.Users)
}

// Decode implements bin.Decoder.
func (i *PhoneInviteToGroupCallRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode phone.inviteToGroupCall#7b393160 to nil")
	}
	if err := b.ConsumeID(PhoneInviteToGroupCallRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.inviteToGroupCall#7b393160: %w", err)
	}
	{
		if err := i.Call.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.inviteToGroupCall#7b393160: field call: %w", err)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode phone.inviteToGroupCall#7b393160: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode phone.inviteToGroupCall#7b393160: field users: %w", err)
			}
			i.Users = append(i.Users, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for PhoneInviteToGroupCallRequest.
var (
	_ bin.Encoder = &PhoneInviteToGroupCallRequest{}
	_ bin.Decoder = &PhoneInviteToGroupCallRequest{}
)

// PhoneInviteToGroupCall invokes method phone.inviteToGroupCall#7b393160 returning error if any.
//
// See https://core.telegram.org/method/phone.inviteToGroupCall for reference.
func (c *Client) PhoneInviteToGroupCall(ctx context.Context, request *PhoneInviteToGroupCallRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
