package main

import "fmt"


func xorEncrypt(text string, key byte) string {
	result := ""
	for i := 0; i < len(text); i++ {
		result += string(text[i] ^ key) 


	}
	return result
}

func main() {
	var message string
	var keyInput string
	fmt.Print("Enter plaintext: ")
	fmt.Scan(&message)

	fmt.Print("Enter a single character key: ")
	fmt.Scan(&keyInput)

	key := keyInput[0]

	cipher := xorEncrypt(message, key)
	fmt.Println("Ciphertext (raw):", cipher)


	fmt.Print("Ciphertext bytes: ")
	for i := 0; i < len(cipher); i++ {
		fmt.Printf("%02x ", cipher[i])
	}
	fmt.Println()


	original := xorEncrypt(cipher, key)
	fmt.Println("Decrypted:", original)
}
