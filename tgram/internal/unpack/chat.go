package unpack

import (
	"github.com/gotd/td/tg"
	"golang.org/x/xerrors"
)

func ChatOnly(u tg.UpdatesClass) (tg.ChatClass, error) {
	switch u := u.(type) {
	case *tg.Updates:
		if len(u.Chats) != 1 {
			return nil, xerrors.Errorf("expected single chat, got %d", len(u.Chats))
		}

		return u.Chats[0], nil
	case *tg.UpdatesCombined:
		if len(u.Chats) != 1 {
			return nil, xerrors.Errorf("expected single chat, got %d", len(u.Chats))
		}

		return u.Chats[0], nil
	default:
		return nil, xerrors.Errorf("unexpected update type: %T", u)
	}
}


func Chat(chats []tg.ChatClass, chatID int) (*tg.Chat, bool) {
	for _, c := range chats {
		if c.GetID() != chatID {
			continue
		}

		switch c := c.(type) {
		case *tg.Chat:
			return c, true
		default:
			return nil, false
		}
	}

	return nil, false
}
