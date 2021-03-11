// Code generated by gotdgen, DO NOT EDIT.

package mt

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
)

// RPCAnswerUnknown represents TL type `rpc_answer_unknown#5e2ad36e`.
type RPCAnswerUnknown struct {
}

// RPCAnswerUnknownTypeID is TL type id of RPCAnswerUnknown.
const RPCAnswerUnknownTypeID = 0x5e2ad36e

func (r *RPCAnswerUnknown) Zero() bool {
	if r == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (r *RPCAnswerUnknown) String() string {
	if r == nil {
		return "RPCAnswerUnknown(nil)"
	}
	type Alias RPCAnswerUnknown
	return fmt.Sprintf("RPCAnswerUnknown%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RPCAnswerUnknown) TypeID() uint32 {
	return RPCAnswerUnknownTypeID
}

// TypeName returns name of type in TL schema.
func (*RPCAnswerUnknown) TypeName() string {
	return "rpc_answer_unknown"
}

// TypeInfo returns info about TL type.
func (r *RPCAnswerUnknown) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "rpc_answer_unknown",
		ID:   RPCAnswerUnknownTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (r *RPCAnswerUnknown) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode rpc_answer_unknown#5e2ad36e as nil")
	}
	b.PutID(RPCAnswerUnknownTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (r *RPCAnswerUnknown) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode rpc_answer_unknown#5e2ad36e to nil")
	}
	if err := b.ConsumeID(RPCAnswerUnknownTypeID); err != nil {
		return fmt.Errorf("unable to decode rpc_answer_unknown#5e2ad36e: %w", err)
	}
	return nil
}

// construct implements constructor of RPCDropAnswerClass.
func (r RPCAnswerUnknown) construct() RPCDropAnswerClass { return &r }

// Ensuring interfaces in compile-time for RPCAnswerUnknown.
var (
	_ bin.Encoder = &RPCAnswerUnknown{}
	_ bin.Decoder = &RPCAnswerUnknown{}

	_ RPCDropAnswerClass = &RPCAnswerUnknown{}
)

// RPCAnswerDroppedRunning represents TL type `rpc_answer_dropped_running#cd78e586`.
type RPCAnswerDroppedRunning struct {
}

// RPCAnswerDroppedRunningTypeID is TL type id of RPCAnswerDroppedRunning.
const RPCAnswerDroppedRunningTypeID = 0xcd78e586

func (r *RPCAnswerDroppedRunning) Zero() bool {
	if r == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (r *RPCAnswerDroppedRunning) String() string {
	if r == nil {
		return "RPCAnswerDroppedRunning(nil)"
	}
	type Alias RPCAnswerDroppedRunning
	return fmt.Sprintf("RPCAnswerDroppedRunning%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RPCAnswerDroppedRunning) TypeID() uint32 {
	return RPCAnswerDroppedRunningTypeID
}

// TypeName returns name of type in TL schema.
func (*RPCAnswerDroppedRunning) TypeName() string {
	return "rpc_answer_dropped_running"
}

// TypeInfo returns info about TL type.
func (r *RPCAnswerDroppedRunning) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "rpc_answer_dropped_running",
		ID:   RPCAnswerDroppedRunningTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (r *RPCAnswerDroppedRunning) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode rpc_answer_dropped_running#cd78e586 as nil")
	}
	b.PutID(RPCAnswerDroppedRunningTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (r *RPCAnswerDroppedRunning) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode rpc_answer_dropped_running#cd78e586 to nil")
	}
	if err := b.ConsumeID(RPCAnswerDroppedRunningTypeID); err != nil {
		return fmt.Errorf("unable to decode rpc_answer_dropped_running#cd78e586: %w", err)
	}
	return nil
}

// construct implements constructor of RPCDropAnswerClass.
func (r RPCAnswerDroppedRunning) construct() RPCDropAnswerClass { return &r }

// Ensuring interfaces in compile-time for RPCAnswerDroppedRunning.
var (
	_ bin.Encoder = &RPCAnswerDroppedRunning{}
	_ bin.Decoder = &RPCAnswerDroppedRunning{}

	_ RPCDropAnswerClass = &RPCAnswerDroppedRunning{}
)

