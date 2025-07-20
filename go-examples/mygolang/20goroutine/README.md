# Goroutines with WaitGroup Example

This example demonstrates how to use goroutines in Go along with a `sync.WaitGroup` to wait for multiple concurrent tasks to finish.

## Example Code (`main.go`)
```go
package main

import (
    "fmt"
    "sync"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Printf("Worker %d starting\n", id)
    // Simulate some work
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }

    wg.Wait()
    fmt.Println("All workers finished")
}
```

### Explanation

- We define a `worker` function that takes an ID and a pointer to a `sync.WaitGroup`. It calls `wg.Done()` when finished.
- In `main`, we create a `WaitGroup` and start three goroutines, each running `worker`.
- `wg.Add(1)` increments the WaitGroup counter for each goroutine.
- `go worker(i, &wg)` starts the goroutine.
- `wg.Wait()` blocks until all goroutines have called `wg.Done()`.
- This ensures the main program waits for all workers to finish before exiting.

