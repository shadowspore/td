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

// UploadFile represents TL type `upload.file#96a18d5`.
// File content.
//
// See https://core.telegram.org/constructor/upload.file for reference.
type UploadFile struct {
	// File type
	Type StorageFileTypeClass
	// Modification type
	Mtime int
	// Binary data, file content
	Bytes []byte
}

// UploadFileTypeID is TL type id of UploadFile.
const UploadFileTypeID = 0x96a18d5

func (f *UploadFile) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.Type == nil) {
		return false
	}
	if !(f.Mtime == 0) {
		return false
	}
	if !(f.Bytes == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *UploadFile) String() string {
	if f == nil {
		return "UploadFile(nil)"
	}
	type Alias UploadFile
	return fmt.Sprintf("UploadFile%+v", Alias(*f))
}

// FillFrom fills UploadFile from given interface.
func (f *UploadFile) FillFrom(from interface {
	GetType() (value StorageFileTypeClass)
	GetMtime() (value int)
	GetBytes() (value []byte)
}) {
	f.Type = from.GetType()
	f.Mtime = from.GetMtime()
	f.Bytes = from.GetBytes()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UploadFile) TypeID() uint32 {
	return UploadFileTypeID
}

// TypeName returns name of type in TL schema.
func (*UploadFile) TypeName() string {
	return "upload.file"
}

// TypeInfo returns info about TL type.
func (f *UploadFile) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "upload.file",
		ID:   UploadFileTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Type",
			SchemaName: "type",
		},
		{
			Name:       "Mtime",
			SchemaName: "mtime",
		},
		{
			Name:       "Bytes",
			SchemaName: "bytes",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (f *UploadFile) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode upload.file#96a18d5 as nil")
	}
	b.PutID(UploadFileTypeID)
	if f.Type == nil {
		return fmt.Errorf("unable to encode upload.file#96a18d5: field type is nil")
	}
	if err := f.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode upload.file#96a18d5: field type: %w", err)
	}
	b.PutInt(f.Mtime)
	b.PutBytes(f.Bytes)
	return nil
}

// GetType returns value of Type field.
func (f *UploadFile) GetType() (value StorageFileTypeClass) {
	return f.Type
}

// GetMtime returns value of Mtime field.
func (f *UploadFile) GetMtime() (value int) {
	return f.Mtime
}

// GetBytes returns value of Bytes field.
func (f *UploadFile) GetBytes() (value []byte) {
	return f.Bytes
}

// Decode implements bin.Decoder.
func (f *UploadFile) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode upload.file#96a18d5 to nil")
	}
	if err := b.ConsumeID(UploadFileTypeID); err != nil {
		return fmt.Errorf("unable to decode upload.file#96a18d5: %w", err)
	}
	{
		value, err := DecodeStorageFileType(b)
		if err != nil {
			return fmt.Errorf("unable to decode upload.file#96a18d5: field type: %w", err)
		}
		f.Type = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode upload.file#96a18d5: field mtime: %w", err)
		}
		f.Mtime = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode upload.file#96a18d5: field bytes: %w", err)
		}
		f.Bytes = value
	}
	return nil
}

// construct implements constructor of UploadFileClass.
func (f UploadFile) construct() UploadFileClass { return &f }

// Ensuring interfaces in compile-time for UploadFile.
var (
	_ bin.Encoder = &UploadFile{}
	_ bin.Decoder = &UploadFile{}

	_ UploadFileClass = &UploadFile{}
)

// UploadFileCDNRedirect represents TL type `upload.fileCdnRedirect#f18cda44`.
// The file must be downloaded from a CDN DC¹.
//
// Links:
//  1) https://core.telegram.org/cdn
//
// See https://core.telegram.org/constructor/upload.fileCdnRedirect for reference.
type UploadFileCDNRedirect struct {
	// CDN DC¹ ID
	//
	// Links:
	//  1) https://core.telegram.org/cdn
	DCID int
	// File token (see CDN files¹)
	//
	// Links:
	//  1) https://core.telegram.org/cdn
	FileToken []byte
	// Encryption key (see CDN files¹)
	//
	// Links:
	//  1) https://core.telegram.org/cdn
	EncryptionKey []byte
	// Encryption IV (see CDN files¹)
	//
	// Links:
	//  1) https://core.telegram.org/cdn
	EncryptionIv []byte
	// File hashes (see CDN files¹)
	//
	// Links:
	//  1) https://core.telegram.org/cdn
	FileHashes []FileHash
}

// UploadFileCDNRedirectTypeID is TL type id of UploadFileCDNRedirect.
const UploadFileCDNRedirectTypeID = 0xf18cda44

