package main

import "fmt"

func main() {

	// creating a map

	m := make(map[string]string)

	// setting an element in the map

	m["name"] = "John"
	m["age"] = "25"

	// getting an element from the map

	name := m["name"]
	age := m["age"]

	fmt.Println("Name:", name, "Age:", age, "Phone:", m["phone"]) // IMP: If the key is not present in the map, it will return the zero value of the value type

}
