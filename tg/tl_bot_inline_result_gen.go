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

// BotInlineResult represents TL type `botInlineResult#11965f3a`.
// Generic result
//
// See https://core.telegram.org/constructor/botInlineResult for reference.
type BotInlineResult struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Result ID
	ID string
	// Result type (see bot API docs¹)
	//
	// Links:
	//  1) https://core.telegram.org/bots/api#inlinequeryresult
	Type string
	// Result title
	//
	// Use SetTitle and GetTitle helpers.
	Title string
	// Result description
	//
	// Use SetDescription and GetDescription helpers.
	Description string
	// URL of article or webpage
	//
	// Use SetURL and GetURL helpers.
	URL string
	// Thumbnail for the result
	//
	// Use SetThumb and GetThumb helpers.
	Thumb WebDocumentClass
	// Content of the result
	//
	// Use SetContent and GetContent helpers.
	Content WebDocumentClass
	// Message to send
	SendMessage BotInlineMessageClass
}

// BotInlineResultTypeID is TL type id of BotInlineResult.
const BotInlineResultTypeID = 0x11965f3a

func (b *BotInlineResult) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.Flags.Zero()) {
		return false
	}
	if !(b.ID == "") {
		return false
	}
	if !(b.Type == "") {
		return false
	}
	if !(b.Title == "") {
		return false
	}
	if !(b.Description == "") {
		return false
	}
	if !(b.URL == "") {
		return false
	}
	if !(b.Thumb == nil) {
		return false
	}
	if !(b.Content == nil) {
		return false
	}
	if !(b.SendMessage == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BotInlineResult) String() string {
	if b == nil {
		return "BotInlineResult(nil)"
	}
	type Alias BotInlineResult
	return fmt.Sprintf("BotInlineResult%+v", Alias(*b))
}

// FillFrom fills BotInlineResult from given interface.
func (b *BotInlineResult) FillFrom(from interface {
	GetID() (value string)
	GetType() (value string)
	GetTitle() (value string, ok bool)
	GetDescription() (value string, ok bool)
	GetURL() (value string, ok bool)
	GetThumb() (value WebDocumentClass, ok bool)
	GetContent() (value WebDocumentClass, ok bool)
	GetSendMessage() (value BotInlineMessageClass)
}) {
	b.ID = from.GetID()
	b.Type = from.GetType()
	if val, ok := from.GetTitle(); ok {
		b.Title = val
	}

	if val, ok := from.GetDescription(); ok {
		b.Description = val
	}

	if val, ok := from.GetURL(); ok {
		b.URL = val
	}

	if val, ok := from.GetThumb(); ok {
		b.Thumb = val
	}

	if val, ok := from.GetContent(); ok {
		b.Content = val
	}

	b.SendMessage = from.GetSendMessage()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotInlineResult) TypeID() uint32 {
	return BotInlineResultTypeID
}

// TypeName returns name of type in TL schema.
func (*BotInlineResult) TypeName() string {
	return "botInlineResult"
}

