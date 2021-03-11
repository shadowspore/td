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

// HelpGetConfigRequest represents TL type `help.getConfig#c4f9186b`.
// Returns current configuration, including data center configuration.
//
// See https://core.telegram.org/method/help.getConfig for reference.
type HelpGetConfigRequest struct {
}

// HelpGetConfigRequestTypeID is TL type id of HelpGetConfigRequest.
const HelpGetConfigRequestTypeID = 0xc4f9186b

func (g *HelpGetConfigRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *HelpGetConfigRequest) String() string {
	if g == nil {
		return "HelpGetConfigRequest(nil)"
	}
	type Alias HelpGetConfigRequest
	return fmt.Sprintf("HelpGetConfigRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpGetConfigRequest) TypeID() uint32 {
	return HelpGetConfigRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpGetConfigRequest) TypeName() string {
	return "help.getConfig"
}

// TypeInfo returns info about TL type.
func (g *HelpGetConfigRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.getConfig",
		ID:   HelpGetConfigRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *HelpGetConfigRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getConfig#c4f9186b as nil")
	}
	b.PutID(HelpGetConfigRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (g *HelpGetConfigRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getConfig#c4f9186b to nil")
	}
	if err := b.ConsumeID(HelpGetConfigRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.getConfig#c4f9186b: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpGetConfigRequest.
var (
	_ bin.Encoder = &HelpGetConfigRequest{}
	_ bin.Decoder = &HelpGetConfigRequest{}
)

// HelpGetConfig invokes method help.getConfig#c4f9186b returning error if any.
// Returns current configuration, including data center configuration.
//
// Possible errors:
//  400 CONNECTION_API_ID_INVALID: The provided API id is invalid
//  400 CONNECTION_APP_VERSION_EMPTY: App version is empty
//  400 CONNECTION_DEVICE_MODEL_EMPTY: Device model empty
//  400 CONNECTION_LANG_PACK_INVALID: Language pack invalid
//  400 CONNECTION_LAYER_INVALID: Layer invalid
//  400 CONNECTION_NOT_INITED: Connection not initialized
//  400 CONNECTION_SYSTEM_EMPTY: Connection system empty
//  400 CONNECTION_SYSTEM_LANG_CODE_EMPTY: The system_lang_code field is empty
//  400 DATA_INVALID: Encrypted data invalid
//  400 INPUT_LAYER_INVALID: The provided layer is invalid
//  400 MSG_ID_INVALID: Invalid message ID provided
//
// See https://core.telegram.org/method/help.getConfig for reference.
// Can be used by bots.
func (c *Client) HelpGetConfig(ctx context.Context) (*Config, error) {
	var result Config

	request := &HelpGetConfigRequest{}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
