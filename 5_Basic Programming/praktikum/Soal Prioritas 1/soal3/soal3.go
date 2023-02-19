package main

import "fmt"

func main() {
	fmt.Println("====== Program Grade Nilai ====== ")
	var grade int64
	fmt.Print("Masukkan nilai : ")
	_, cekGrade := fmt.Scan(&grade)
	if cekGrade != nil {
		fmt.Println("Input yang dimasukkan harus berupa angka!")
		return
	}
	if grade >= 80 && grade <= 100 {
		fmt.Println("A")
	} else if grade >= 65 && grade <= 79 {
		fmt.Println("B")
	} else if grade >= 50 && grade <= 64 {
		fmt.Println("C")
	} else if grade >= 35 && grade <= 49 {
		fmt.Println("D")
	} else if grade >= 0 && grade <= 34 {
		fmt.Println("E")
	} else {
		fmt.Println("Nilai Invalid")
	}
}
