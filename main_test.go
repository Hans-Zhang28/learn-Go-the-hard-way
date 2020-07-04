package main

import (
	"testing"
)

func TestGame(t *testing.T) {
	win := Game()
	counter := 2
	for _, w := range win {
		if !w {
			counter--
		}
		if counter == 0 {
			t.Fail()
		}
	}
}
