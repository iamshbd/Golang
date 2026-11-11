package crack

import (
	"bufio"
	"crypto/sha1" 
	"fmt"
	"os"
)

func SHA1Cracker(filePath string, targetHash string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Could not open file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		hashed := fmt.Sprintf("%x", sha1.Sum([]byte(word)))
		
		fmt.Println("[SHA1] trying:", word, "->", hashed)

		if hashed == targetHash {
			fmt.Println("\n>>> found SHA1 password:", word, "<<<")
			return
		}
	}
	fmt.Println("\n>>> SHA1 not found <<<")
}
