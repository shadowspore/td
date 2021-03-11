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

// MessagesFeaturedStickersNotModified represents TL type `messages.featuredStickersNotModified#c6dc0c66`.
// Featured stickers haven't changed
//
// See https://core.telegram.org/constructor/messages.featuredStickersNotModified for reference.
type MessagesFeaturedStickersNotModified struct {
	// Total number of featured stickers
	Count int
}

// MessagesFeaturedStickersNotModifiedTypeID is TL type id of MessagesFeaturedStickersNotModified.
const MessagesFeaturedStickersNotModifiedTypeID = 0xc6dc0c66

func (f *MessagesFeaturedStickersNotModified) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.Count == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *MessagesFeaturedStickersNotModified) String() string {
	if f == nil {
		return "MessagesFeaturedStickersNotModified(nil)"
	}
	type Alias MessagesFeaturedStickersNotModified
	return fmt.Sprintf("MessagesFeaturedStickersNotModified%+v", Alias(*f))
}

// FillFrom fills MessagesFeaturedStickersNotModified from given interface.
func (f *MessagesFeaturedStickersNotModified) FillFrom(from interface {
	GetCount() (value int)
}) {
	f.Count = from.GetCount()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesFeaturedStickersNotModified) TypeID() uint32 {
	return MessagesFeaturedStickersNotModifiedTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesFeaturedStickersNotModified) TypeName() string {
	return "messages.featuredStickersNotModified"
}

// TypeInfo returns info about TL type.
func (f *MessagesFeaturedStickersNotModified) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.featuredStickersNotModified",
		ID:   MessagesFeaturedStickersNotModifiedTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Count",
			SchemaName: "count",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (f *MessagesFeaturedStickersNotModified) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode messages.featuredStickersNotModified#c6dc0c66 as nil")
	}
	b.PutID(MessagesFeaturedStickersNotModifiedTypeID)
	b.PutInt(f.Count)
	return nil
}

// GetCount returns value of Count field.
func (f *MessagesFeaturedStickersNotModified) GetCount() (value int) {
	return f.Count
}

// Decode implements bin.Decoder.
func (f *MessagesFeaturedStickersNotModified) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode messages.featuredStickersNotModified#c6dc0c66 to nil")
	}
	if err := b.ConsumeID(MessagesFeaturedStickersNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.featuredStickersNotModified#c6dc0c66: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.featuredStickersNotModified#c6dc0c66: field count: %w", err)
		}
		f.Count = value
	}
	return nil
}

// construct implements constructor of MessagesFeaturedStickersClass.
func (f MessagesFeaturedStickersNotModified) construct() MessagesFeaturedStickersClass { return &f }

// Ensuring interfaces in compile-time for MessagesFeaturedStickersNotModified.
var (
	_ bin.Encoder = &MessagesFeaturedStickersNotModified{}
	_ bin.Decoder = &MessagesFeaturedStickersNotModified{}

	_ MessagesFeaturedStickersClass = &MessagesFeaturedStickersNotModified{}
)

// MessagesFeaturedStickers represents TL type `messages.featuredStickers#b6abc341`.
// Featured stickersets
//
// See https://core.telegram.org/constructor/messages.featuredStickers for reference.
type MessagesFeaturedStickers struct {
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int
	// Total number of featured stickers
	Count int
	// Featured stickersets
	Sets []StickerSetCoveredClass
	// IDs of new featured stickersets
	Unread []int64
}

// MessagesFeaturedStickersTypeID is TL type id of MessagesFeaturedStickers.
const MessagesFeaturedStickersTypeID = 0xb6abc341

func (f *MessagesFeaturedStickers) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.Hash == 0) {
		return false
	}
	if !(f.Count == 0) {
		return false
	}
	if !(f.Sets == nil) {
		return false
	}
	if !(f.Unread == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *MessagesFeaturedStickers) String() string {
	if f == nil {
		return "MessagesFeaturedStickers(nil)"
	}
	type Alias MessagesFeaturedStickers
	return fmt.Sprintf("MessagesFeaturedStickers%+v", Alias(*f))
}

