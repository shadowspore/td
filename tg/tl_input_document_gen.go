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

// InputDocumentEmpty represents TL type `inputDocumentEmpty#72f0eaae`.
// Empty constructor.
//
// See https://core.telegram.org/constructor/inputDocumentEmpty for reference.
type InputDocumentEmpty struct {
}

// InputDocumentEmptyTypeID is TL type id of InputDocumentEmpty.
const InputDocumentEmptyTypeID = 0x72f0eaae

func (i *InputDocumentEmpty) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputDocumentEmpty) String() string {
	if i == nil {
		return "InputDocumentEmpty(nil)"
	}
	type Alias InputDocumentEmpty
	return fmt.Sprintf("InputDocumentEmpty%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputDocumentEmpty) TypeID() uint32 {
	return InputDocumentEmptyTypeID
}

// TypeName returns name of type in TL schema.
func (*InputDocumentEmpty) TypeName() string {
	return "inputDocumentEmpty"
}

// TypeInfo returns info about TL type.
func (i *InputDocumentEmpty) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputDocumentEmpty",
		ID:   InputDocumentEmptyTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputDocumentEmpty) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputDocumentEmpty#72f0eaae as nil")
	}
	b.PutID(InputDocumentEmptyTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputDocumentEmpty) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputDocumentEmpty#72f0eaae to nil")
	}
	if err := b.ConsumeID(InputDocumentEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode inputDocumentEmpty#72f0eaae: %w", err)
	}
	return nil
}

// construct implements constructor of InputDocumentClass.
func (i InputDocumentEmpty) construct() InputDocumentClass { return &i }

// Ensuring interfaces in compile-time for InputDocumentEmpty.
var (
	_ bin.Encoder = &InputDocumentEmpty{}
	_ bin.Decoder = &InputDocumentEmpty{}

	_ InputDocumentClass = &InputDocumentEmpty{}
)

// InputDocument represents TL type `inputDocument#1abfb575`.
// Defines a video for subsequent interaction.
//
// See https://core.telegram.org/constructor/inputDocument for reference.
type InputDocument struct {
	// Document ID
	ID int64
	// access_hash parameter from the document¹ constructor
	//
	// Links:
	//  1) https://core.telegram.org/constructor/document
	AccessHash int64
	// File reference¹
	//
	// Links:
	//  1) https://core.telegram.org/api/file_reference
	FileReference []byte
}

// InputDocumentTypeID is TL type id of InputDocument.
const InputDocumentTypeID = 0x1abfb575

func (i *InputDocument) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ID == 0) {
		return false
	}
	if !(i.AccessHash == 0) {
		return false
	}
	if !(i.FileReference == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputDocument) String() string {
	if i == nil {
		return "InputDocument(nil)"
	}
	type Alias InputDocument
	return fmt.Sprintf("InputDocument%+v", Alias(*i))
}

// FillFrom fills InputDocument from given interface.
func (i *InputDocument) FillFrom(from interface {
	GetID() (value int64)
	GetAccessHash() (value int64)
	GetFileReference() (value []byte)
}) {
	i.ID = from.GetID()
	i.AccessHash = from.GetAccessHash()
	i.FileReference = from.GetFileReference()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputDocument) TypeID() uint32 {
	return InputDocumentTypeID
}

// TypeName returns name of type in TL schema.
func (*InputDocument) TypeName() string {
	return "inputDocument"
}

// TypeInfo returns info about TL type.
func (i *InputDocument) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputDocument",
		ID:   InputDocumentTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "AccessHash",
			SchemaName: "access_hash",
		},
		{
			Name:       "FileReference",
			SchemaName: "file_reference",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputDocument) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputDocument#1abfb575 as nil")
	}
	b.PutID(InputDocumentTypeID)
	b.PutLong(i.ID)
	b.PutLong(i.AccessHash)
	b.PutBytes(i.FileReference)
	return nil
}

// GetID returns value of ID field.
func (i *InputDocument) GetID() (value int64) {
	return i.ID
}

// GetAccessHash returns value of AccessHash field.
func (i *InputDocument) GetAccessHash() (value int64) {
	return i.AccessHash
}

// GetFileReference returns value of FileReference field.
func (i *InputDocument) GetFileReference() (value []byte) {
	return i.FileReference
}

// Decode implements bin.Decoder.
func (i *InputDocument) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputDocument#1abfb575 to nil")
	}
	if err := b.ConsumeID(InputDocumentTypeID); err != nil {
		return fmt.Errorf("unable to decode inputDocument#1abfb575: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputDocument#1abfb575: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputDocument#1abfb575: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode inputDocument#1abfb575: field file_reference: %w", err)
		}
		i.FileReference = value
	}
	return nil
}

// construct implements constructor of InputDocumentClass.
func (i InputDocument) construct() InputDocumentClass { return &i }

// Ensuring interfaces in compile-time for InputDocument.
var (
	_ bin.Encoder = &InputDocument{}
	_ bin.Decoder = &InputDocument{}

	_ InputDocumentClass = &InputDocument{}
)

