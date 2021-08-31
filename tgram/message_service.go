package tgram

import (
	"github.com/gotd/td/tg"
	"github.com/gotd/td/tgram/internal/unpack"
	"golang.org/x/xerrors"
)

type ActionChatUsersJoined struct {
	Users []*User
	Chat  *Chat
}

func (c *Client) handleServiceMessage(
	msg *tg.MessageService,
	users []tg.UserClass,
	chats []tg.ChatClass,
) error {
	switch action := msg.Action.(type) {
	// case *tg.MessageActionEmpty:
	// case *tg.MessageActionChatCreate:
	// case *tg.MessageActionChatEditTitle:
	// case *tg.MessageActionChatEditPhoto:
	// case *tg.MessageActionChatDeletePhoto:
	case *tg.MessageActionChatAddUser:
		return c.handleActionChatAddUser(action, users, chats)
	// case *tg.MessageActionChatDeleteUser:
	// case *tg.MessageActionChatJoinedByLink:
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
		return c.handlers.fallback(&tg.Updates{
			Updates: []tg.UpdateClass{
				&tg.UpdateNewMessage{
					Message: msg,
				},
			},
			Users: users,
			Chats: chats,
		})
	}
}

func (c *Client) handleActionChatAddUser(
	action *tg.MessageActionChatAddUser,
	users []tg.UserClass,
	chats []tg.ChatClass,
) error {
	if len(chats) != 1 {
		return xerrors.Errorf("expected one chat, got %d", len(chats))
	}

	chat, ok := chats[0].(*tg.Chat)
	if !ok {
		return xerrors.Errorf("unexpected chat type: %T", chats[0])
	}

	usersSugared := make([]*User, 0, len(users))
	for _, id := range action.Users {
		u, ok := unpack.User(users, id)
		if !ok {
			return xerrors.Errorf("user %d not found", id)
		}

		usersSugared = append(usersSugared, &User{
			user:   u,
			client: c,
		})
	}

	return c.handlers.onChatUsersJoined(&ActionChatUsersJoined{
		Chat: &Chat{
			chat:   chat,
			client: c,
		},
		Users: usersSugared,
	})
}
