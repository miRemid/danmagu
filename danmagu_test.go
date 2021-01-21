package danmagu_test

import (
	"context"
	"log"
	"testing"

	"github.com/miRemid/danmagu/client"
	"github.com/miRemid/danmagu/message"
)

func TestDanmaku(t *testing.T) {
	cli := client.NewClient(271744, &client.ClientConfig{
		HeartBeatTime: 30,
	})

	cli.Handler(message.DANMU_MSG, func(ctx context.Context, danmaku message.Danmaku) {
		log.Println(danmaku.Content)
	})

	if err := cli.Listen(); err != nil {
		log.Println(err)
	}
}
