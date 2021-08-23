package danmagu_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"
	"time"

	"github.com/miRemid/danmagu"
	"github.com/miRemid/danmagu/message"
)

func TestDanmaku(t *testing.T) {
	cli := danmagu.NewClient(6136246, &danmagu.ClientConfig{
		HeartBeatTime: 30 * time.Second,
		HttpTimeout:   10 * time.Second,
	})

	cli.Handler(message.DANMU_MSG, func(ctx context.Context, danmaku message.Danmaku) {
		log.Println(danmaku.Content)
	})

	cli.RawHandler("ROOM_BLOCK_MSG", func(ctx context.Context, msg *message.Context) {
		var mapper = make(map[string]interface{})
		json.Unmarshal(msg.Buffer, &mapper)
		log.Println(mapper)
	})

	if err := cli.Listen(); err != nil {
		log.Println(err)
	}
}
