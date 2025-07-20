package main

// lets learn, go routines in this example

import (
	"fmt"
	"sync"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Notify the WaitGroup that this task is done, defer -- function completion
	fmt.Println("Task", id, "is running")
}

func main() {

	var wg sync.WaitGroup
	i := 0
	for range 10 {
		wg.Add(1)
		go task(i, &wg) // Start a goroutine for each task runs concurrently
		i++
	}
	wg.Wait()
	fmt.Println("Main function is running")
}
