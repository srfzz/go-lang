package main

import "fmt"

func main() {
	var accountBalance float64 = 10000.00
	fmt.Println("************WELCOME TO GO BANK*************")
	fmt.Println("What do you wana do ")
	fmt.Println("1.Check Balance")
	fmt.Println("2.Deposit Money")
	fmt.Println("3.Withdraw Money")
	fmt.Println("4.Exit")
	var choice int32
	fmt.Println("Please Select Your choice :")
	fmt.Scan(&choice)
	fmt.Printf("You Have Selectd %d \n", choice)
	if choice == 1 {
		fmt.Printf("Your Balance is : %.2f \n", accountBalance)
	} else if choice == 2 {
		fmt.Println("Enter Money You wanna deposit")
		var moneyToBeDeposited float64
		fmt.Scan(&moneyToBeDeposited)
		if moneyToBeDeposited > 0 {
			accountBalance += moneyToBeDeposited
			fmt.Printf("New Balanace is : %.3f", accountBalance)
		} else {
			fmt.Println("Money to be Deposited Should be Greater Than Zero !")
		}
	} else if choice == 3 {
		var withdrawAmount float64
		fmt.Println("Enter Money You Wnna Withdraw !")
		fmt.Scan(&withdrawAmount)
		if withdrawAmount > accountBalance {
			fmt.Println("You Dont Have Enough Amount !")
		} else {
			accountBalance -= withdrawAmount
			fmt.Printf("Succesfull and New AMount Balance is : %.2f", accountBalance)
		}
	} else {
		fmt.Println("Thank you !")
	}
}
