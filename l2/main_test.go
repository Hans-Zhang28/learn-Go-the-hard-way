package main

import (
	"testing"
)

func TestParallelSum(t *testing.T) {
	var slice1 = []int{0, 1, 2, 3}
	var slice2 = []int{2, -1, 7, 4}
	var slice4 = []int{2, -1, 7, 5}
	result := ParallelSum(slice1, slice2, slice4)
	if result[0] != 4 || result[1] != -1 || result[2] != 16 || result[3] != 12 || len(result) != 4 {
		t.Fail()
	}
}
