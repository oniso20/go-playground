package main

import "fmt"

var work string = "Jonathan"

func calculateTax(price float32) (float32, float32) {
	return price * 0.99, price * 0.32
}

func calculateTax2(price float32) (tax3 float32, tax4 float32) {
	tax3 = price * 0.99
	tax4 = price * 0.32
	return
}

func birthday(age *int) int {
	*age++

	// using Printf
	fmt.Printf("The pointer is %v and the value is %v\n", age, *age)

	return *age
}

func anotherBirthday(oldAge int) int {
	oldAge++

	return oldAge
}

func init() {
	// use the underscore to avoid reading some values
	_, tax2 := calculateTax(100)
	tax3, tax4 := calculateTax2(200)

	fmt.Println(tax2, tax3, tax4)
}

func main() {
	tax1, tax2 := calculateTax(100)
	age := 40
	oldAge := 22
	// Copy needing a pointer
	birthday(&age)

	// store in a variable and need no pointer
	newAge := anotherBirthday(oldAge)

	if age > 40 {
		panic("Too old for this class")
	}

	defer fmt.Println("Thank you")

	fmt.Println(tax1, tax2, work, age, newAge)

}
