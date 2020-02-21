package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/miRemid/danmagu"
	"github.com/miRemid/danmagu/model"
)

func parse(message []byte) {
	h := message[:16]
	var head model.RHeader
	buf := bytes.NewReader(h)
	binary.Read(buf, binary.BigEndian, &head)
	body := message[16:head.Length]

	switch head.Type {
	case danmagu.WsHeartbeatReply:
		var rqz model.Population
		binary.Read(bytes.NewReader(body), binary.BigEndian, &rqz)
		if rqz.Value > 1 {
			fmt.Print(color.BlueString("当前人气值=%v\n", rqz.Value))
		}
	case danmagu.WsMessage:
		var cmd model.CMD
		_ = json.Unmarshal(body, &cmd)
		switch cmd.Cmd {
		case "DANMU_MSG":
			var danmaku model.Danmaku
			json.Unmarshal(body, &danmaku)
			log.Println(string(body))
			fmt.Print(color.GreenString("%s: %s\n", cmd.Info[2][1], danmaku.Info[1]))
		case "SEND_GIFT":
			var g model.Gift
			json.Unmarshal(body, &g)
			log.Println(string(body))
			fmt.Print(color.RedString("%s: %s (%s) x %d\n", g.Data.Uname, g.Data.GiftName, g.Data.CoinType, g.Data.Num))
		case "GUARD_BUY":
			var g model.Guard
			json.Unmarshal(body, &g)
			fmt.Print(color.YellowString("%s: %s (%s) x %d\n", g.Data.Username, g.Data.GiftName, "gold", g.Data.Num))
		}
	}
	next := message[head.Length:]
	if binary.Size(next) != 0 {
		parse(next)
	}
}

func main() {
	client := danmagu.NewClient(0)
	client.DebugMode = true
	client.BeforeListen = func() {
		fmt.Println("\033[2J\033[100A")
	}
	client.Enter(1029)
	client.OnMessage(parse)
	client.Listen(30)
}
