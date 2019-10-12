# danmagu
一个b站直播间websocket包
# 快速使用
```go
package main
import "github.com/miRemid/danmagu"

func parse(data []byte){
    // Parse data
}

func main() {
    // 创建一个连接，参数为本人uid，可以为0
    client := danmagu.NewClient(0)
    // 认证房间
    client.Auth(roomid)
    // 设置消息处理函数
    client.OnMessage(parse)
    // 开始监听，设置心跳响应间隔，70s以下建议30s
    client.Listen(heart beat time)
}
```