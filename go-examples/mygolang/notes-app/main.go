package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/prateekdegaons1991/notes-app/note"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error creating note:", err)
		return
	}

	userNote.Display()
	userNote.SaveToFile()
	if err != nil {
		fmt.Println("Error saving note:", err)
		return
	}

}

func getNoteData() (string, string) {

	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")
	return title, content

}
func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")
	return input

}
