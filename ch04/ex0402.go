// Ex 14.2
//
// Notes:
// Cannot use function pointer in go.

package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	line := input.Text()

	algorithm := "sha256"

	if len(os.Args) > 1 {
		algorithm = os.Args[1]
	}

	switch algorithm {
	case "sha384":
		fmt.Println("sha384")
		hash := sha512.Sum384([]byte(line))
		fmt.Printf("%x\n", hash)
	case "sha512":
		fmt.Println("sha512")
		hash := sha512.Sum512([]byte(line))
		fmt.Printf("%x\n", hash)
	case "sha256":
		fmt.Println("sha256")
		hash := sha256.Sum256([]byte(line))
		fmt.Printf("%x\n", hash)
	default:
		fmt.Printf("Unknown algorithm specified: %s\n", algorithm)
	}
}
