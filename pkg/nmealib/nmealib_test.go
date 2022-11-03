package nmealib

import (
	"testing"
	"time"
)

func TestFromTimeToZDA(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			"2006-01-02T15:04:05Z",
			"$GPZDA,150405.000,02,01,2006,,\r\n",
		},
	}

	for _, tt := range tests {
		time, err := time.Parse(time.RFC3339, tt.input)
		if err != nil {
			t.Errorf("invalid input: %v", err)
		}
		got := FromTimeToZDA(time)
		if got != tt.want {
			t.Errorf("got = \"%s\", want = \"%s\"", got, tt.want)
		}
	}
}
