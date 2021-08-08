package danmagu

import "time"

type ClientConfig struct {
	HeartBeatTime time.Duration
	HttpTimeout   time.Duration
}
