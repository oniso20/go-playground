package main

import (
	"fmt"

	"onis-emem.com/go/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// Parsing values into the struct
	// appUser := user{
	// 	firstName: userFirstName,
	// 	lastName:  userLastName,
	// 	birthDate: userBirthdate,
	// 	createdAt: time.Now(),
	// 	updatedAt: time.Now(),
	// }

	// Using the new constructor
	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("test@example.com", "test123")

	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

	// the above can still work like this if the variable follow the struct structure set in type user for instance
	// appUser = user{
	// 	userFirstName,
	// 	userLastName,
	// 	userBirthdate,
	// 	time.Now(),
	// 	time.Now(),
	// }

	// For a null value for our struct we can use the code below. This means all fields will be null
	// appUser = user{}

	// outputUserDetails(&appUser) // Calling as a function
	appUser.OutputUserDetails() // using as a method. You don't need to pass any parameter like a function above.

	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	// fmt.Scan(&value)   // does not stop waiting for new entry when you hit enter key
	fmt.Scanln(&value) // Hitting the enter key will complete the input action
	return value
}
