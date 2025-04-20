package main

import "fmt"

func main() {

	fmt.Println("Loops in Go")

	days := []string{"Sunday", "Tuesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	// For loop
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	// For range loop
	for i := range days {
		fmt.Println(i, days[i])
	}

	// For range loop with value
	for i, d := range days {
		fmt.Printf("index is: %d: and value is %s\n", i, d)
	}

	for _, d := range days {
		fmt.Printf("value is %s\n", d)
	}

	// Iterartion over values

	rougueValue := 1

	for rougueValue <= 10 {

		if rougueValue == 2 {

			goto goTo
		}

		if rougueValue == 5 {
			rougueValue++
			continue
		}
		fmt.Println(rougueValue)
		rougueValue++
	}

	// goto statement
goTo:
	fmt.Println("goto statement")
	

}