// InputDocumentClass represents InputDocument generic type.
//
// See https://core.telegram.org/type/InputDocument for reference.
//
// Example:
//  g, err := tg.DecodeInputDocument(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.InputDocumentEmpty: // inputDocumentEmpty#72f0eaae
//  case *tg.InputDocument: // inputDocument#1abfb575
//  default: panic(v)
//  }
type InputDocumentClass interface {
	bin.Encoder
	bin.Decoder
	construct() InputDocumentClass

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

	// AsNotEmpty tries to map InputDocumentClass to InputDocument.
	AsNotEmpty() (*InputDocument, bool)
}

// AsInputDocumentFileLocation tries to map InputDocument to InputDocumentFileLocation.
func (i *InputDocument) AsInputDocumentFileLocation() *InputDocumentFileLocation {
	value := new(InputDocumentFileLocation)
	value.ID = i.GetID()
	value.AccessHash = i.GetAccessHash()
	value.FileReference = i.GetFileReference()

	return value
}

// AsNotEmpty tries to map InputDocumentEmpty to InputDocument.
func (i *InputDocumentEmpty) AsNotEmpty() (*InputDocument, bool) {
	return nil, false
}

// AsNotEmpty tries to map InputDocument to InputDocument.
func (i *InputDocument) AsNotEmpty() (*InputDocument, bool) {
	return i, true
}

// DecodeInputDocument implements binary de-serialization for InputDocumentClass.
func DecodeInputDocument(buf *bin.Buffer) (InputDocumentClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputDocumentEmptyTypeID:
		// Decoding inputDocumentEmpty#72f0eaae.
		v := InputDocumentEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputDocumentClass: %w", err)
		}
		return &v, nil
	case InputDocumentTypeID:
		// Decoding inputDocument#1abfb575.
		v := InputDocument{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputDocumentClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputDocumentClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputDocument boxes the InputDocumentClass providing a helper.
type InputDocumentBox struct {
	InputDocument InputDocumentClass
}

// Decode implements bin.Decoder for InputDocumentBox.
func (b *InputDocumentBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputDocumentBox to nil")
	}
	v, err := DecodeInputDocument(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputDocument = v
	return nil
}

// Encode implements bin.Encode for InputDocumentBox.
func (b *InputDocumentBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputDocument == nil {
		return fmt.Errorf("unable to encode InputDocumentClass as nil")
	}
	return b.InputDocument.Encode(buf)
}

// InputDocumentClassArray is adapter for slice of InputDocumentClass.
type InputDocumentClassArray []InputDocumentClass

// Sort sorts slice of InputDocumentClass.
func (s InputDocumentClassArray) Sort(less func(a, b InputDocumentClass) bool) InputDocumentClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputDocumentClass.
func (s InputDocumentClassArray) SortStable(less func(a, b InputDocumentClass) bool) InputDocumentClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputDocumentClass.
func (s InputDocumentClassArray) Retain(keep func(x InputDocumentClass) bool) InputDocumentClassArray {
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
func (s InputDocumentClassArray) First() (v InputDocumentClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputDocumentClassArray) Last() (v InputDocumentClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputDocumentClassArray) PopFirst() (v InputDocumentClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputDocumentClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputDocumentClassArray) Pop() (v InputDocumentClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsInputDocument returns copy with only InputDocument constructors.
func (s InputDocumentClassArray) AsInputDocument() (to InputDocumentArray) {
	for _, elem := range s {
		value, ok := elem.(*InputDocument)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AppendOnlyNotEmpty appends only NotEmpty constructors to
// given slice.
func (s InputDocumentClassArray) AppendOnlyNotEmpty(to []*InputDocument) []*InputDocument {
	for _, elem := range s {
		value, ok := elem.AsNotEmpty()
		if !ok {
			continue
		}
		to = append(to, value)
	}

	return to
}

// AsNotEmpty returns copy with only NotEmpty constructors.
func (s InputDocumentClassArray) AsNotEmpty() (to []*InputDocument) {
	return s.AppendOnlyNotEmpty(to)
}

// FirstAsNotEmpty returns first element of slice (if exists).
func (s InputDocumentClassArray) FirstAsNotEmpty() (v *InputDocument, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// LastAsNotEmpty returns last element of slice (if exists).
func (s InputDocumentClassArray) LastAsNotEmpty() (v *InputDocument, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopFirstAsNotEmpty returns element of slice (if exists).
func (s *InputDocumentClassArray) PopFirstAsNotEmpty() (v *InputDocument, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopAsNotEmpty returns element of slice (if exists).
func (s *InputDocumentClassArray) PopAsNotEmpty() (v *InputDocument, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// InputDocumentArray is adapter for slice of InputDocument.
type InputDocumentArray []InputDocument

// Sort sorts slice of InputDocument.
func (s InputDocumentArray) Sort(less func(a, b InputDocument) bool) InputDocumentArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputDocument.
func (s InputDocumentArray) SortStable(less func(a, b InputDocument) bool) InputDocumentArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputDocument.
func (s InputDocumentArray) Retain(keep func(x InputDocument) bool) InputDocumentArray {
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
func (s InputDocumentArray) First() (v InputDocument, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputDocumentArray) Last() (v InputDocument, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputDocumentArray) PopFirst() (v InputDocument, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputDocument
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputDocumentArray) Pop() (v InputDocument, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
