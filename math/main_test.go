package main

import "testing"

// go test -timeout 30s -run TestSumExpectNoError github.com/NoSkillGuy/go-test/math
func TestSumExpectNoError(t *testing.T) {
	got := sum(1, 2)
	if got != 3 {
		t.Errorf("expected 3 got %d", got)
	}
}

func TestSumExpectError(t *testing.T) {
	got := sum(1, 2)
	if got != 2 {
		t.Errorf("expected 2 got %d", got)
	}
}

// go test -timeout 30s -run TestSumExpectError github.com/NoSkillGuy/go-test/math

// === RUN   TestSumExpectError
//     go-test/basic_test/main_test.go:15: expected 2 got 3
// --- FAIL: TestSumExpectError (0.00s)
// FAIL
// FAIL    github.com/NoSkillGuy/go-test/math        0.332s
