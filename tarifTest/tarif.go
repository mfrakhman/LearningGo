package main

import "fmt"

func main() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	fmt.Println("1. HALO Max")
	fmt.Println("2. HALO Nasional")
	fmt.Println("3. SUPER KUOTA")
	fmt.Println("4. default")
	var packet int
	fmt.Print("Enter your packet: ")
	fmt.Scan(&packet)

	var duration int
	fmt.Print("Enter Duration: ")
	fmt.Scan(&duration)

	fmt.Println("1. Local")
	fmt.Println("2. Inter")
	var locPhone int
	fmt.Print("Enter your location phone number: ")
	fmt.Scan(&locPhone)

	var firstThirty = 0.0
	var secondNinty = 0.0
	var thirdOneTwenty = 0.0

	if duration <= 30 {
		firstThirty = float64(duration * 25)
	} else if duration <= 120 {
		firstThirty = float64(30 * 25)
		secondNinty = float64((duration - 30) * 20)
	} else if duration > 120 {
		firstThirty = float64(30 * 25)
		secondNinty = float64(90 * 20)
		thirdOneTwenty = float64((duration - 120) * 15)
	}
	if locPhone == 2 {
		firstThirty *= 1.25
		secondNinty *= 1.25
		thirdOneTwenty *= 1.25
	}
	var baseCost = firstThirty + secondNinty + thirdOneTwenty
	fmt.Println("Cost:", baseCost)

	if packet == 1 {
		if locPhone == 1 {
			firstThirty = 0.0
			secondNinty -= secondNinty * 0.1
			thirdOneTwenty -= thirdOneTwenty * 0.1
			fmt.Println("discounted by packet HALO MAX")
		} else if locPhone == 2 {
			firstThirty -= firstThirty * 0.1
			secondNinty -= secondNinty * 0.1
			thirdOneTwenty -= thirdOneTwenty * 0.1
		}
	} else if packet == 2 {
		if locPhone == 1 {
			firstThirty -= firstThirty * 0.2
			secondNinty -= secondNinty * 0.2
			thirdOneTwenty -= thirdOneTwenty * 0.2
		} else if locPhone == 2 {
			firstThirty = 0.0
			secondNinty -= secondNinty * 0.2
			thirdOneTwenty -= thirdOneTwenty * 0.2
			fmt.Println("discounted by packet HALO NASIONAL")
		}
	}

	var discountedCost = firstThirty + secondNinty + thirdOneTwenty
	fmt.Println("Discounted Cost:", discountedCost)
}
