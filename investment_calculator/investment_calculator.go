package main

import (
	"fmt"
	"math"
)

func main() {

	const inflationRate = 6.5

	// investmentAmount, years := 1000.0, 10.0
	// years := 10.0
	// expectedReturnRate := 5.5

	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	//fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	//fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	// fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// fmt.Println("Future Value", futureValue)
	// fmt.Println("Future Real Value", futureRealValue)
	// fmt.Printf("Future Value: %.1f\nFuture Real Value: %.0f", futureValue, futureRealValue)

	// Using Sprint when storing values in variables
	// formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	// formattedFRV := fmt.Sprintf("Future Real Value: %.0f", futureRealValue)
	// fmt.Print(formattedFV, formattedFRV)

	// Print using multi lines with back tick
	fmt.Printf(`
	Future Value: %.1f
	Future Real Value: %.0f
	`, futureValue, futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}
