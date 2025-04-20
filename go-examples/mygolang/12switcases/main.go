package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println("Switches in Go")

	diceNumber := rand.Intn(6) + 1
	switch diceNumber {
	case 1:
		fmt.Println("You rolled a 1, move forward 1 space")
		fallthrough // fallthrough allows the execution to continue to the next case
	case 2:
		fmt.Println("You rolled a 2")
		fallthrough
	case 3:
		fmt.Println("You rolled a 3")
		fallthrough
	case 4:
		fmt.Println("You rolled a 4")
		fallthrough
	case 5:
		fmt.Println("You rolled a 5")
		fallthrough
	case 6:
		fmt.Println("You rolled a 6")
	default:
		fmt.Println("Invalid number")

	}

}
