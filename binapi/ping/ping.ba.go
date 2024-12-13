// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.11.0
//  VPP:              24.06-release
// source: plugins/ping.api.json

// Package ping contains generated bindings for API file ping.api.
//
// Contents:
// -  3 messages
package ping

import (
	api "go.fd.io/govpp/api"
	_ "go.fd.io/govpp/binapi/interface_types"
	ip_types "go.fd.io/govpp/binapi/ip_types"
	codec "go.fd.io/govpp/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "ping"
	APIVersion = "0.1.0"
	VersionCrc = 0x5b568e46
)

// PingFinishedEvent defines message 'ping_finished_event'.
type PingFinishedEvent struct {
	RequestCount uint32 `binapi:"u32,name=request_count" json:"request_count,omitempty"`
	ReplyCount   uint32 `binapi:"u32,name=reply_count" json:"reply_count,omitempty"`
}

func (m *PingFinishedEvent) Reset()               { *m = PingFinishedEvent{} }
func (*PingFinishedEvent) GetMessageName() string { return "ping_finished_event" }
func (*PingFinishedEvent) GetCrcString() string   { return "397ccf72" }
func (*PingFinishedEvent) GetMessageType() api.MessageType {
	return api.EventMessage
}

func (m *PingFinishedEvent) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.RequestCount
	size += 4 // m.ReplyCount
	return size
}
func (m *PingFinishedEvent) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.RequestCount)
	buf.EncodeUint32(m.ReplyCount)
	return buf.Bytes(), nil
}
func (m *PingFinishedEvent) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.RequestCount = buf.DecodeUint32()
	m.ReplyCount = buf.DecodeUint32()
	return nil
}

// /*
//   - Copyright (c) 2023 Cisco and/or its affiliates.
//   - Licensed under the Apache License, Version 2.0 (the "License");
//   - you may not use this file except in compliance with the License.
//   - You may obtain a copy of the License at:
//     *
//   - http://www.apache.org/licenses/LICENSE-2.0
//     *
//   - Unless required by applicable law or agreed to in writing, software
//   - distributed under the License is distributed on an "AS IS" BASIS,
//   - WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   - See the License for the specific language governing permissions and
//   - limitations under the License.
//
// WantPingFinishedEvents defines message 'want_ping_finished_events'.
type WantPingFinishedEvents struct {
	Address  ip_types.Address `binapi:"address,name=address" json:"address,omitempty"`
	Repeat   uint32           `binapi:"u32,name=repeat,default=1" json:"repeat,omitempty"`
	Interval float64          `binapi:"f64,name=interval,default=1" json:"interval,omitempty"`
}

func (m *WantPingFinishedEvents) Reset()               { *m = WantPingFinishedEvents{} }
func (*WantPingFinishedEvents) GetMessageName() string { return "want_ping_finished_events" }
func (*WantPingFinishedEvents) GetCrcString() string   { return "e79ee58b" }
func (*WantPingFinishedEvents) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *WantPingFinishedEvents) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.Address.Af
	size += 1 * 16 // m.Address.Un
	size += 4      // m.Repeat
	size += 8      // m.Interval
	return size
}
func (m *WantPingFinishedEvents) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(uint8(m.Address.Af))
	buf.EncodeBytes(m.Address.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(m.Repeat)
	buf.EncodeFloat64(m.Interval)
	return buf.Bytes(), nil
}
func (m *WantPingFinishedEvents) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Address.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Address.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Repeat = buf.DecodeUint32()
	m.Interval = buf.DecodeFloat64()
	return nil
}

// WantPingFinishedEventsReply defines message 'want_ping_finished_events_reply'.
type WantPingFinishedEventsReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *WantPingFinishedEventsReply) Reset()               { *m = WantPingFinishedEventsReply{} }
func (*WantPingFinishedEventsReply) GetMessageName() string { return "want_ping_finished_events_reply" }
func (*WantPingFinishedEventsReply) GetCrcString() string   { return "e8d4e804" }
func (*WantPingFinishedEventsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *WantPingFinishedEventsReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *WantPingFinishedEventsReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *WantPingFinishedEventsReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_ping_binapi_init() }
func file_ping_binapi_init() {
	api.RegisterMessage((*PingFinishedEvent)(nil), "ping_finished_event_397ccf72")
	api.RegisterMessage((*WantPingFinishedEvents)(nil), "want_ping_finished_events_e79ee58b")
	api.RegisterMessage((*WantPingFinishedEventsReply)(nil), "want_ping_finished_events_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*PingFinishedEvent)(nil),
		(*WantPingFinishedEvents)(nil),
		(*WantPingFinishedEventsReply)(nil),
	}
}
