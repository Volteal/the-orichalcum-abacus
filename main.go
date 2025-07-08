package main

import (
	"fmt"

	"github.com/volteal/the-orichalcum-abacus/attacks"
)

func main() {
	var attkType int
	fmt.Println("What kind of attack is it?")
	fmt.Println("1. A Withering Attack")
	fmt.Println("2. A Decisive Attack")
	fmt.Scan(&attkType)

	switch attkType {
	case 1:
		attacks.WitheringAttack()
	case 2:
		fmt.Println("Error: Withering Damage Not yet Implemented.")
	case 3:
		fmt.Println("Error: Decisive Attacks Not yet Implemented.")
	case 4:
		fmt.Println("Error: Decisive Damage Not yet Implemented.")
	default:
		fmt.Println("That option does not exist.")
	}
}
