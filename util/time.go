package util

import "time"

const layout = "2006-01-02T15:04:05"

// "2021-11-28T11:00:00"
func StringToTime(value string) time.Time {
	t, _ := time.Parse(layout, value)
	return t
}
