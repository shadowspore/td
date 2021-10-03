package tgram

import "github.com/gotd/td/tg"

type MessageActionChatAddUser struct {
	Users []*User
	Chat  Entity
}

type MessageActionChatDeleteUser struct {
	User *User
	Chat Entity
}

func (c *Client) handleServiceMessage(msg *tg.MessageService, ents *Entities) error {
	switch action := msg.Action.(type) {
	// case *tg.MessageActionEmpty:
	// case *tg.MessageActionChatCreate:
	// case *tg.MessageActionChatEditTitle:
	// case *tg.MessageActionChatEditPhoto:
	// case *tg.MessageActionChatDeletePhoto:
	case *tg.MessageActionChatAddUser:
		chat, err := ents.ChatOrChannelOnly()
		if err != nil {
			return err
		}

		var users []*User
		for _, userID := range action.Users {
			u, err := ents.User(userID)
			if err != nil {
				return err
			}

			users = append(users, u)
		}

		return c.handlers.onChatAddUser(&MessageActionChatAddUser{
			Chat:  chat,
			Users: users,
		})
	case *tg.MessageActionChatDeleteUser:
		chat, err := ents.ChatOrChannelOnly()
		if err != nil {
			return err
		}

		u, err := ents.User(action.UserID)
		if err != nil {
			return err
		}

		return c.handlers.onChatDeleteUser(&MessageActionChatDeleteUser{
			Chat: chat,
			User: u,
		})
	//case *tg.MessageActionChatJoinedByLink:
	// case *tg.MessageActionChannelCreate:
	// case *tg.MessageActionChatMigrateTo:
	// case *tg.MessageActionChannelMigrateFrom:
	// case *tg.MessageActionPinMessage:
	// case *tg.MessageActionHistoryClear:
	// case *tg.MessageActionGameScore:
	// case *tg.MessageActionPaymentSentMe:
	// case *tg.MessageActionPaymentSent:
	// case *tg.MessageActionPhoneCall:
	// case *tg.MessageActionScreenshotTaken:
	// case *tg.MessageActionCustomAction:
	// case *tg.MessageActionBotAllowed:
	// case *tg.MessageActionSecureValuesSentMe:
	// case *tg.MessageActionSecureValuesSent:
	// case *tg.MessageActionContactSignUp:
	// case *tg.MessageActionGeoProximityReached:
	// case *tg.MessageActionGroupCall:
	// case *tg.MessageActionInviteToGroupCall:
	// case *tg.MessageActionSetMessagesTTL:
	// case *tg.MessageActionGroupCallScheduled:
	default:
		return c.handlers.onUpdate(&tg.UpdateNewMessage{
			Message:  msg,
			Pts:      -1,
			PtsCount: -1,
		}, ents)
	}
}
