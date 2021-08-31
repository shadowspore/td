package unpack

import (
	"github.com/gotd/td/tg"
	"golang.org/x/xerrors"
)

func Messages(u tg.UpdatesClass, randomIDs []int64) ([]tg.MessageClass, error) {
	switch u := u.(type) {
	case *tg.Updates:
		return mapRandomIDs(u.Updates, randomIDs)
	case *tg.UpdatesCombined:
		return mapRandomIDs(u.Updates, randomIDs)
	default:
		return nil, xerrors.Errorf("unexpected updates class: %T", u)
	}
}

func mapRandomIDs(updates []tg.UpdateClass, randomIDs []int64) ([]tg.MessageClass, error) {
	var (
		// randomID -> messageID
		mapping  = make(map[int64]int)
		messages []tg.MessageClass
	)
	for _, u := range updates {
		switch u := u.(type) {
		case *tg.UpdateMessageID:
			mapping[u.RandomID] = u.ID
		case *tg.UpdateNewMessage:
			messages = append(messages, u.Message)
		case *tg.UpdateNewChannelMessage:
			messages = append(messages, u.Message)
		}
	}

	var result []tg.MessageClass
	for _, randomID := range randomIDs {
		targetMessageID, ok := mapping[randomID]
		if !ok {
			return nil, xerrors.Errorf("messageID for randomID %d not found", randomID)
		}

		found := false
		for _, msg := range messages {
			if msg.GetID() == targetMessageID {
				result = append(result, msg)
				found = true
				break
			}
		}

		if !found {
			return nil, xerrors.Errorf("message with ID %d not found", targetMessageID)
		}
	}

	return result, nil
}