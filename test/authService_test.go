package test

import "testing"

func TestInmemoryUserService(t *testing.T) {
	got := 332
	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}
