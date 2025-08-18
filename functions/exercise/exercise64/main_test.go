package main

import "testing"

func TestAdd(t *testing.T) {
	total := add(2, 2)
	if total != 4 {
		t.Errorf("Output should be 4 but it is isntead %d", total)
	}
}
