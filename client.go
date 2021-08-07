package danmagu

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"time"

	"github.com/asmcos/requests"
	"github.com/gorilla/websocket"
	"github.com/miRemid/danmagu/message"
	"github.com/miRemid/danmagu/tools"
	"github.com/tidwall/gjson"
)

const (
	ROOMINFO_URL = "https://api.live.bilibili.com/xlive/web-room/v1/index/getDanmuInfo?id=%d&type=0"
)

// 直播间Client
type LiveClient struct {
	cfg    *ClientConfig
	conn   *websocket.Conn
	roomid uint32

	heartBeatErr chan error
	recieveErr   chan error
	close        chan struct{}

	rawQueue chan *message.Message
	msgQueue chan *message.Context

	funcs map[string]HandlerFunc
}

func NewClient(roomid uint32, opt *ClientConfig) *LiveClient {
	var cli LiveClient
	cli.cfg = opt
	cli.roomid = roomid

	cli.heartBeatErr = make(chan error, 1)
	cli.recieveErr = make(chan error, 1)
	cli.close = make(chan struct{}, 1)
	cli.rawQueue = make(chan *message.Message)
	cli.msgQueue = make(chan *message.Context)

	cli.funcs = defaultFuncs
	return &cli
}

func (cli *LiveClient) Close() {
	close(cli.heartBeatErr)
	close(cli.recieveErr)
	close(cli.rawQueue)
	close(cli.msgQueue)
	cli.conn.Close()
	cli.close <- struct{}{}
}

func (cli *LiveClient) Handler(cmd string, handler HandlerFunc) {
	cli.funcs[cmd] = handler
}

func (cli *LiveClient) GetTokenAndURLS() (string, []string, error) {
	r, err := requests.Get(fmt.Sprintf(ROOMINFO_URL, cli.roomid))
	if err != nil {
		return "", nil, err
	}
	var urls = make([]string, 0)
	token := gjson.Get(r.Text(), "data.token").String()
	gjson.Get(r.Text(), "data.host_list").ForEach(func(key, value gjson.Result) bool {
		urls = append(urls, value.Get("host").String())
		return true
	})
	return token, urls, nil
}

func (cli *LiveClient) Send(data []byte, operation uint32) error {
	header := message.Header{
		PacketLen: uint32(binary.Size(data) + 16),
		HeaderLen: 16,
		Version:   1,
		Operation: operation,
		Sequence:  1,
	}
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, header)
	sendData := append(buf.Bytes(), data...)
	return cli.conn.WriteMessage(websocket.BinaryMessage, sendData)
}

func (cli *LiveClient) heartBeat(ctx context.Context) {
	for {
		DPrintf("Send Heart Beat Packet")
		select {
		case <-ctx.Done():
		case <-cli.close:
			return
		default:
			if err := cli.Send([]byte(""), message.WsHeartbeatSent); err != nil {
				cli.heartBeatErr <- err
				return
			}
		}
		time.Sleep(time.Second * cli.cfg.HeartBeatTime)
	}
}

func (cli *LiveClient) recieve(ctx context.Context) {
	count := 0
	closeConn := make(chan struct{}, 1)
	for {
		select {
		case <-ctx.Done():
		case <-cli.close:
			return
		case <-closeConn:
			DPrintf("reconnecting...")
			cli.connect()
			count = 0
		default:
			_, body, err := cli.conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					DPrintf("Unexpected Close Error")
					closeConn <- struct{}{}
					continue
				}
				if count != 5 {
					count++
					continue
				}
				DPrintf("RoomID=%d Read Message Error...", cli.roomid)
				cli.recieveErr <- err
				return
			}
			msg := message.NewMessage(body)
			cli.rawQueue <- msg
		}
	}
}

func (cli *LiveClient) split(ctx context.Context) {
	var (
		msg          *message.Message
		header       message.Header
		headerBuffer *bytes.Reader
		buffer       []byte
	)
	for {
		msg = <-cli.rawQueue
		for len(msg.Body) > 0 {
			select {
			case <-ctx.Done():
			case <-cli.close:
				return
			default:
			}

			headerBuffer = bytes.NewReader(msg.Body[:message.HeaderLength])
			_ = binary.Read(headerBuffer, binary.BigEndian, &header)
			buffer = msg.Body[message.HeaderLength:int(header.PacketLen)]
			msg.Body = msg.Body[int(header.PacketLen):]

			if header.PacketLen == message.HeaderLength {
				continue
			}

			if header.Version == message.WsNeedZlib {
				msg.Body = tools.ZlibInflate(buffer)
				continue
			}

			cli.msgQueue <- &message.Context{Context: ctx, Operation: header.Operation, Buffer: buffer}
		}
	}
}

