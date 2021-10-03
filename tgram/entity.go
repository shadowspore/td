package tgram

import (
	"context"
	"io"

	"github.com/gotd/td/tg"
)

// Entity is a telegram chat, channel or user.
type Entity interface {
	// ID of the entity.
	ID() int64
	// Name of the entity.
	Name() string
	// SendText sends text message to the entity.
	SendText(ctx context.Context, text string) (*Message, error)
	// ForwardMessages forwards messages to the entity.
	ForwardMessages(ctx context.Context, from Entity, msgs []*Message) ([]*Message, error)
	// Messages returns entity messages iterator.
	Messages(opts MessageIterOptions) *MessageIterator
	// HasProfilePhoto shows whether the profile has a photo. 
	HasProfilePhoto() bool
	// DownloadPhoto downloads entity's profile photo.
	//
	// big - Whether to download the high-quality version of the picture.
	DownloadProfilePhoto(ctx context.Context, big bool, w io.Writer) error

	AsInputPeer() tg.InputPeerClass
}
