package tgram

import (
	"context"

	"github.com/gotd/td/tg"
)

type Entity interface {
	ID() int
	Name() string
	SendText(ctx context.Context, text string) (*Message, error)
	ForwardMessages(ctx context.Context, from Entity, msgs []*Message) ([]*Message, error)
	AsInputPeer() tg.InputPeerClass
}
