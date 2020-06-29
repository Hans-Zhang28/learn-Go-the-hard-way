package main

import (
	"reflect"
)

//Reverse reverses a slice.
func Reverse(slice interface{}) {
	n := reflect.ValueOf(slice).Len()
	swap := reflect.Swapper(slice)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

func main() {
}
