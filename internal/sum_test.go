package internal

import (
	"testing"
)

func TestSum(t *testing.T) {

	s := Sum(1, 2)
	if s != 3 {
		t.Errorf("Expected 3, but got %d", s)
	}
}

func TestAvg(t *testing.T) {

	avg := Avg(2, 2)
	if avg != 2 {
		t.Errorf("Expected 2, but got %d", avg)
	}
}
