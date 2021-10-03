package tgram

import (
	"context"

	"github.com/gotd/td/telegram/query/messages"
	"github.com/gotd/td/tg"
)

// MessageIterOptions contains message iteration options.
type MessageIterOptions struct {
	BatchSize  int
	OffsetDate int
	OffsetID   int
}

// MessageIterator is the dialog messages iterator.
type MessageIterator struct {
	iter   *messages.Iterator
	client *Client
}

// ForEach iterates over dialog's messages.
func (m *MessageIterator) ForEach(ctx context.Context, f func(ctx context.Context, msg *Message) error) error {
	for m.iter.Next(ctx) {
		elem := m.iter.Value()
		switch msg := elem.Msg.(type) {
		case *tg.Message:
			m, err := m.client.newMessageFromMessage(
				msg,
				m.client.ents(elem.Users, elem.Chats),
			)
			if err != nil {
				return err
			}

			if err := f(ctx, m); err != nil {
				return err
			}
		}
	}

	return nil
}
