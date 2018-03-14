package common

var globalVal int

// SetVal ./...
func SetVal(v int) {
	globalVal = v
}

// GetVal ./...
func GetVal() int {
	return globalVal
}
