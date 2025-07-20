package main

import (
	"fmt"
)

type paymenter interface {
	pay(amount float32)
	refund(amount float32)
}

type payment struct {
	// This struct can be extended to include more fields related to payment if needed
	gateway paymenter
}

// Open close principle: The payment struct can be extended to support new payment gateways without modifying existing code.
func (p payment) makePayment(amount float32) {
	// Process payment logic here
	// razorpayGw := razorpay{}
	// razorpayGw.pay(amount)
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	// Process payment logic here
	fmt.Println("Payment made using Razorpay:", amount)
}

func (r razorpay) refund(amount float32) {
	// Process refund logic here
	fmt.Println("Refund initiated using Razorpay:", amount)
}

type phonePe struct{}

func (p phonePe) pay(amount float32) {
	// Process payment logic here
	fmt.Println("Payment made using PhonePe:", amount)
}

func (p phonePe) refund(amount float32) {
	// Process refund logic here
	fmt.Println("Refund initiated using PhonePe:", amount)
}

func main() {

	phonePeGw := phonePe{}
	razorpayGw := razorpay{}

	phonePeGw.pay(100.50)
	razorpayGw.pay(150.75)
	phonePeGw.refund(50.00)
	razorpayGw.refund(75.00)

}
