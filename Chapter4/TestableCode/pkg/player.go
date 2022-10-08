package player

import "fmt"

type Player string

const (
	PlayerOne Player = "Remi"
	PlayerTwo Player = "Yvonne"
)

func PlayerSelectPure(i int) (Player, error) {
	switch i {
	case 0:
		return PlayerOne, nil
	case 1:
		return PlayerTwo, nil
	}
	return Player(""), fmt.Errorf("no player matching input %v", i)
}
