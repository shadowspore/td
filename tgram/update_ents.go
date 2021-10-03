package tgram

import (
	"fmt"

	"github.com/gotd/td/tg"
	"golang.org/x/xerrors"
)

type Entities struct {
	client *Client
	users  []tg.UserClass
	chats  []tg.ChatClass
}

func (c *Client) entsFromUpdates(updates tg.UpdatesClass) *Entities {
	switch u := updates.(type) {
	case *tg.Updates:
		return c.ents(u.Users, u.Chats)
	case *tg.UpdatesCombined:
		return c.ents(u.Users, u.Chats)
	default:
		return c.ents(nil, nil)
	}
}

func (c *Client) ents(users []tg.UserClass, chats []tg.ChatClass) *Entities {
	return &Entities{
		client: c,
		users:  users,
		chats:  chats,
	}
}

func (e *Entities) user(u tg.UserClass) (*User, error) {
	switch u := u.(type) {
	case *tg.User:
		return &User{
			user:   u,
			client: e.client,
		}, nil
	case *tg.UserEmpty:
		return nil, fmt.Errorf("user empty")
	default:
		panic(fmt.Sprintf("unexpected tg.UserClass: %T", u))
	}
}

func (e *Entities) chat(c tg.ChatClass) (*Chat, error) {
	switch c := c.(type) {
	case *tg.ChatEmpty:
		return nil, fmt.Errorf("chat empty")
	case *tg.Chat:
		return &Chat{
			chat:   c,
			client: e.client,
		}, nil
	case *tg.ChatForbidden:
		return nil, fmt.Errorf("chat forbidden")
	case *tg.Channel:
		return nil, fmt.Errorf("expected chat, got channel")
	case *tg.ChannelForbidden:
		return nil, fmt.Errorf("expected chat, got channel")
	default:
		panic(fmt.Sprintf("unexpected tg.ChatClass: %T", c))
	}
}

func (e *Entities) channel(c tg.ChatClass) (*Channel, error) {
	switch c := c.(type) {
	case *tg.ChatEmpty:
		return nil, fmt.Errorf("expected channel, but got chat")
	case *tg.Chat:
		return nil, fmt.Errorf("expected channel, but got chat")
	case *tg.ChatForbidden:
		return nil, fmt.Errorf("expected channel, but got chat")
	case *tg.Channel:
		return &Channel{
			channel: c,
			client:  e.client,
		}, nil
	case *tg.ChannelForbidden:
		return nil, fmt.Errorf("channel forbidden")
	default:
		panic(fmt.Sprintf("unexpected tg.ChatClass: %T", c))
	}
}

func (e *Entities) User(userID int64) (*User, error) {
	for _, u := range e.users {
		if u.GetID() == userID {
			return e.user(u)
		}
	}
	return nil, fmt.Errorf("user not found")
}

func (e *Entities) Chat(chatID int64) (*Chat, error) {
	for _, c := range e.chats {
		if c.GetID() == chatID {
			return e.chat(c)
		}
	}
	return nil, fmt.Errorf("chat not found")
}

func (e *Entities) Channel(channelID int64) (*Channel, error) {
	for _, c := range e.chats {
		if c.GetID() == channelID {
			return e.channel(c)
		}
	}
	return nil, fmt.Errorf("channel not found")
}

func (e *Entities) Peer(peer tg.PeerClass) (Entity, error) {
	switch p := peer.(type) {
	case *tg.PeerChannel:
		return e.Channel(p.ChannelID)
	case *tg.PeerChat:
		return e.Chat(p.ChatID)
	case *tg.PeerUser:
		return e.User(p.UserID)
	default:
		return nil, xerrors.Errorf("unexpected peer type: %T", p)
	}
}

func (e *Entities) UserOnly() (*User, error) {
	switch len(e.users) {
	case 0:
		return nil, fmt.Errorf("no users")
	case 1:
		return e.user(e.users[0])
	default:
		return nil, fmt.Errorf("expected one user, but have many")
	}
}

func (e *Entities) ChatOnly() (*Chat, error) {
	switch len(e.chats) {
	case 0:
		return nil, fmt.Errorf("no chats")
	case 1:
		return e.chat(e.chats[0])
	default:
		return nil, fmt.Errorf("expected one chat, but have many")
	}
}

func (e *Entities) ChannelOnly() (*Channel, error) {
	switch len(e.chats) {
	case 0:
		return nil, fmt.Errorf("no channels")
	case 1:
		return e.channel(e.chats[0])
	default:
		return nil, fmt.Errorf("expected one channel, but have many")
	}
}

func (e *Entities) ChatOrChannelOnly() (Entity, error) {
	switch len(e.chats) {
	case 0:
		return nil, fmt.Errorf("no chats or channels")
	case 1:
		switch c := e.chats[0].(type) {
		case *tg.Chat, *tg.ChatEmpty, *tg.ChatForbidden:
			return e.chat(c)
		case *tg.Channel, *tg.ChannelForbidden:
			return e.channel(c)
		default:
			panic(fmt.Sprintf("unexpected tg.ChatClass: %T", c))
		}
	default:
		return nil, fmt.Errorf("expected one chat or channel, but have many")
	}
}
