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

	urls := make(map[string]string)
	urls["google"] = "https://www.google.com"
	urls["en"] = "https://www.google.com/en"
	urls["es"] = "https://www.google.com/es"
	urls["fr"] = "https://www.google.com/fr"
	urls["de"] = "https://www.google.com/de"
	// Printing the map
	fmt.Println("List of urls: ", urls)
	// Accessing a map
	// Get German url

	fmt.Println("Get value german: ", urls["de"])

	for key, value := range urls {
		fmt.Printf("For Key: %v"+" Value: %v\n", key, value)
	}

}
