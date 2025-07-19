package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {
	fmt.Println("go interfaces")

	// sb := spanishBot{}

	eb := englishBot{}
	printGreeting(eb)
	// printGreeting(sb)
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func (englishBot) getGreeting() string {
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "¡Hola!"
}
