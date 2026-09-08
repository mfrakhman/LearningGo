package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "secondLecture/Balance.txt"

func writeBalanceToFile(balance float64) {
	balanceTxt := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceTxt), 0644)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		fmt.Println(err)
		fmt.Println("init new account balance...")
		return 0.0, err
	}
	balanceTxt := string(data)
	balance, err := strconv.ParseFloat(balanceTxt, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return 0.0, err
	}
	return balance, nil
}

func main() {
	var accountBalance, err = getBalanceFromFile()
	if err != nil {
		fmt.Println("Failed to get account balance:", err)
		return
	}
	var depositAmount float64
	var withdrawAmount float64
	var option int
	var isOpen bool = true

	for isOpen {
		presentOption()

		fmt.Print("Input Option: ")
		_, err := fmt.Scan(&option)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

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
			writeBalanceToFile(accountBalance)
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
			writeBalanceToFile(accountBalance)
		} else if option == 4 {
			isOpen = false
			fmt.Println("---Exited, Bye!---")
		} else {
			fmt.Println("---Invalid Option!---")
		}
	}
}
