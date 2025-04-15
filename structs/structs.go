package main

import (
	"fmt"
	"time"
)

type UserData struct {
	FirstName   string
	LastName    string
	DateOfBirth string
	CreatedAt   time.Time
	Role        string
}

func getUserData(promtText string) string {
	fmt.Print(promtText)
	var input string
	fmt.Scan(&input)
	return input
}

func outputUserData(u *UserData) {

	fmt.Println(u.FirstName, u.LastName, u.DateOfBirth, u.CreatedAt)
	if u.Role != "" {
		fmt.Println("Role:", u.Role)
	}
}

func main() {
	firtName := getUserData("Enter your first name: ")
	lastName := getUserData("Enter your last name: ")
	dateOfBirth := getUserData("Enter your date of birth (YYYY-MM-DD): ")

	appUser := UserData{
		FirstName:   firtName,
		LastName:    lastName,
		DateOfBirth: dateOfBirth,
		CreatedAt:   time.Now(),
	}

	outputUserData(&appUser)

}
