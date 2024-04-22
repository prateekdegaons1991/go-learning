package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your username: ")

	// comma ok || err ok

	input, _ := reader.ReadString('\n')

	fmt.Println("your name is :", input)
	fmt.Printf("Type of input is %T \n", input)

}
