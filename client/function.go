package client

import (
	"context"
	"log"

	"github.com/fatih/color"
	"github.com/miRemid/danmagu/message"
)

type danmakuTemp struct {
	Cmd  string        `json:"cmd"`
	Info []interface{} `json:"info"`
}

var (
	yellow = color.New(color.FgYellow).SprintFunc()
	green  = color.New(color.FgGreen).SprintFunc()
	red    = color.New(color.FgRed).SprintFunc()
	blue   = color.New(color.FgBlue).SprintFunc()
)

type HandlerFunc interface{}

var (
	defaultFuncs = map[string]HandlerFunc{
		"DANMU_MSG": func(ctx context.Context, msg message.Danmaku) {
			log.Printf("%s, %s: %s", green(msg.UID), yellow(msg.Username), msg.Content)
		},
		"RQZ": func(ctx context.Context, rqz uint32) {
			log.Printf("%s: %d", green("人气值"), rqz)
		},
		"DEFAULT": func(c1 context.Context, c2 *message.Context) {
			log.Println(string(c2.Buffer))
		},
		"SEND_GIFT": func(ctx context.Context, gift message.SendGift) {
			log.Printf("%s = %s", green("礼物名称"), blue(gift.Data.GiftName))
		},
		"INTERACT_WORD": func(ctx context.Context, word message.InteractWord) {
		},
		"ONLINE_RANK_V2": func(ctx context.Context, rankV2 message.OnlineRankV2) {

		},
		"ONLINE_RANK_TOP3": func(ctx context.Context, rankTOP3 message.OnlineRankTOP3) {

		},
		"COMBO_SEND": func(ctx context.Context, combo message.ComboSend) {

		},
		"WIDGET_BANNER": func(ctx context.Context, banner message.WidgetBanner) {

		},
		"ENTRY_EFFECT": func(ctx context.Context, entry message.EntryEffect) {

		},
		"ONLINE_RANK_COUNT": func(ctx context.Context, count message.OnlineRankCount) {

		},
		"ROOM_RANK": func(ctx context.Context, rank message.RoomRank) {
			log.Printf("%v: %v", green("直播间排名"), rank.Data.RankDesc)
		},
	}
)

type DanmakuHandler = func(context.Context, message.Danmaku)

type RQZHandler = func(context.Context, uint32)

type GiftHandler = func(ctx context.Context, gift message.SendGift)

type InteractWordHandler = func(ctx context.Context, word message.InteractWord)

type RankV2Handler = func(ctx context.Context, rankV2 message.OnlineRankV2)

type RankTOP3Handler = func(ctx context.Context, rankTOP3 message.OnlineRankTOP3)

type ComboSendHandler = func(ctx context.Context, combo message.ComboSend)

type WidgetBannerHandler = func(ctx context.Context, banner message.WidgetBanner)

type EntryEffectHandler = func(ctx context.Context, entry message.EntryEffect)

type RankCountHandler = func(ctx context.Context, count message.OnlineRankCount)

type DefaultHandler = func(context.Context, *message.Context)

type RoomRankHandler = func(context.Context, message.RoomRank)
