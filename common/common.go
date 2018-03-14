package common

var globalVal int

func SetVal(v int) {
	globalVal = v
}

// GetVal ./...
func GetVal() int {
	return globalVal
}
