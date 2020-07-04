package main

import (
	"testing"
)

func TestGame(t *testing.T) {
	win := Game()
	counter := 2 // win if someone won 2 out 3 games
	for _, w := range win {
		if !w {
			counter--
		}
		if counter == 0 {
			t.Fail()
		}
	}
}
