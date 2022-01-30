package internal

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// ParseTime parses the time string to time.Time
// Layout is "2006-01-02" means "YYYY-MM-DD"
func ParseTime(t string) (time.Time, error) {
	ymd := strings.Split(t, "-")
	if len(ymd) != 3 {
		return time.Time{}, errors.New("invalid time format")
	}

	year, err := strconv.Atoi(ymd[0])
	if err != nil || year < 0 || year > 9999 {
		return time.Time{}, errors.New("invalid year")
	}

	month, err := strconv.Atoi(ymd[1])
	if err != nil || month < 1 || month > 12 {
		return time.Time{}, errors.New("invalid month")
	}

	day, err := strconv.Atoi(ymd[2])
	if err != nil || day > 31 {
		return time.Time{}, errors.New("invalid day")
	}

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC), nil
}
