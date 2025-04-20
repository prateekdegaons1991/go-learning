package main

import "fmt"

func main() {
	fmt.Println("Methods in Go	")
	Pratik := User{
		Name:   "Pratik",
		Email:  "prateekde@gmail.com",
		Status: true,
		Age:    22,
	}

	fmt.Println("Name: ", Pratik.Name, "Email: ", Pratik.Email, "Status: ", Pratik.Status, "Age: ", Pratik.Age)
	Pratik.GetStatus()
	Pratik.NewEmail("pratik@go.dev")
	fmt.Println("Name: ", Pratik.Name, "Email: ", Pratik.Email, "Status: ", Pratik.Status, "Age: ", Pratik.Age)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// methdo receiver
func (u *User) GetStatus() {
	fmt.Println("User Status: ", u.Status)
}

// method to pass new email

func (u *User) NewEmail(email string) {
	u.Email = email
	fmt.Println("New Email: ", u.Email)
}
