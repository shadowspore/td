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

// ContactsUnblockRequest represents TL type `contacts.unblock#bea65d50`.
// Deletes the user from the blacklist.
//
// See https://core.telegram.org/method/contacts.unblock for reference.
type ContactsUnblockRequest struct {
	// User ID
	ID InputPeerClass
}

// ContactsUnblockRequestTypeID is TL type id of ContactsUnblockRequest.
const ContactsUnblockRequestTypeID = 0xbea65d50

func (u *ContactsUnblockRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.ID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *ContactsUnblockRequest) String() string {
	if u == nil {
		return "ContactsUnblockRequest(nil)"
	}
	type Alias ContactsUnblockRequest
	return fmt.Sprintf("ContactsUnblockRequest%+v", Alias(*u))
}

// FillFrom fills ContactsUnblockRequest from given interface.
func (u *ContactsUnblockRequest) FillFrom(from interface {
	GetID() (value InputPeerClass)
}) {
	u.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ContactsUnblockRequest) TypeID() uint32 {
	return ContactsUnblockRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ContactsUnblockRequest) TypeName() string {
	return "contacts.unblock"
}

// TypeInfo returns info about TL type.
func (u *ContactsUnblockRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "contacts.unblock",
		ID:   ContactsUnblockRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *ContactsUnblockRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode contacts.unblock#bea65d50 as nil")
	}
	b.PutID(ContactsUnblockRequestTypeID)
	if u.ID == nil {
		return fmt.Errorf("unable to encode contacts.unblock#bea65d50: field id is nil")
	}
	if err := u.ID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode contacts.unblock#bea65d50: field id: %w", err)
	}
	return nil
}

// GetID returns value of ID field.
func (u *ContactsUnblockRequest) GetID() (value InputPeerClass) {
	return u.ID
}

// Decode implements bin.Decoder.
func (u *ContactsUnblockRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode contacts.unblock#bea65d50 to nil")
	}
	if err := b.ConsumeID(ContactsUnblockRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.unblock#bea65d50: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode contacts.unblock#bea65d50: field id: %w", err)
		}
		u.ID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ContactsUnblockRequest.
var (
	_ bin.Encoder = &ContactsUnblockRequest{}
	_ bin.Decoder = &ContactsUnblockRequest{}
)

// ContactsUnblock invokes method contacts.unblock#bea65d50 returning error if any.
// Deletes the user from the blacklist.
//
// Possible errors:
//  400 CONTACT_ID_INVALID: The provided contact ID is invalid
//
// See https://core.telegram.org/method/contacts.unblock for reference.
func (c *Client) ContactsUnblock(ctx context.Context, id InputPeerClass) (bool, error) {
	var result BoolBox

	request := &ContactsUnblockRequest{
		ID: id,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
