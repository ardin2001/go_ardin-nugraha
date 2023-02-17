package main

import "fmt"

func main() {
	var words string
	var result string = "Kata Palindrome"
	fmt.Println("Apakah Palindrome?")
	fmt.Print("Masukkan kata : ")
	_, cekWords := fmt.Scanln(&words)
	if cekWords != nil {
		fmt.Println("Input yang dimasukkan harus berupa kata!")
		return
	}
	for i := 0; i < len(words); i++ {
		for j := len(words) - 1 - i; j >= 0; j-- {
			if string(words[i]) != string(words[j]) {
				fmt.Println("Kata tidak palindrome")
				return
			}
			break
		}
	}
	fmt.Println(result)
}
