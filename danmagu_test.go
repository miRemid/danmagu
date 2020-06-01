package danmagu

import (
	"fmt"
	"testing"
	"time"

	"github.com/miRemid/danmagu/model"
)

var cnt chan int

func test(danmaku model.Danmaku) {
	fmt.Println(danmaku.Nickname, ":", danmaku.Content)
	cnt <- 1
}
func TestDanmagu(t *testing.T) {
	cnt = make(chan int)
	count := 0
	client := NewClient(0)
	client.DebugMode = true
	client.BeforeListen = func() {
		fmt.Println("\033[2J\033[100A")
	}
	client.DanmakuHandler = test
	client.Enter(12235923)
	go client.Listen(30)
	timeout := time.After(5 * time.Minute)
	for {
		select {
		case <-timeout:
			client.Cancle()
			fmt.Printf("共耗时5分钟，共收到弹幕: %v条\n", count)
			return
		case <-cnt:
			count++
			break
		}
	}
}
