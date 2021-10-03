package tgram

import (
	"context"
	"io"

	"github.com/gotd/td/telegram/downloader"
	"github.com/gotd/td/telegram/query/messages"
	"github.com/gotd/td/telegram/thumbnail"
	"github.com/gotd/td/tg"
	"golang.org/x/xerrors"
)

var _ Entity = (*User)(nil)

// User is a telegram user.
type User struct {
	user   *tg.User
	client *Client
}

// ID of the user.
func (u *User) ID() int64 { return u.user.ID }

// Name of the user.
func (u *User) Name() string {
	fname, fok := u.user.GetFirstName()
	lname, lok := u.user.GetLastName()
	uname, uok := u.user.GetUsername()
	switch {
	case fok && lok:
		return fname + " " + lname
	case fok:
		return fname
	case lok:
		return lname
	case uok:
		return uname
	default:
		return ""
	}
}

// Username of the user.
func (u *User) Username() (string, bool) { return u.user.GetUsername() }

// IsBot whether user is bot.
func (u *User) IsBot() bool { return u.user.Bot }

// SendText sends text message the user.
func (u *User) SendText(ctx context.Context, text string) (*Message, error) {
	return u.client.SendText(ctx, u, text)
}

// ForwardMessages forwards messages to the user.
func (u *User) ForwardMessages(ctx context.Context, from Entity, msgs []*Message) ([]*Message, error) {
	msgIDs := make([]int, 0, len(msgs))
	for _, msg := range msgs {
		msgIDs = append(msgIDs, msg.ID())
	}

	return u.client.ForwardMessages(ctx, from, u, msgIDs)
}

// CommonChats returns common chats with the user.
func (u *User) CommonChats(ctx context.Context, maxID int64, limit int) ([]Entity, error) {
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

// Thumbnail returns User's profile thumbnail.
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

// Messages returns messages iterator.
func (u *User) Messages(opts MessageIterOptions) *MessageIterator {
	q := messages.NewQueryBuilder(u.client.telegram.API()).
		GetHistory(u.AsInputPeer())
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
		client: u.client,
	}
}

// func (u *User) SearchMessages() {
// 	q := messages.NewQueryBuilder(u.client.telegram.API()).
// 		Search(u.AsInputPeer()).
// }

func (u *User) HasProfilePhoto() bool {
	p, ok := u.user.GetPhoto()
	if !ok {
		return false
	}

	switch p.(type) {
	case *tg.UserProfilePhoto:
		return true
	default:
		return false
	}
}

// DownloadProfilePhoto downloads User's profile photo.
//
// big - Whether to download the high-quality version of the picture.
func (u *User) DownloadProfilePhoto(ctx context.Context, big bool, w io.Writer) error {
	p, ok := u.user.GetPhoto()
	if !ok {
		return xerrors.Errorf("no photo")
	}

	switch p := p.(type) {
	case *tg.UserProfilePhoto:
		conn, err := u.client.telegram.ConnectDC(ctx, p.DCID)
		if err != nil {
			return err
		}
		defer func() { _ = conn.Close() }()

		_, err = downloader.NewDownloader().
			Download(
				tg.NewClient(conn),
				&tg.InputPeerPhotoFileLocation{
					Big:     big,
					Peer:    u.AsInputPeer(),
					PhotoID: p.PhotoID,
				},
			).Stream(ctx, w)
		return err
	case *tg.UserProfilePhotoEmpty:
		return xerrors.Errorf("photo empty")
	default:
		panic("unreachable")
	}
}
