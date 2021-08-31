package tgram

import (
	"github.com/gotd/td/tg"
	"github.com/gotd/td/tgram/internal/unpack"
	"golang.org/x/xerrors"
)

func (c *Client) peerSugared(peer tg.PeerClass, users []tg.UserClass, chats []tg.ChatClass) (Entity, error) {
	switch p := peer.(type) {
	case *tg.PeerChannel:
		channel, ok := unpack.Channel(chats, p.ChannelID)
		if !ok {
			return nil, xerrors.Errorf("channel %d not found", p.ChannelID)
		}
		return &Channel{channel: channel, client: c}, nil
	case *tg.PeerChat:
		chat, ok := unpack.Chat(chats, p.ChatID)
		if !ok {
			return nil, xerrors.Errorf("chat %d not found", p.ChatID)
		}
		return &Chat{chat: chat, client: c}, nil
	case *tg.PeerUser:
		user, ok := unpack.User(users, p.UserID)
		if !ok {
			return nil, xerrors.Errorf("user %d not found", p.UserID)
		}
		return &User{user: user, client: c}, nil
	default:
		return nil, xerrors.Errorf("unexpected peer type: %T", p)
	}
}
