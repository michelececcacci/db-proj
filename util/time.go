package util

import "time"
const shortForm = "2006-Jan-02"

func ParseTime(s string) (time.Time, error) {
	return time.Parse(shortForm, s)
}