// TypeInfo returns info about TL type.
func (b *BotInlineResult) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "botInlineResult",
		ID:   BotInlineResultTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Type",
			SchemaName: "type",
		},
		{
			Name:       "Title",
			SchemaName: "title",
			Null:       !b.Flags.Has(1),
		},
		{
			Name:       "Description",
			SchemaName: "description",
			Null:       !b.Flags.Has(2),
		},
		{
			Name:       "URL",
			SchemaName: "url",
			Null:       !b.Flags.Has(3),
		},
		{
			Name:       "Thumb",
			SchemaName: "thumb",
			Null:       !b.Flags.Has(4),
		},
		{
			Name:       "Content",
			SchemaName: "content",
			Null:       !b.Flags.Has(5),
		},
		{
			Name:       "SendMessage",
			SchemaName: "send_message",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *BotInlineResult) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botInlineResult#11965f3a as nil")
	}
	buf.PutID(BotInlineResultTypeID)
	if !(b.Title == "") {
		b.Flags.Set(1)
	}
	if !(b.Description == "") {
		b.Flags.Set(2)
	}
	if !(b.URL == "") {
		b.Flags.Set(3)
	}
	if !(b.Thumb == nil) {
		b.Flags.Set(4)
	}
	if !(b.Content == nil) {
		b.Flags.Set(5)
	}
	if err := b.Flags.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode botInlineResult#11965f3a: field flags: %w", err)
	}
	buf.PutString(b.ID)
	buf.PutString(b.Type)
	if b.Flags.Has(1) {
		buf.PutString(b.Title)
	}
	if b.Flags.Has(2) {
		buf.PutString(b.Description)
	}
	if b.Flags.Has(3) {
		buf.PutString(b.URL)
	}
	if b.Flags.Has(4) {
		if b.Thumb == nil {
			return fmt.Errorf("unable to encode botInlineResult#11965f3a: field thumb is nil")
		}
		if err := b.Thumb.Encode(buf); err != nil {
			return fmt.Errorf("unable to encode botInlineResult#11965f3a: field thumb: %w", err)
		}
	}
	if b.Flags.Has(5) {
		if b.Content == nil {
			return fmt.Errorf("unable to encode botInlineResult#11965f3a: field content is nil")
		}
		if err := b.Content.Encode(buf); err != nil {
			return fmt.Errorf("unable to encode botInlineResult#11965f3a: field content: %w", err)
		}
	}
	if b.SendMessage == nil {
		return fmt.Errorf("unable to encode botInlineResult#11965f3a: field send_message is nil")
	}
	if err := b.SendMessage.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode botInlineResult#11965f3a: field send_message: %w", err)
	}
	return nil
}

// GetID returns value of ID field.
func (b *BotInlineResult) GetID() (value string) {
	return b.ID
}

// GetType returns value of Type field.
func (b *BotInlineResult) GetType() (value string) {
	return b.Type
}

// SetTitle sets value of Title conditional field.
func (b *BotInlineResult) SetTitle(value string) {
	b.Flags.Set(1)
	b.Title = value
}

// GetTitle returns value of Title conditional field and
// boolean which is true if field was set.
func (b *BotInlineResult) GetTitle() (value string, ok bool) {
	if !b.Flags.Has(1) {
		return value, false
	}
	return b.Title, true
}

// SetDescription sets value of Description conditional field.
func (b *BotInlineResult) SetDescription(value string) {
	b.Flags.Set(2)
	b.Description = value
}

// GetDescription returns value of Description conditional field and
// boolean which is true if field was set.
func (b *BotInlineResult) GetDescription() (value string, ok bool) {
	if !b.Flags.Has(2) {
		return value, false
	}
	return b.Description, true
}

// SetURL sets value of URL conditional field.
func (b *BotInlineResult) SetURL(value string) {
	b.Flags.Set(3)
	b.URL = value
}

// GetURL returns value of URL conditional field and
// boolean which is true if field was set.
func (b *BotInlineResult) GetURL() (value string, ok bool) {
	if !b.Flags.Has(3) {
		return value, false
	}
	return b.URL, true
}

// SetThumb sets value of Thumb conditional field.
func (b *BotInlineResult) SetThumb(value WebDocumentClass) {
	b.Flags.Set(4)
	b.Thumb = value
}

// GetThumb returns value of Thumb conditional field and
// boolean which is true if field was set.
func (b *BotInlineResult) GetThumb() (value WebDocumentClass, ok bool) {
	if !b.Flags.Has(4) {
		return value, false
	}
	return b.Thumb, true
}

// SetContent sets value of Content conditional field.
func (b *BotInlineResult) SetContent(value WebDocumentClass) {
	b.Flags.Set(5)
	b.Content = value
}

// GetContent returns value of Content conditional field and
// boolean which is true if field was set.
func (b *BotInlineResult) GetContent() (value WebDocumentClass, ok bool) {
	if !b.Flags.Has(5) {
		return value, false
	}
	return b.Content, true
}

// GetSendMessage returns value of SendMessage field.
func (b *BotInlineResult) GetSendMessage() (value BotInlineMessageClass) {
	return b.SendMessage
}

