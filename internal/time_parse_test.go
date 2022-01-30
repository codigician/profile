package internal_test

import (
	"errors"
	"testing"
	"time"

	"github.com/codigician/profile/internal"
	"github.com/stretchr/testify/assert"
)

func TestParseTime(t *testing.T) {
	testCases := []struct {
		scenario        string
		givenStringTime string
		expectedTime    time.Time
		expectedErr     error
	}{
		{
			scenario:        "invalid string format more than 3 parts return error",
			givenStringTime: "2020-11-11-22",
			expectedErr:     errors.New("invalid time format"),
		},
		{
			scenario:        "invalid string format less than 3 parts return error",
			givenStringTime: "2020-11",
			expectedErr:     errors.New("invalid time format"),
		},
		{
			scenario:        "invalid string format invalid year return error",
			givenStringTime: "202s-11-11",
			expectedErr:     errors.New("invalid year"),
		},
		{
			scenario:        "invalid string format invalid month return error",
			givenStringTime: "2020-139-11",
			expectedErr:     errors.New("invalid month"),
		},
		{
			scenario:        "invalid string format invalid day return error",
			givenStringTime: "2020-11-23s",
			expectedErr:     errors.New("invalid day"),
		},
		{
			scenario:        "valid string format return time",
			givenStringTime: "2020-11-11",
			expectedTime:    time.Date(2020, 11, 11, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.scenario, func(t *testing.T) {
			tm, err := internal.ParseTime(tC.givenStringTime)

			assert.Equal(t, tC.expectedTime, tm)
			assert.Equal(t, tC.expectedErr, err)
		})
	}
}
