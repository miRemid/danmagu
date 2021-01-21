package message

import (
	"context"
)

type Header struct {
	PacketLen uint32 // 4
	HeaderLen uint16 // 2
	Version   uint16 // 2
	Operation uint32 // 4
	Sequence  uint32 // 4
}

type Auth struct {
	UID       uint8  `json:"uid"`
	Roomid    uint32 `json:"roomid"`
	Protover  uint8  `json:"protover"`
	Platform  string `json:"platform"`
	Clientver string `json:"clientver"`
	Type      uint8  `json:"type"`
	Key       string `json:"key"`
}

type Message struct {
	Header

	Body []byte
}

type Context struct {
	context.Context
	Operation uint32
	Buffer    []byte
}

func NewMessage(data []byte) *Message {
	var msg Message
	msg.Body = data
	return &msg
}

const (
	WsNeedZlib = 2
	// WsHeartbeatSent 心跳
	WsHeartbeatSent = 2
	// WsHeartbeatReply 心跳响应
	WsHeartbeatReply = 3
	// WsMessage 消息
	WsMessage = 5
	// WsAuth 认证
	WsAuth = 7
	// WsAuthSuccess 认证成功
	WsAuthSuccess = 8
)

const (
	HeaderLength = 16
)
