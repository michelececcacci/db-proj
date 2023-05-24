package util

import "time"

var (
	validShortForms = []string{"2006-Jan-02"} // treated as const
)

func ParseTime(s string) (time.Time, error) {
	var err error
	for _, shortForm := range validShortForms {
		t, err := time.Parse(shortForm, s)
		if err == nil {
			return t, nil
		}
	}
	return time.Time{}, err
}
