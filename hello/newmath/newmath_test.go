package newmath

// Run: go test -v go-samples/hello/newmath
// From GOPATH (note: no src/)

import "testing"

func TestMul(t *testing.T) {
	if Mul(2.0, 3.0) != 6.0 {
		t.Errorf("Mul failed")
	}
}

// Can test private methods too because it's in the same package.
func TestMinus(t *testing.T) {
	if minus(4.0, 2.0) != 2.0 {
		t.Error("minus failed")
	}
}
