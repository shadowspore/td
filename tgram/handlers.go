package tgram

import "github.com/gotd/td/tg"

type handlers struct {
	onNewMessage      func(*Message) error
	onEditMessage     func(*Message) error
	onDeleteMessage   func(*Message) error
	onChatUsersJoined func(*ActionChatUsersJoined) error
	fallback          func(tg.UpdatesClass) error
}

func newHandlers() *handlers {
	return &handlers{
		onNewMessage:      func(*Message) error { return nil },
		onEditMessage:     func(*Message) error { return nil },
		onDeleteMessage:   func(*Message) error { return nil },
		onChatUsersJoined: func(*ActionChatUsersJoined) error { return nil },
		fallback:          func(tg.UpdatesClass) error { return nil },
	}
}

func (c *Client) OnNewMessage(f func(msg *Message) error) {
	c.handlers.onNewMessage = f
}

func (c *Client) OnEditMessage(f func(msg *Message) error) {
	c.handlers.onEditMessage = f
}

func (c *Client) OnChatUsersJoined(f func(joined *ActionChatUsersJoined) error) {
	c.handlers.onChatUsersJoined = f
}

func (c *Client) OnFallback(f func(updates tg.UpdatesClass) error) {
	c.handlers.fallback = f
}
