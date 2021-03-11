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

// MessagesMessageViews represents TL type `messages.messageViews#b6c4f543`.
// View, forward counter + info about replies
//
// See https://core.telegram.org/constructor/messages.messageViews for reference.
type MessagesMessageViews struct {
	// View, forward counter + info about replies
	Views []MessageViews
	// Chats mentioned in constructor
	Chats []ChatClass
	// Users mentioned in constructor
	Users []UserClass
}

// MessagesMessageViewsTypeID is TL type id of MessagesMessageViews.
const MessagesMessageViewsTypeID = 0xb6c4f543

func (m *MessagesMessageViews) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.Views == nil) {
		return false
	}
	if !(m.Chats == nil) {
		return false
	}
	if !(m.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessagesMessageViews) String() string {
	if m == nil {
		return "MessagesMessageViews(nil)"
	}
	type Alias MessagesMessageViews
	return fmt.Sprintf("MessagesMessageViews%+v", Alias(*m))
}

// FillFrom fills MessagesMessageViews from given interface.
func (m *MessagesMessageViews) FillFrom(from interface {
	GetViews() (value []MessageViews)
	GetChats() (value []ChatClass)
	GetUsers() (value []UserClass)
}) {
	m.Views = from.GetViews()
	m.Chats = from.GetChats()
	m.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesMessageViews) TypeID() uint32 {
	return MessagesMessageViewsTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesMessageViews) TypeName() string {
	return "messages.messageViews"
}

// TypeInfo returns info about TL type.
func (m *MessagesMessageViews) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.messageViews",
		ID:   MessagesMessageViewsTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Views",
			SchemaName: "views",
		},
		{
			Name:       "Chats",
			SchemaName: "chats",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessagesMessageViews) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messages.messageViews#b6c4f543 as nil")
	}
	b.PutID(MessagesMessageViewsTypeID)
	b.PutVectorHeader(len(m.Views))
	for idx, v := range m.Views {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.messageViews#b6c4f543: field views element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(m.Chats))
	for idx, v := range m.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode messages.messageViews#b6c4f543: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.messageViews#b6c4f543: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(m.Users))
	for idx, v := range m.Users {
		if v == nil {
			return fmt.Errorf("unable to encode messages.messageViews#b6c4f543: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.messageViews#b6c4f543: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetViews returns value of Views field.
func (m *MessagesMessageViews) GetViews() (value []MessageViews) {
	return m.Views
}

// GetChats returns value of Chats field.
func (m *MessagesMessageViews) GetChats() (value []ChatClass) {
	return m.Chats
}

// MapChats returns field Chats wrapped in ChatClassArray helper.
func (m *MessagesMessageViews) MapChats() (value ChatClassArray) {
	return ChatClassArray(m.Chats)
}

// GetUsers returns value of Users field.
func (m *MessagesMessageViews) GetUsers() (value []UserClass) {
	return m.Users
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (m *MessagesMessageViews) MapUsers() (value UserClassArray) {
	return UserClassArray(m.Users)
}

// Decode implements bin.Decoder.
func (m *MessagesMessageViews) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messages.messageViews#b6c4f543 to nil")
	}
	if err := b.ConsumeID(MessagesMessageViewsTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.messageViews#b6c4f543: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.messageViews#b6c4f543: field views: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value MessageViews
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode messages.messageViews#b6c4f543: field views: %w", err)
			}
			m.Views = append(m.Views, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.messageViews#b6c4f543: field chats: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.messageViews#b6c4f543: field chats: %w", err)
			}
			m.Chats = append(m.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.messageViews#b6c4f543: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.messageViews#b6c4f543: field users: %w", err)
			}
			m.Users = append(m.Users, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesMessageViews.
var (
	_ bin.Encoder = &MessagesMessageViews{}
	_ bin.Decoder = &MessagesMessageViews{}
)
