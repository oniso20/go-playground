package main

import "fmt"

var Countries [10]string

// var slice []int
var Codes map[int]bool

func init() {
	Countries[0] = "Argentina"
	Countries[1] = "Brazil"
	Countries[2] = "USA"
	Countries[3] = "Canada"
	Countries[4] = "Collumbia"

	num := len(Countries)

	fmt.Println(Countries, num)
}