func (f *UploadFileCDNRedirect) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.DCID == 0) {
		return false
	}
	if !(f.FileToken == nil) {
		return false
	}
	if !(f.EncryptionKey == nil) {
		return false
	}
	if !(f.EncryptionIv == nil) {
		return false
	}
	if !(f.FileHashes == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *UploadFileCDNRedirect) String() string {
	if f == nil {
		return "UploadFileCDNRedirect(nil)"
	}
	type Alias UploadFileCDNRedirect
	return fmt.Sprintf("UploadFileCDNRedirect%+v", Alias(*f))
}

// FillFrom fills UploadFileCDNRedirect from given interface.
func (f *UploadFileCDNRedirect) FillFrom(from interface {
	GetDCID() (value int)
	GetFileToken() (value []byte)
	GetEncryptionKey() (value []byte)
	GetEncryptionIv() (value []byte)
	GetFileHashes() (value []FileHash)
}) {
	f.DCID = from.GetDCID()
	f.FileToken = from.GetFileToken()
	f.EncryptionKey = from.GetEncryptionKey()
	f.EncryptionIv = from.GetEncryptionIv()
	f.FileHashes = from.GetFileHashes()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UploadFileCDNRedirect) TypeID() uint32 {
	return UploadFileCDNRedirectTypeID
}

// TypeName returns name of type in TL schema.
func (*UploadFileCDNRedirect) TypeName() string {
	return "upload.fileCdnRedirect"
}

// TypeInfo returns info about TL type.
func (f *UploadFileCDNRedirect) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "upload.fileCdnRedirect",
		ID:   UploadFileCDNRedirectTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "DCID",
			SchemaName: "dc_id",
		},
		{
			Name:       "FileToken",
			SchemaName: "file_token",
		},
		{
			Name:       "EncryptionKey",
			SchemaName: "encryption_key",
		},
		{
			Name:       "EncryptionIv",
			SchemaName: "encryption_iv",
		},
		{
			Name:       "FileHashes",
			SchemaName: "file_hashes",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (f *UploadFileCDNRedirect) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode upload.fileCdnRedirect#f18cda44 as nil")
	}
	b.PutID(UploadFileCDNRedirectTypeID)
	b.PutInt(f.DCID)
	b.PutBytes(f.FileToken)
	b.PutBytes(f.EncryptionKey)
	b.PutBytes(f.EncryptionIv)
	b.PutVectorHeader(len(f.FileHashes))
	for idx, v := range f.FileHashes {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode upload.fileCdnRedirect#f18cda44: field file_hashes element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetDCID returns value of DCID field.
func (f *UploadFileCDNRedirect) GetDCID() (value int) {
	return f.DCID
}

// GetFileToken returns value of FileToken field.
func (f *UploadFileCDNRedirect) GetFileToken() (value []byte) {
	return f.FileToken
}

// GetEncryptionKey returns value of EncryptionKey field.
func (f *UploadFileCDNRedirect) GetEncryptionKey() (value []byte) {
	return f.EncryptionKey
}

// GetEncryptionIv returns value of EncryptionIv field.
func (f *UploadFileCDNRedirect) GetEncryptionIv() (value []byte) {
	return f.EncryptionIv
}

// GetFileHashes returns value of FileHashes field.
func (f *UploadFileCDNRedirect) GetFileHashes() (value []FileHash) {
	return f.FileHashes
}

// Decode implements bin.Decoder.
func (f *UploadFileCDNRedirect) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode upload.fileCdnRedirect#f18cda44 to nil")
	}
	if err := b.ConsumeID(UploadFileCDNRedirectTypeID); err != nil {
		return fmt.Errorf("unable to decode upload.fileCdnRedirect#f18cda44: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode upload.fileCdnRedirect#f18cda44: field dc_id: %w", err)
		}
		f.DCID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode upload.fileCdnRedirect#f18cda44: field file_token: %w", err)
		}
		f.FileToken = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode upload.fileCdnRedirect#f18cda44: field encryption_key: %w", err)
		}
		f.EncryptionKey = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode upload.fileCdnRedirect#f18cda44: field encryption_iv: %w", err)
		}
		f.EncryptionIv = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode upload.fileCdnRedirect#f18cda44: field file_hashes: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value FileHash
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode upload.fileCdnRedirect#f18cda44: field file_hashes: %w", err)
			}
			f.FileHashes = append(f.FileHashes, value)
		}
	}
	return nil
}

// construct implements constructor of UploadFileClass.
func (f UploadFileCDNRedirect) construct() UploadFileClass { return &f }

// Ensuring interfaces in compile-time for UploadFileCDNRedirect.
var (
	_ bin.Encoder = &UploadFileCDNRedirect{}
	_ bin.Decoder = &UploadFileCDNRedirect{}

	_ UploadFileClass = &UploadFileCDNRedirect{}
)

