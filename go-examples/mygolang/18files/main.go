package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Files in Go")

	content := GetContent()
	filename := GetFileName()

	WriteToFile(filename, content)

}

// Function to creat file and typ of file based on extension

// Function to write content to file
func WriteToFile(filename string, content string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		panic(err)
	}

	fmt.Println("File created successfully: ", filename)
}

// function to get content by user

func GetContent() string {
	fmt.Println("Enter the content (press Enter twice to finish):")
	var content string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		content += line + "\n"
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
	return content
}

// function to get file name by user
func GetFileName() string {
	var filename string
	fmt.Println("Enter the file name: ")
	fmt.Scanln(&filename)
	return filename
}
