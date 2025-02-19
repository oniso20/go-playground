package main

import (
	"fmt"

	"onis-emem.com/go/fem-server/data"
)

func main() {

	// create an instructor
	//var max data.Instructor

	max := data.Instructor{Id: 1, LastName: "Jonny", Score: 5} // new instructor
	max.FirstName = "Maximiliano"

	doe := data.NewInstructor("Jane", "Doe")

	goCourse := data.Course{Id: 2, Name: "Go for Beginners", Instructor: max}

	fmt.Printf("%v", goCourse)

	// swiftWS := data.Workshop{
	// 	course: data.Course{Name: "Swift for IOS", Instructor: max}, Date: time.Now(),
	// }

	swiftWS := data.NewWorkshop("Swift with iOS", max)

	fmt.Printf("%v", swiftWS)
	print(max.Print())
	print(doe.Print())
}
