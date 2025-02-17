package data

import "fmt"

type Instructor struct {
	Id        int
	FirstName string
	LastName  string
	Score     int
}

func (i Instructor) Print() string {
	return fmt.Sprintf("%v, %v, (%d)", i.FirstName, i.LastName, i.Score)
}
