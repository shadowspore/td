package tgram

import (
	"context"
	"fmt"

	"github.com/gotd/td/tg"
	"golang.org/x/xerrors"
)

type Message struct {
	msg    *tg.Message
	client *Client
	peer   Entity
	from   Entity
}

func (c *Client) newMessageFromMessage(
	msg *tg.Message, users []tg.UserClass, chats []tg.ChatClass,
) (*Message, error) {
	peer, err := c.peerSugared(msg.PeerID, users, chats)
	if err != nil {
		return nil, xerrors.Errorf("get msg.PeerID: %w", err)
	}

	m := &Message{
		msg:    msg,
		client: c,
		peer:   peer,
	}

	if from, ok := msg.GetFromID(); ok {
		fromEnt, err := c.peerSugared(from, users, chats)
		if err != nil {
			return nil, xerrors.Errorf("get msg.FromID: %w", err)
		}
		m.from = fromEnt
	}

	return m, nil
}

// func (c *Client) newMessageFromService(
// 	msg *tg.MessageService,
// 	users []tg.UserClass,
// 	chats []tg.ChatClass,
// ) (*Message, error) {
// 	peer, err := c.peerSugared(msg.PeerID, users, chats)
// 	if err != nil {
// 		return nil, xerrors.Errorf("get msg.PeerID: %w", err)
// 	}

// 	m := &Message{
// 		msg: &tg.Message{
// 			Out:         msg.Out,
// 			Mentioned:   msg.Mentioned,
// 			MediaUnread: msg.MediaUnread,
// 			Silent:      msg.Silent,
// 			Post:        msg.Post,
// 			Legacy:      msg.Legacy,
// 			ID:          msg.ID,
// 			PeerID:      msg.PeerID,
// 			Date:        msg.Date,
// 		},
// 		client: c,
// 		peer:   peer,
// 		action: msg.Action,
// 	}

// 	if from, ok := msg.GetFromID(); ok {
// 		m.msg.SetFromID(from)
// 		fromEnt, err := c.peerSugared(from, users, chats)
// 		if err != nil {
// 			return nil, xerrors.Errorf("get msg.FromID: %w", err)
// 		}
// 		m.from = fromEnt
// 	}

// 	if replyTo, ok := msg.GetReplyTo(); ok {
// 		m.msg.SetReplyTo(replyTo)
// 	}
// 	if ttl, ok := msg.GetTTLPeriod(); ok {
// 		m.msg.SetTTLPeriod(ttl)
// 	}

// 	return m, nil
// }

func (c *Client) newMessageFromShortSent(dst Entity, text string, sent *tg.UpdateShortSentMessage) (*Message, error) {
	msg := &tg.Message{
		Message: text,
		Out:     sent.Out,
		ID:      sent.ID,
		Date:    sent.Date,
	}

	self, ok := c.self.Load().(*tg.User)
	if !ok {
		return nil, xerrors.Errorf("not authenticated")
	}

	if media, ok := sent.GetMedia(); ok {
		msg.SetMedia(media)
	}
	if ents, ok := sent.GetEntities(); ok {
		msg.SetEntities(ents)
	}
	if ttl, ok := sent.GetTTLPeriod(); ok {
		msg.SetTTLPeriod(ttl)
	}

	if ch, ok := dst.(*Channel); ok {
		msg.Post = ch.channel.Broadcast
	}

	return &Message{
		msg:    msg,
		client: c,
		peer:   dst,
		from:   &User{user: self, client: c},
	}, nil
}

func (m *Message) ID() int { return m.msg.ID }

func (m *Message) Text() string { return m.msg.Message }

func (m *Message) Out() bool { return m.msg.Out }

func (m *Message) From() Entity {
	if m.from != nil {
		return m.from
	}

	return m.Peer()
}

func (m *Message) Personal() bool { _, ok := m.peer.(*User); return ok }

func (m *Message) Peer() Entity {
	return m.peer
}

func (m *Message) EditText(ctx context.Context, text string) error {
	return m.client.EditMessageText(ctx, m, text)
}

func (m *Message) Delete(ctx context.Context) error {
	return m.client.DeleteMessages(ctx, m.From(), []int{m.msg.ID})
}

func (m *Message) Raw() *tg.Message { return m.msg }

func (m *Message) HasDocument() (*Document, bool) {
	media, ok := m.msg.GetMedia()
	if !ok {
		return nil, false
	}

	doc, ok := media.(*tg.MessageMediaDocument)
	if !ok {
		return nil, false
	}

	d, ok := doc.GetDocument()
	if !ok {
		return nil, false
	}

	switch d := d.(type) {
	case *tg.Document:
		sugared := &Document{
			doc:    d,
			client: m.client,
		}

		if ttl, ok := doc.GetTTLSeconds(); ok {
			sugared.ttlPeriod = ttl
		}

		return sugared, true
	default:
		panic(fmt.Sprintf("unexpected document type: %T", d))
	}
}

func (m *Message) HasPhoto() (*Photo, bool) {
	media, ok := m.msg.GetMedia()
	if !ok {
		return nil, false
	}

	photo, ok := media.(*tg.MessageMediaPhoto)
	if !ok {
		return nil, false
	}

	p, ok := photo.GetPhoto()
	if !ok {
		return nil, false
	}

	switch p := p.(type) {
	case *tg.Photo:
		sugared := &Photo{
			photo:  p,
			client: m.client,
		}

		if ttl, ok := photo.GetTTLSeconds(); ok {
			sugared.ttlPeriod = ttl
		}

		return sugared, true
	default:
		return nil, false
	}
}

func (m *Message) HasMedia() (tg.MessageMediaClass, bool) {
	return m.msg.GetMedia()
}
