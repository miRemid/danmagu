package danmagu

import "log"

var (
	LOG_TOGGLE = true
)

func DPrintf(format string, args ...interface{}) {
	if LOG_TOGGLE {
		log.Printf("[DEBUG] "+format, args...)
	}
}
