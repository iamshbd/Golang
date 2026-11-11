package crack
import (
	"bufio"
	"crypto/md5"
	"fmt"
	"os"
	"strings"
)

func MD5Cracker(filePath string, targetHash string) {
	file, err:=os.Open(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer file.Close()

	targetHash = strings.ToLower(targetHash)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word:= scanner.Text()
		hashed := fmt.Sprintf("%x", md5.Sum([]byte(word)))

		fmt.Println("[MD5] trying:", word, "->", hashed)

		if hashed == targetHash {
			fmt.Println("\n>>> found MD5 password:", word, "<<<")
			return
		}
	}
	fmt.Println("\n>>> MD5 not found <<<")
	
}
