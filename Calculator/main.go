package main

import "fmt"

func main() {
	var operator string
	var firstNumber, secondNumber int

	fmt.Println("Welcome to Onis first Calculator on Go")
	fmt.Println("======================================")

	fmt.Println("What operation would you want to perform? (+,-,*,/)")
	fmt.Scanf("%s", &operator)

	fmt.Println("Please enter the First number:")
	fmt.Scanf("%d", &firstNumber)

	fmt.Println("Please enter the Second number:")
	fmt.Scanf("%d", &secondNumber)

	fmt.Println("=========Calculating=============")
	fmt.Println("=============3===================")
	fmt.Println("=============2===================")
	fmt.Println("=============1===================")
	fmt.Println("")

	switch operator {
	case "+":
		fmt.Println("The total is: ", firstNumber+secondNumber)
	case "-":
		fmt.Println("The answer is: ", firstNumber-secondNumber)
	case "*":
		fmt.Println("The product is: ", firstNumber*secondNumber)
	case "/":
		fmt.Println("The answer is: ", firstNumber/secondNumber)
	}

}
