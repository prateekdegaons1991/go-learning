package main

import (
	"fmt"
)

func main() {

	// colors := make(map[string]string)

	colors := map[string]string{
		"Red":    "#ff00",
		"Green":  "#ffkh",
		"Yellow": "#ffoo1",
		"Blue":   "#ff0002",
	}
	// colors["white"] = "#000000"

	fmt.Printf("colors: %+v", colors)
}
