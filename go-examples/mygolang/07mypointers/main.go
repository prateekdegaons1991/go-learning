package main

import "fmt"

func main() {
	fmt.Println("Welcome to pinters")

	myNumber := 23

	var ptr = &myNumber

	*ptr = *ptr + 2

	fmt.Println("Value of ptr is: ", *ptr)

	myPhoneNum := 7411455742

	myPhoneNumPtr := &myPhoneNum

	fmt.Println("Value of myPhoneNumPtr is: ", *myPhoneNumPtr)

	getMyPhoneNum(myPhoneNumPtr)
	fmt.Println("Value of myPhoneNum is: ", myPhoneNum)

}

func getMyPhoneNum(myPhoneNum *int) int {

	return *myPhoneNum
}