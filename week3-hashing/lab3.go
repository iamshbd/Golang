package crack
import (
	"bufio"
	"crypto/sha512"
	"fmt"
	"os"
)

func SHA512Cracker(filePath string, targetHash string) {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("File error")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		word := scanner.Text()
		hashed := fmt.Sprintf("%x", sha512.Sum512([]byte(word)))
		
		fmt.Println("[SHA512] trying:", word, "->", hashed)

		if hashed == targetHash {
			fmt.Println("\n>>> found SHA512 password:", word, "<<<")
			return
		}
	}
	fmt.Println("\n>>> SHA512 not found <<<")
}
