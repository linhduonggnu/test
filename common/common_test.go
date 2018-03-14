package common

import (
	"testing"
)

func init() {
	SetVal(12)
}

func TestSetVal(t *testing.T) {
	if globalVal != 12 {
		t.FailNow()
	}
}
