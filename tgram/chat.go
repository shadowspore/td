package tgram

import (
	"context"

	"github.com/gotd/td/tg"
	"github.com/gotd/td/tgram/internal/unpack"
	"golang.org/x/xerrors"
)

var _ Entity = (*Chat)(nil)

type Chat struct {
	chat   *tg.Chat
	client *Client
}

func (c *Chat) ID() int { return c.chat.ID }

func (c *Chat) Name() string { return c.chat.Title }

func (c *Chat) SendText(ctx context.Context, text string) (*Message, error) {
	return c.client.SendText(ctx, c, text)
}

func (c *Chat) ForwardMessages(ctx context.Context, from Entity, msgs []*Message) ([]*Message, error) {
	msgIDs := make([]int, 0, len(msgs))
	for _, msg := range msgs {
		msgIDs = append(msgIDs, msg.ID())
	}

	return c.client.ForwardMessages(ctx, from, c, msgIDs)
}

func (c *Chat) Invite(ctx context.Context, user *User, forwardLimit int) error {
	// TODO: Check response.
	_, err := c.client.telegram.API().MessagesAddChatUser(ctx, &tg.MessagesAddChatUserRequest{
		ChatID:   c.chat.ID,
		UserID:   user.asInputUser(),
		FwdLimit: forwardLimit,
	})
	return err
}

func (c *Chat) AsInputPeer() tg.InputPeerClass {
	return &tg.InputPeerChat{ChatID: c.chat.ID}
}

type ChatParticipantType string

const (
	ChatParticipantTypeParticipant ChatParticipantType = "participant"
	ChatParticipantTypeAdmin       ChatParticipantType = "admin"
	ChatParticipantTypeCreator     ChatParticipantType = "creator"
)

type ChatParticipant struct {
	User *User
	Type ChatParticipantType
}

func (c *Chat) Participants(ctx context.Context) ([]*ChatParticipant, error) {
	resp, err := c.client.telegram.API().MessagesGetFullChat(ctx, c.chat.ID)
	if err != nil {
		return nil, err
	}

	full, ok := resp.FullChat.(*tg.ChatFull)
	if !ok {
		return nil, xerrors.Errorf("expected ChatFull, got: %T", err)
	}

	var result []*ChatParticipant
	switch p := full.Participants.(type) {
	case *tg.ChatParticipants:
		for _, part := range p.Participants {
			u, ok := unpack.User(resp.Users, part.GetUserID())
			if !ok {
				return nil, xerrors.Errorf("user %d not found", part.GetUserID())
			}

			switch part := part.(type) {
			case *tg.ChatParticipant:
				result = append(result, &ChatParticipant{
					User: &User{user: u, client: c.client},
					Type: ChatParticipantTypeParticipant,
				})
			case *tg.ChatParticipantAdmin:
				result = append(result, &ChatParticipant{
					User: &User{user: u, client: c.client},
					Type: ChatParticipantTypeAdmin,
				})
			case *tg.ChatParticipantCreator:
				result = append(result, &ChatParticipant{
					User: &User{user: u, client: c.client},
					Type: ChatParticipantTypeCreator,
				})
			default:
				return nil, xerrors.Errorf("unexpected chat participant type: %T", part)
			}
		}
	}

	return result, nil
}

func (c *Chat) Photo(ctx context.Context) (*Photo, error) {
	resp, err := c.client.telegram.API().MessagesGetFullChat(ctx, c.chat.ID)
	if err != nil {
		return nil, err
	}

	full, ok := resp.FullChat.(*tg.ChatFull)
	if !ok {
		return nil, xerrors.Errorf("expected ChatFull, got: %T", err)
	}

	photo, ok := full.GetChatPhoto()
	if !ok {
		return nil, xerrors.New("chat has no photo")
	}

	switch photo := photo.(type) {
	case *tg.Photo:
		return &Photo{
			photo:  photo,
			client: c.client,
		}, nil
	case *tg.PhotoEmpty:
		return nil, xerrors.New("chat has no photo")
	default:
		return nil, xerrors.Errorf("unexpected photo type: %T", photo)
	}
}
