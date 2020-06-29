package main

import (
	"testing"
)

type myType int

func TestReverse(t *testing.T) {
	slice := []myType{0, 1, 2}
	Reverse(slice)
	if slice[0] != myType(2) || slice[1] != myType(1) || slice[2] != myType(0) || len(slice) != 3 {
		t.Fail()
	}
	anotherSlice := []interface{}{1, "2", uint(3), byte(4), float64(5)}
	Reverse(anotherSlice)
	if anotherSlice[0] != float64(5) || anotherSlice[1] != byte(4) || anotherSlice[2] != uint(3) || anotherSlice[3] != "2" || anotherSlice[4] != 1 || len(anotherSlice) != 5 {
		t.Fail()
	}
}
