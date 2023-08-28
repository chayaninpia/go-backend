package timex

import "time"

func TimeNowPtr() *time.Time {
	timeNow := time.Now()
	return &timeNow
}

func TimeNowPtrUTC() *time.Time {
	timeNow := time.Now().UTC()
	return &timeNow
}

func DateNow() time.Time {
	timeNow := time.Now()
	return time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 0, 0, 0, 0, time.Local)
}

func DateNowUTC() time.Time {
	timeNow := time.Now()
	return time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 0, 0, 0, 0, time.UTC)
}

func DateNowPtr() *time.Time {
	timeNow := time.Now()
	dateNow := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 0, 0, 0, 0, time.Local)
	return &dateNow
}

func DateNowPtrUTC() *time.Time {
	timeNow := time.Now()
	dateNow := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 0, 0, 0, 0, time.UTC)
	return &dateNow
}
