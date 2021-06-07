package intutils

import (
	"testing"
)

func TestMinOf(t *testing.T) {
	if MinOf(5, 1, 4) != 1 {
		t.FailNow()
	}
}

func TestTrackDurationFromString(t *testing.T) {
	if NewDurationFromString("37") != 0 {
		t.Fail()
	}
	if NewDurationFromString("1:25:37") != (3600+25*60+37)*1000 {
		t.Fail()
	}
}

func TestTrackDurationToString(t *testing.T) {
	dur := Duration((3600 + 25*60 + 37) * 1000)
	if dur.String() != "1:25:37" {
		t.Fail()
	}
}
