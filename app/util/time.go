package util

import "time"

func DateFromMidnight(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}