// RPCAnswerDropped represents TL type `rpc_answer_dropped#a43ad8b7`.
type RPCAnswerDropped struct {
	// MsgID field of RPCAnswerDropped.
	MsgID int64
	// SeqNo field of RPCAnswerDropped.
	SeqNo int
	// Bytes field of RPCAnswerDropped.
	Bytes int
}

// RPCAnswerDroppedTypeID is TL type id of RPCAnswerDropped.
const RPCAnswerDroppedTypeID = 0xa43ad8b7

func (r *RPCAnswerDropped) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.MsgID == 0) {
		return false
	}
	if !(r.SeqNo == 0) {
		return false
	}
	if !(r.Bytes == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RPCAnswerDropped) String() string {
	if r == nil {
		return "RPCAnswerDropped(nil)"
	}
	type Alias RPCAnswerDropped
	return fmt.Sprintf("RPCAnswerDropped%+v", Alias(*r))
}

// FillFrom fills RPCAnswerDropped from given interface.
func (r *RPCAnswerDropped) FillFrom(from interface {
	GetMsgID() (value int64)
	GetSeqNo() (value int)
	GetBytes() (value int)
}) {
	r.MsgID = from.GetMsgID()
	r.SeqNo = from.GetSeqNo()
	r.Bytes = from.GetBytes()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RPCAnswerDropped) TypeID() uint32 {
	return RPCAnswerDroppedTypeID
}

// TypeName returns name of type in TL schema.
func (*RPCAnswerDropped) TypeName() string {
	return "rpc_answer_dropped"
}

// TypeInfo returns info about TL type.
func (r *RPCAnswerDropped) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "rpc_answer_dropped",
		ID:   RPCAnswerDroppedTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "MsgID",
			SchemaName: "msg_id",
		},
		{
			Name:       "SeqNo",
			SchemaName: "seq_no",
		},
		{
			Name:       "Bytes",
			SchemaName: "bytes",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *RPCAnswerDropped) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode rpc_answer_dropped#a43ad8b7 as nil")
	}
	b.PutID(RPCAnswerDroppedTypeID)
	b.PutLong(r.MsgID)
	b.PutInt(r.SeqNo)
	b.PutInt(r.Bytes)
	return nil
}

// GetMsgID returns value of MsgID field.
func (r *RPCAnswerDropped) GetMsgID() (value int64) {
	return r.MsgID
}

// GetSeqNo returns value of SeqNo field.
func (r *RPCAnswerDropped) GetSeqNo() (value int) {
	return r.SeqNo
}

// GetBytes returns value of Bytes field.
func (r *RPCAnswerDropped) GetBytes() (value int) {
	return r.Bytes
}

// Decode implements bin.Decoder.
func (r *RPCAnswerDropped) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode rpc_answer_dropped#a43ad8b7 to nil")
	}
	if err := b.ConsumeID(RPCAnswerDroppedTypeID); err != nil {
		return fmt.Errorf("unable to decode rpc_answer_dropped#a43ad8b7: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode rpc_answer_dropped#a43ad8b7: field msg_id: %w", err)
		}
		r.MsgID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode rpc_answer_dropped#a43ad8b7: field seq_no: %w", err)
		}
		r.SeqNo = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode rpc_answer_dropped#a43ad8b7: field bytes: %w", err)
		}
		r.Bytes = value
	}
	return nil
}

// construct implements constructor of RPCDropAnswerClass.
func (r RPCAnswerDropped) construct() RPCDropAnswerClass { return &r }

// Ensuring interfaces in compile-time for RPCAnswerDropped.
var (
	_ bin.Encoder = &RPCAnswerDropped{}
	_ bin.Decoder = &RPCAnswerDropped{}

	_ RPCDropAnswerClass = &RPCAnswerDropped{}
)