// Decode implements bin.Decoder.
func (b *BotInlineResult) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botInlineResult#11965f3a to nil")
	}
	if err := buf.ConsumeID(BotInlineResultTypeID); err != nil {
		return fmt.Errorf("unable to decode botInlineResult#11965f3a: %w", err)
	}
	{
		if err := b.Flags.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode botInlineResult#11965f3a: field flags: %w", err)
		}
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineResult#11965f3a: field id: %w", err)
		}
		b.ID = value
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineResult#11965f3a: field type: %w", err)
		}
		b.Type = value
	}
	if b.Flags.Has(1) {
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineResult#11965f3a: field title: %w", err)
		}
		b.Title = value
	}
	if b.Flags.Has(2) {
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineResult#11965f3a: field description: %w", err)
		}
		b.Description = value
	}
	if b.Flags.Has(3) {
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineResult#11965f3a: field url: %w", err)
		}
		b.URL = value
	}
	if b.Flags.Has(4) {
		value, err := DecodeWebDocument(buf)
		if err != nil {
			return fmt.Errorf("unable to decode botInlineResult#11965f3a: field thumb: %w", err)
		}
		b.Thumb = value
	}
	if b.Flags.Has(5) {
		value, err := DecodeWebDocument(buf)
		if err != nil {
			return fmt.Errorf("unable to decode botInlineResult#11965f3a: field content: %w", err)
		}
		b.Content = value
	}
	{
		value, err := DecodeBotInlineMessage(buf)
		if err != nil {
			return fmt.Errorf("unable to decode botInlineResult#11965f3a: field send_message: %w", err)
		}
		b.SendMessage = value
	}
	return nil
}

// construct implements constructor of BotInlineResultClass.
func (b BotInlineResult) construct() BotInlineResultClass { return &b }

// Ensuring interfaces in compile-time for BotInlineResult.
var (
	_ bin.Encoder = &BotInlineResult{}
	_ bin.Decoder = &BotInlineResult{}

	_ BotInlineResultClass = &BotInlineResult{}
)

// BotInlineMediaResult represents TL type `botInlineMediaResult#17db940b`.
// Media result
//
// See https://core.telegram.org/constructor/botInlineMediaResult for reference.
type BotInlineMediaResult struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Result ID
	ID string
	// Result type (see bot API docs¹)
	//
	// Links:
	//  1) https://core.telegram.org/bots/api#inlinequeryresult
	Type string
	// If type is photo, the photo to send
	//
	// Use SetPhoto and GetPhoto helpers.
	Photo PhotoClass
	// If type is document, the document to send
	//
	// Use SetDocument and GetDocument helpers.
	Document DocumentClass
	// Result title
	//
	// Use SetTitle and GetTitle helpers.
	Title string
	// Description
	//
	// Use SetDescription and GetDescription helpers.
	Description string
	// Depending on the type and on the constructor¹, contains the caption of the media or the content of the message to be sent instead of the media
	//
	// Links:
	//  1) https://core.telegram.org/type/BotInlineMessage
	SendMessage BotInlineMessageClass
}

// BotInlineMediaResultTypeID is TL type id of BotInlineMediaResult.
const BotInlineMediaResultTypeID = 0x17db940b

func (b *BotInlineMediaResult) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.Flags.Zero()) {
		return false
	}
	if !(b.ID == "") {
		return false
	}
	if !(b.Type == "") {
		return false
	}
	if !(b.Photo == nil) {
		return false
	}
	if !(b.Document == nil) {
		return false
	}
	if !(b.Title == "") {
		return false
	}
	if !(b.Description == "") {
		return false
	}
	if !(b.SendMessage == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BotInlineMediaResult) String() string {
	if b == nil {
		return "BotInlineMediaResult(nil)"
	}
	type Alias BotInlineMediaResult
	return fmt.Sprintf("BotInlineMediaResult%+v", Alias(*b))
}

// FillFrom fills BotInlineMediaResult from given interface.
func (b *BotInlineMediaResult) FillFrom(from interface {
	GetID() (value string)
	GetType() (value string)
	GetPhoto() (value PhotoClass, ok bool)
	GetDocument() (value DocumentClass, ok bool)
	GetTitle() (value string, ok bool)
	GetDescription() (value string, ok bool)
	GetSendMessage() (value BotInlineMessageClass)
}) {
	b.ID = from.GetID()
	b.Type = from.GetType()
	if val, ok := from.GetPhoto(); ok {
		b.Photo = val
	}

	if val, ok := from.GetDocument(); ok {
		b.Document = val
	}

	if val, ok := from.GetTitle(); ok {
		b.Title = val
	}

	if val, ok := from.GetDescription(); ok {
		b.Description = val
	}

	b.SendMessage = from.GetSendMessage()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotInlineMediaResult) TypeID() uint32 {
	return BotInlineMediaResultTypeID
}

