package main

import "fmt"

func main() {
	var accountBalance float64 = 200
	var depositAmount float64
	var withdrawAmount float64
	var option int
	var isOpen bool = true

	for isOpen {
		fmt.Println("=====Hello Bank=====")
		fmt.Println("Please Select Command Option:")
		fmt.Println("1. Check Account Balance")
		fmt.Println("2. Deposit Funds")
		fmt.Println("3. Withdraw Funds")
		fmt.Println("4. Exit")

		fmt.Print("Input Option: ")
		fmt.Scan(&option)
		fmt.Println("Option Selected:", option)
		if option == 1 {
			fmt.Println("Account Balance: ", accountBalance)
		} else if option == 2 {
			fmt.Print("Please Enter the Deposit Amount: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("---Must greater than zero!---")
				continue
			}
			fmt.Println("Deposited Funds: ", depositAmount)
			accountBalance += depositAmount
			fmt.Println("Account Balance ", accountBalance)
		} else if option == 3 {
			fmt.Print("Please Enter the Withdraw amount: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount > accountBalance {
				fmt.Println("---Not Enough Balance!---")
				continue
			}
			if withdrawAmount <= 0 {
				fmt.Println("---Must greater than zero!---")
				continue
			}
			fmt.Println("Withdraw Funds: ", withdrawAmount)
			accountBalance -= withdrawAmount
			fmt.Println("Account Balance ", accountBalance)
		} else if option == 4 {
			isOpen = false
			fmt.Println("Exited")
		} else {
			fmt.Println("---Invalid Option!---")
		}
	}
}
