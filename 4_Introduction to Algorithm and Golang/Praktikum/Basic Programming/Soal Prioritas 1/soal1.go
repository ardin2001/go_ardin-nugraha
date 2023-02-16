package main

import "fmt"

func main() {
	fmt.Println("====== Program Menghitung Luas Trapesium ======")
	var A float64
	var B float64
	var t float64

	fmt.Print("Masukkan sisi sejajar(A) : ")
	_, cekA := fmt.Scan(&A)
	if cekA != nil {
		fmt.Println("Inputan A harus berupa angka!")
		return
	}
	fmt.Print("Masukkan sisi sejajar(B) : ")
	_, cekB := fmt.Scan(&B)
	if cekB != nil {
		fmt.Println("Inputan B harus berupa angka!")
		return
	}

	fmt.Print("Masukkan tinggi : ")
	_, cekC := fmt.Scan(&t)
	if cekC != nil {
		fmt.Println("Inputan tinggi harus berupa angka!")
		return
	}

	luasTrapesium := (A + B) * t / 2
	fmt.Println(luasTrapesium)
}