// TypeName returns name of type in TL schema.
func (*BotInlineMediaResult) TypeName() string {
	return "botInlineMediaResult"
}

// TypeInfo returns info about TL type.
func (b *BotInlineMediaResult) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "botInlineMediaResult",
		ID:   BotInlineMediaResultTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Type",
			SchemaName: "type",
		},
		{
			Name:       "Photo",
			SchemaName: "photo",
			Null:       !b.Flags.Has(0),
		},
		{
			Name:       "Document",
			SchemaName: "document",
			Null:       !b.Flags.Has(1),
		},
		{
			Name:       "Title",
			SchemaName: "title",
			Null:       !b.Flags.Has(2),
		},
		{
			Name:       "Description",
			SchemaName: "description",
			Null:       !b.Flags.Has(3),
		},
		{
			Name:       "SendMessage",
			SchemaName: "send_message",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *BotInlineMediaResult) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botInlineMediaResult#17db940b as nil")
	}
	buf.PutID(BotInlineMediaResultTypeID)
	if !(b.Photo == nil) {
		b.Flags.Set(0)
	}
	if !(b.Document == nil) {
		b.Flags.Set(1)
	}
	if !(b.Title == "") {
		b.Flags.Set(2)
	}
	if !(b.Description == "") {
		b.Flags.Set(3)
	}
	if err := b.Flags.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode botInlineMediaResult#17db940b: field flags: %w", err)
	}
	buf.PutString(b.ID)
	buf.PutString(b.Type)
	if b.Flags.Has(0) {
		if b.Photo == nil {
			return fmt.Errorf("unable to encode botInlineMediaResult#17db940b: field photo is nil")
		}
		if err := b.Photo.Encode(buf); err != nil {
			return fmt.Errorf("unable to encode botInlineMediaResult#17db940b: field photo: %w", err)
		}
	}
	if b.Flags.Has(1) {
		if b.Document == nil {
			return fmt.Errorf("unable to encode botInlineMediaResult#17db940b: field document is nil")
		}
		if err := b.Document.Encode(buf); err != nil {
			return fmt.Errorf("unable to encode botInlineMediaResult#17db940b: field document: %w", err)
		}
	}
	if b.Flags.Has(2) {
		buf.PutString(b.Title)
	}
	if b.Flags.Has(3) {
		buf.PutString(b.Description)
	}
	if b.SendMessage == nil {
		return fmt.Errorf("unable to encode botInlineMediaResult#17db940b: field send_message is nil")
	}
	if err := b.SendMessage.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode botInlineMediaResult#17db940b: field send_message: %w", err)
	}
	return nil
}

// GetID returns value of ID field.
func (b *BotInlineMediaResult) GetID() (value string) {
	return b.ID
}

// GetType returns value of Type field.
func (b *BotInlineMediaResult) GetType() (value string) {
	return b.Type
}

// SetPhoto sets value of Photo conditional field.
func (b *BotInlineMediaResult) SetPhoto(value PhotoClass) {
	b.Flags.Set(0)
	b.Photo = value
}

// GetPhoto returns value of Photo conditional field and
// boolean which is true if field was set.
func (b *BotInlineMediaResult) GetPhoto() (value PhotoClass, ok bool) {
	if !b.Flags.Has(0) {
		return value, false
	}
	return b.Photo, true
}

// SetDocument sets value of Document conditional field.
func (b *BotInlineMediaResult) SetDocument(value DocumentClass) {
	b.Flags.Set(1)
	b.Document = value
}

// GetDocument returns value of Document conditional field and
// boolean which is true if field was set.
func (b *BotInlineMediaResult) GetDocument() (value DocumentClass, ok bool) {
	if !b.Flags.Has(1) {
		return value, false
	}
	return b.Document, true
}