// UploadFileClass represents upload.File generic type.
//
// See https://core.telegram.org/type/upload.File for reference.
//
// Example:
//  g, err := tg.DecodeUploadFile(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.UploadFile: // upload.file#96a18d5
//  case *tg.UploadFileCDNRedirect: // upload.fileCdnRedirect#f18cda44
//  default: panic(v)
//  }
type UploadFileClass interface {
	bin.Encoder
	bin.Decoder
	construct() UploadFileClass

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
}

// DecodeUploadFile implements binary de-serialization for UploadFileClass.
func DecodeUploadFile(buf *bin.Buffer) (UploadFileClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case UploadFileTypeID:
		// Decoding upload.file#96a18d5.
		v := UploadFile{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UploadFileClass: %w", err)
		}
		return &v, nil
	case UploadFileCDNRedirectTypeID:
		// Decoding upload.fileCdnRedirect#f18cda44.
		v := UploadFileCDNRedirect{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UploadFileClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode UploadFileClass: %w", bin.NewUnexpectedID(id))
	}
}

// UploadFile boxes the UploadFileClass providing a helper.
type UploadFileBox struct {
	File UploadFileClass
}

// Decode implements bin.Decoder for UploadFileBox.
func (b *UploadFileBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode UploadFileBox to nil")
	}
	v, err := DecodeUploadFile(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.File = v
	return nil
}

// Encode implements bin.Encode for UploadFileBox.
func (b *UploadFileBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.File == nil {
		return fmt.Errorf("unable to encode UploadFileClass as nil")
	}
	return b.File.Encode(buf)
}

// UploadFileClassArray is adapter for slice of UploadFileClass.
type UploadFileClassArray []UploadFileClass

// Sort sorts slice of UploadFileClass.
func (s UploadFileClassArray) Sort(less func(a, b UploadFileClass) bool) UploadFileClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of UploadFileClass.
func (s UploadFileClassArray) SortStable(less func(a, b UploadFileClass) bool) UploadFileClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of UploadFileClass.
func (s UploadFileClassArray) Retain(keep func(x UploadFileClass) bool) UploadFileClassArray {
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
func (s UploadFileClassArray) First() (v UploadFileClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s UploadFileClassArray) Last() (v UploadFileClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *UploadFileClassArray) PopFirst() (v UploadFileClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero UploadFileClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *UploadFileClassArray) Pop() (v UploadFileClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsUploadFile returns copy with only UploadFile constructors.
func (s UploadFileClassArray) AsUploadFile() (to UploadFileArray) {
	for _, elem := range s {
		value, ok := elem.(*UploadFile)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsUploadFileCDNRedirect returns copy with only UploadFileCDNRedirect constructors.
func (s UploadFileClassArray) AsUploadFileCDNRedirect() (to UploadFileCDNRedirectArray) {
	for _, elem := range s {
		value, ok := elem.(*UploadFileCDNRedirect)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// UploadFileArray is adapter for slice of UploadFile.
type UploadFileArray []UploadFile

// Sort sorts slice of UploadFile.
func (s UploadFileArray) Sort(less func(a, b UploadFile) bool) UploadFileArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of UploadFile.
func (s UploadFileArray) SortStable(less func(a, b UploadFile) bool) UploadFileArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of UploadFile.
func (s UploadFileArray) Retain(keep func(x UploadFile) bool) UploadFileArray {
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
func (s UploadFileArray) First() (v UploadFile, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s UploadFileArray) Last() (v UploadFile, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *UploadFileArray) PopFirst() (v UploadFile, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero UploadFile
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *UploadFileArray) Pop() (v UploadFile, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// UploadFileCDNRedirectArray is adapter for slice of UploadFileCDNRedirect.
type UploadFileCDNRedirectArray []UploadFileCDNRedirect

// Sort sorts slice of UploadFileCDNRedirect.
func (s UploadFileCDNRedirectArray) Sort(less func(a, b UploadFileCDNRedirect) bool) UploadFileCDNRedirectArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of UploadFileCDNRedirect.
func (s UploadFileCDNRedirectArray) SortStable(less func(a, b UploadFileCDNRedirect) bool) UploadFileCDNRedirectArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of UploadFileCDNRedirect.
func (s UploadFileCDNRedirectArray) Retain(keep func(x UploadFileCDNRedirect) bool) UploadFileCDNRedirectArray {
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
func (s UploadFileCDNRedirectArray) First() (v UploadFileCDNRedirect, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s UploadFileCDNRedirectArray) Last() (v UploadFileCDNRedirect, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *UploadFileCDNRedirectArray) PopFirst() (v UploadFileCDNRedirect, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero UploadFileCDNRedirect
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *UploadFileCDNRedirectArray) Pop() (v UploadFileCDNRedirect, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
