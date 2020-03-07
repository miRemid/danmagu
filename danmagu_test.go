package danmagu

import (
	"fmt"
	"testing"

	"github.com/miRemid/danmagu/model"
)

func test(danmaku model.Danmaku) {
	fmt.Println(danmaku.Content)
}

func TestDanmagu(t *testing.T) {
	client := NewClient(0)
	client.DebugMode = true
	client.BeforeListen = func() {
		fmt.Println("\033[2J\033[100A")
	}
	client.DanmakuHandler = test
	client.Enter(56159)
	client.Listen(30)
}
