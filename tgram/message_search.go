package tgram

type MessagesSearchOptions struct {
	BatchSize int
	Documents bool

	From *User
}

type MessagesSearchIterator struct {
}
