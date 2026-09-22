package fileops

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "Balance.txt"

func WriteBalanceToFile(balance float64) {
	balanceTxt := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceTxt), 0644)
}

func GetBalanceFromFile() (float64, error) {
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
