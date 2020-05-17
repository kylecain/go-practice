package main

import (
	"testing"
)

func TestIsUgly(t *testing.T) {

	if !isUgly(1) {
		t.Errorf("Ugly(1) returned false; wanted true")
	}

	if !isUgly(6) {
		t.Errorf("Ugly(6) returned false; wanted true")
	}

	if !isUgly(8) {
		t.Errorf("Ugly(8) returned false; wanted true")
	}

	if !isUgly(10) {
		t.Errorf("Ugly(10) returned false; wanted true")
	}

	if isUgly(14) {
		t.Errorf("Ugly(14) returned true; wanted false")
	}

	if isUgly(-6) {
		t.Errorf("Ugly(-6) returned true; wanted false")
	}

}
