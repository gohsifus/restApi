package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestEvent_Validate(t *testing.T) {
	testTable := []struct {
		name  string
		event Event
		want  bool
	}{
		{
			name:  "valid event",
			event: *NewEvent(strToTime("2022-09-09"), "HeaderOfEvent", "DescriptionOfEvent"),
			want:  true,
		},
		{
			name:  "empty name of event",
			event: *NewEvent(strToTime("2022-09-09"), "", "DescriptionOfEvent"),
			want:  false,
		},
		{
			name:  "empty description of event",
			event: *NewEvent(strToTime("2022-09-09"), "HeaderOfEvent", ""),
			want:  false,
		},
	}

	for _, testCase := range testTable {
		got := testCase.event.Validate()
		assert.Equal(t, testCase.want, got)
	}
}

func strToTime(str string) time.Time {
	date, _ := time.Parse("2006-01-02", str)
	return date
}
