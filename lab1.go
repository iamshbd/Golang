package main

import "fmt"

func main() {
	var a int
	var b int

	fmt.Print("Enter first number: ")
	fmt.Scan(&a)
	fmt.Print("Enter second number: ")
	fmt.Scan(&b)
	fmt.Println("Starting values: a =", a, "b =", b)

	c := a
	fmt.Println("After c = a : c =", c)

	c += b
	fmt.Println("After c += b : c =", c)

	c -= b
	fmt.Println("After c -= b : c =", c)

	c *= b
	fmt.Println("After c *= b : c =", c)

	if b != 0 {
		c /= b
		fmt.Println("After c /= b : c =", c)

		c %= b
		fmt.Println("After c %= b : c =", c)
	} else {
		fmt.Println("Can't divide or modulo by 0")
	}
}
