package main

import (
	"fmt"
	"hashlab/utils/crack" 
)

func main() {

	wordlist := "../nord_vpn.txt" 

	fmt.Println("--- Lab 1 ---")
	crack.MD5Cracker(wordlist, "6a85dfd77d9cb35770c9dc6728d73d3f")

	fmt.Println("--- Lab 2 ---")
	crack.SHA1Cracker(wordlist, "aa1c7d931cf140bb35a5a16adeb83a551649c3b9")

	fmt.Println("--- Lab 3 ---")
	hash512 := "485f5c36c6f8474f53a3b0e361369ee3e32ee0de2f368b87b847dd23cb284b316bb0f026ada27df76c31ae8fa8696708d14b4d8fa352dbd8a31991b90ca5dd38"
	crack.SHA512Cracker(wordlist, hash512)
}
