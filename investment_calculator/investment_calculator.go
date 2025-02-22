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

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
