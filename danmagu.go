package danmaku

import (
	"log"
	"fmt"
	"time"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/websocket"
)

const (
	tokenURL  = "https://api.live.bilibili.com/room/v1/Danmu/getConf"
	socketURL = "wss://broadcastlv.chat.bilibili.com/sub"
)

const (
	// WsHeartbeatSent 心跳
	WsHeartbeatSent 	= 2
	// WsHeartbeatReply 心跳响应
	WsHeartbeatReply 	= 3
	// WsMessage 消息
	WsMessage 			= 5
	// WsAuth 认证
	WsAuth				= 7
	// WsAuthSuccess 认证成功
	WsAuthSuccess 		= 8
)

// EventHandler 处理函数
type EventHandler func(body []byte)

// Client 客户端
type Client struct {
	uid    int
	roomid int
	token  string
	conn   *websocket.Conn
	message EventHandler
}

type auth struct {
	UID       int    `json:"uid"`
	RoomID    int    `json:"roomid"`
	Protover  int    `json:"protover"`
	PlatForm  string `json:"platform"`
	ClientVer string `json:"clientver"`
	// Type      int    `json:"type"`
	// Token     string `json:"token"`
}

type header struct {
	Length 		uint32
	LBody	 	uint16
	Ver 		uint16
	Opcode		uint32
	Normal		uint32
}

// NewClient 建立新客户端
func NewClient(uid int) *Client {
	res := &Client{
		uid: uid,
	}
	res.StartSocket()
	return res
}

// GetToken 获取房间token
func GetToken(roomid int) string {
	// 1. 请求tokenURL
	url := fmt.Sprintf("%s?roomid=%d&platform=pc&player=web", tokenURL, roomid)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("New Request Failed=%s\n", err.Error())
		return ""
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Request Failed=%s\n", err.Error())
	}
	defer res.Body.Close()
	// 2. 获取token
	bytedata, _ := ioutil.ReadAll(res.Body)
	var msg map[string]interface{}
	json.Unmarshal(bytedata, &msg)
	data := msg["data"].(map[string]interface{})
	token := data["token"].(string)
	return token
}

func (c *Client) send(data []byte, operation uint32) {
	h := header{uint32(binary.Size(data) + 16), 16, 1, operation, 1}
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, h)
	c.conn.WriteMessage(websocket.BinaryMessage, []byte(fmt.Sprintf("%s%s", buf, data)))
}

// StartSocket 连接websocket服务器
func (c *Client) StartSocket() {
	log.Printf("try to connect...\n")
	conn, _, err := websocket.DefaultDialer.Dial(socketURL, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	c.conn = conn
	log.Printf("connect successful\n")
}

// Listen 开始监听数据
// t 心跳包发送间隔，30s以下
func (c *Client) Listen(t time.Duration) {
	log.Printf("start Listening roomid=%d\n", c.roomid)
	c.startHeart(t)
	for {
		_, body, err := c.conn.ReadMessage()
		if err != nil {
			log.Printf("connect failed... reconnect after 3s\n")
			c.conn.Close()
			time.Sleep(time.Second * 3)
			c.StartSocket()
		}
		go c.message(body)
	}
}

// Auth 进入房间
func (c *Client) Auth(roomid int) {
	c.roomid = roomid	
	body, _ := json.Marshal(auth{
		UID:       c.uid,
		RoomID:    c.roomid,
		Protover:  1,
		ClientVer: "1.4.0",
		PlatForm:  "web",
	})
	c.send(body, WsAuth)
}

func (c *Client) startHeart(t time.Duration) {
	log.Printf("心跳包发送间隔为%ds\n", t)
	go func(t time.Duration) {
		for {
			time.Sleep(time.Second * 5)
			c.send([]byte(""), WsHeartbeatSent)			
			time.Sleep(time.Second * (t - 5))
		}
	}(t)
}

// OnMessage 消息处理函数
func (c *Client) OnMessage(handler EventHandler) {
	c.message = handler
}