package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
	updatedAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
	u.birthDate = ""
}

// creating a constructor to create the struct rather than using the appUser with user struct method below.
// The constructor helps to setup reusable logic.
func New(firstName, lastName, birthDate string) (*User, error) {

	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last name and birthdate are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil // nil for the error
}

func NewAdmin(email, password string) *Admin {
	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName:  "Admin",
			birthDate: "---",
			createdAt: time.Now(),
			updatedAt: time.Now(),
		},
	}
}

// Using the struct as a method
func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt, u.updatedAt)
}

// Solution without structs
// func outputUserDetails(firstName, lastName, birthdate string) {
// 	fmt.Println(firstName, lastName, birthdate)
// }

// With structs where "u" is the parameter name and "user" is the parameter type
// func outputUserDetails(u *user) {
// 	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt, u.updatedAt)
// }
