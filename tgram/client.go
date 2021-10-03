package tgram

import (
	"context"
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"

	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/updates"
	"github.com/gotd/td/telegram/updates/hook"
	"github.com/gotd/td/tg"
	"github.com/gotd/td/tgram/internal/bg"
	"github.com/gotd/td/tgram/internal/deeplink"
	"go.uber.org/zap"
	"golang.org/x/xerrors"
)

// Client is the Telegram client.
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

// New creates new telegram client.
func New(appID int, appHash string, opts Options) *Client {
	opts.setDefaults()
	c := &Client{
		handlers: newHandlers(),
		rand:     rand.New(rand.NewSource(time.Now().Unix())),
		log:      opts.Logger,
	}

	gapsConfig := updates.Config{
		Handler:       telegram.UpdateHandlerFunc(c.handleUpdates),
		Logger:        opts.Logger.Named("gaps"),
		DisableShorts: true,
	}

	var session telegram.SessionStorage
	if opts.Storage != nil {
		m := newSessionManager(opts.Storage)
		gapsConfig.Storage = m
		session = m
	} else {
		session = &telegram.FileSessionStorage{
			Path: "./session.json",
		}
	}

	c.gaps = updates.New(gapsConfig)
	c.telegram = telegram.NewClient(appID, appHash, telegram.Options{
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
		SessionStorage: session,
		Logger:         opts.Logger.Named("telegram"),
	})
	return c
}

// Connect connects client to the telegram server.
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

// SendText sends text message to the provided entity.
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

	msgs, err := mapMessages(updates, []int64{randomID})
	if err != nil {
		return nil, xerrors.Errorf("map messages: %w", err)
	}

	if len(msgs) != 1 {
		return nil, xerrors.Errorf("expected one message, got %d", len(msgs))
	}

	switch msg := msgs[0].(type) {
	case *tg.Message:
		return c.newMessageFromMessage(msg, c.entsFromUpdates(updates))
	default:
		return nil, xerrors.Errorf("unexpected message type: %T", msg)
	}
}

// ForwardMessage forwards messages from one entity to another.
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

	msgs, err := mapMessages(updates, randomIDs)
	if err != nil {
		return nil, xerrors.Errorf("map messages: %w", err)
	}

	var (
		result []*Message
		ents   = c.entsFromUpdates(updates)
	)
	for _, msg := range msgs {
		switch msg := msg.(type) {
		case *tg.Message:
			m, err := c.newMessageFromMessage(msg, ents)
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

// EditMessageText edits text of the message.
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

// DeleteMessage deletes messages with provided IDs.
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

// JoinLink joins to the chat or channel using telegram invite link.
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

	return c.entsFromUpdates(updates).ChatOrChannelOnly()
}

// ResolveEntity searches the User, Chat or Channel by its username.
func (c *Client) ResolveEntity(ctx context.Context, username string) (Entity, error) {
	result, err := c.telegram.API().ContactsResolveUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	return c.ents(result.Users, result.Chats).Peer(result.Peer)
}

// ResolveChannel searches the channel by its username.
// Returns error if channel not found, or it's a chat or user.
func (c *Client) ResolveChannel(ctx context.Context, username string) (*Channel, error) {
	ent, err := c.ResolveEntity(ctx, username)
	if err != nil {
		return nil, err
	}

	if c, ok := ent.(*Channel); ok {
		return c, nil
	}
	return nil, fmt.Errorf("%s is a %T, not channel", username, ent)
}

// ResolveChat searches the chat by its username.
// Returns error if chat not found, or it's a user or channel.
func (c *Client) ResolveChat(ctx context.Context, username string) (*Chat, error) {
	ent, err := c.ResolveEntity(ctx, username)
	if err != nil {
		return nil, err
	}

	if c, ok := ent.(*Chat); ok {
		return c, nil
	}
	return nil, fmt.Errorf("%s is a %T, not chat", username, ent)
}

// ResolveUser searches the user by its username.
// Returns error if user not found, or it's a chat or channel.
func (c *Client) ResolveUser(ctx context.Context, username string) (*User, error) {
	ent, err := c.ResolveEntity(ctx, username)
	if err != nil {
		return nil, err
	}

	if c, ok := ent.(*User); ok {
		return c, nil
	}
	return nil, fmt.Errorf("%s is a %T, not user", username, ent)
}

// Close closes the client.
func (c *Client) Close() error { return c.close() }
