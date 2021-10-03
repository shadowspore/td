package tgram

import (
	"context"
	"io"

	"github.com/gotd/td/telegram/downloader"
	"github.com/gotd/td/telegram/query/channels/participants"
	"github.com/gotd/td/telegram/query/messages"
	"github.com/gotd/td/tg"
	"golang.org/x/xerrors"
)

var _ Entity = (*Channel)(nil)

// Channel is a telegram channel.
type Channel struct {
	channel *tg.Channel
	client  *Client
}

// ID of the channel.
func (c *Channel) ID() int64 { return c.channel.ID }

// Name of the channel.
func (c *Channel) Name() string { return c.channel.Title }

// IsBroadcast - whether the channel is broadcast channel.
func (c *Channel) IsBroadcast() bool { return c.channel.Broadcast }

// SendText sends text message to this channel.
func (c *Channel) SendText(ctx context.Context, text string) (*Message, error) {
	return c.client.SendText(ctx, c, text)
}

// ForwardMessages forwards messages to this channel.
func (c *Channel) ForwardMessages(ctx context.Context, from Entity, msgs []*Message) ([]*Message, error) {
	msgIDs := make([]int, 0, len(msgs))
	for _, msg := range msgs {
		msgIDs = append(msgIDs, msg.ID())
	}

	return c.client.ForwardMessages(ctx, from, c, msgIDs)
}

// Invite invites users to this channel.
func (c *Channel) Invite(ctx context.Context, users []*User) error {
	inputs := make([]tg.InputUserClass, 0, len(users))
	for _, u := range users {
		inputs = append(inputs, u.asInputUser())
	}

	// TODO: Check response.
	_, err := c.client.telegram.API().ChannelsInviteToChannel(ctx, &tg.ChannelsInviteToChannelRequest{
		Channel: c.asInputChannel(),
		Users:   inputs,
	})
	return err
}

func (c *Channel) HasProfilePhoto() bool {
	switch c.channel.Photo.(type) {
	case *tg.ChatPhoto:
		return true
	default:
		return false
	}
}

// DownloadProfilePhoto downloads Channel's profile photo.
//
// big - Whether to download the high-quality version of the picture.
func (c *Channel) DownloadProfilePhoto(ctx context.Context, big bool, w io.Writer) error {
	switch p := c.channel.Photo.(type) {
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
		return xerrors.Errorf("chat photo empty")
	default:
		panic("unreachable")
	}
}

func (c *Channel) AsInputPeer() tg.InputPeerClass {
	return &tg.InputPeerChannel{
		ChannelID:  c.channel.ID,
		AccessHash: c.channel.AccessHash,
	}
}

func (c *Channel) asInputChannel() tg.InputChannelClass {
	return &tg.InputChannel{
		ChannelID:  c.channel.ID,
		AccessHash: c.channel.AccessHash,
	}
}

// ChannelParticipant represents the participant of the channel.
type ChannelParticipant struct {
	User *User
	Type tg.ChannelParticipantClass
}

// Participants iterates over channel participants.
//
// TODO: Filter options.
func (c *Channel) Participants(ctx context.Context, f func(*ChannelParticipant) error) error {
	return participants.NewQueryBuilder(c.client.telegram.API()).
		GetParticipants(c.asInputChannel()).BatchSize(50).
		ForEach(ctx, func(ctx context.Context, e participants.Elem) error {
			user, ok := e.User()
			if !ok {
				return nil
			}

			return f(&ChannelParticipant{
				User: &User{user: user, client: c.client},
				Type: e.Participant,
			})
		})
}

// Messages returns channel message iterator.
func (c *Channel) Messages(opts MessageIterOptions) *MessageIterator {
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
