# danmagu
一个b站直播间websocket封装包
# 快速使用
```go
package main
import (
	"log"
	"time"
	"context"
	"encoding/json"

    "github.com/miRemid/danmagu"
    "github.com/miRemid/danmagu/message"
)
func main() {
	cli := danmagu.NewClient(271744, &danmagu.ClientConfig{
		HeartBeatTime: 30 * time.Second,
		HttpTimeout: 10 * time.Second,
    })
	
	// 注册已序列化函数
    cli.Handler(message.DANMU_MSG, func(ctx context.Context, danmaku message.Danmaku) {
		log.Println(danmaku.Content)
	})

	// 注册未序列化函数
	cli.RawHandler("ROOM_BLOCK_MSG", func(ctx context.Context, msg *message.Context) {
		var mapper = make(map[string]interface{})
		json.Unmarshal(msg.Buffer, &mapper)
		log.Println(mapper)
	})
    
	if err := cli.Listen(); err != nil {
		log.Println(err)
	}
}
```

# Handler
Handler函数是针对不同消息的处理函数，具体方法的参数请看`function.go`

# RawHandler
RawHandler函数是针对所有消息的处理函数，参数为`*message.Context`其中包含了消息的`操作码(uint32)`和`消息本体([]byte)`，具体方法请查看`function.go`