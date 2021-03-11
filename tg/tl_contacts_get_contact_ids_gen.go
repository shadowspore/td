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

// ContactsGetContactIDsRequest represents TL type `contacts.getContactIDs#2caa4a42`.
// Get contact by telegram IDs
//
// See https://core.telegram.org/method/contacts.getContactIDs for reference.
type ContactsGetContactIDsRequest struct {
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int
}

// ContactsGetContactIDsRequestTypeID is TL type id of ContactsGetContactIDsRequest.
const ContactsGetContactIDsRequestTypeID = 0x2caa4a42

func (g *ContactsGetContactIDsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *ContactsGetContactIDsRequest) String() string {
	if g == nil {
		return "ContactsGetContactIDsRequest(nil)"
	}
	type Alias ContactsGetContactIDsRequest
	return fmt.Sprintf("ContactsGetContactIDsRequest%+v", Alias(*g))
}

// FillFrom fills ContactsGetContactIDsRequest from given interface.
func (g *ContactsGetContactIDsRequest) FillFrom(from interface {
	GetHash() (value int)
}) {
	g.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ContactsGetContactIDsRequest) TypeID() uint32 {
	return ContactsGetContactIDsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ContactsGetContactIDsRequest) TypeName() string {
	return "contacts.getContactIDs"
}

// TypeInfo returns info about TL type.
func (g *ContactsGetContactIDsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "contacts.getContactIDs",
		ID:   ContactsGetContactIDsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *ContactsGetContactIDsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode contacts.getContactIDs#2caa4a42 as nil")
	}
	b.PutID(ContactsGetContactIDsRequestTypeID)
	b.PutInt(g.Hash)
	return nil
}

// GetHash returns value of Hash field.
func (g *ContactsGetContactIDsRequest) GetHash() (value int) {
	return g.Hash
}

// Decode implements bin.Decoder.
func (g *ContactsGetContactIDsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode contacts.getContactIDs#2caa4a42 to nil")
	}
	if err := b.ConsumeID(ContactsGetContactIDsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.getContactIDs#2caa4a42: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.getContactIDs#2caa4a42: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ContactsGetContactIDsRequest.
var (
	_ bin.Encoder = &ContactsGetContactIDsRequest{}
	_ bin.Decoder = &ContactsGetContactIDsRequest{}
)

// ContactsGetContactIDs invokes method contacts.getContactIDs#2caa4a42 returning error if any.
// Get contact by telegram IDs
//
// See https://core.telegram.org/method/contacts.getContactIDs for reference.
func (c *Client) ContactsGetContactIDs(ctx context.Context, hash int) ([]int, error) {
	var result IntVector

	request := &ContactsGetContactIDsRequest{
		Hash: hash,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return []int(result.Elems), nil
}
