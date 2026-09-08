package main

import (
	"fmt"
	"math"
)

func main() {
	var years float64
	var expectedReturnRate float64
	var investmentAmount float64
	const inflationRate = 2.8

	fmt.Print("Enter Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter Years of Investment: ")
	fmt.Scan(&years)

	fmt.Print("Enter Expected Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calcInvestmentValue(investmentAmount, years, expectedReturnRate, inflationRate)
	//fmt.Println("by function calculation value: ", byfunctionValue)

	//futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	//futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	//fmt.Println("Future Value: ", futureValue)
	//fmt.Println("Future Real Value: ", futureRealValue)

	//fmt.Printf("Future Value: %.2f\nFuture Adjusted Inflation Value: %.4f\n", futureValue, futureRealValue)

	formattedValue := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedRealValue := fmt.Sprintf("Future Adjusted Inflation Value: %.4f\n", futureRealValue)

	fmt.Print(formattedValue, formattedRealValue)
}

// function example
func calcInvestmentValue(amount, duration, returnRate, inflation float64) (float64, float64) {
	value := amount * math.Pow(1+returnRate/100, duration)
	realValue := value / math.Pow(1+inflation/100, duration)
	return value, realValue
}

// method example
