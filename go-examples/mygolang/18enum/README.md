# Go Enums Example

This folder demonstrates how to implement and use enums in Go, since Go does not have native enum types like some other languages. Instead, Go uses `const` declarations with custom types.

## Files

- `main.go`: Entry point of the example. Shows how to use the defined enum type.
- `status.go`: Contains the enum definition using `iota` and a custom type.
- `README`: This file.

## Example: Defining and Using Enums in Go

### status.go
```go
// This is a simple example of using enums in Go.
package main

import "fmt"

// Define a custom type for order status.
type OrderStatus int

// Declare possible values using iota.
const (
    Received OrderStatus = iota
    Processing
    Shipped
    Delivered
    Cancelled
    Returned
)

// Function to change and print order status.
func changeOrderStatus(status OrderStatus) {
    fmt.Println("Order status changed to:", status)
}

func main() {
    fmt.Println("Lets learn, enum in Go!")
    changeOrderStatus(Returned)
}
```

**Explanation:**

- Go does not have built-in enums, but you can use `const` with a custom type and `iota` to create a set of related constants.
- `OrderStatus` is a custom type representing possible order states.
- The `const` block assigns sequential integer values to each status.
- The `changeOrderStatus` function demonstrates how to use the enum type in practice.
- The `main` function prints a message and changes the order status to `Returned`.
