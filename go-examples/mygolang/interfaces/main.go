package main

import "fmt"

type payment struct {
	// This struct can be extended to include more fields related to payment if needed
	gateway razorpay
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

type phonePe struct{}

func (p phonePe) pay(amount float32) {
	// Process payment logic here
	fmt.Println("Payment made using PhonePe:", amount)
}

func main() {

	razorpayGw := razorpay{}
	newPayment := payment{gateway: razorpayGw}
	newPayment.makePayment(100.50)
	phonePeGw := phonePe{}
	newPayment.gateway = phonePeGw
	newPayment.makePayment(200.75)
	fmt.Println("Payment processing completed.")

}
