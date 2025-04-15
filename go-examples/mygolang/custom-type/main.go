package main

import "fmt"

// Custom type creation

type str string

func (text str) log() {

	fmt.Println(text)
}

func main() {
	// Custom type creation
	var logMessage str = "Info: this is a custom type!"
	logMessage.log()

}
