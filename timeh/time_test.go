package timeh

import (
	"testing"
	"time"
)

func TestUnixEpoch(t *testing.T) {
	s := "01 Jan 70 00:00 -0600"
	time1, err := time.Parse(time.RFC822Z, s)
	if err != nil {
		t.Errorf("Parsing string:%v\ngot err: %v", s, err)
	}
	time2 := UnixEpoch()
	if time2.Equal(time1) {
		t.Errorf("Wrong time. Expexted: %v, got: %v", time1, time2)
	}
}
