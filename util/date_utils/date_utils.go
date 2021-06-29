package date_utils

import "time"

const (
	formatString = "2006-01-02T15:0405Z"
	apiDbDateLayout = "2006-01-02 15:04:05"
)

func GetTimeNow() time.Time {
	return time.Now().UTC()
}

func GetTimeNowString() string {
	now := GetTimeNow()
	return now.Format(formatString)
}

func GetDbTimeNowString() string {
	now := GetTimeNow()
	return now.Format(apiDbDateLayout)
}