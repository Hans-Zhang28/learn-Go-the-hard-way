package main

import (
	"fmt"
	"math/rand"
)

const (
	ROCK int = iota
	PAPER
	SCISSORS
)

type Choice struct {
	Who   int //0 you 1 your opponent
	Guess int
}

//Win returns true if you win.
func Win(you, he int) bool {
	if you == ROCK && he == SCISSORS {
		return true
	}
	if you == PAPER && he == ROCK {
		return true
	}
	if you == SCISSORS && he == PAPER {
		return true
	}
	return false
}

func Opponent(guess chan Choice, please chan struct{}) {
	for i := 0; i < 3; i++ {
		<-please
		choice := rand.Intn(3)
		who := 1
		guess <- Choice{who, choice}
		please <- struct{}{}
	}
}

func Me(guess chan Choice, please chan struct{}) {
	for i := 0; i < 3; i++ {
		<-please
		choice := rand.Intn(3)
		who := 0
		guess <- Choice{who, choice}
		please <- struct{}{}
	}
}

func Game() []bool {
	guess := make(chan Choice)
	//please sync 2 goroutines.
	please := make(chan struct{})
	go func() { please <- struct{}{} }()
	go Opponent(guess, please)
	go Me(guess, please)
	guess = Cheat(guess)
	var wins []bool

	for i := 0; i < 3; i++ {
		g1 := <-guess
		g2 := <-guess
		fmt.Println("g1", g1)
		fmt.Println("g2", g2)
		win := false
		if g1.Who == 0 {
			win = Win(g1.Guess, g2.Guess)
		} else {
			win = Win(g2.Guess, g1.Guess)
		}
		fmt.Println("game result", win)
		wins = append(wins, win)
	}

	return wins
}

func Cheat(guess chan Choice) chan Choice {
	out := make(chan Choice)
	go func() {
		myNewGuess := 0
		for choice := range guess {
			if choice.Who == 1 {
				// save my new choice after looking at the guess from opponent
				myNewGuess = (choice.Guess + 1) % 3
			} else {
				// update my guess if the choice is from me
				choice.Guess = myNewGuess
			}
			out <- choice
		}
	}()
	return out
}

func main() {
	println("Now let's play a game 'rock-paper-scissors',there are 2 players-you and a goroutine!\nTo be bound to win,you should call a goroutine to help you to peer what your opponent choose.\nTwo out of three.\nPlease edit main.go to complete func 'Cheat' to win!")
}