// FillFrom fills MessagesFeaturedStickers from given interface.
func (f *MessagesFeaturedStickers) FillFrom(from interface {
	GetHash() (value int)
	GetCount() (value int)
	GetSets() (value []StickerSetCoveredClass)
	GetUnread() (value []int64)
}) {
	f.Hash = from.GetHash()
	f.Count = from.GetCount()
	f.Sets = from.GetSets()
	f.Unread = from.GetUnread()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesFeaturedStickers) TypeID() uint32 {
	return MessagesFeaturedStickersTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesFeaturedStickers) TypeName() string {
	return "messages.featuredStickers"
}

// TypeInfo returns info about TL type.
func (f *MessagesFeaturedStickers) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.featuredStickers",
		ID:   MessagesFeaturedStickersTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
		{
			Name:       "Count",
			SchemaName: "count",
		},
		{
			Name:       "Sets",
			SchemaName: "sets",
		},
		{
			Name:       "Unread",
			SchemaName: "unread",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (f *MessagesFeaturedStickers) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode messages.featuredStickers#b6abc341 as nil")
	}
	b.PutID(MessagesFeaturedStickersTypeID)
	b.PutInt(f.Hash)
	b.PutInt(f.Count)
	b.PutVectorHeader(len(f.Sets))
	for idx, v := range f.Sets {
		if v == nil {
			return fmt.Errorf("unable to encode messages.featuredStickers#b6abc341: field sets element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.featuredStickers#b6abc341: field sets element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(f.Unread))
	for _, v := range f.Unread {
		b.PutLong(v)
	}
	return nil
}

// GetHash returns value of Hash field.
func (f *MessagesFeaturedStickers) GetHash() (value int) {
	return f.Hash
}

// GetCount returns value of Count field.
func (f *MessagesFeaturedStickers) GetCount() (value int) {
	return f.Count
}

// GetSets returns value of Sets field.
func (f *MessagesFeaturedStickers) GetSets() (value []StickerSetCoveredClass) {
	return f.Sets
}

// MapSets returns field Sets wrapped in StickerSetCoveredClassArray helper.
func (f *MessagesFeaturedStickers) MapSets() (value StickerSetCoveredClassArray) {
	return StickerSetCoveredClassArray(f.Sets)
}

// GetUnread returns value of Unread field.
func (f *MessagesFeaturedStickers) GetUnread() (value []int64) {
	return f.Unread
}

// Decode implements bin.Decoder.
func (f *MessagesFeaturedStickers) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode messages.featuredStickers#b6abc341 to nil")
	}
	if err := b.ConsumeID(MessagesFeaturedStickersTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.featuredStickers#b6abc341: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.featuredStickers#b6abc341: field hash: %w", err)
		}
		f.Hash = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.featuredStickers#b6abc341: field count: %w", err)
		}
		f.Count = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.featuredStickers#b6abc341: field sets: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeStickerSetCovered(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.featuredStickers#b6abc341: field sets: %w", err)
			}
			f.Sets = append(f.Sets, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.featuredStickers#b6abc341: field unread: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode messages.featuredStickers#b6abc341: field unread: %w", err)
			}
			f.Unread = append(f.Unread, value)
		}
	}
	return nil
}

// construct implements constructor of MessagesFeaturedStickersClass.
func (f MessagesFeaturedStickers) construct() MessagesFeaturedStickersClass { return &f }

// Ensuring interfaces in compile-time for MessagesFeaturedStickers.
var (
	_ bin.Encoder = &MessagesFeaturedStickers{}
	_ bin.Decoder = &MessagesFeaturedStickers{}

	_ MessagesFeaturedStickersClass = &MessagesFeaturedStickers{}
)

