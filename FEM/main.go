package main

import (
	"onis-emem.com/go/fem-server/data"
)

func main() {

	// create an instructor
	//var max data.Instructor

	max := data.Instructor{Id: 1, LastName: "Jonny", Score: 5} // new instructor
	max.FirstName = "Maximiliano"

	print(max.Print())
}
