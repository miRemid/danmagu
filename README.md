# danmagu
一个b站直播间websocket封装包
# 快速使用
```go
package main
import (
	"log"

    "github.com/miRemid/danmagu"
    "github.com/miRemid/danmagu/message"
)
func main() {
	cli := danmagu.NewClient(271744, &danmagu.ClientConfig{
		HeartBeatTime: 30,
    })

    cli.Handler(message.DANMU_MSG, func(ctx context.Context, danmaku message.Danmaku) {
		log.Println(danmaku.Content)
	})
    
	if err := cli.Listen(); err != nil {
		log.Println(err)
	}
}
```

# Handler
Handler函数是针对不同消息的处理函数，具体方法的参数请看`function.go`