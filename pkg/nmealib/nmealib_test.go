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
			"$GPZDA,150405.000,02,01,2006,,*54\r\n",
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

func TestGetChecksum(t *testing.T) {
	inputs := []struct {
		input, want string
	}{
		{
			"GPAAM,A,A,0.10,N,WPTNME",
			"32",
		},
		{
			"GPRMC,085120.307,A,3541.1493,N,13945.3994,E,000.0,240.3,181211,,,A",
			"6A",
		},
		{
			"GPVTG,240.3,T,,M,000.0,N,000.0,K,A",
			"08",
		},
	}
	for _, tt := range inputs {
		if getChecksum(tt.input) != tt.want {
			t.Errorf("getChecksum(%s) = %s, want = %s", tt.input, getChecksum(tt.input), tt.want)
		}
	}
}
