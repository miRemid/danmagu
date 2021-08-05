package danmagu_test

import (
	"context"
	"log"
	"testing"

	"github.com/miRemid/danmagu"
	"github.com/miRemid/danmagu/message"
)

func TestDanmaku(t *testing.T) {
	cli := danmagu.NewClient(56159, &danmagu.ClientConfig{
		HeartBeatTime: 30,
	})

	cli.Handler(message.DANMU_MSG, func(ctx context.Context, danmaku message.Danmaku) {
		log.Println(danmaku.Content)
	})

	if err := cli.Listen(); err != nil {
		log.Println(err)
	}
}
