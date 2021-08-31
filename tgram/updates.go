package tgram

import (
	"context"

	"github.com/gotd/td/tg"
	"go.uber.org/zap"
)

func (c *Client) handleUpdates(ctx context.Context, u tg.UpdatesClass) error {
	switch u := u.(type) {
	case *tg.Updates:
		for _, upd := range u.Updates {
			if err := c.handleUpdate(upd, u.Users, u.Chats); err != nil {
				c.log.Error("Handle update error", zap.Error(err))
			}
		}
	case *tg.UpdatesCombined:
		for _, upd := range u.Updates {
			if err := c.handleUpdate(upd, u.Users, u.Chats); err != nil {
				c.log.Error("Handle update error", zap.Error(err))
			}
		}

	default:
		return c.handlers.fallback(u)
	}

	return nil
}

func (c *Client) handleUpdate(u tg.UpdateClass, users []tg.UserClass, chats []tg.ChatClass) error {
	onNew := func(msg tg.MessageClass) error {
		switch msg := msg.(type) {
		case *tg.Message:
			m, err := c.newMessageFromMessage(msg, users, chats)
			if err != nil {
				return err
			}
			return c.handlers.onNewMessage(m)
		case *tg.MessageService:
			return c.handleServiceMessage(msg, users, chats)
		default:
			c.log.Debug("Unexpected message", zap.Any("msg", msg))
			return nil
		}
	}

	onEdit := func(msg tg.MessageClass) error {
		switch msg := msg.(type) {
		case *tg.Message:
			m, err := c.newMessageFromMessage(msg, users, chats)
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
		return c.handlers.fallback(&tg.Updates{
			Updates: []tg.UpdateClass{u},
			Users:   users,
			Chats:   chats,
		})
	}
}
