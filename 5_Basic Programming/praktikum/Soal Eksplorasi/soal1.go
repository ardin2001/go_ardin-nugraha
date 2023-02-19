package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Apakah Palindrome?")
	fmt.Print("Masukkan kata : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	var container string
	for i := 0; i < len(text); i++ {
		if string(text[i]) == " " {
			continue
		}
		container += string(text[i])
	}
	for i := 0; i < len(text); i++ {
		for j := len(text) - 1 - i; j >= 0; j-- {
			if string(text[i]) != string(text[j]) {
				fmt.Println("Captured : ", text)
				fmt.Println("Tidak Palindrome")
				return
			}
			break
		}
	}
	fmt.Println("Captured :", text)
	fmt.Println("Palindrome")
}
