package common

import "time"

type DateTime struct {
	time time.Time
}

func Now() DateTime {
	return NewDateTime(time.Now())
}

func NewDateTime(t time.Time) DateTime {
	return DateTime{time: time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)}
}
