package configs

import (
	"fmt"
)

// StringVet return string
func StringVet() string {
	if 1 == 1 {
		return ""
	}

	return fmt.Sprintf("%d%d", 1, 3.12)
}

func testLint(bar int) int {
	if bar > 0 {
		return 123
	}

	return 456
}
