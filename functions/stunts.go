package functions

import "fmt"

func Stunt() (int, int) {
	var level int
	fmt.Println("What level stunt did you get? 1, 2 or 3? Enter 0 for none.")
	fmt.Scan(&level)

	var dice, successes int

	switch level {
	case 1:
		dice = 2
		successes = 0
	case 2:
		dice = 2
		successes = 1
	case 3:
		dice = 2
		successes = 2
	default:
		dice = 0
		successes = 0
	}

	return dice, successes
}
