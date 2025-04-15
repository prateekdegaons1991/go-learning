package main

import (
	"fmt"
)

func main() {

	fmt.Println("Maps in Go")

	// Creating a map

	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"
	languages["GO"] = "Go"

	// Printing the map
	fmt.Println("List of languages: ", languages)
	// Accessing a map
	fmt.Println("Get value python: ", languages["PY"])
	// Deleting a map
	delete(languages, "RB")
	fmt.Println("List of languages after deleting Ruby: ", languages)
	// Checking if a key exists
	_, exists := languages["RB"]
	if exists {
		fmt.Println("Key exists")
	} else {
		fmt.Println("Key does not exist")
	}
	// Iterating over a map // loops over a map
	for key, value := range languages {
		fmt.Printf("For Key: %v\n"+" Value: %v\n", value, key)
	}

}
