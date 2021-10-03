package tgram

import "github.com/gotd/td/tg"

type handlers struct {
	onNewMessage     func(*Message) error
	onEditMessage    func(*Message) error
	onDeleteMessage  func(*Message) error
	onChatAddUser    func(*MessageActionChatAddUser) error
	onChatDeleteUser func(*MessageActionChatDeleteUser) error
	onUpdate         func(update tg.UpdateClass, ents *Entities) error
}

func newHandlers() *handlers {
	return &handlers{
		onNewMessage:    func(*Message) error { return nil },
		onEditMessage:   func(*Message) error { return nil },
		onDeleteMessage: func(*Message) error { return nil },
		onChatAddUser:   func(*MessageActionChatAddUser) error { return nil },
		onUpdate:        func(update tg.UpdateClass, ents *Entities) error { return nil },
	}
}

func (c *Client) OnNewMessage(f func(msg *Message) error) {
	c.handlers.onNewMessage = f
}

func (c *Client) OnEditMessage(f func(msg *Message) error) {
	c.handlers.onEditMessage = f
}

func (c *Client) OnChatAddUser(f func(joined *MessageActionChatAddUser) error) {
	c.handlers.onChatAddUser = f
}

func (c *Client) OnChatDeleteUser(f func(joined *MessageActionChatDeleteUser) error) {
	c.handlers.onChatDeleteUser = f
}

func (c *Client) OnUpdate(f func(update tg.UpdateClass, ents *Entities) error) {
	c.handlers.onUpdate = f
}
