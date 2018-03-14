package configs

import (
	"fmt"
	"testing"
)

func TestStringVet(t *testing.T) {
	if StringVet() != "" {
		t.FailNow()
	}
	fmt.Println("done")
}
