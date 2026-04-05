package main

import "fmt"

func main() {
	fmt.Println("************WELCOME TO GO BANK*************")
	fmt.Println("What do you wana do ")
	fmt.Println("1.Check Balance")
	fmt.Println("2.Deposit Money")
	fmt.Println("3.Withdraw Money")
	fmt.Println("4.Exit")
	var choice int32
	fmt.Println("Please Select Your choice")
	fmt.Scan(&choice)
	fmt.Printf("You Have Selectd %d", choice)
}
