package unpack

import "github.com/gotd/td/tg"

func UsersChats(u tg.UpdatesClass) (users []tg.UserClass, chats []tg.ChatClass) {
	switch u := u.(type) {
	case *tg.Updates:
		return u.Users, u.Chats
	case *tg.UpdatesCombined:
		return u.Users, u.Chats
	default:
		return nil, nil
	}
}
