package mstotime

import (
	"time"
)

func MsToTime(ms uint64) time.Time {
	var msInt = int64(ms)
	return time.Unix(0, msInt*int64(time.Millisecond))
}
