package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func main() {

	// const inflationRate = 6.5

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

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	// solution before switching to the calculateFutureValues function
	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	//futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

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

// Functions can work without the arguments e.g. outputText() without the text, string.
func outputText(text string) {
	fmt.Print(text)
}

// This is one way of returning the fv and frv values in Go see other
// func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
// 	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
// 	frv := fv / math.Pow(1+inflationRate/100, years)

// 	return fv, frv
// }

// Notice how fv and frv are declared with the type from the start; hence, no need to create the variable anymore
func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, frv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	frv = fv / math.Pow(1+inflationRate/100, years)

	return fv, frv

	// return alone can work but it's difficult to read.
}
