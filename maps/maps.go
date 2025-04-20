package main

import "fmt"

func main() {

	// creating a map

	m := make(map[string]string)
	m2 := map[string]int{"price": 990, "age": 33}

	fmt.Println("m2:", m2)
	// setting an element in the map

	m["name"] = "John"
	m["age"] = "25"

	// getting an element from the map

	name := m["name"]
	age := m["age"]

	fmt.Println("Name:", name, "Age:", age, "Phone:", m["phone"]) // IMP: If the key is not present in the map, it will return the zero value of the value type

	clear(m)

	fmt.Println(m)

	_, ok := m2["price"]

	if ok {
		fmt.Println("all ok", ok)

	} else {
		fmt.Println("not ok", ok)
	}

}
