package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	if sum([]int{1, 1, 1}) != 3 {
		t.Error("Expected 3")
	}

	if sum([]int{1, 1, -2}) != 0 {
		t.Error("Expected 0")
	}

	if sum([]int{-1, -2, -3}) != -6 {
		t.Error("Expected -6")
	}
}

func TestTwice(t *testing.T) {
	if twice([]int{1, -1}) != 0 {
		t.Error("Expected 0")
	}

	if twice([]int{3, 3, 4, -2, -4}) != 10 {
		t.Error("Expected 10")
	}

	if twice([]int{-6, 3, 8, 5, -6}) != 5 {
		t.Error("Expected 5")
	}

	if twice([]int{7, 7, -2, -7, -4}) != 14 {
		t.Error("Expected 14")
	}
}
