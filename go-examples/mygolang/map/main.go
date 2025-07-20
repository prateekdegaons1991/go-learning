package main

import (
	"fmt"
)

func main() {

	// colors := make(map[string]string)

	colors := map[string]string{
		"Red":    "#ff0000",
		"Green":  "#00ff00",
		"Yellow": "#ffff00",
		"Blue":   "#0000ff",
		"Black":  "#000000",
	}
	// colors["white"] = "#000000"

	fmt.Printf("colors: %+v", colors)

	delete(colors, "Red")

	fmt.Println("After deleting Red color:", colors)

	fmt.Println(colors["Green"]) // accessing value of key Green)

	// using make function to create a map

	colors2 := make(map[string]string)

	fmt.Println("colors2:", colors2)

	// iterate over a map

	for colour := range colors {
		fmt.Println("Colour:", colour, "Value:", colors[colour])
	}

	printMap(colors)
	manipulateMap(colors)

}

func printMap(c map[string]string) {
	for colour, hex := range c {
		fmt.Println("Colour:", colour, "Hex:", hex)

	}
}

// manupulate map

func manipulateMap(c map[string]string) {
	// create a map
	c = make(map[string]string)

	// add key-value pairs
	c["Red"] = "#ff0000"
	c["Green"] = "#00ff00"
	c["Blue"] = "#0000ff"

	// update a value
	c["Red"] = "#ff0001"

	// delete a key-value pair
	delete(c, "Green")

	// check if a key exists
	if hex, exists := c["Blue"]; exists {
		fmt.Println("Blue color hex code:", hex)
	} else {
		fmt.Println("Blue color not found")
	}
}
