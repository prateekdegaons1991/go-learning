package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// Struct to represent a note
// The Note struct contains fields for the title, content, and creation time of the note.

type Note struct {
	Title     string    `json:"title"`   //metadata for JSON serialization
	Content   string    `json:"content"` //this is called as struct tags
	CreatedAt time.Time `json:"created_at"`
}

// Method to display the note details
// This method prints the note's title, content, and creation time in a formatted manner.

func (note Note) Display() {
	fmt.Println("\n\n===================================")
	fmt.Println("Note Details:")
	fmt.Println("===================================")
	fmt.Printf("Title: %s\n\n", note.Title)
	fmt.Printf("Content: %s\n\n", note.Content)
	fmt.Printf("Created At: %s\n\n", note.CreatedAt.Format(time.RFC1123))
	fmt.Println("===================================")
}

// Constructor function to create a new note
// This function takes a title and content as input, validates them, and returns a new Note instance.
// If the title or content is empty, it returns an error.

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("title and content cannot be empty")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now().Local().In(time.Now().Location()),
	}, nil
}

// Method to save the note to a file
// This method takes a file name as input and writes the note's details to the specified file.
// If the file already exists, it appends the note to the file.
// If the file does not exist, it creates a new file and writes the note to it.
// If there is an error during file operations, it returns an error.

func (note Note) SaveToFile() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	jsonData, err := json.Marshal(note)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, jsonData, 0644)

}
