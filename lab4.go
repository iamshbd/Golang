package main

import "fmt"


func add(a, b float64) float64 {
	return a + b
}
func sub(a, b float64) float64 {

	return a - b
}

func mul(a, b float64) float64 {
	return a * b
}
func div(a, b float64) float64 {
	if b == 0 {
		fmt.Println("Error: division by zero")
		return 0
	}
	return a / b


}

func mod(a, b int) int {

	return a % b
}

func main() {
	var choice int
	for {
		fmt.Println("===== Mini Calculator =====")
		fmt.Println("1) Add")
		fmt.Println("2) Sub")

		fmt.Println("3) Mul")
		fmt.Println("4) Div")
		fmt.Println("5) Mod")
		fmt.Println("6) Exit")
		fmt.Print("Choose: ")
		fmt.Scan(&choice)

		if choice == 6 {
			fmt.Println("Bye!")
			break
		}

		var a, b int
		fmt.Print("Enter a: ")
		fmt.Scan(&a)
		fmt.Print("Enter b: ")
		fmt.Scan(&b)


		switch choice {
		case 1:
			fmt.Println("Result:", add(float64(a), float64(b)))
		case 2:

			fmt.Println("Result:", sub(float64(a), float64(b)))
		case 3:
			fmt.Println("Result:", mul(float64(a), float64(b)))
		case 4:
			if b == 0 {
				fmt.Println("Error: division by zero")
			} else {
				fmt.Println("Result:", div(float64(a), float64(b)))
			}
		case 5:
			if b == 0 {
				fmt.Println("Error: cannot mod by zero")
			} else {
				fmt.Println("Result:", mod(a, b))
			}
		default:
			fmt.Println("Invalid operator!")
		}
		fmt.Println()
	}
}
