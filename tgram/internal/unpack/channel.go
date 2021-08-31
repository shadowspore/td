package unpack

import "github.com/gotd/td/tg"

func Channel(chats []tg.ChatClass, channelID int) (*tg.Channel, bool) {
	for _, c := range chats {
		if c.GetID() != channelID {
			continue
		}

		switch c := c.(type) {
		case *tg.Channel:
			return c, true
		default:
			return nil, false
		}
	}

	return nil, false
}
