package danmagu

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
	"github.com/miRemid/danmagu/model"
)

const (
	tokenURL  = "https://api.live.bilibili.com/room/v1/Danmu/getConf"
	socketURL = "wss://broadcastlv.chat.bilibili.com/sub"
)

const (
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

// HandlerFunc 钩子函数
type HandlerFunc func()

func defaultHandlerFunc() {}

// Client 客户端
type Client struct {
	uid    int
	roomid int
	token  string
	conn   *websocket.Conn

	DebugMode bool

	AfterConnect  HandlerFunc
	BeforeConnect HandlerFunc
	AfterEnter    HandlerFunc
	BeforeEnter   HandlerFunc
	Listening     HandlerFunc
	BeforeListen  HandlerFunc

	DanmakuHandler   func(danmaku model.Danmaku)
	HeartBeatHandler func(heart model.HeartBeat)
	GiftHandler      func(git model.Gift)
	GuardHandler     func(guard model.Guard)

	stop bool
}

type auth struct {
	UID       int    `json:"uid"`
	RoomID    int    `json:"roomid"`
	Protover  int    `json:"protover"`
	PlatForm  string `json:"platform"`
	ClientVer string `json:"clientver"`
}

type header struct {
	Length uint32
	LBody  uint16
	Ver    uint16
	Opcode uint32
	Normal uint32
}

// NewClient 建立新客户端
func NewClient(uid int) *Client {
	res := &Client{
		uid: uid,
	}

	res.BeforeConnect = defaultHandlerFunc
	res.AfterConnect = defaultHandlerFunc
	res.BeforeEnter = defaultHandlerFunc
	res.AfterEnter = defaultHandlerFunc
	res.BeforeListen = defaultHandlerFunc

	res.HeartBeatHandler = func(hear model.HeartBeat) {
		fmt.Println(hear.Value)
	}
	res.DanmakuHandler = func(danmaku model.Danmaku) {
		fmt.Printf("%v:%v\n", danmaku.Nickname, danmaku.Content)
	}
	res.GiftHandler = func(gift model.Gift) {
		fmt.Println(gift.Uname)
	}

	return res
}

func (client *Client) debug(format string, a ...interface{}) {
	if client.DebugMode {
		log.Printf(format, a...)
	}
}

// GetToken 获取房间token
func GetToken(roomid int) (string, error) {
	// 1. 请求tokenURL
	url := fmt.Sprintf("%s?roomid=%d&platform=pc&player=web", tokenURL, roomid)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	// 2. 获取token
	bytedata, _ := ioutil.ReadAll(res.Body)
	var msg map[string]interface{}
	json.Unmarshal(bytedata, &msg)
	data := msg["data"].(map[string]interface{})
	token := data["token"].(string)
	return token, nil
}

func (client *Client) send(data []byte, operation uint32) {
	h := header{uint32(binary.Size(data) + 16), 16, 1, operation, 1}
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, h)
	client.conn.WriteMessage(websocket.BinaryMessage, []byte(fmt.Sprintf("%s%s", buf, data)))
}

func (client *Client) connect() {
	client.BeforeConnect()
	client.debug("try to connect...\n")
	conn, _, err := websocket.DefaultDialer.Dial(socketURL, nil)
	if err != nil {
		client.debug(err.Error())
		os.Exit(1)
	}
	client.conn = conn
	client.debug("connect successful\n")
	client.AfterConnect()
}

// Listen 开始监听数据
// t 心跳包发送间隔，30s以下
func (client *Client) Listen(t time.Duration) {
	client.connect()
	client.auth()
	client.BeforeListen()
	client.debug("start Listening roomid=%d\n", client.roomid)
	client.startHeart(t)
	for {
		if client.stop {
			client.stop = false
			client.debug("stop listening")
			return
		}
		_, body, err := client.conn.ReadMessage()
		if err != nil {
			log.Printf("connect failed... reconnect after 3s\n")
			client.conn.Close()
			time.Sleep(time.Second * 3)
			client.connect()
		}
		go client.parse(body)
	}
}

// Enter 进入房间
func (client *Client) Enter(roomid int) {
	client.roomid = roomid
}

func (client *Client) auth() {
	client.BeforeEnter()
	body, _ := json.Marshal(auth{
		UID:       client.uid,
		RoomID:    client.roomid,
		Protover:  1,
		ClientVer: "1.4.0",
		PlatForm:  "web",
	})
	client.send(body, WsAuth)
	client.AfterEnter()
}

func (client *Client) startHeart(t time.Duration) {
	client.debug("心跳包发送间隔为%ds\n", t)
	go func(t time.Duration) {
		for {
			time.Sleep(time.Second * 5)
			client.send([]byte(""), WsHeartbeatSent)
			time.Sleep(time.Second * (t - 5))
		}
	}(t)
}

// Cancle 取消监听
func (client *Client) Cancle() {
	client.stop = true
}

func (client *Client) parse(message []byte) {
	h := message[:16]
	var head model.RHeader
	buf := bytes.NewReader(h)
	binary.Read(buf, binary.BigEndian, &head)
	body := message[16:head.Length]

	switch head.Type {
	case WsHeartbeatReply:
		var rqz model.HeartBeat
		binary.Read(bytes.NewReader(body), binary.BigEndian, &rqz)
		go client.HeartBeatHandler(rqz)
	case WsMessage:
		var cmd model.CMD
		fmt.Println(string(body))
		_ = json.Unmarshal(body, &cmd)
		switch cmd.Cmd {
		case "DANMU_MSG":
			var info model.Infomation
			if err := json.Unmarshal(body, &info); err != nil {
				log.Println(err)
			} else {
				var danmaku model.Danmaku
				danmaku.Nickname = cmd.Info[2][1].(string)
				danmaku.Content = info.Info[1].(string)
				v, ok := cmd.Info[2][0].(float64)
				if !ok {
					log.Println("获取uid错误")
				} else {
					danmaku.UID = int(v)
					go client.DanmakuHandler(danmaku)
				}
			}
		case "SEND_GIFT":
			var g model.GiftCmd
			if err := json.Unmarshal(body, &g); err != nil {
				log.Println(err)
			} else {
				go client.GiftHandler(g.Data)
			}
		case "GUARD_BUY":
			var g model.GuardCmd
			if err := json.Unmarshal(body, &g); err != nil {
				log.Println(err)
			} else {
				go client.GuardHandler(g.Data)
			}
		}
	}
	next := message[head.Length:]
	if binary.Size(next) != 0 {
		go client.parse(next)
	}
}
