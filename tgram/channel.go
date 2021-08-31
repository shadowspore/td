package tgram

import (
	"context"

	"github.com/gotd/td/telegram/query/channels/participants"
	"github.com/gotd/td/telegram/query/messages"
	"github.com/gotd/td/tg"
)

var _ Entity = (*Channel)(nil)

type Channel struct {
	channel *tg.Channel
	client  *Client
}

func (c *Channel) ID() int { return c.channel.ID }

func (c *Channel) Name() string { return c.channel.Title }

func (c *Channel) SendText(ctx context.Context, text string) (*Message, error) {
	return c.client.SendText(ctx, c, text)
}

func (c *Channel) ForwardMessages(ctx context.Context, from Entity, msgs []*Message) ([]*Message, error) {
	msgIDs := make([]int, 0, len(msgs))
	for _, msg := range msgs {
		msgIDs = append(msgIDs, msg.ID())
	}

	return c.client.ForwardMessages(ctx, from, c, msgIDs)
}

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

type ChannelParticipant struct {
	User *User
	Type tg.ChannelParticipantClass
}

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

func (c *Channel) Messages() *MessageIterator {
	return &MessageIterator{
		iter: messages.NewQueryBuilder(c.client.telegram.API()).
			GetHistory(c.AsInputPeer()).Iter(),
		client: c.client,
	}
}
