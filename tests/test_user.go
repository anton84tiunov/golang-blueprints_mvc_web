package tests

import (
	"math"
	"testing"
)

func Test_user(t *testing.T) {
	got := math.Abs(-2)
	if got != 1 {
		t.Errorf("Abs(-1) = %f; want 1", got)
	}
}
