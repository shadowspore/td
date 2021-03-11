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

// PaymentsGetSavedInfoRequest represents TL type `payments.getSavedInfo#227d824b`.
// Get saved payment information
//
// See https://core.telegram.org/method/payments.getSavedInfo for reference.
type PaymentsGetSavedInfoRequest struct {
}

// PaymentsGetSavedInfoRequestTypeID is TL type id of PaymentsGetSavedInfoRequest.
const PaymentsGetSavedInfoRequestTypeID = 0x227d824b

func (g *PaymentsGetSavedInfoRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *PaymentsGetSavedInfoRequest) String() string {
	if g == nil {
		return "PaymentsGetSavedInfoRequest(nil)"
	}
	type Alias PaymentsGetSavedInfoRequest
	return fmt.Sprintf("PaymentsGetSavedInfoRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentsGetSavedInfoRequest) TypeID() uint32 {
	return PaymentsGetSavedInfoRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentsGetSavedInfoRequest) TypeName() string {
	return "payments.getSavedInfo"
}

// TypeInfo returns info about TL type.
func (g *PaymentsGetSavedInfoRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "payments.getSavedInfo",
		ID:   PaymentsGetSavedInfoRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *PaymentsGetSavedInfoRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode payments.getSavedInfo#227d824b as nil")
	}
	b.PutID(PaymentsGetSavedInfoRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (g *PaymentsGetSavedInfoRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode payments.getSavedInfo#227d824b to nil")
	}
	if err := b.ConsumeID(PaymentsGetSavedInfoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.getSavedInfo#227d824b: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for PaymentsGetSavedInfoRequest.
var (
	_ bin.Encoder = &PaymentsGetSavedInfoRequest{}
	_ bin.Decoder = &PaymentsGetSavedInfoRequest{}
)

// PaymentsGetSavedInfo invokes method payments.getSavedInfo#227d824b returning error if any.
// Get saved payment information
//
// See https://core.telegram.org/method/payments.getSavedInfo for reference.
func (c *Client) PaymentsGetSavedInfo(ctx context.Context) (*PaymentsSavedInfo, error) {
	var result PaymentsSavedInfo

	request := &PaymentsGetSavedInfoRequest{}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
