package main

import "testing"

func TestSum(t *testing.T) {
	got := sum(1, 2)
	if got != 3 {
		t.Errorf("expected 3 got %d", got)
	}
}

