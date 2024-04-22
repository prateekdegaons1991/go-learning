package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of go language")

	presentTime := time.Now()
	fmt.Println(presentTime.Format("2006-01-02 15:04:05 Monday Jan"))

	fmt.Println("Welcome to pizza shop")
	

}
