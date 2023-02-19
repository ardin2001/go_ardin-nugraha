package main

import "fmt"

func main() {
	fmt.Println("Sample Test Cases")
	var value int
	fmt.Print("angka : ")
	_, cekValue := fmt.Scan(&value)
	if cekValue != nil {
		fmt.Println("Input : ")
		return
	}
	fmt.Println("Output :")
	for i := 1; i <= value; i++ {
		for spasi := 0; spasi < value-i; spasi++ {
			fmt.Print(" ")
		}
		for j := value; j >= 1; j-- {
			if i >= j {
				fmt.Print("* ")
			}
		}
		fmt.Print("\n")
	}
}
