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

// AccountUpdateDeviceLockedRequest represents TL type `account.updateDeviceLocked#38df3532`.
// When client-side passcode lock feature is enabled, will not show message texts in incoming PUSH notifications¹.
//
// Links:
//  1) https://core.telegram.org/api/push-updates
//
// See https://core.telegram.org/method/account.updateDeviceLocked for reference.
type AccountUpdateDeviceLockedRequest struct {
	// Inactivity period after which to start hiding message texts in PUSH notifications¹.
	//
	// Links:
	//  1) https://core.telegram.org/api/push-updates
	Period int
}

// AccountUpdateDeviceLockedRequestTypeID is TL type id of AccountUpdateDeviceLockedRequest.
const AccountUpdateDeviceLockedRequestTypeID = 0x38df3532

func (u *AccountUpdateDeviceLockedRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Period == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *AccountUpdateDeviceLockedRequest) String() string {
	if u == nil {
		return "AccountUpdateDeviceLockedRequest(nil)"
	}
	type Alias AccountUpdateDeviceLockedRequest
	return fmt.Sprintf("AccountUpdateDeviceLockedRequest%+v", Alias(*u))
}

// FillFrom fills AccountUpdateDeviceLockedRequest from given interface.
func (u *AccountUpdateDeviceLockedRequest) FillFrom(from interface {
	GetPeriod() (value int)
}) {
	u.Period = from.GetPeriod()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountUpdateDeviceLockedRequest) TypeID() uint32 {
	return AccountUpdateDeviceLockedRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountUpdateDeviceLockedRequest) TypeName() string {
	return "account.updateDeviceLocked"
}

// TypeInfo returns info about TL type.
func (u *AccountUpdateDeviceLockedRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.updateDeviceLocked",
		ID:   AccountUpdateDeviceLockedRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Period",
			SchemaName: "period",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *AccountUpdateDeviceLockedRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.updateDeviceLocked#38df3532 as nil")
	}
	b.PutID(AccountUpdateDeviceLockedRequestTypeID)
	b.PutInt(u.Period)
	return nil
}

// GetPeriod returns value of Period field.
func (u *AccountUpdateDeviceLockedRequest) GetPeriod() (value int) {
	return u.Period
}

// Decode implements bin.Decoder.
func (u *AccountUpdateDeviceLockedRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.updateDeviceLocked#38df3532 to nil")
	}
	if err := b.ConsumeID(AccountUpdateDeviceLockedRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.updateDeviceLocked#38df3532: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode account.updateDeviceLocked#38df3532: field period: %w", err)
		}
		u.Period = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountUpdateDeviceLockedRequest.
var (
	_ bin.Encoder = &AccountUpdateDeviceLockedRequest{}
	_ bin.Decoder = &AccountUpdateDeviceLockedRequest{}
)

// AccountUpdateDeviceLocked invokes method account.updateDeviceLocked#38df3532 returning error if any.
// When client-side passcode lock feature is enabled, will not show message texts in incoming PUSH notifications¹.
//
// Links:
//  1) https://core.telegram.org/api/push-updates
//
// See https://core.telegram.org/method/account.updateDeviceLocked for reference.
func (c *Client) AccountUpdateDeviceLocked(ctx context.Context, period int) (bool, error) {
	var result BoolBox

	request := &AccountUpdateDeviceLockedRequest{
		Period: period,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