// SetTitle sets value of Title conditional field.
func (b *BotInlineMediaResult) SetTitle(value string) {
	b.Flags.Set(2)
	b.Title = value
}

// GetTitle returns value of Title conditional field and
// boolean which is true if field was set.
func (b *BotInlineMediaResult) GetTitle() (value string, ok bool) {
	if !b.Flags.Has(2) {
		return value, false
	}
	return b.Title, true
}

// SetDescription sets value of Description conditional field.
func (b *BotInlineMediaResult) SetDescription(value string) {
	b.Flags.Set(3)
	b.Description = value
}

// GetDescription returns value of Description conditional field and
// boolean which is true if field was set.
func (b *BotInlineMediaResult) GetDescription() (value string, ok bool) {
	if !b.Flags.Has(3) {
		return value, false
	}
	return b.Description, true
}

// GetSendMessage returns value of SendMessage field.
func (b *BotInlineMediaResult) GetSendMessage() (value BotInlineMessageClass) {
	return b.SendMessage
}

// Decode implements bin.Decoder.
func (b *BotInlineMediaResult) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botInlineMediaResult#17db940b to nil")
	}
	if err := buf.ConsumeID(BotInlineMediaResultTypeID); err != nil {
		return fmt.Errorf("unable to decode botInlineMediaResult#17db940b: %w", err)
	}
	{
		if err := b.Flags.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode botInlineMediaResult#17db940b: field flags: %w", err)
		}
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMediaResult#17db940b: field id: %w", err)
		}
		b.ID = value
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMediaResult#17db940b: field type: %w", err)
		}
		b.Type = value
	}
	if b.Flags.Has(0) {
		value, err := DecodePhoto(buf)
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMediaResult#17db940b: field photo: %w", err)
		}
		b.Photo = value
	}
	if b.Flags.Has(1) {
		value, err := DecodeDocument(buf)
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMediaResult#17db940b: field document: %w", err)
		}
		b.Document = value
	}
	if b.Flags.Has(2) {
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMediaResult#17db940b: field title: %w", err)
		}
		b.Title = value
	}
	if b.Flags.Has(3) {
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMediaResult#17db940b: field description: %w", err)
		}
		b.Description = value
	}
	{
		value, err := DecodeBotInlineMessage(buf)
		if err != nil {
			return fmt.Errorf("unable to decode botInlineMediaResult#17db940b: field send_message: %w", err)
		}
		b.SendMessage = value
	}
	return nil
}

// construct implements constructor of BotInlineResultClass.
func (b BotInlineMediaResult) construct() BotInlineResultClass { return &b }

// Ensuring interfaces in compile-time for BotInlineMediaResult.
var (
	_ bin.Encoder = &BotInlineMediaResult{}
	_ bin.Decoder = &BotInlineMediaResult{}

	_ BotInlineResultClass = &BotInlineMediaResult{}
)

// BotInlineResultClass represents BotInlineResult generic type.
//
// See https://core.telegram.org/type/BotInlineResult for reference.
//
// Example:
//  g, err := tg.DecodeBotInlineResult(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.BotInlineResult: // botInlineResult#11965f3a
//  case *tg.BotInlineMediaResult: // botInlineMediaResult#17db940b
//  default: panic(v)
//  }
type BotInlineResultClass interface {
	bin.Encoder
	bin.Decoder
	construct() BotInlineResultClass

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

	// Result ID
	GetID() (value string)
	// Result type (see bot API docs¹)
	//
	// Links:
	//  1) https://core.telegram.org/bots/api#inlinequeryresult
	GetType() (value string)
	// Result title
	GetTitle() (value string, ok bool)
	// Result description
	GetDescription() (value string, ok bool)
	// Message to send
	GetSendMessage() (value BotInlineMessageClass)
}

// DecodeBotInlineResult implements binary de-serialization for BotInlineResultClass.
func DecodeBotInlineResult(buf *bin.Buffer) (BotInlineResultClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case BotInlineResultTypeID:
		// Decoding botInlineResult#11965f3a.
		v := BotInlineResult{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BotInlineResultClass: %w", err)
		}
		return &v, nil
	case BotInlineMediaResultTypeID:
		// Decoding botInlineMediaResult#17db940b.
		v := BotInlineMediaResult{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BotInlineResultClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode BotInlineResultClass: %w", bin.NewUnexpectedID(id))
	}
}

