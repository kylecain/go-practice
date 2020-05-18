package main

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	ret := removeDuplicates("abbazy")
	if ret != "zy" {
		t.Errorf(`removeDuplicates("abbazy") returned %v; wanted zy`, ret)
	}
}
