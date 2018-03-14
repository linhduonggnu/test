package configs

import (
	"testing"
)

func TestStringVet(t *testing.T) {
	if StringVet() != "" {
		t.FailNow()
	}
}
