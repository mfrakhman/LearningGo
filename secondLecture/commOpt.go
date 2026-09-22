package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

func presentOption() {
	fmt.Println("=====Hello Bank=====")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())
	fmt.Println("Please Select Command Option:")
	fmt.Println("1. Check Account Balance")
	fmt.Println("2. Deposit Funds")
	fmt.Println("3. Withdraw Funds")
	fmt.Println("4. Exit")
}
