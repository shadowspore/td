package tgram

import (
	"context"
	"fmt"

	"github.com/gotd/td/tg"
	"go.uber.org/zap"
)

func (c *Client) handleUpdates(ctx context.Context, u tg.UpdatesClass) error {
	switch u := u.(type) {
	case *tg.Updates:
		ents := c.ents(u.Users, u.Chats)
		for _, upd := range u.Updates {
			if err := c.handleUpdate(upd, ents); err != nil {
				c.log.Error("Handle update error", zap.Error(err))
			}
		}
	case *tg.UpdatesCombined:
		ents := c.ents(u.Users, u.Chats)
		for _, upd := range u.Updates {
			if err := c.handleUpdate(upd, ents); err != nil {
				c.log.Error("Handle update error", zap.Error(err))
			}
		}

	default:
		c.log.Warn("Received unexpected type of update",
			zap.String("update_type", fmt.Sprintf("%T", u)),
		)
	}

	return nil
}

func (c *Client) handleUpdate(u tg.UpdateClass, ents *Entities) error {
	onNew := func(msg tg.MessageClass) error {
		switch msg := msg.(type) {
		case *tg.Message:
			m, err := c.newMessageFromMessage(msg, ents)
			if err != nil {
				return err
			}
			return c.handlers.onNewMessage(m)
		case *tg.MessageService:
			return c.handleServiceMessage(msg, ents)
		default:
			c.log.Debug("Unexpected message", zap.Any("msg", msg))
			return nil
		}
	}

	onEdit := func(msg tg.MessageClass) error {
		switch msg := msg.(type) {
		case *tg.Message:
			m, err := c.newMessageFromMessage(msg, ents)
			if err != nil {
				return err
			}
			return c.handlers.onEditMessage(m)
		case *tg.MessageService:
			c.log.Warn("tg.MessageService edited")
			return nil
		default:
			c.log.Debug("Unexpected message", zap.Any("msg", msg))
			return nil
		}
	}

	switch u := u.(type) {
	case *tg.UpdateNewMessage:
		return onNew(u.Message)
	case *tg.UpdateEditMessage:
		return onEdit(u.Message)
	case *tg.UpdateNewChannelMessage:
		return onNew(u.Message)
	case *tg.UpdateEditChannelMessage:
		return onEdit(u.Message)
	default:
		return c.handlers.onUpdate(u, ents)
	}
}
