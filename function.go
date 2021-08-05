package danmagu

import (
	"context"
	"log"

	"github.com/fatih/color"
	"github.com/miRemid/danmagu/message"
)

var (
	yellow = color.New(color.FgYellow).SprintFunc()
	green  = color.New(color.FgGreen).SprintFunc()
	blue   = color.New(color.FgBlue).SprintFunc()
)

type HandlerFunc interface{}

var (
	defaultFuncs = map[string]HandlerFunc{
		message.DANMU_MSG: func(ctx context.Context, msg message.Danmaku) {
			DPrintf("%s, %s: %s", green(msg.UID), yellow(msg.Username), msg.Content)
		},
		message.RQZ: func(ctx context.Context, rqz uint32) {
			DPrintf("%s: %d", green("人气值"), rqz)
		},
		message.DEFAULT: func(c1 context.Context, c2 *message.Context) {
			log.Println(string(c2.Buffer))
		},
		message.SEND_GIFT: func(ctx context.Context, gift message.SendGift) {
			DPrintf("%s = %s", green("礼物名称"), blue(gift.Data.GiftName))
		},
		message.INTERACT_WORD: func(ctx context.Context, word message.InteractWord) {
		},
		message.ONLINE_RANK_V2: func(ctx context.Context, rankV2 message.OnlineRankV2) {

		},
		message.ONLINE_RANK_TOP3: func(ctx context.Context, rankTOP3 message.OnlineRankTOP3) {

		},
		message.COMBO_SEND: func(ctx context.Context, combo message.ComboSend) {

		},
		message.WIDGET_BANNER: func(ctx context.Context, banner message.WidgetBanner) {

		},
		message.ENTRY_EFFECT: func(ctx context.Context, entry message.EntryEffect) {

		},
		message.ONLINE_RANK_COUNT: func(ctx context.Context, count message.OnlineRankCount) {

		},
		message.ROOM_RANK: func(ctx context.Context, rank message.RoomRank) {
			DPrintf("%v: %v", green("直播间排名"), rank.Data.RankDesc)
		},
		message.LIVE: func(ctx context.Context, live message.Live) {
			DPrintf("[ROOM] %d starts living", live.Roomid)
		},
		message.PREPARING: func(ctx context.Context, pre message.Preparing) {
			DPrintf("[ROOM] %s stop living", pre.RoomID)
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

type LiveHandler = func(ctx context.Context, live message.Live)

type PreparingHandler = func(ctx context.Context, pre message.Preparing)
