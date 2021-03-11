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

// AccountGetContactSignUpNotificationRequest represents TL type `account.getContactSignUpNotification#9f07c728`.
// Whether the user will receive notifications when contacts sign up
//
// See https://core.telegram.org/method/account.getContactSignUpNotification for reference.
type AccountGetContactSignUpNotificationRequest struct {
}

// AccountGetContactSignUpNotificationRequestTypeID is TL type id of AccountGetContactSignUpNotificationRequest.
const AccountGetContactSignUpNotificationRequestTypeID = 0x9f07c728

func (g *AccountGetContactSignUpNotificationRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *AccountGetContactSignUpNotificationRequest) String() string {
	if g == nil {
		return "AccountGetContactSignUpNotificationRequest(nil)"
	}
	type Alias AccountGetContactSignUpNotificationRequest
	return fmt.Sprintf("AccountGetContactSignUpNotificationRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountGetContactSignUpNotificationRequest) TypeID() uint32 {
	return AccountGetContactSignUpNotificationRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountGetContactSignUpNotificationRequest) TypeName() string {
	return "account.getContactSignUpNotification"
}

// TypeInfo returns info about TL type.
func (g *AccountGetContactSignUpNotificationRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.getContactSignUpNotification",
		ID:   AccountGetContactSignUpNotificationRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *AccountGetContactSignUpNotificationRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getContactSignUpNotification#9f07c728 as nil")
	}
	b.PutID(AccountGetContactSignUpNotificationRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (g *AccountGetContactSignUpNotificationRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getContactSignUpNotification#9f07c728 to nil")
	}
	if err := b.ConsumeID(AccountGetContactSignUpNotificationRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getContactSignUpNotification#9f07c728: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountGetContactSignUpNotificationRequest.
var (
	_ bin.Encoder = &AccountGetContactSignUpNotificationRequest{}
	_ bin.Decoder = &AccountGetContactSignUpNotificationRequest{}
)

// AccountGetContactSignUpNotification invokes method account.getContactSignUpNotification#9f07c728 returning error if any.
// Whether the user will receive notifications when contacts sign up
//
// See https://core.telegram.org/method/account.getContactSignUpNotification for reference.
func (c *Client) AccountGetContactSignUpNotification(ctx context.Context) (bool, error) {
	var result BoolBox

	request := &AccountGetContactSignUpNotificationRequest{}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
