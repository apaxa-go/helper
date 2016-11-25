package httph

import (
	"github.com/apaxa-go/helper/timeh"
	"testing"
)

func TestCookieDeleteTime(t *testing.T) {
	if r := CookieDeleteTime().UnixNano(); r != 1*timeh.NanosecsInSec {
		t.Errorf("TestCookieDeleteTime: expected %v, got %v", timeh.NanosecsInSec, r)
	}
}
