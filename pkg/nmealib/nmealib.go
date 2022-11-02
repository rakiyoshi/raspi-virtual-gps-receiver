package nmealib

import (
	"fmt"
	"strings"
	"time"
)

func FromTime(t time.Time) string {
	var data []string
	data = append(data, "$GPZDA")
	data = append(data, fmt.Sprintf("%02d%02d%02d.%03d", t.Hour(), t.Minute(), t.Second(), t.Nanosecond()/1000/1000)) // UTC
	data = append(data, fmt.Sprintf("%02d", t.Day()))
	data = append(data, fmt.Sprintf("%02d", t.Month()))
	data = append(data, fmt.Sprintf("%02d", t.Year()))
	data = append(data, "")
	data = append(data, "")
	return strings.Join(data, ",")
}
