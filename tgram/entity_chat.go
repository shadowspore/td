package tgram

import (
	"context"
	"fmt"
	"io"

	"github.com/gotd/td/telegram/downloader"
	"github.com/gotd/td/telegram/query/messages"
	"github.com/gotd/td/tg"
	"golang.org/x/xerrors"
)

var _ Entity = (*Chat)(nil)

// Chat is a telegram chat.
type Chat struct {
	chat   *tg.Chat
	client *Client
}

// ID of the chat.
func (c *Chat) ID() int64 { return c.chat.ID }

// Name of the chat.
func (c *Chat) Name() string { return c.chat.Title }

// SendText sends text message to the chat.
func (c *Chat) SendText(ctx context.Context, text string) (*Message, error) {
	return c.client.SendText(ctx, c, text)
}

// ForwardMessages forwards messages to the chat.
func (c *Chat) ForwardMessages(ctx context.Context, from Entity, msgs []*Message) ([]*Message, error) {
	msgIDs := make([]int, 0, len(msgs))
	for _, msg := range msgs {
		msgIDs = append(msgIDs, msg.ID())
	}

	return c.client.ForwardMessages(ctx, from, c, msgIDs)
}

// Invite invites user to the chat.
//
// forwardLimit is the number of messages which invited user can see.
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

// ChatParticipantType is the type of chat participant.
type ChatParticipantType string

const (
	ChatParticipantTypeParticipant ChatParticipantType = "participant"
	ChatParticipantTypeAdmin       ChatParticipantType = "admin"
	ChatParticipantTypeCreator     ChatParticipantType = "creator"
)

// ChatParticipant represents Telegram's chat participant.
type ChatParticipant struct {
	User *User
	Type ChatParticipantType
}

// Participants iterates over participants of the chat.
func (c *Chat) Participants(ctx context.Context) ([]*ChatParticipant, error) {
	resp, err := c.client.telegram.API().MessagesGetFullChat(ctx, c.chat.ID)
	if err != nil {
		return nil, err
	}

	full, ok := resp.FullChat.(*tg.ChatFull)
	if !ok {
		return nil, xerrors.Errorf("expected ChatFull, got: %T", err)
	}

	var (
		ents = c.client.ents(
			resp.Users,
			resp.Chats,
		)

		result []*ChatParticipant
	)

	switch p := full.Participants.(type) {
	case *tg.ChatParticipants:
		for _, part := range p.Participants {
			u, err := ents.User(part.GetUserID())
			if err != nil {
				return nil, fmt.Errorf("get participant %d: %w", part.GetUserID(), err)
			}

			var typ ChatParticipantType
			switch part := part.(type) {
			case *tg.ChatParticipant:
				typ = ChatParticipantTypeParticipant
			case *tg.ChatParticipantAdmin:
				typ = ChatParticipantTypeAdmin
			case *tg.ChatParticipantCreator:
				typ = ChatParticipantTypeCreator
			default:
				return nil, xerrors.Errorf("unexpected chat participant type: %T", part)
			}

			result = append(result, &ChatParticipant{
				User: u,
				Type: typ,
			})
		}
	}

	return result, nil
}

// Messages returns chat messages iterator.
func (c *Chat) Messages(opts MessageIterOptions) *MessageIterator {
	q := messages.NewQueryBuilder(c.client.telegram.API()).
		GetHistory(c.AsInputPeer())
	if opts.BatchSize > 1 {
		q = q.BatchSize(opts.BatchSize)
	}
	if opts.OffsetDate > 0 {
		q = q.OffsetDate(opts.OffsetDate)
	}
	if opts.OffsetID > 0 {
		q = q.OffsetID(opts.OffsetID)
	}

	return &MessageIterator{
		iter:   q.Iter(),
		client: c.client,
	}
}

func (c *Chat) HasProfilePhoto() bool {
	switch c.chat.Photo.(type) {
	case *tg.ChatPhoto:
		return true
	default:
		return false
	}
}

// DownloadPhoto downloads chat profile photo.
//
// big - Whether to download the high-quality version of the picture.
func (c *Chat) DownloadProfilePhoto(ctx context.Context, big bool, w io.Writer) error {
	switch p := c.chat.Photo.(type) {
	case *tg.ChatPhoto:
		conn, err := c.client.telegram.ConnectDC(ctx, p.DCID)
		if err != nil {
			return err
		}
		defer func() { _ = conn.Close() }()

		_, err = downloader.NewDownloader().
			Download(
				tg.NewClient(conn),
				&tg.InputPeerPhotoFileLocation{
					Big:     big,
					Peer:    c.AsInputPeer(),
					PhotoID: p.PhotoID,
				},
			).Stream(ctx, w)
		return err
	case *tg.ChatPhotoEmpty:
		return xerrors.Errorf("photo empty")
	default:
		panic("unreachable")
	}
}
