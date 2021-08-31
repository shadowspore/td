package tgram

import (
	"context"

	"github.com/gotd/td/telegram/thumbnail"
	"github.com/gotd/td/tg"
	"golang.org/x/xerrors"
)

var _ Entity = (*User)(nil)

type User struct {
	user   *tg.User
	client *Client
}

func (u *User) ID() int { return u.user.ID }

func (u *User) Name() string {
	fname, fok := u.user.GetFirstName()
	lname, lok := u.user.GetLastName()
	switch {
	case fok && lok:
		return fname + " " + lname
	case fok:
		return fname
	case lok:
		return lname
	default:
		return ""
	}
}

func (u *User) IsBot() bool { return u.user.Bot }

func (u *User) SendText(ctx context.Context, text string) (*Message, error) {
	return u.client.SendText(ctx, u, text)
}

func (u *User) ForwardMessages(ctx context.Context, from Entity, msgs []*Message) ([]*Message, error) {
	msgIDs := make([]int, 0, len(msgs))
	for _, msg := range msgs {
		msgIDs = append(msgIDs, msg.ID())
	}

	return u.client.ForwardMessages(ctx, from, u, msgIDs)
}

func (u *User) CommonChats(ctx context.Context, maxID, limit int) ([]Entity, error) {
	chats, err := u.client.telegram.API().MessagesGetCommonChats(ctx, &tg.MessagesGetCommonChatsRequest{
		UserID: u.asInputUser(),
		MaxID:  maxID,
		Limit:  limit,
	})
	if err != nil {
		return nil, err
	}

	var result []Entity
	for _, c := range chats.GetChats() {
		switch c := c.(type) {
		case *tg.Channel:
			result = append(result, &Channel{
				channel: c,
				client:  u.client,
			})
		case *tg.Chat:
			result = append(result, &Chat{
				chat:   c,
				client: u.client,
			})
		default:
			return nil, xerrors.Errorf("unexpected chat type: %T", c)
		}
	}

	return result, nil
}

func (u *User) asInputUser() tg.InputUserClass {
	return &tg.InputUser{
		UserID:     u.user.ID,
		AccessHash: u.user.AccessHash,
	}
}

func (u *User) AsInputPeer() tg.InputPeerClass {
	return &tg.InputPeerUser{
		UserID:     u.user.ID,
		AccessHash: u.user.AccessHash,
	}
}

func (u *User) Thumbnail() (jpeg []byte, exist bool, err error) {
	switch photo := u.user.Photo.(type) {
	case *tg.UserProfilePhoto:
		t, ok := photo.GetStrippedThumb()
		if !ok {
			return nil, false, nil
		}

		bs, err := thumbnail.Expand(t)
		if err != nil {
			return nil, false, err
		}

		return bs, true, nil
	case *tg.UserProfilePhotoEmpty:
		return nil, false, nil
	default:
		return nil, false, xerrors.Errorf("unexpected photo type: %T", photo)
	}
}
