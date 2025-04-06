package user

import (
	"errors"
	"fmt"
	"time"
)

type Admin struct {
	password string
	isActive bool
	Person
}

type contactInfo struct {
	email string
	phone string
}

type Person struct {
	firstName string
	lastName  string
	age       int
	createdAt time.Time
	contactInfo
}

// Create constructor for Person or utility function

func NewPerson(firstName string, lastName string, age int, email string, phone string) (*Person, error) {

	if firstName == "" || lastName == "" {
		fmt.Println("All fields are required.")
		return nil, errors.New("all fields are required")
	}
	if age < 0 {
		fmt.Println("Age cannot be negative.")
		return nil, errors.New("age cannot be negative")
	}
	if len(phone) < 10 {
		fmt.Println("Phone number must be at least 10 digits long.")
		return nil, errors.New("phone number must be at least 10 digits long")
	}

	return &Person{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
		createdAt: time.Now(),
		contactInfo: contactInfo{
			email: email,
			phone: phone,
		},
	}, nil
}

func GetUserData(promtText string) string {
	var userInput string
	fmt.Print(promtText)
	fmt.Scanln(&userInput)
	return userInput
}

//write getuserdetails outputnction, and this fucntion will take input fro struct

func (u *Person) OutputUserDetails() {
	fmt.Println("User details:", u.firstName, u.lastName, u.age, u.createdAt, u.contactInfo.email, u.contactInfo.phone)
}
