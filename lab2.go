package main

import (
	"fmt"
	"os"
)

func main() {
	var num1 int
	var num2 int

	fmt.Print("Enter first integer: ")
	fmt.Scan(&num1)
	fmt.Print("Enter second integer: ")
	fmt.Scan(&num2)

	fmt.Println("You entered:", num1, "and", num2)

	
	num1 = num1
	num2 = num2

	
	fmt.Println("--- Checking && (both positive) ---")
	if num1 > 0 && num2 > 0 {
		fmt.Println("Result:", true)
	} else {
		fmt.Println("Result:", false)
	}

	fmt.Println("--- Checking || (one greater than the other) ---")
	if num1 > num2 || num2 > num1 {
		fmt.Println("Result:", true)
	} else {
		fmt.Println("Result:", false)
	}

	fmt.Println("--- Checking ! (not equal) ---")
	isEqual := num1 == num2
	if !isEqual {
		fmt.Println("Result: They are not equal! (True)")
	} else {
		fmt.Println("Result: They are equal! (False)")
	}
}
