package main

import (
	"fmt"
	"math/rand"
)

type Player string

const (
	PlayerOne Player = "Remi"
	PlayerTwo Player = "Yvonne"
)

func main() {
}

func selectStartingPlayer() Player {
	randomized := rand.Intn(2)
	switch i {
	case 0:
		return PlayerOne, nil
	case 1:
		return PlayerTwo, nil
	}
	panic("No further player available")
}

func selectStartingPlayerPure(i int) (Player, error) {
	switch i {
	case 0:
		return PlayerOne, nil
	case 1:
		return PlayerTwo, nil
	}
	return Player(""), fmt.Errorf("no player matchin %v", i)
}
