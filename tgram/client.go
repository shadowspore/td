package tgram

import (
	"context"
	"math/rand"
	"sync/atomic"
	"time"

	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/updates"
	"github.com/gotd/td/telegram/updates/hook"
	"github.com/gotd/td/tg"
	"github.com/gotd/td/tgram/internal/bg"
	"github.com/gotd/td/tgram/internal/deeplink"
	"github.com/gotd/td/tgram/internal/unpack"
	"go.uber.org/zap"
	"golang.org/x/xerrors"
)

type Client struct {
	telegram *telegram.Client
	rand     *rand.Rand
	handlers *handlers
	gaps     *updates.Manager
	log      *zap.Logger

	self  atomic.Value
	ctx   context.Context
	close func() error
}

func New(appID int, appHash string, opts Options) *Client {
	opts.setDefaults()
	c := &Client{
		handlers: newHandlers(),
		rand:     rand.New(rand.NewSource(time.Now().Unix())),
		log:      opts.Logger,
	}
	c.gaps = updates.New(updates.Config{
		Handler: telegram.UpdateHandlerFunc(c.handleUpdates),
		Logger:  opts.Logger.Named("gaps"),
	})

	telegram := telegram.NewClient(appID, appHash, telegram.Options{
		UpdateHandler: c.gaps,
		Middlewares: []telegram.Middleware{
			hook.UpdateHook(telegram.UpdateHandlerFunc(func(ctx context.Context, u tg.UpdatesClass) error {
				go func() {
					if err := c.gaps.Handle(ctx, u); err != nil {
						c.log.Error("Handle updates error", zap.Error(err))
					}
				}()
				return nil
			})),
		},
		SessionStorage: &telegram.FileSessionStorage{
			Path: "./session.json",
		},
		Logger: opts.Logger.Named("telegram"),
	})

	c.telegram = telegram
	return c
}

func (c *Client) Connect(ctx context.Context) error {
	ctx, close, err := bg.Connect(ctx, c.telegram)
	if err != nil {
		return err
	}

	c.ctx = ctx
	c.close = close

	status, err := c.telegram.Auth().Status(ctx)
	if err != nil {
		c.close()
		return err
	}

	if !status.Authorized {
		return nil
	}

	c.self.Store(status.User)
	return c.gaps.Auth(c.ctx, c.telegram.API(), status.User.ID, status.User.Bot, false)
}

func (c *Client) SendText(ctx context.Context, to Entity, text string) (*Message, error) {
	randomID := c.rand.Int63()
	updates, err := c.telegram.API().MessagesSendMessage(ctx, &tg.MessagesSendMessageRequest{
		Peer:     to.AsInputPeer(),
		Message:  text,
		RandomID: randomID,
	})
	if err != nil {
		return nil, err
	}

	if sent, ok := updates.(*tg.UpdateShortSentMessage); ok {
		return c.newMessageFromShortSent(to, text, sent)
	}

	msgs, err := unpack.Messages(updates, []int64{randomID})
	if err != nil {
		return nil, err
	}

	if len(msgs) != 1 {
		return nil, xerrors.Errorf("expected one message, got %d", len(msgs))
	}

	switch msg := msgs[0].(type) {
	case *tg.Message:
		users, chats := unpack.UsersChats(updates)
		return c.newMessageFromMessage(msg, users, chats)
	default:
		return nil, xerrors.Errorf("unexpected message type: %T", msg)
	}
}

func (c *Client) ForwardMessages(ctx context.Context, from, to Entity, msgIDs []int) ([]*Message, error) {
	randomIDs := make([]int64, 0, len(msgIDs))
	for i := 0; i < len(msgIDs); i++ {
		randomIDs = append(randomIDs, c.rand.Int63())
	}

	updates, err := c.telegram.API().MessagesForwardMessages(ctx, &tg.MessagesForwardMessagesRequest{
		ID:       msgIDs,
		RandomID: randomIDs,
		FromPeer: from.AsInputPeer(),
		ToPeer:   to.AsInputPeer(),
	})
	if err != nil {
		return nil, err
	}

	msgs, err := unpack.Messages(updates, randomIDs)
	if err != nil {
		return nil, err
	}

	var result []*Message
	users, chats := unpack.UsersChats(updates)
	for _, msg := range msgs {
		switch msg := msg.(type) {
		case *tg.Message:
			m, err := c.newMessageFromMessage(msg, users, chats)
			if err != nil {
				return nil, err
			}
			result = append(result, m)
		default:
			return nil, xerrors.Errorf("unexpected message type: %T", msg)
		}
	}

	return result, nil
}

func (c *Client) EditMessageText(ctx context.Context, msg *Message, text string) error {
	_, err := c.telegram.API().MessagesEditMessage(ctx, &tg.MessagesEditMessageRequest{
		Peer:    msg.Peer().AsInputPeer(),
		ID:      msg.ID(),
		Message: text,
	})
	if err != nil {
		return err
	}

	// TODO: Set edit date from response.
	msg.msg.Message = text
	return nil
}

func (c *Client) DeleteMessages(ctx context.Context, from Entity, msgIDs []int) error {
	if channel, ok := from.(*Channel); ok {
		_, err := c.telegram.API().ChannelsDeleteMessages(ctx, &tg.ChannelsDeleteMessagesRequest{
			Channel: channel.asInputChannel(),
			ID:      msgIDs,
		})
		return err
	}

	_, err := c.telegram.API().MessagesDeleteMessages(ctx, &tg.MessagesDeleteMessagesRequest{
		Revoke: true,
		ID:     msgIDs,
	})
	return err
}

func (c *Client) JoinLink(ctx context.Context, link string) (Entity, error) {
	hash := link
	if deeplink.IsDeeplinkLike(link) {
		l, err := deeplink.Expect(link, deeplink.Join)
		if err != nil {
			return nil, err
		}

		hash = l.Args.Get("invite")
	}

	updates, err := c.telegram.API().MessagesImportChatInvite(ctx, hash)
	if err != nil {
		return nil, err
	}

	chat, err := unpack.ChatOnly(updates)
	if err != nil {
		return nil, xerrors.Errorf("unpack chat: %w", err)
	}

	switch chat := chat.(type) {
	case *tg.Chat:
		return &Chat{
			chat:   chat,
			client: c,
		}, nil
	case *tg.Channel:
		return &Channel{
			channel: chat,
			client:  c,
		}, nil
	default:
		return nil, xerrors.Errorf("unexpected chat type: %T", chat)
	}
}

func (c *Client) ResolveUsername(ctx context.Context, username string) (Entity, error) {
	result, err := c.telegram.API().ContactsResolveUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	switch peer := result.Peer.(type) {
	case *tg.PeerUser:
		user, ok := unpack.User(result.Users, peer.UserID)
		if !ok {
			return nil, xerrors.Errorf("incomplete response")
		}

		return &User{
			user:   user,
			client: c,
		}, nil
	case *tg.PeerChannel:
		channel, ok := unpack.Channel(result.Chats, peer.ChannelID)
		if !ok {
			return nil, xerrors.Errorf("incomplete response")
		}

		return &Channel{
			channel: channel,
			client:  c,
		}, nil
	case *tg.PeerChat:
		chat, ok := unpack.Chat(result.Chats, peer.ChatID)
		if !ok {
			return nil, xerrors.Errorf("incomplete response")
		}

		return &Chat{
			chat:   chat,
			client: c,
		}, nil
	default:
		return nil, xerrors.Errorf("unexpected peer type: %T", peer)
	}
}

func (c *Client) Close() error { return c.close() }
