package main

import "fmt"

const LoginToken string = "somfiabegibeberishvalaer" // Public 

func main() {
	var username string = "Pratik"
	fmt.Println(username)
	fmt.Printf("Veriable is of type : %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Veriable is of type : %T \n", isLoggedIn)

	var smallVal uint8 = 254
	fmt.Println(smallVal)
	fmt.Printf("Veriable is of type : %T \n", smallVal)


	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Veriable is of type : %T \n", anotherVariable)

	// implicit type

	var website = "demo website hi"
	fmt.Println(website)
	fmt.Printf("Veriable is of type : %T \n", website)

	// no var style

	numberOfUser := 24
	fmt.Println(numberOfUser)
	fmt.Printf("Veriable is of type : %T \n", numberOfUser)

	fmt.Printf(LoginToken  + website)
	fmt.Printf("Veriable is of type : %T \n", LoginToken)



}