// BotInlineResult boxes the BotInlineResultClass providing a helper.
type BotInlineResultBox struct {
	BotInlineResult BotInlineResultClass
}

// Decode implements bin.Decoder for BotInlineResultBox.
func (b *BotInlineResultBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode BotInlineResultBox to nil")
	}
	v, err := DecodeBotInlineResult(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.BotInlineResult = v
	return nil
}

// Encode implements bin.Encode for BotInlineResultBox.
func (b *BotInlineResultBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.BotInlineResult == nil {
		return fmt.Errorf("unable to encode BotInlineResultClass as nil")
	}
	return b.BotInlineResult.Encode(buf)
}

// BotInlineResultClassArray is adapter for slice of BotInlineResultClass.
type BotInlineResultClassArray []BotInlineResultClass

// Sort sorts slice of BotInlineResultClass.
func (s BotInlineResultClassArray) Sort(less func(a, b BotInlineResultClass) bool) BotInlineResultClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of BotInlineResultClass.
func (s BotInlineResultClassArray) SortStable(less func(a, b BotInlineResultClass) bool) BotInlineResultClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of BotInlineResultClass.
func (s BotInlineResultClassArray) Retain(keep func(x BotInlineResultClass) bool) BotInlineResultClassArray {
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
func (s BotInlineResultClassArray) First() (v BotInlineResultClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s BotInlineResultClassArray) Last() (v BotInlineResultClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *BotInlineResultClassArray) PopFirst() (v BotInlineResultClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero BotInlineResultClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *BotInlineResultClassArray) Pop() (v BotInlineResultClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsBotInlineResult returns copy with only BotInlineResult constructors.
func (s BotInlineResultClassArray) AsBotInlineResult() (to BotInlineResultArray) {
	for _, elem := range s {
		value, ok := elem.(*BotInlineResult)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsBotInlineMediaResult returns copy with only BotInlineMediaResult constructors.
func (s BotInlineResultClassArray) AsBotInlineMediaResult() (to BotInlineMediaResultArray) {
	for _, elem := range s {
		value, ok := elem.(*BotInlineMediaResult)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// BotInlineResultArray is adapter for slice of BotInlineResult.
type BotInlineResultArray []BotInlineResult

// Sort sorts slice of BotInlineResult.
func (s BotInlineResultArray) Sort(less func(a, b BotInlineResult) bool) BotInlineResultArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of BotInlineResult.
func (s BotInlineResultArray) SortStable(less func(a, b BotInlineResult) bool) BotInlineResultArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of BotInlineResult.
func (s BotInlineResultArray) Retain(keep func(x BotInlineResult) bool) BotInlineResultArray {
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
func (s BotInlineResultArray) First() (v BotInlineResult, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s BotInlineResultArray) Last() (v BotInlineResult, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *BotInlineResultArray) PopFirst() (v BotInlineResult, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero BotInlineResult
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *BotInlineResultArray) Pop() (v BotInlineResult, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// BotInlineMediaResultArray is adapter for slice of BotInlineMediaResult.
type BotInlineMediaResultArray []BotInlineMediaResult

// Sort sorts slice of BotInlineMediaResult.
func (s BotInlineMediaResultArray) Sort(less func(a, b BotInlineMediaResult) bool) BotInlineMediaResultArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of BotInlineMediaResult.
func (s BotInlineMediaResultArray) SortStable(less func(a, b BotInlineMediaResult) bool) BotInlineMediaResultArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of BotInlineMediaResult.
func (s BotInlineMediaResultArray) Retain(keep func(x BotInlineMediaResult) bool) BotInlineMediaResultArray {
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
func (s BotInlineMediaResultArray) First() (v BotInlineMediaResult, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s BotInlineMediaResultArray) Last() (v BotInlineMediaResult, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *BotInlineMediaResultArray) PopFirst() (v BotInlineMediaResult, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero BotInlineMediaResult
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *BotInlineMediaResultArray) Pop() (v BotInlineMediaResult, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
