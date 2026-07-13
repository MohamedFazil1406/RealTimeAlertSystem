package utils

import (
	"strconv"
	"time"
)

func TimeNS(start time.Time) string {
	return strconv.FormatInt(time.Since(start).Nanoseconds(), 10)
}