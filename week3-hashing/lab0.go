package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"golang.org/x/crypto/sha3"
)

func hashMD5(s string) string    {
	return fmt.Sprintf("%x", md5.Sum([]byte(s))) }
func hashSHA1(s string) string   {
	return fmt.Sprintf("%x", sha1.Sum([]byte(s))) }

func hashSHA256(s string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s))) }
func hashSHA512(s string) string {
	return fmt.Sprintf("%x", sha512.Sum512([]byte(s))) }

func hashSHA3(s string) string   { 
	return fmt.Sprintf("%x", sha3.Sum256([]byte(s))) }

func main() {
	var input1, input2 string

	fmt.Println("Name + Hashing Program")
	fmt.Print("Please input value 1: ")
	fmt.Scanln(&input1)
	
	fmt.Print("Please input value 2: ")
	fmt.Scanln(&input2)

	fmt.Println("\n--- Results ---")

	//md5 
	a, b := hashMD5(input1), hashMD5(input2)

	if a == b {
		fmt.Printf("Hash (MD5) : %s / %s => Match!\n", a, b)
	} else {
		fmt.Printf("Hash (MD5) : %s / %s => No Match!\n", a, b)
	}

	// sha1
	a, b = hashSHA1(input1), hashSHA1(input2)
	if a == b {
		fmt.Printf("Hash (SHA1) : %s / %s => Match!\n", a, b)
	} else {
		fmt.Printf("Hash (SHA1) : %s / %s => No Match!\n", a, b)
	}

	// sha256
	a, b = hashSHA256(input1), hashSHA256(input2)
	if a == b {
		fmt.Printf("Hash (SHA256) : %s / %s => Match!\n", a, b)
	} else {

		fmt.Printf("Hash (SHA256) : %s / %s => No Match!\n", a, b)
	}

	//sha512
	a, b = hashSHA512(input1), hashSHA512(input2)
	if a == b {
		fmt.Printf("Hash (SHA512) : %s / %s => Match!\n", a, b)

	} else {
		fmt.Printf("Hash (SHA512) : %s / %s => No Match!\n", a, b)
	}

	//sha3
	a, b = hashSHA3(input1), hashSHA3(input2)
	if a == b {
		fmt.Printf("Hash (SHA3) : %s / %s => Match!\n", a, b)
	} else {
		fmt.Printf("Hash (SHA3) : %s / %s => No Match!\n", a, b)
	}
}
