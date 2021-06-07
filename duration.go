package intutils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Duration - тип для хранения длительность медийных треков.
type Duration int64

// NewDurationFromString converts string expression like "hh:mm:ss" to integer value with miliseconds.
// If error was occured zero is returned.
func NewDurationFromString(val string) Duration {
	flds := strings.Split(val, ":")
	if len(flds) < 2 {
		return 0
	}
	seconds, err := strconv.Atoi(flds[len(flds)-1])
	minutes, err := strconv.Atoi(flds[len(flds)-2])
	hours := 0
	if len(flds) == 3 {
		hours, err = strconv.Atoi(flds[0])
	}
	if err != nil {
		return 0
	}
	return Duration(int64(1000 * (seconds + 60*minutes + 3600*hours)))
}

// String converts miliseconds to string representation of duration like "hh:mm:ss".
func (dur Duration) String() string {
	d := time.Duration(dur * 1000 * 1000)
	hours := int(d.Hours())
	d -= time.Duration(hours) * time.Hour
	minutes := int(d.Minutes())
	d -= time.Duration(minutes) * time.Minute
	seconds := int(d.Seconds())
	if hours > 0 {
		return fmt.Sprintf("%d:%02d:%02d", hours, minutes, seconds)
	}
	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}
