# danmagu
一个b站直播间websocket封装包
# 快速使用
```go
package main
import "github.com/miRemid/danmagu"
import "github.com/miRemid/danmagu/model"
import "fmt"

func main() {
    // 创建一个连接，参数为本人uid，可以为0(功能暂停)
    client := danmagu.NewClient(0)
    // 设置Debug模式
    client.DebugMode = true
    // 钩子函数
    client.BeforeListen = func() {
        fmt.Println("Begin!")
    }
    // 设置弹幕处理函数
    client.DanmakuHandler = func(danmaku model.Danmaku) {
        fmt.Println(danmaku.Content)
    }
    // 设置房间id，仅支持长id
    client.Enter(roomid)
    // 开始监听，设置心跳响应间隔，70s以下建议30s
    go client.Listen(30)
    select {
    case <- time.After(time.Second * time.Duration(60)):
        // 60s后取消监听
        client.Cancle()
    }
}
```

# 钩子函数
```golang
通过钩子函数可以在建立danmagu的生命周期中添加少许操作
client.BeforeConnect
client.AfterConnect
...

// 钩子函数结构
type HandlerFunc func()
```