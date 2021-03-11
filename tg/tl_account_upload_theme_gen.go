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

// AccountUploadThemeRequest represents TL type `account.uploadTheme#1c3db333`.
// Upload theme
//
// See https://core.telegram.org/method/account.uploadTheme for reference.
type AccountUploadThemeRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Theme file uploaded as described in files »¹
	//
	// Links:
	//  1) https://core.telegram.org/api/files
	File InputFileClass
	// Thumbnail
	//
	// Use SetThumb and GetThumb helpers.
	Thumb InputFileClass
	// File name
	FileName string
	// MIME type, must be application/x-tgtheme-{format}, where format depends on the client
	MimeType string
}

// AccountUploadThemeRequestTypeID is TL type id of AccountUploadThemeRequest.
const AccountUploadThemeRequestTypeID = 0x1c3db333

func (u *AccountUploadThemeRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Flags.Zero()) {
		return false
	}
	if !(u.File == nil) {
		return false
	}
	if !(u.Thumb == nil) {
		return false
	}
	if !(u.FileName == "") {
		return false
	}
	if !(u.MimeType == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *AccountUploadThemeRequest) String() string {
	if u == nil {
		return "AccountUploadThemeRequest(nil)"
	}
	type Alias AccountUploadThemeRequest
	return fmt.Sprintf("AccountUploadThemeRequest%+v", Alias(*u))
}

// FillFrom fills AccountUploadThemeRequest from given interface.
func (u *AccountUploadThemeRequest) FillFrom(from interface {
	GetFile() (value InputFileClass)
	GetThumb() (value InputFileClass, ok bool)
	GetFileName() (value string)
	GetMimeType() (value string)
}) {
	u.File = from.GetFile()
	if val, ok := from.GetThumb(); ok {
		u.Thumb = val
	}

	u.FileName = from.GetFileName()
	u.MimeType = from.GetMimeType()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountUploadThemeRequest) TypeID() uint32 {
	return AccountUploadThemeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountUploadThemeRequest) TypeName() string {
	return "account.uploadTheme"
}

// TypeInfo returns info about TL type.
func (u *AccountUploadThemeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.uploadTheme",
		ID:   AccountUploadThemeRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "File",
			SchemaName: "file",
		},
		{
			Name:       "Thumb",
			SchemaName: "thumb",
			Null:       !u.Flags.Has(0),
		},
		{
			Name:       "FileName",
			SchemaName: "file_name",
		},
		{
			Name:       "MimeType",
			SchemaName: "mime_type",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *AccountUploadThemeRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.uploadTheme#1c3db333 as nil")
	}
	b.PutID(AccountUploadThemeRequestTypeID)
	if !(u.Thumb == nil) {
		u.Flags.Set(0)
	}
	if err := u.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.uploadTheme#1c3db333: field flags: %w", err)
	}
	if u.File == nil {
		return fmt.Errorf("unable to encode account.uploadTheme#1c3db333: field file is nil")
	}
	if err := u.File.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.uploadTheme#1c3db333: field file: %w", err)
	}
	if u.Flags.Has(0) {
		if u.Thumb == nil {
			return fmt.Errorf("unable to encode account.uploadTheme#1c3db333: field thumb is nil")
		}
		if err := u.Thumb.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.uploadTheme#1c3db333: field thumb: %w", err)
		}
	}
	b.PutString(u.FileName)
	b.PutString(u.MimeType)
	return nil
}

// GetFile returns value of File field.
func (u *AccountUploadThemeRequest) GetFile() (value InputFileClass) {
	return u.File
}

// SetThumb sets value of Thumb conditional field.
func (u *AccountUploadThemeRequest) SetThumb(value InputFileClass) {
	u.Flags.Set(0)
	u.Thumb = value
}

// GetThumb returns value of Thumb conditional field and
// boolean which is true if field was set.
func (u *AccountUploadThemeRequest) GetThumb() (value InputFileClass, ok bool) {
	if !u.Flags.Has(0) {
		return value, false
	}
	return u.Thumb, true
}

// GetFileName returns value of FileName field.
func (u *AccountUploadThemeRequest) GetFileName() (value string) {
	return u.FileName
}

// GetMimeType returns value of MimeType field.
func (u *AccountUploadThemeRequest) GetMimeType() (value string) {
	return u.MimeType
}

// Decode implements bin.Decoder.
func (u *AccountUploadThemeRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.uploadTheme#1c3db333 to nil")
	}
	if err := b.ConsumeID(AccountUploadThemeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.uploadTheme#1c3db333: %w", err)
	}
	{
		if err := u.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.uploadTheme#1c3db333: field flags: %w", err)
		}
	}
	{
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.uploadTheme#1c3db333: field file: %w", err)
		}
		u.File = value
	}
	if u.Flags.Has(0) {
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.uploadTheme#1c3db333: field thumb: %w", err)
		}
		u.Thumb = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.uploadTheme#1c3db333: field file_name: %w", err)
		}
		u.FileName = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.uploadTheme#1c3db333: field mime_type: %w", err)
		}
		u.MimeType = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountUploadThemeRequest.
var (
	_ bin.Encoder = &AccountUploadThemeRequest{}
	_ bin.Decoder = &AccountUploadThemeRequest{}
)

// AccountUploadTheme invokes method account.uploadTheme#1c3db333 returning error if any.
// Upload theme
//
// Possible errors:
//  400 THEME_FILE_INVALID: Invalid theme file provided
//
// See https://core.telegram.org/method/account.uploadTheme for reference.
func (c *Client) AccountUploadTheme(ctx context.Context, request *AccountUploadThemeRequest) (DocumentClass, error) {
	var result DocumentBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Document, nil
}
