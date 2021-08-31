package unpack

import "github.com/gotd/td/tg"

func User(users []tg.UserClass, userID int) (*tg.User, bool) {
	for _, u := range users {
		if u.GetID() != userID {
			continue
		}

		switch u := u.(type) {
		case *tg.User:
			return u, true
		default:
			return nil, false
		}
	}

	return nil, false
}