// MessagesFeaturedStickersClass represents messages.FeaturedStickers generic type.
//
// See https://core.telegram.org/type/messages.FeaturedStickers for reference.
//
// Example:
//  g, err := tg.DecodeMessagesFeaturedStickers(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.MessagesFeaturedStickersNotModified: // messages.featuredStickersNotModified#c6dc0c66
//  case *tg.MessagesFeaturedStickers: // messages.featuredStickers#b6abc341
//  default: panic(v)
//  }
type MessagesFeaturedStickersClass interface {
	bin.Encoder
	bin.Decoder
	construct() MessagesFeaturedStickersClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	// Total number of featured stickers
	GetCount() (value int)
	// AsModified tries to map MessagesFeaturedStickersClass to MessagesFeaturedStickers.
	AsModified() (*MessagesFeaturedStickers, bool)
}

// AsModified tries to map MessagesFeaturedStickersNotModified to MessagesFeaturedStickers.
func (f *MessagesFeaturedStickersNotModified) AsModified() (*MessagesFeaturedStickers, bool) {
	return nil, false
}

// AsModified tries to map MessagesFeaturedStickers to MessagesFeaturedStickers.
func (f *MessagesFeaturedStickers) AsModified() (*MessagesFeaturedStickers, bool) {
	return f, true
}

// DecodeMessagesFeaturedStickers implements binary de-serialization for MessagesFeaturedStickersClass.
func DecodeMessagesFeaturedStickers(buf *bin.Buffer) (MessagesFeaturedStickersClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessagesFeaturedStickersNotModifiedTypeID:
		// Decoding messages.featuredStickersNotModified#c6dc0c66.
		v := MessagesFeaturedStickersNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesFeaturedStickersClass: %w", err)
		}
		return &v, nil
	case MessagesFeaturedStickersTypeID:
		// Decoding messages.featuredStickers#b6abc341.
		v := MessagesFeaturedStickers{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesFeaturedStickersClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessagesFeaturedStickersClass: %w", bin.NewUnexpectedID(id))
	}
}

// MessagesFeaturedStickers boxes the MessagesFeaturedStickersClass providing a helper.
type MessagesFeaturedStickersBox struct {
	FeaturedStickers MessagesFeaturedStickersClass
}

// Decode implements bin.Decoder for MessagesFeaturedStickersBox.
func (b *MessagesFeaturedStickersBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessagesFeaturedStickersBox to nil")
	}
	v, err := DecodeMessagesFeaturedStickers(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.FeaturedStickers = v
	return nil
}

// Encode implements bin.Encode for MessagesFeaturedStickersBox.
func (b *MessagesFeaturedStickersBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.FeaturedStickers == nil {
		return fmt.Errorf("unable to encode MessagesFeaturedStickersClass as nil")
	}
	return b.FeaturedStickers.Encode(buf)
}

// MessagesFeaturedStickersClassArray is adapter for slice of MessagesFeaturedStickersClass.
type MessagesFeaturedStickersClassArray []MessagesFeaturedStickersClass

