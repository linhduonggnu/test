package common

import (
	"testing"
)

func TestGetVal(t *testing.T) {
	if GetVal() != 12 {
		t.FailNow()
	}
}