// RPCDropAnswerClass represents RpcDropAnswer generic type.
//
// Example:
//  g, err := mt.DecodeRPCDropAnswer(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *mt.RPCAnswerUnknown: // rpc_answer_unknown#5e2ad36e
//  case *mt.RPCAnswerDroppedRunning: // rpc_answer_dropped_running#cd78e586
//  case *mt.RPCAnswerDropped: // rpc_answer_dropped#a43ad8b7
//  default: panic(v)
//  }
type RPCDropAnswerClass interface {
	bin.Encoder
	bin.Decoder
	construct() RPCDropAnswerClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeRPCDropAnswer implements binary de-serialization for RPCDropAnswerClass.
func DecodeRPCDropAnswer(buf *bin.Buffer) (RPCDropAnswerClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case RPCAnswerUnknownTypeID:
		// Decoding rpc_answer_unknown#5e2ad36e.
		v := RPCAnswerUnknown{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RPCDropAnswerClass: %w", err)
		}
		return &v, nil
	case RPCAnswerDroppedRunningTypeID:
		// Decoding rpc_answer_dropped_running#cd78e586.
		v := RPCAnswerDroppedRunning{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RPCDropAnswerClass: %w", err)
		}
		return &v, nil
	case RPCAnswerDroppedTypeID:
		// Decoding rpc_answer_dropped#a43ad8b7.
		v := RPCAnswerDropped{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RPCDropAnswerClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode RPCDropAnswerClass: %w", bin.NewUnexpectedID(id))
	}
}

// RPCDropAnswer boxes the RPCDropAnswerClass providing a helper.
type RPCDropAnswerBox struct {
	RpcDropAnswer RPCDropAnswerClass
}

// Decode implements bin.Decoder for RPCDropAnswerBox.
func (b *RPCDropAnswerBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode RPCDropAnswerBox to nil")
	}
	v, err := DecodeRPCDropAnswer(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.RpcDropAnswer = v
	return nil
}

// Encode implements bin.Encode for RPCDropAnswerBox.
func (b *RPCDropAnswerBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.RpcDropAnswer == nil {
		return fmt.Errorf("unable to encode RPCDropAnswerClass as nil")
	}
	return b.RpcDropAnswer.Encode(buf)
}

// RPCDropAnswerClassArray is adapter for slice of RPCDropAnswerClass.
type RPCDropAnswerClassArray []RPCDropAnswerClass

// Sort sorts slice of RPCDropAnswerClass.
func (s RPCDropAnswerClassArray) Sort(less func(a, b RPCDropAnswerClass) bool) RPCDropAnswerClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of RPCDropAnswerClass.
func (s RPCDropAnswerClassArray) SortStable(less func(a, b RPCDropAnswerClass) bool) RPCDropAnswerClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of RPCDropAnswerClass.
func (s RPCDropAnswerClassArray) Retain(keep func(x RPCDropAnswerClass) bool) RPCDropAnswerClassArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s RPCDropAnswerClassArray) First() (v RPCDropAnswerClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s RPCDropAnswerClassArray) Last() (v RPCDropAnswerClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *RPCDropAnswerClassArray) PopFirst() (v RPCDropAnswerClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero RPCDropAnswerClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *RPCDropAnswerClassArray) Pop() (v RPCDropAnswerClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsRPCAnswerDropped returns copy with only RPCAnswerDropped constructors.
func (s RPCDropAnswerClassArray) AsRPCAnswerDropped() (to RPCAnswerDroppedArray) {
	for _, elem := range s {
		value, ok := elem.(*RPCAnswerDropped)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// RPCAnswerDroppedArray is adapter for slice of RPCAnswerDropped.
type RPCAnswerDroppedArray []RPCAnswerDropped

// Sort sorts slice of RPCAnswerDropped.
func (s RPCAnswerDroppedArray) Sort(less func(a, b RPCAnswerDropped) bool) RPCAnswerDroppedArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of RPCAnswerDropped.
func (s RPCAnswerDroppedArray) SortStable(less func(a, b RPCAnswerDropped) bool) RPCAnswerDroppedArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of RPCAnswerDropped.
func (s RPCAnswerDroppedArray) Retain(keep func(x RPCAnswerDropped) bool) RPCAnswerDroppedArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s RPCAnswerDroppedArray) First() (v RPCAnswerDropped, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s RPCAnswerDroppedArray) Last() (v RPCAnswerDropped, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *RPCAnswerDroppedArray) PopFirst() (v RPCAnswerDropped, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero RPCAnswerDropped
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *RPCAnswerDroppedArray) Pop() (v RPCAnswerDropped, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
