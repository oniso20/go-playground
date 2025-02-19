package data

import "fmt"

type Instructor struct {
	Id        int
	FirstName string
	LastName  string
	Score     int
}

func NewInstructor(name string, surname string) Instructor {
	return Instructor{FirstName: "John", LastName: "Doe"}
}

func (i Instructor) Print() string {
	return fmt.Sprintf("%v, %v, (%d)", i.FirstName, i.LastName, i.Score)
}
