package tgram

import (
	"context"
	"encoding/binary"
	"strconv"
	"strings"

	"github.com/gotd/td/session"
	"github.com/gotd/td/telegram/updates"
	"golang.org/x/xerrors"
)

var (
	_ session.Storage      = (*sessionManager)(nil)
	_ updates.StateStorage = (*sessionManager)(nil)
)

type sessionManager struct {
	storage Storage
}

func newSessionManager(storage Storage) *sessionManager {
	return &sessionManager{storage}
}

func (s *sessionManager) LoadSession(ctx context.Context) ([]byte, error) {
	b, found, err := s.storage.Bucket([]byte("mtproto"))
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, session.ErrNotFound
	}

	data, found, err := b.Get([]byte("session"))
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, session.ErrNotFound
	}

	return data, nil
}

func (s *sessionManager) StoreSession(ctx context.Context, value []byte) error {
	b, err := s.storage.GetOrCreateBucket([]byte("mtproto"))
	if err != nil {
		return err
	}

	return b.Set([]byte("session"), value)
}

func (s *sessionManager) SetState(userID int, state updates.State) error {
	stateBucket, err := s.storage.GetOrCreateBucket(utab(userID, "state"))
	if err != nil {
		return err
	}

	tx, err := stateBucket.Tx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := tx.Set([]byte("pts"), i2b(state.Pts)); err != nil {
		return err
	}
	if err := tx.Set([]byte("qts"), i2b(state.Qts)); err != nil {
		return err
	}
	if err := tx.Set([]byte("date"), i2b(state.Date)); err != nil {
		return err
	}
	if err := tx.Set([]byte("seq"), i2b(state.Seq)); err != nil {
		return err
	}
	return tx.Commit()
}

func (s *sessionManager) GetState(userID int) (state updates.State, found bool, err error) {
	stateTable, found, err := s.storage.Bucket(utab(userID, "state"))
	if err != nil {
		return updates.State{}, false, err
	}

	if !found {
		return updates.State{}, false, nil
	}

	tx, err := stateTable.Tx()
	if err != nil {
		return updates.State{}, false, err
	}
	defer tx.Rollback()

	forceNotFound := false
	loadValue := func(name string) int {
		if err != nil || forceNotFound {
			return 0
		}

		val, found, getErr := tx.Get([]byte(name))
		if getErr != nil {
			err = getErr
			return 0
		}

		if !found {
			// Bucket created, but have no state values.
			// Corrupted state?
			forceNotFound = true
			return 0
		}

		return b2i(val)
	}

	state.Pts = loadValue("pts")
	state.Qts = loadValue("qts")
	state.Seq = loadValue("seq")
	state.Date = loadValue("date")
	if err != nil {
		return updates.State{}, false, err
	}

	if forceNotFound {
		return updates.State{}, true, nil
	}

	return state, false, nil
}

func (s *sessionManager) SetPts(userID, pts int) error {
	stateBucket, found, err := s.storage.Bucket(utab(userID, "state"))
	if err != nil {
		return err
	}

	if !found {
		return xerrors.Errorf("state not found")
	}

	return stateBucket.Set([]byte("pts"), i2b(pts))
}

func (s *sessionManager) SetQts(userID, qts int) error {
	stateBucket, found, err := s.storage.Bucket(utab(userID, "state"))
	if err != nil {
		return err
	}

	if !found {
		return xerrors.Errorf("state not found")
	}

	return stateBucket.Set([]byte("qts"), i2b(qts))
}

func (s *sessionManager) SetDate(userID, date int) error {
	stateBucket, found, err := s.storage.Bucket(utab(userID, "state"))
	if err != nil {
		return err
	}

	if !found {
		return xerrors.Errorf("state not found")
	}

	return stateBucket.Set([]byte("date"), i2b(date))
}

func (s *sessionManager) SetSeq(userID, seq int) error {
	stateBucket, found, err := s.storage.Bucket(utab(userID, "state"))
	if err != nil {
		return err
	}

	if !found {
		return xerrors.Errorf("state not found")
	}

	return stateBucket.Set([]byte("seq"), i2b(seq))
}

func (s *sessionManager) SetDateSeq(userID, date, seq int) error {
	stateBucket, found, err := s.storage.Bucket(utab(userID, "state"))
	if err != nil {
		return err
	}

	if !found {
		return xerrors.Errorf("state not found")
	}

	tx, err := stateBucket.Tx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := tx.Set([]byte("date"), i2b(date)); err != nil {
		return err
	}
	if err := tx.Set([]byte("seq"), i2b(seq)); err != nil {
		return err
	}
	return tx.Commit()
}

func (s *sessionManager) ForEachChannels(userID int, f func(channelID, pts int) error) error {
	channels, found, err := s.storage.Bucket(utab(userID, "state", "channels"))
	if err != nil {
		return err
	}

	if !found {
		return xerrors.Errorf("channels state not found")
	}

	return channels.ForEach(func(key, value []byte) error {
		return f(b2i(key), b2i(value))
	})
}

func (s *sessionManager) SetChannelPts(userID, channelID, pts int) error {
	channels, found, err := s.storage.Bucket(utab(userID, "state", "channels"))
	if err != nil {
		return err
	}

	if !found {
		return xerrors.Errorf("channels state not found")
	}

	return channels.Set(i2b(channelID), i2b(pts))
}

func (s *sessionManager) GetChannelPts(userID, channelID int) (pts int, found bool, err error) {
	channels, found, err := s.storage.Bucket(utab(userID, "state", "channels"))
	if err != nil {
		return 0, false, err
	}

	if !found {
		return 0, false, nil
	}

	val, found, err := channels.Get(i2b(channelID))
	if err != nil {
		return 0, false, err
	}

	if !found {
		return 0, false, nil
	}

	return b2i(val), true, nil
}

func i2b(v int) []byte { b := make([]byte, 8); binary.LittleEndian.PutUint64(b, uint64(v)); return b }

func b2i(b []byte) int { return int(binary.LittleEndian.Uint64(b)) }

func utab(userID int, tablePath ...string) []byte {
	s := append([]string{"user", strconv.Itoa(userID)}, tablePath...)
	return []byte(strings.Join(s, "_"))
}
