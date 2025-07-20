package main

import "fmt"

// This is a simple example of using enums in Go.
type OrderStatus int

const (
	Received OrderStatus = iota
	Processing
	Shipped
	Delivered
	Cancelled
	Returned
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Order status changed to:", status)
}

func main() {
	fmt.Println("Lets learn, enum in Go!")
	changeOrderStatus(Returned)
}
