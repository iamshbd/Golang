package main

import "fmt"

func myAND(a, b int) int {
	return a & b
}

func myOR(a, b int) int {
	return a | b


}

func myXor(a, b int) int {
	return a ^ b
}

func myNOT(a int) int {
	return ^a
}

func main() {
	var x int
	var y int

	fmt.Print("Enter first number: ")
	fmt.Scan(&x)
	fmt.Print("Enter second number: ")
	fmt.Scan(&y)

	fmt.Println("AND:", myAND(x, y))
	fmt.Println("OR:", myOR(x, y))
	fmt.Println("XOR:", myXor(x, y))

	fmt.Println("NOT of first:", myNOT(x))


	fmt.Println("Left shift (x << 1):", x<<1)
	fmt.Println("Right shift (x >> 1):", x>>1)


	fmt.Println("--- Assignment effects ---")
	z := x
	fmt.Println("z starts as:", z)

	z &= y
	fmt.Println("z &= y :", z)
	z |= y
	fmt.Println("z |= y :", z)
	z ^= y
	fmt.Println("z ^= y :", z)

	z <<= 2
	fmt.Println("z <<= 2 :", z)
	z >>= 2
	fmt.Println("z >>= 2 :", z)
}
