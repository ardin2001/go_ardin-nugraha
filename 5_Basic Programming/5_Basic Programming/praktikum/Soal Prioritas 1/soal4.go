package main

import "fmt"

func main() {
	fmt.Println("====== Program Ganti Sebutan ====== ")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Printf("fizz ")
		} else if i%5 == 0 {
			fmt.Print("buzz ")
		} else {
			fmt.Print(i, " ")
		}
	}
}
