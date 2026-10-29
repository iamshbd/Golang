package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var text string
	fmt.Print("Enter a string: ")
	fmt.Scan(&text)
	fmt.Println("--- Binary ---")
	for _, ch := range text {

		fmt.Printf("%08b ", ch)
	}
	fmt.Println()


	fmt.Println("--- Hex ---")
	for _, ch := range text {
		fmt.Printf("%02x ", ch)
	}
	fmt.Println()


	fmt.Println("--- Base64 ---")

	encoded := base64.StdEncoding.EncodeToString([]byte(text))
	fmt.Println(encoded)
}
