package tgram

import (
	"context"

	"github.com/gotd/td/telegram/query/messages"
	"github.com/gotd/td/tg"
)

type MessageIterator struct {
	iter   *messages.Iterator
	client *Client
}

func (m *MessageIterator) Next(ctx context.Context) *Message {
	if !m.iter.Next(ctx) {
		return nil
	}

	elem := m.iter.Value()
	switch msg := elem.Msg.(type) {
	case *tg.Message:
		return &Message{
			msg:    msg,
			client: m.client,
		}
	default:
		return nil
	}
}

func (m *MessageIterator) Error() error {
	return m.iter.Err()
}
