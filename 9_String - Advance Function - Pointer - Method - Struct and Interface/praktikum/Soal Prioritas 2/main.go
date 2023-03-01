package main

import (
	"fmt"
	"strings"
)

func caesar(offset int, input string) string {
	databaseAlphabet := "abcdefghijklmnopqrstuvwxyz"
	chiperText := ""
	for _, alphabet := range input {
		indexChiper := offset + strings.Index(databaseAlphabet, string(alphabet))
		if indexChiper > 25 {
			indexChiper %= 26
			chiperText += string(databaseAlphabet[indexChiper])
		} else {
			chiperText += string(databaseAlphabet[indexChiper])
		}

	}
	return chiperText
}

func main() {
	fmt.Println(caesar(3, "abc"))
	fmt.Println(caesar(10, "alterraacademy"))
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(caesar(4, "xyz"))
}