func (cli *LiveClient) parse(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
		case <-cli.close:
			return
		default:
		}

		msg := <-cli.msgQueue
		buffer := bytes.NewReader(msg.Buffer)
		switch msg.Operation {
		case message.WsHeartbeatReply:
			var rqz uint32
			binary.Read(buffer, binary.BigEndian, &rqz)
			handler := cli.funcs["RQZ"].(RQZHandler)
			go handler(ctx, rqz)
		case message.WsMessage:
			cmd := gjson.GetBytes(msg.Buffer, "cmd").String()
			switch cmd {
			case message.DANMU_MSG:
				res := gjson.GetBytes(msg.Buffer, "info").Array()
				var danmaku message.Danmaku
				danmaku.Content = res[1].String()
				userInfo := res[2].Array()
				danmaku.UID = uint(userInfo[0].Uint())
				danmaku.Username = userInfo[1].String()
				handler := cli.funcs[cmd].(DanmakuHandler)
				go handler(ctx, danmaku)
			case message.SEND_GIFT:
				var gift message.SendGift
				if err := tools.Unmarshal(msg.Buffer, &gift); err != nil {
					cli.errorHandler(cmd, err)
					break
				}
				handler := cli.funcs[cmd].(GiftHandler)
				go handler(ctx, gift)
			case message.INTERACT_WORD:
				var word message.InteractWord
				if err := tools.Unmarshal(msg.Buffer, &word); err != nil {
					cli.errorHandler(cmd, err)
					break
				}
				handler := cli.funcs[cmd].(InteractWordHandler)
				go handler(ctx, word)
			case message.ONLINE_RANK_V2:
				var item message.OnlineRankV2
				if err := tools.Unmarshal(msg.Buffer, &item); err != nil {
					cli.errorHandler(cmd, err)
					break
				}
				handler := cli.funcs[cmd].(RankV2Handler)
				go handler(ctx, item)
			case message.ONLINE_RANK_TOP3:
				var item message.OnlineRankTOP3
				if err := tools.Unmarshal(msg.Buffer, &item); err != nil {
					cli.errorHandler(cmd, err)
					break
				}
				handler := cli.funcs[cmd].(RankTOP3Handler)
				go handler(ctx, item)
			case message.COMBO_SEND:
				var item message.ComboSend
				if err := tools.Unmarshal(msg.Buffer, &item); err != nil {
					cli.errorHandler(cmd, err)
					break
				}
				handler := cli.funcs[cmd].(ComboSendHandler)
				go handler(ctx, item)
			case message.WIDGET_BANNER:
				var item message.WidgetBanner
				if err := tools.Unmarshal(msg.Buffer, &item); err != nil {
					cli.errorHandler(cmd, err)
					break
				}
				handler := cli.funcs[cmd].(WidgetBannerHandler)
				go handler(ctx, item)
			case message.ENTRY_EFFECT:
				var item message.EntryEffect
				if err := tools.Unmarshal(msg.Buffer, &item); err != nil {
					cli.errorHandler(cmd, err)
					break
				}
				handler := cli.funcs[cmd].(EntryEffectHandler)
				go handler(ctx, item)
			case message.ONLINE_RANK_COUNT:
				var item message.OnlineRankCount
				if err := tools.Unmarshal(msg.Buffer, &item); err != nil {
					cli.errorHandler(cmd, err)
					break
				}
				handler := cli.funcs[cmd].(RankCountHandler)
				go handler(ctx, item)
			case message.ROOM_RANK:
				var item message.RoomRank
				if err := tools.Unmarshal(msg.Buffer, &item); err != nil {
					cli.errorHandler(cmd, err)
					break
				}
				handler := cli.funcs[cmd].(RoomRankHandler)
				go handler(ctx, item)
			case message.LIVE:
				var item message.Live
				if err := tools.Unmarshal(msg.Buffer, &item); err != nil {
					cli.errorHandler(cmd, err)
				}
				handler := cli.funcs[cmd].(LiveHandler)
				go handler(ctx, item)
			case message.PREPARING:
				var item message.Preparing
				if err := tools.Unmarshal(msg.Buffer, &item); err != nil {
					cli.errorHandler(cmd, err)
				}
				handler := cli.funcs[cmd].(PreparingHandler)
				go handler(ctx, item)
			default:
				handler := cli.funcs["DEFAULT"].(DefaultHandler)
				go handler(ctx, msg)
			}
		}
	}
}

func (cli *LiveClient) errorHandler(cmd string, err error) {
	DPrintf("[ERROR] %s: %v", cmd, err)
}

func (cli *LiveClient) connect() (string, error) {
	token, urls, err := cli.GetTokenAndURLS()
	if err != nil {
		return "", err
	}
	for _, url := range urls {
		cli.conn, _, err = websocket.DefaultDialer.Dial(fmt.Sprintf("wss://%s:443/sub", url), nil)
		if err != nil {
			DPrintf("connect to the %s failed... Trying next one.", url)
			continue
		}
	}
	if err != nil {
		return "", err
	}
	return token, nil
}

func (cli *LiveClient) Listen() error {
	token, err := cli.connect()
	if err != nil {
		return err
	}
	auth := message.Auth{
		UID:       0,
		Roomid:    cli.roomid,
		Protover:  2,
		Platform:  "web",
		Clientver: "1.10.2",
		Type:      2,
		Key:       token,
	}
	// send auth message
	data, _ := json.Marshal(auth)
	if err := cli.Send(data, message.WsAuth); err != nil {
		cli.conn.Close()
		return err
	}

	DPrintf("Start Listening RoomID = %d", cli.roomid)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go cli.heartBeat(ctx)
	go cli.recieve(ctx)
	go cli.split(ctx)
	go cli.parse(ctx)

	select {
	case err := <-cli.heartBeatErr:
		return err
	case err := <-cli.recieveErr:
		return err
	case <-ctx.Done():
	case <-cli.close:
		return nil
	}
	return nil
}
