package main

import "fmt"

func main() {
	var value int
	fmt.Print("angka : ")
	_, cekValue := fmt.Scan(&value)
	if cekValue != nil {
		fmt.Println("Input yang dimasukkan harus berupa angka!")
		return
	}
	for i := 1; i <= value; i++ {
		if value%i == 0 {
			fmt.Println(i)
		}
	}
}
