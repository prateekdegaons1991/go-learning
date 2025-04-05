package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

func main() {

	fmt.Println("Constructor in golang")

	firstName, err := getUserData("Enter your first name: ")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	lastName, err := getUserData("Enter your last name: ")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	ageStr, err := getUserData("Enter your age: ")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		fmt.Println("Error: Invalid age input")
		return
	}
	email, err := getUserData("Enter your email: ")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	appUser, err := newUser(firstName, lastName, email, age)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	appUser.createdAt = time.Now().Truncate(time.Second)
	outputUserDetails(appUser)

}

type user struct {
	Name      string
	LastName  string
	Email     string
	Status    bool
	Age       int
	createdAt time.Time
}

func newUser(name, lastName, email string, age int) (*user, error) {

	if name == "" || lastName == "" || email == "" {
		fmt.Println("Error: All fields are required")
		return nil, errors.New("all fields are required")
	}
	if age <= 0 {
		fmt.Println("Error: Age must be a positive number")
		return nil, errors.New("age must be a positive number")
	}
	return &user{
		Name:     name,
		LastName: lastName,
		Email:    email,
		Status:   true,
		Age:      age,
	}, nil
}

// getUserData prompts the user for input and returns the input as a string.

func getUserData(prompt string) (string, error) {
	var input string
	print(prompt)
	_, err := fmt.Scanln(&input)
	if err != nil {
		return "", err
	}
	return input, nil
}

func outputUserDetails(u *user) {
	fmt.Println("User Details: ")
	fmt.Println("Name: ", u.Name, "", u.LastName, " Age: ", u.Age, " Email: ", u.Email, " Status: ", u.Status, " Created At: ", u.createdAt)
}
