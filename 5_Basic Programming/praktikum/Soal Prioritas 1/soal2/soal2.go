package main

import "fmt"

func main() {
	fmt.Println("====== Program Deteksi Tipe Bilangan ====== ")

	var value int64
	fmt.Print("Masukkan angka : ")
	_, cekValue := fmt.Scan(&value)
	if cekValue != nil {
		fmt.Println("Input yang dimasukkan harus berupa angka!")
		return
	}
	if value%2 == 0 {
		fmt.Printf("Angka %v adalah bilangan GENAP.", value)
	} else {
		fmt.Printf("Angka %v adalah bilangan GANJIL.", value)
	}
}