// Sort sorts slice of MessagesFeaturedStickersClass.
func (s MessagesFeaturedStickersClassArray) Sort(less func(a, b MessagesFeaturedStickersClass) bool) MessagesFeaturedStickersClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessagesFeaturedStickersClass.
func (s MessagesFeaturedStickersClassArray) SortStable(less func(a, b MessagesFeaturedStickersClass) bool) MessagesFeaturedStickersClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessagesFeaturedStickersClass.
func (s MessagesFeaturedStickersClassArray) Retain(keep func(x MessagesFeaturedStickersClass) bool) MessagesFeaturedStickersClassArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s MessagesFeaturedStickersClassArray) First() (v MessagesFeaturedStickersClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessagesFeaturedStickersClassArray) Last() (v MessagesFeaturedStickersClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessagesFeaturedStickersClassArray) PopFirst() (v MessagesFeaturedStickersClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessagesFeaturedStickersClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessagesFeaturedStickersClassArray) Pop() (v MessagesFeaturedStickersClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsMessagesFeaturedStickersNotModified returns copy with only MessagesFeaturedStickersNotModified constructors.
func (s MessagesFeaturedStickersClassArray) AsMessagesFeaturedStickersNotModified() (to MessagesFeaturedStickersNotModifiedArray) {
	for _, elem := range s {
		value, ok := elem.(*MessagesFeaturedStickersNotModified)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsMessagesFeaturedStickers returns copy with only MessagesFeaturedStickers constructors.
func (s MessagesFeaturedStickersClassArray) AsMessagesFeaturedStickers() (to MessagesFeaturedStickersArray) {
	for _, elem := range s {
		value, ok := elem.(*MessagesFeaturedStickers)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AppendOnlyModified appends only Modified constructors to
// given slice.
func (s MessagesFeaturedStickersClassArray) AppendOnlyModified(to []*MessagesFeaturedStickers) []*MessagesFeaturedStickers {
	for _, elem := range s {
		value, ok := elem.AsModified()
		if !ok {
			continue
		}
		to = append(to, value)
	}

	return to
}

// AsModified returns copy with only Modified constructors.
func (s MessagesFeaturedStickersClassArray) AsModified() (to []*MessagesFeaturedStickers) {
	return s.AppendOnlyModified(to)
}

// FirstAsModified returns first element of slice (if exists).
func (s MessagesFeaturedStickersClassArray) FirstAsModified() (v *MessagesFeaturedStickers, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsModified()
}

// LastAsModified returns last element of slice (if exists).
func (s MessagesFeaturedStickersClassArray) LastAsModified() (v *MessagesFeaturedStickers, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsModified()
}

// PopFirstAsModified returns element of slice (if exists).
func (s *MessagesFeaturedStickersClassArray) PopFirstAsModified() (v *MessagesFeaturedStickers, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsModified()
}

// PopAsModified returns element of slice (if exists).
func (s *MessagesFeaturedStickersClassArray) PopAsModified() (v *MessagesFeaturedStickers, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsModified()
}

// MessagesFeaturedStickersNotModifiedArray is adapter for slice of MessagesFeaturedStickersNotModified.
type MessagesFeaturedStickersNotModifiedArray []MessagesFeaturedStickersNotModified

// Sort sorts slice of MessagesFeaturedStickersNotModified.
func (s MessagesFeaturedStickersNotModifiedArray) Sort(less func(a, b MessagesFeaturedStickersNotModified) bool) MessagesFeaturedStickersNotModifiedArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessagesFeaturedStickersNotModified.
func (s MessagesFeaturedStickersNotModifiedArray) SortStable(less func(a, b MessagesFeaturedStickersNotModified) bool) MessagesFeaturedStickersNotModifiedArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessagesFeaturedStickersNotModified.
func (s MessagesFeaturedStickersNotModifiedArray) Retain(keep func(x MessagesFeaturedStickersNotModified) bool) MessagesFeaturedStickersNotModifiedArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s MessagesFeaturedStickersNotModifiedArray) First() (v MessagesFeaturedStickersNotModified, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessagesFeaturedStickersNotModifiedArray) Last() (v MessagesFeaturedStickersNotModified, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessagesFeaturedStickersNotModifiedArray) PopFirst() (v MessagesFeaturedStickersNotModified, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessagesFeaturedStickersNotModified
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessagesFeaturedStickersNotModifiedArray) Pop() (v MessagesFeaturedStickersNotModified, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// MessagesFeaturedStickersArray is adapter for slice of MessagesFeaturedStickers.
type MessagesFeaturedStickersArray []MessagesFeaturedStickers

// Sort sorts slice of MessagesFeaturedStickers.
func (s MessagesFeaturedStickersArray) Sort(less func(a, b MessagesFeaturedStickers) bool) MessagesFeaturedStickersArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessagesFeaturedStickers.
func (s MessagesFeaturedStickersArray) SortStable(less func(a, b MessagesFeaturedStickers) bool) MessagesFeaturedStickersArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessagesFeaturedStickers.
func (s MessagesFeaturedStickersArray) Retain(keep func(x MessagesFeaturedStickers) bool) MessagesFeaturedStickersArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s MessagesFeaturedStickersArray) First() (v MessagesFeaturedStickers, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessagesFeaturedStickersArray) Last() (v MessagesFeaturedStickers, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessagesFeaturedStickersArray) PopFirst() (v MessagesFeaturedStickers, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessagesFeaturedStickers
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessagesFeaturedStickersArray) Pop() (v MessagesFeaturedStickers, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
