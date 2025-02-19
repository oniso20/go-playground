package data

import "time"

type Workshop struct {
	// Course Course
	//Let's use embed the course type
	Course
	Instructor
	Date time.Time
}

func NewWorkshop(name string, instructor Instructor) Workshop {
	w := Workshop{}
	w.Name = name
	w.Instructor = instructor

	return w
}
